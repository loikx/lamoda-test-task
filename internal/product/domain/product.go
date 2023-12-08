package domain

import (
	"github.com/gofrs/uuid"
	"github.com/lamoda-tech/loikx/internal/product/pkg"
	"github.com/lamoda-tech/loikx/pkg/text"
)

type Product struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Count int       `json:"count"`
	Size  pkg.Size  `json:"size"`
}

func NewProduct(ID uuid.UUID, name string, count int, size pkg.Size) *Product {
	return &Product{
		ID:    ID,
		Name:  name,
		Count: count,
		Size:  size,
	}
}

func (p *Product) Purify() {
	p.Name = text.Purify(p.Name)
}
