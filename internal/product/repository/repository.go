package repository

import (
	"context"
	"sync"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5"
)

type ProductRepository struct {
	conn *pgx.Conn

	mu sync.Mutex
}

func NewProductRepository(conn *pgx.Conn) *ProductRepository {
	return &ProductRepository{
		conn: conn,
		mu:   sync.Mutex{},
	}
}

func (r *ProductRepository) Reserve(ctx context.Context, ids []uuid.UUID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, err := r.conn.Exec(
		ctx,
		"update product.product set is_reserved=true where id=any($1)",
		ids,
	)

	return err
}

func (r *ProductRepository) Release(ctx context.Context, ids []uuid.UUID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, err := r.conn.Exec(
		ctx,
		"update product.product set is_reserved=false where id=any($1)",
		ids,
	)

	return err
}
