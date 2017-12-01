package scheme

type Vector2D struct {
	X float64 `xml:"x,attr"`
	Y float64 `xml:"y,attr"`
	Space string `xml:"space,attr,omitempty"`
}

func (v *Vector2D) GetProperties() *map[string]interface{} {
	return &map[string]interface{} {
		"x":v.X,
		"y":v.Y,
		"space":v.Space,
	}
}

type Vector1D struct {
	X float64 `xml:"x,attr"`
	Space string `xml:"space,attr,omitempty"`
}

func (v *Vector1D) GetProperties() *map[string]interface{} {
	return &map[string]interface{} {
		"x":v.X,
		"space":v.Space,
	}
}