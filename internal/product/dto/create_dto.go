package dto

import (
	"github.com/gofrs/uuid"
	"github.com/lamoda-tech/loikx/internal/product/pkg"
)

type CreateCommandDto struct {
	Name        string    `json:"name"`
	Count       int       `json:"count"`
	Size        pkg.Size  `json:"size"`
	WarehouseID uuid.UUID `json:"warehouse"`
}
