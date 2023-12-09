package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/lamoda-tech/loikx/internal/warehouse/dto"
	"github.com/lamoda-tech/loikx/internal/warehouse/usecases"
)

type CreateWarehouseHandler struct {
	useCase *usecases.CreateWarehouseUseCase
}

func NewCreateWarehouseHandler(useCase *usecases.CreateWarehouseUseCase) *CreateWarehouseHandler {
	return &CreateWarehouseHandler{useCase: useCase}
}

func (handler *CreateWarehouseHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var createDto dto.CreateCommandDto
	if err := json.NewDecoder(request.Body).Decode(&createDto); err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write(
			[]byte(err.Error()),
		)
		return
	}

	command := usecases.CreateWarehouseCommand{}
	command.Name = createDto.Name
	command.Availability = createDto.Availability

	warehouse, err := handler.useCase.Handle(request.Context(), command)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	warehouseMarshaled, err := json.Marshal(warehouse)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write(warehouseMarshaled)
}