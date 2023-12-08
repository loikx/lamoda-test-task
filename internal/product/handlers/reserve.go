package handlers

import (
	"github.com/lamoda-tech/loikx/internal/product/usecases"
	"net/http"
)

type ReserveProductHandler struct {
	useCase *usecases.ReserveUseCase
}

func NewReserveProductHandler(useCase *usecases.ReserveUseCase) *ReserveProductHandler {
	return &ReserveProductHandler{useCase: useCase}
}

func (handler *ReserveProductHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	panic("implement me")
}
