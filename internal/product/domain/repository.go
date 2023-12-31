package domain

import (
	"context"
	"github.com/gofrs/uuid"
)

type ProductRepository interface {
	Release(ctx context.Context, ids []uuid.UUID) error
	Reserve(ctx context.Context, ids []uuid.UUID) error
	FindByWarehouse(ctx context.Context, id uuid.UUID) ([]*Product, error)
	Save(ctx context.Context, product *Product) error
	Delete(ctx context.Context, id uuid.UUID) error
}
