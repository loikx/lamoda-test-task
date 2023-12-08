package repository

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5"
)

type ProductRepository struct {
	conn *pgx.Conn
}

func (r *ProductRepository) Reserve(ctx context.Context, ids []uuid.UUID) error {
	_, err := r.conn.Exec(
		ctx,
		"update product set is_reserved=true where id=any($1)",
		ids,
	)

	return err
}

func (r *ProductRepository) Release(ctx context.Context, ids []uuid.UUID) error {
	_, err := r.conn.Exec(
		ctx,
		"update product set is_reserved=false where id=any($1)",
		ids,
	)

	return err
}
