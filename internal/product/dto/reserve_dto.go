package dto

import "github.com/gofrs/uuid"

type ReserveDto struct {
	IDs []uuid.UUID `json:"ids"`
}
