package usecases

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/lamoda-tech/loikx/internal/product/domain"
)

type ReleaseUseCase struct {
	products domain.ProductRepository
}

type ReleaseCommand struct {
	IDs []uuid.UUID
}

func NewReleaseUseCase(products domain.ProductRepository) *ReleaseUseCase {
	return &ReleaseUseCase{products: products}
}

func (useCase *ReleaseUseCase) Handle(ctx context.Context, command ReleaseCommand) error {
	return useCase.products.Release(ctx, command.IDs)
}
