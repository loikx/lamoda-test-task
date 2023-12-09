package domain

import (
	"context"

	"github.com/gofrs/uuid"
)

type WarehouseRepository interface {
	Save(ctx context.Context, warehouse *Warehouse) error
	Delete(ctx context.Context, id uuid.UUID) error
}
