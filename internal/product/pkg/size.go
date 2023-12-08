package pkg

type Unit string

type Size struct {
	Length float32 `json:"length"`
	Width  float32 `json:"width"`
	Height float32 `json:"height"`
	Unit   Unit    `json:"unit"`
}
