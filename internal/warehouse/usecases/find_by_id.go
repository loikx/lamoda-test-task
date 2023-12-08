package usecases

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/lamoda-tech/loikx/internal/product/domain"
	"github.com/lamoda-tech/loikx/internal/warehouse/repository"
)

type WareHouseResponse struct {
	Items []domain.Product `json:"items"`
	Count int              `json:"count"`
}

type FindByIDUseCase struct {
	warehouse *repository.WareHouseRepository
}

func NewFindByIDUseCase(warehouse *repository.WareHouseRepository) *FindByIDUseCase {
	return &FindByIDUseCase{warehouse: warehouse}
}

func (useCase *FindByIDUseCase) Handle(ctx context.Context, id uuid.UUID) (*WareHouseResponse, error) {
	products, err := useCase.warehouse.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &WareHouseResponse{
		Items: products,
		Count: len(products),
	}, nil
}
