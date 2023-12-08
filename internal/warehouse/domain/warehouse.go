package domain

import (
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
