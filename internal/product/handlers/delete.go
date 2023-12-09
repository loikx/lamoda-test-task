package handlers

import (
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"
	"github.com/lamoda-tech/loikx/internal/product/usecases"
)

type DeleteProductHandler struct {
	useCase *usecases.DeleteProductUseCase
}

func NewDeleteProductHandler(useCase *usecases.DeleteProductUseCase) *DeleteProductHandler {
	return &DeleteProductHandler{useCase: useCase}
}

func (handler *DeleteProductHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
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
