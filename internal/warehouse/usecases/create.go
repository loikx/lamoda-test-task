package usecases

import (
	"context"
	"fmt"

	"github.com/lamoda-tech/loikx/internal/warehouse/domain"
)

type CreateWarehouseUseCase struct {
	warehouse domain.WarehouseRepository
}

type CreateWarehouseCommand struct {
	Name         string
	Availability bool
}

func NewCreateWarehouseUseCase(warehouse domain.WarehouseRepository) *CreateWarehouseUseCase {
	return &CreateWarehouseUseCase{warehouse: warehouse}
}

func (useCase *CreateWarehouseUseCase) Handle(
	ctx context.Context,
	command CreateWarehouseCommand,
) (*domain.Warehouse, error) {
	warehouse := domain.NewWarehouse(
		command.Name,
		command.Availability,
	)

	if err := warehouse.Validation(); err != nil {
		return nil, fmt.Errorf("warehouse: create warehouse %w", err)
	}

	err := useCase.warehouse.Save(ctx, warehouse)
	if err != nil {
		return nil, fmt.Errorf("warehouse: save warehouse %w", err)
	}

	return warehouse, nil
}
