package repository

import (
	"context"
	"sync"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/lamoda-tech/loikx/internal/product/domain"
)

type WareHouseRepository struct {
	con *pgx.Conn

	mu sync.Mutex
}

func (r *WareHouseRepository) FindByID(ctx context.Context, id uuid.UUID) ([]domain.Product, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var items []domain.Product

	err := r.con.QueryRow(
		ctx,
		"select * from product p where exists(select 1 from warehouse w where w.id = p.warehouse_id and not p.is_reserved)",
	).Scan(&items)
	if err != nil {
		return nil, err
	}

	return items, nil
}
