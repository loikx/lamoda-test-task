package handlers

import (
	"encoding/json"
	"github.com/lamoda-tech/loikx/internal/warehouse/dto"
	"net/http"

	"github.com/lamoda-tech/loikx/internal/warehouse/usecases"
)

type DeleteWarehouseHandler struct {
	useCase *usecases.DeleteWarehouseUseCase
}

func NewDeleteWarehouseHandler(useCase *usecases.DeleteWarehouseUseCase) *DeleteWarehouseHandler {
	return &DeleteWarehouseHandler{useCase: useCase}
}

func (handler *DeleteWarehouseHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var deleteDto dto.DeleteDto
	if err := json.NewDecoder(request.Body).Decode(&deleteDto); err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write(
			[]byte(err.Error()),
		)
		return
	}

	command := usecases.DeleteWarehouseCommand{}
	command.ID = deleteDto.ID

	err := handler.useCase.Handle(request.Context(), command)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	writer.WriteHeader(http.StatusNoContent)
}
