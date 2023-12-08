package domain

import (
	"context"

	"github.com/muonsoft/validation"
	"github.com/muonsoft/validation/it"
)

func (w *Warehouse) Validation(ctx context.Context, validator *validation.Validator) error {
	return validator.Validate(
		ctx,
		validation.StringProperty(
			"name",
			w.Name,
			it.HasLengthBetween(2, 100),
		),
	)
}
