package usecases

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/lamoda-tech/loikx/internal/product/domain"
)

type FindByWarehouseResponse struct {
	Items []domain.Product `json:"items"`
	Count int              `json:"count"`
}

type FindByWareHouseUseCase struct {
	products domain.ProductRepository
}

func NewFindByWarehouseUseCase(products domain.ProductRepository) *FindByWareHouseUseCase {
	return &FindByWareHouseUseCase{products: products}
}

func (useCase *FindByWareHouseUseCase) Handle(ctx context.Context, id uuid.UUID) (*FindByWarehouseResponse, error) {
	products, err := useCase.products.FindByWarehouse(ctx, id)
	if err != nil {
		return nil, err
	}

	return &FindByWarehouseResponse{
		Items: products,
		Count: len(products),
	}, nil
}
