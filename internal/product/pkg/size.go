package pkg

import (
	"fmt"
)

type Unit string

type Size struct {
	Length float32 `json:"length"`
	Width  float32 `json:"width"`
	Height float32 `json:"height"`
	Unit   Unit    `json:"unit"`
}

func (s *Size) Validate() error {
	if s.Length <= 0 {
		return fmt.Errorf("size: length must be more than 0")
	}

	if s.Width <= 0 {
		return fmt.Errorf("size: width must be more than 0")
	}

	if s.Height <= 0 {
		return fmt.Errorf("size: height must be more than 0")
	}

	return nil
}
