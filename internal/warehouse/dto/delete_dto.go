package dto

import "github.com/gofrs/uuid"

type DeleteDto struct {
	ID uuid.UUID `json:"warehouse"`
}
