package usecases

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/lamoda-tech/loikx/internal/product/domain"
)

type DeleteProductUseCase struct {
	products domain.ProductRepository
}

func NewDeleteProductUseCase(products domain.ProductRepository) *DeleteProductUseCase {
	return &DeleteProductUseCase{products: products}
}

func (useCase *DeleteProductUseCase) Handle(ctx context.Context, id uuid.UUID) error {
	return useCase.products.Delete(ctx, id)
}
