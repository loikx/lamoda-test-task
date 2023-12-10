package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"
	"github.com/lamoda-tech/loikx/internal/product/domain"
	"github.com/lamoda-tech/loikx/internal/product/usecases"
	"github.com/lamoda-tech/loikx/pkg/errors"
)

type JsonFindByWarehouseResponse struct {
	Items []*domain.Product `json:"data"`
	Count int               `json:"count"`
}

type FindByWarehouseHandler struct {
	useCase *usecases.FindByWarehouseUseCase
}

func NewFindByWarehouseHandler(useCase *usecases.FindByWarehouseUseCase) *FindByWarehouseHandler {
	return &FindByWarehouseHandler{useCase: useCase}
}

func (handler *FindByWarehouseHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	id, ok := mux.Vars(request)["id"]
	if !ok {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	uuidID, err := uuid.FromString(id)
	if err != nil {
		customError := errors.NewError(err)
		marshaledError, _ := json.Marshal(customError)

		writer.WriteHeader(http.StatusBadRequest)
		writer.Write(marshaledError)
		return
	}

	products, err := handler.useCase.Handle(request.Context(), uuidID)
	if err != nil {
		customError := errors.NewError(err)
		marshaledError, _ := json.Marshal(customError)

		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write(marshaledError)
		return
	}

	response := JsonFindByWarehouseResponse{
		Items: products.Items,
		Count: products.Count,
	}

	marshaledResponse, err := json.Marshal(response)
	if err != nil {
		customError := errors.NewError(err)
		marshaledError, _ := json.Marshal(customError)

		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write(marshaledError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write(marshaledResponse)
}
