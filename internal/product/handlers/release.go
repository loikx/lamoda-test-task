package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/lamoda-tech/loikx/internal/product/dto"
	"github.com/lamoda-tech/loikx/internal/product/usecases"
)

type ReleaseProductHandler struct {
	useCase *usecases.ReleaseUseCase
}

func NewReleaseProductHandler(useCase *usecases.ReleaseUseCase) *ReleaseProductHandler {
	return &ReleaseProductHandler{useCase: useCase}
}

func (handler *ReleaseProductHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var releaseDto dto.ReleaseDto
	if err := json.NewDecoder(request.Body).Decode(&releaseDto); err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	command := usecases.ReleaseCommand{}
	command.IDs = releaseDto.IDs

	if err := handler.useCase.Handle(request.Context(), command); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
}
