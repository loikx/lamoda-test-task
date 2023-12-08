package handlers

import (
	"net/http"

	"github.com/lamoda-tech/loikx/internal/product/usecases"
)

type ReleaseProductHandler struct {
	useCase *usecases.ReleaseUseCase
}

func NewReleaseProductHandler(useCase *usecases.ReleaseUseCase) *ReleaseProductHandler {
	return &ReleaseProductHandler{useCase: useCase}
}

func (handler *ReleaseProductHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	panic("implement me")
}
