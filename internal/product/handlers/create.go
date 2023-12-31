package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/lamoda-tech/loikx/internal/product/domain"
	"github.com/lamoda-tech/loikx/internal/product/dto"
	"github.com/lamoda-tech/loikx/internal/product/usecases"
	"github.com/lamoda-tech/loikx/pkg/errors"
)

type JsonCreateResponse struct {
	Product *domain.Product `json:"data"`
}

type CreateProductHandler struct {
	useCase *usecases.CreateProductUseCase
}

func NewCreateProductHandler(useCase *usecases.CreateProductUseCase) *CreateProductHandler {
	return &CreateProductHandler{useCase: useCase}
}

func (handler *CreateProductHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var createCommandDto dto.CreateCommandDto
	if err := json.NewDecoder(request.Body).Decode(&createCommandDto); err != nil {
		customError := errors.NewError(err)
		marshaledError, _ := json.Marshal(customError)

		writer.WriteHeader(http.StatusBadRequest)
		writer.Write(marshaledError)
		return
	}

	command := usecases.CreateProductCommand{}
	command.Name = createCommandDto.Name
	command.Count = createCommandDto.Count
	command.Size = createCommandDto.Size
	command.WarehouseID = createCommandDto.WarehouseID

	product, err := handler.useCase.Handle(request.Context(), command)
	if err != nil {
		customError := errors.NewError(err)
		marshaledError, _ := json.Marshal(customError)

		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write(marshaledError)
		return
	}

	response := JsonCreateResponse{
		Product: product,
	}

	marshaledResponse, err := json.Marshal(response)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write(marshaledResponse)
}
