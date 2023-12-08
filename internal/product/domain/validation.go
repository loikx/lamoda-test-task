package domain

import (
	"context"

	"github.com/muonsoft/validation"
	"github.com/muonsoft/validation/it"
)

func (p *Product) Validate(ctx context.Context, validator *validation.Validator) error {
	return validator.Validate(
		ctx,
		validation.StringProperty(
			"name",
			p.Name,
			it.HasLengthBetween(2, 100),
		),
		validation.CountableProperty(
			"count",
			p.Count,
			it.HasMaxCount(1000),
		),
	)
}
