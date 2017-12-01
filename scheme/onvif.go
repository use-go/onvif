package scheme

import "encoding/xml"

//Types

type TypeWorker interface {
	GetProperties() *map[string] interface{}
}


//Start of ContinuousMove
type Velocity struct {
	xmlName xml.Name `xml:"wsdl:velocity"`
	PanTilt  Vector2D `xml:"tt:PanTilt,omitempty"`
	Zoom Vector1D `xml:"tt:Zoom,omitempty"`
}

type Zoom struct {
	XMLName xml.Name
	Zoom int `xml:"x,attr"`
}

type ContinuousMove struct {
	XMLName xml.Name `xml:"wsdl:ContinuousMove"`
	ReferenceToken string 	`xml:"wsdl:ProfileToken"`
	Velocity 		Velocity 	`xml:"wsdl:Velocity"`
	Timeout 		int 		`xml:"wsdl:Timeout"`   //In seconds. Maybe need set more possibilities
}

func (t *Velocity)GetProperties() *map[string] interface{} {
	return &map[string]interface{}{
		"PanTilt":t.PanTilt,
		"Zoom":t.Zoom,
	}
}

func (t *Zoom)GetProperties() *map[string] interface{} {
	return &map[string]interface{} {
		"Zoom": t.Zoom,
	}
}


func (t *ContinuousMove) GetProperties() *map[string] interface{} {
	return &map[string]interface{}{
		"ProfileToken": t.ReferenceToken,
		"Velocity": t.Velocity,
		"Timeout": t.Timeout,
		}
}
//End of ContinuousMove