package scheme

import (
	"net/url"
)

type Vector2D struct {
	x float64
	y float64
	space url.URL `xml:"omitempty"`
}

func (v *Vector2D) GetProperties() *map[string]interface{} {
	return &map[string]interface{} {
		"x":v.x,
		"y":v.y,
		"space":v.space,
	}
}
