package repository

import (
	"context"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5"
)

type ProductRepository struct {
	conn *pgx.Conn
}

func NewProductRepository(conn *pgx.Conn) *ProductRepository {
	return &ProductRepository{
		conn: conn,
	}
}

func (r *ProductRepository) Reserve(ctx context.Context, ids []uuid.UUID) error {
	tx, err := r.conn.Begin(ctx)
	if err != nil {
		return fmt.Errorf("product: reserve fail start transaction %w", err)
	}

	_, err = tx.Exec(
		ctx,
		"update product.product set is_reserved=true "+
			"where id=any($1) and warehouse_id=any(select id from product.warehouse where availability=true)",
		ids,
	)
	if err != nil {
		if err := tx.Rollback(ctx); err != nil {
			return fmt.Errorf("product: reserve rollback fail %w", err)
		}

		return fmt.Errorf("product: reserve fail %w", err)
	}

	if err = tx.Commit(ctx); err != nil {
		return fmt.Errorf("product: reserve fail commit %w", err)
	}

	return nil
}

func (r *ProductRepository) Release(ctx context.Context, ids []uuid.UUID) error {
	_, err := r.conn.Exec(
		ctx,
		"update product.product set is_reserved=false "+
			"where id=any($1) and warehouse_id=any(select id from product.warehouse where availability=true)",
		ids,
	)

	return err
}
