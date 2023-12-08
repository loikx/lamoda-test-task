package repository

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/lamoda-tech/loikx/internal/product/domain"
)

type WareHouseRepository struct {
	con *pgx.Conn
}

func (r *WareHouseRepository) FindByID(ctx context.Context, id uuid.UUID) ([]domain.Product, error) {
	panic("implement me")
}
