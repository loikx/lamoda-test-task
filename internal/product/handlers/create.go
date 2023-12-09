package handlers

import (
	"encoding/json"
	"github.com/lamoda-tech/loikx/internal/product/dto"
	"github.com/lamoda-tech/loikx/internal/product/usecases"
	"net/http"
)

type CreateProductHandler struct {
	useCase *usecases.CreateProductUseCase
}

func NewCreateProductHandler(useCase *usecases.CreateProductUseCase) *CreateProductHandler {
	return &CreateProductHandler{useCase: useCase}
}

func (handler *CreateProductHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var createCommandDto dto.CreateCommandDto
	if err := json.NewDecoder(request.Body).Decode(&createCommandDto); err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write(
			[]byte(err.Error()),
		)
		return
	}

	command := usecases.CreateProductCommand{}
	command.Name = createCommandDto.Name
	command.Count = createCommandDto.Count
	command.Size = createCommandDto.Size
	command.WarehouseID = createCommandDto.WarehouseID

	product, err := handler.useCase.Handle(request.Context(), command)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	marshaledProduct, err := json.Marshal(product)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write(marshaledProduct)
}