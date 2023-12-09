package repository

import (
	"context"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/lamoda-tech/loikx/internal/warehouse/domain"
)

type WarehouseRepository struct {
	conn *pgx.Conn
}

func NewWarehouseRepository(conn *pgx.Conn) *WarehouseRepository {
	return &WarehouseRepository{conn: conn}
}

func (r *WarehouseRepository) Save(ctx context.Context, warehouse *domain.Warehouse) error {
	tx, err := r.conn.Begin(ctx)
	if err != nil {
		return fmt.Errorf("warehouse: save transaction start fail %w", err)
	}

	_, err = tx.Exec(
		ctx,
		"insert into product.warehouse(id, name, availability) values ($1, $2, $3)",
		warehouse.ID, warehouse.Name, warehouse.Availability,
	)
	if err != nil {
		if err = tx.Rollback(ctx); err != nil {
			return fmt.Errorf("warehouse: save rollback transaction fail %w", err)
		}

		return fmt.Errorf("warehouse: save fail %w", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("warehouse: save commit fail %w", err)
	}

	return nil
}

func (r *WarehouseRepository) Delete(ctx context.Context, id uuid.UUID) error {
	tx, err := r.conn.Begin(ctx)
	if err != nil {
		return fmt.Errorf("warehouse: delete transaction start fail %w", err)
	}

	_, err = tx.Exec(
		ctx,
		"delete from product.warehouse where id=($1)",
		id,
	)
	if err != nil {
		if err = tx.Rollback(ctx); err != nil {
			return fmt.Errorf("warehouse: delete rollback transaction fail %w", err)
		}

		return fmt.Errorf("warehouse: delete fail %w", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("warehouse: delete commit fail %w", err)
	}

	return nil
}
