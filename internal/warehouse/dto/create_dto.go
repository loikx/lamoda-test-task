package dto

type CreateCommandDto struct {
	Name         string `json:"name"`
	Availability bool   `json:"availability"`
}
