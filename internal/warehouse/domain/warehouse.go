package domain

import (
	"encoding/json"

	"github.com/gofrs/uuid"
	"github.com/lamoda-tech/loikx/pkg/text"
)

type Warehouse struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Availability bool      `json:"availability"`
}

func NewWarehouse(name string, availability bool) *Warehouse {
	return &Warehouse{
		ID:           uuid.Must(uuid.NewV7()),
		Name:         name,
		Availability: availability,
	}
}

func (w *Warehouse) Purify() {
	w.Name = text.Purify(w.Name)
}

func (w *Warehouse) MarshalJSON() ([]byte, error) {
	dto := struct {
		Type       string    `json:"type"`
		ID         uuid.UUID `json:"id"`
		Attributes struct {
			Name         string `json:"name"`
			Availability bool   `json:"availability"`
		} `json:"attributes"`
	}{
		Type: "warehouse",
		ID:   w.ID,
		Attributes: struct {
			Name         string `json:"name"`
			Availability bool   `json:"availability"`
		}(struct {
			Name         string
			Availability bool
		}{
			Name:         w.Name,
			Availability: w.Availability,
		}),
	}

	return json.Marshal(dto)
}
