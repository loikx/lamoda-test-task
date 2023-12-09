package domain

import (
	"github.com/gofrs/uuid"
	"github.com/lamoda-tech/loikx/internal/product/pkg"
	"github.com/lamoda-tech/loikx/pkg/text"
)

type Product struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Count       int       `json:"count"`
	Size        pkg.Size  `json:"size"`
	WarehouseID uuid.UUID `json:"warehouse"`
	IsReserved  bool      `json:"is_reserved"`
}

func NewProduct(name string, count int, size pkg.Size, warehouseID uuid.UUID) *Product {
	return &Product{
		ID:          uuid.Must(uuid.NewV7()),
		Name:        name,
		Count:       count,
		Size:        size,
		WarehouseID: warehouseID,
		IsReserved:  false,
	}
}

func (p *Product) Purify() {
	p.Name = text.Purify(p.Name)
}
