package domain

import (
	"encoding/json"

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

func (p *Product) MarshalJSON() ([]byte, error) {
	dto := struct {
		Type       string    `json:"type"`
		ID         uuid.UUID `json:"id"`
		Attributes struct {
			Name       string   `json:"name"`
			Count      int      `json:"count"`
			Size       pkg.Size `json:"size"`
			IsReserved bool     `json:"is_reserved"`
		} `json:"attributes"`
		Relationships struct {
			Warehouse struct {
				Data struct {
					Type string    `json:"type"`
					ID   uuid.UUID `json:"id"`
				} `json:"data"`
			} `json:"warehouse"`
		} `json:"relationships"`
	}{
		Type: "product",
		ID:   p.ID,
		Attributes: struct {
			Name       string   `json:"name"`
			Count      int      `json:"count"`
			Size       pkg.Size `json:"size"`
			IsReserved bool     `json:"is_reserved"`
		}{
			Name:       p.Name,
			Count:      p.Count,
			Size:       p.Size,
			IsReserved: p.IsReserved,
		},
		Relationships: struct {
			Warehouse struct {
				Data struct {
					Type string    `json:"type"`
					ID   uuid.UUID `json:"id"`
				} `json:"data"`
			} `json:"warehouse"`
		}{
			Warehouse: struct {
				Data struct {
					Type string    `json:"type"`
					ID   uuid.UUID `json:"id"`
				} `json:"data"`
			}{
				struct {
					Type string    `json:"type"`
					ID   uuid.UUID `json:"id"`
				}{
					Type: "warehouse",
					ID:   p.WarehouseID,
				},
			},
		},
	}

	return json.Marshal(dto)
}
