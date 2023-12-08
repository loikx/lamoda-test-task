package domain

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/lamoda-tech/loikx/internal/product/domain"
)

type WarehouseRepository interface {
	FindByID(ctx context.Context, id uuid.UUID) ([]domain.Product, error)
}
