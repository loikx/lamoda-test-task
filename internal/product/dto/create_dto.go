package dto

import (
	"github.com/gofrs/uuid"
	"github.com/lamoda-tech/loikx/internal/product/pkg"
)

type CreateCommandDto struct {
	Name        string
	Count       int
	Size        pkg.Size
	WarehouseID uuid.UUID
}
