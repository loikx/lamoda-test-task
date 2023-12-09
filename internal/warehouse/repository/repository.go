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

func NewWareHouseRepository(con *pgx.Conn) *WareHouseRepository {
	return &WareHouseRepository{
		con: con,
		mu:  sync.Mutex{},
	}
}

func (r *WareHouseRepository) FindByID(ctx context.Context, id uuid.UUID) ([]domain.Product, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var items []domain.Product

	rows, err := r.con.Query(
		ctx,
		"select * from product.product p "+
			"where exists(select 1 from product.warehouse w where w.id = p.warehouse_id and not p.is_reserved)",
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var product domain.Product
		err = rows.Scan(
			&product.ID,
			&product.Name,
			&product.Count,
			&product.Size.Length,
			&product.Size.Width,
			&product.Size.Height,
			&product.Size.Unit,
			&product.WarehouseID,
			&product.IsReserved,
		)
		if err != nil {
			return nil, err
		}

		items = append(items, product)
	}

	return items, nil
}
