package domain

import (
	"github.com/gofrs/uuid"
	"github.com/lamoda-tech/loikx/internal/product/pkg"
	"github.com/lamoda-tech/loikx/pkg/text"
)

type Product struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	Count      int       `json:"count"`
	Size       pkg.Size  `json:"size"`
	IsReserved bool      `json:"is_reserved"`
}

func NewProduct(name string, count int, size pkg.Size, isReserved bool) *Product {
	return &Product{
		ID:         uuid.Must(uuid.NewV7()),
		Name:       name,
		Count:      count,
		Size:       size,
		IsReserved: isReserved,
	}
}

func (p *Product) Purify() {
	p.Name = text.Purify(p.Name)
}
