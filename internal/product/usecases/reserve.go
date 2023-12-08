package usecases

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/lamoda-tech/loikx/internal/product/domain"
)

type ReserveUseCase struct {
	products domain.ProductRepository
}

type ReserveCommand struct {
	IDs []uuid.UUID
}

func NewReserveUseCase(repository domain.ProductRepository) *ReserveUseCase {
	return &ReserveUseCase{
		products: repository,
	}
}

func (useCase *ReserveUseCase) Handle(ctx context.Context, command ReserveCommand) error {
	return useCase.products.Reserve(ctx, command.IDs)
}
