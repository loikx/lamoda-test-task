package domain

import (
	"fmt"
)

func (w *Warehouse) Validation() error {
	if len(w.Name) < 2 || len(w.Name) > 100 {
		return fmt.Errorf("warehouse: name length must be between 2-100")
	}

	return nil
}
