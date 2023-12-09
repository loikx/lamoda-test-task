package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/lamoda-tech/loikx/internal/product/dto"
	"github.com/lamoda-tech/loikx/internal/product/usecases"
)

type ReserveProductHandler struct {
	useCase *usecases.ReserveUseCase
}

func NewReserveProductHandler(useCase *usecases.ReserveUseCase) *ReserveProductHandler {
	return &ReserveProductHandler{useCase: useCase}
}

func (handler *ReserveProductHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var reserveDto dto.ReserveDto
	if err := json.NewDecoder(request.Body).Decode(&reserveDto); err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write(
			[]byte(err.Error()),
		)
		return
	}

	command := usecases.ReserveCommand{}
	command.IDs = reserveDto.IDs

	if err := handler.useCase.Handle(request.Context(), command); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
}
