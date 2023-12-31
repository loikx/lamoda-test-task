package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/lamoda-tech/loikx/internal/warehouse/domain"
	"github.com/lamoda-tech/loikx/internal/warehouse/dto"
	"github.com/lamoda-tech/loikx/internal/warehouse/usecases"
	"github.com/lamoda-tech/loikx/pkg/errors"
)

type JsonCreateResponse struct {
	Warehouse *domain.Warehouse `json:"data,omitempty"`
}

type CreateWarehouseHandler struct {
	useCase *usecases.CreateWarehouseUseCase
}

func NewCreateWarehouseHandler(useCase *usecases.CreateWarehouseUseCase) *CreateWarehouseHandler {
	return &CreateWarehouseHandler{useCase: useCase}
}

func (handler *CreateWarehouseHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var createDto dto.CreateCommandDto
	if err := json.NewDecoder(request.Body).Decode(&createDto); err != nil {
		customError := errors.NewError(err)
		marshaledError, _ := json.Marshal(customError)

		writer.WriteHeader(http.StatusBadRequest)
		writer.Write(marshaledError)
		return
	}

	command := usecases.CreateWarehouseCommand{}
	command.Name = createDto.Name
	command.Availability = createDto.Availability

	warehouse, err := handler.useCase.Handle(request.Context(), command)
	if err != nil {
		customError := errors.NewError(err)
		marshaledError, _ := json.Marshal(customError)

		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write(marshaledError)
		return
	}

	response := JsonCreateResponse{
		Warehouse: warehouse,
	}

	marshaledResponse, err := json.Marshal(response)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write(marshaledResponse)
}
