package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/lamoda-tech/loikx/internal/product/domain"
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
	tx, err := r.conn.Begin(ctx)
	if err != nil {
		return fmt.Errorf("product: release fail start transaction %w", err)
	}

	_, err = tx.Exec(
		ctx,
		"update product.product set is_reserved=false "+
			"where id=any($1) and warehouse_id=any(select id from product.warehouse where availability=true)",
		ids,
	)
	if err != nil {
		if err = tx.Rollback(ctx); err != nil {
			return fmt.Errorf("product: release rollback fail %w", err)
		}

		return fmt.Errorf("product: release fail %w", err)
	}

	if err = tx.Commit(ctx); err != nil {
		return fmt.Errorf("product: release fail commit %w", err)
	}

	return nil
}

func (r *ProductRepository) FindByWarehouse(ctx context.Context, id uuid.UUID) ([]*domain.Product, error) {
	var items []*domain.Product

	rows, err := r.conn.Query(
		ctx,
		"select * from product.product p "+
			"where ($1)=p.warehouse_id and not p.is_reserved",
		id,
	)
	if err != nil {
		return nil, fmt.Errorf("product: find by warehouse %w", err)
	}

	if !rows.Next() {
		if rows.Err() == nil {
			return nil, errors.New("product: find by warehouse empty")
		}

		return nil, rows.Err()
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
			return nil, fmt.Errorf("product: find by warehouse failed to scan product %w", err)
		}

		items = append(items, &product)
	}

	return items, nil
}

func (r *ProductRepository) Save(ctx context.Context, product *domain.Product) error {
	tx, err := r.conn.Begin(ctx)
	if err != nil {
		return fmt.Errorf("product: save start transaction fail %w", err)
	}

	_, err = tx.Exec(
		ctx,
		"insert into product.product(id, name, count, length, width, height, unit, warehouse_id, is_reserved) "+
			"values ($1, $2, $3, $4, $5, $6, $7, $8, $9)",
		product.ID, product.Name, product.Count,
		product.Size.Length, product.Size.Width, product.Size.Height, product.Size.Unit,
		product.WarehouseID, product.IsReserved,
	)
	if err != nil {
		if err = tx.Rollback(ctx); err != nil {
			return fmt.Errorf("product: save rollback fail %w", err)
		}

		return fmt.Errorf("product: save fail %w", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("product: save commit fail %w", err)
	}

	return nil
}

func (r *ProductRepository) Delete(ctx context.Context, id uuid.UUID) error {
	tx, err := r.conn.Begin(ctx)
	if err != nil {
		return fmt.Errorf("product: delete start transaction fail %w", err)
	}

	_, err = tx.Exec(
		ctx,
		"delete from product.product where id=($1)",
		id,
	)
	if err != nil {
		if err = tx.Rollback(ctx); err != nil {
			return fmt.Errorf("product: delete rollback fail %w", err)
		}

		return fmt.Errorf("product: delete fail %w", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("product: delete commit fail %w", err)
	}

	return nil
}
