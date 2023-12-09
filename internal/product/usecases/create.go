package usecases

import (
	"context"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/lamoda-tech/loikx/internal/product/domain"
	"github.com/lamoda-tech/loikx/internal/product/pkg"
)

type CreateProductCommand struct {
	Name        string
	Count       int
	Size        pkg.Size
	WarehouseID uuid.UUID
}

type CreateProductUseCase struct {
	products domain.ProductRepository
}

func NewCreateProductUseCase(products domain.ProductRepository) *CreateProductUseCase {
	return &CreateProductUseCase{products: products}
}

func (useCase *CreateProductUseCase) Handle(
	ctx context.Context,
	command CreateProductCommand,
) (*domain.Product, error) {
	product := domain.NewProduct(
		command.Name,
		command.Count,
		command.Size,
		command.WarehouseID,
	)

	err := useCase.products.Save(ctx, product)
	if err != nil {
		return nil, fmt.Errorf("product: save product %w", err)
	}

	return product, nil
}
