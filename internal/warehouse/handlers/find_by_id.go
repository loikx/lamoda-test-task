package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"
	"github.com/lamoda-tech/loikx/internal/warehouse/usecases"
)

type FindByIDHandler struct {
	useCase *usecases.FindByIDUseCase
}

func NewFindByIDHandler(useCase *usecases.FindByIDUseCase) *FindByIDHandler {
	return &FindByIDHandler{useCase: useCase}
}

func (handler *FindByIDHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	id, ok := mux.Vars(request)["id"]
	if !ok {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	uuidID, err := uuid.FromString(id)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write(
			[]byte(fmt.Sprintf("warehouse: id is invalid %s", id)),
		)
		return
	}

	products, err := handler.useCase.Handle(request.Context(), uuidID)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	body, err := json.Marshal(products)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write(body)
}
