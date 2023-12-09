package domain

import (
	"fmt"
)

func (p *Product) Validate() error {
	if len(p.Name) < 2 || len(p.Name) > 100 {
		return fmt.Errorf("product: validation length must be between 2-100")
	}

	if p.Count > 1000 {
		return fmt.Errorf("product: validation product count must be lower than 1000")
	}

	return p.Size.Validate()
}
