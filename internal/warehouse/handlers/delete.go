package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"
	"github.com/lamoda-tech/loikx/internal/warehouse/usecases"
	"github.com/lamoda-tech/loikx/pkg/errors"
)

type DeleteWarehouseHandler struct {
	useCase *usecases.DeleteWarehouseUseCase
}

func NewDeleteWarehouseHandler(useCase *usecases.DeleteWarehouseUseCase) *DeleteWarehouseHandler {
	return &DeleteWarehouseHandler{useCase: useCase}
}

func (handler *DeleteWarehouseHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
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

	err = handler.useCase.Handle(request.Context(), uuidID)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	writer.WriteHeader(http.StatusNoContent)
}
