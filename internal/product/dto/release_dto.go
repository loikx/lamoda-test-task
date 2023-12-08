package dto

import "github.com/gofrs/uuid"

type ReleaseDto struct {
	IDs []uuid.UUID `json:"ids"`
}
