package usecases

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/lamoda-tech/loikx/internal/warehouse/domain"
)

type DeleteWarehouseCommand struct {
	ID uuid.UUID
}

type DeleteWarehouseUseCase struct {
	warehouse domain.WarehouseRepository
}

func NewDeleteWarehouseUseCase(warehouse domain.WarehouseRepository) *DeleteWarehouseUseCase {
	return &DeleteWarehouseUseCase{warehouse: warehouse}
}

func (useCase *DeleteWarehouseUseCase) Handle(ctx context.Context, command DeleteWarehouseCommand) error {
	return useCase.warehouse.Delete(ctx, command.ID)
}
