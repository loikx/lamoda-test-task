package handlers

import (
	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"
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
	id, ok := mux.Vars(request)["id"]
	if !ok {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	uuidID, err := uuid.FromString(id)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write(
			[]byte(err.Error()),
		)
		return
	}

	err = handler.useCase.Handle(request.Context(), uuidID)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	writer.WriteHeader(http.StatusNoContent)
}
