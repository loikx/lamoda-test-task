package pkg

import (
	"context"

	"github.com/muonsoft/validation"
)

type Unit string

type Size struct {
	Length float32 `json:"length"`
	Width  float32 `json:"width"`
	Height float32 `json:"height"`
	Unit   Unit    `json:"unit"`
}

func (s *Size) Validate(ctx context.Context, validator *validation.Validator) error {
	return validator.Validate(
		ctx,
		validation.AtProperty(
			"length",
			validation.Check(s.Length > 0),
		),
		validation.AtProperty(
			"width",
			validation.Check(s.Width > 0),
		),
		validation.AtProperty(
			"height",
			validation.Check(s.Height > 0),
		),
	)
}
