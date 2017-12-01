package scheme

import "encoding/xml"

//Types

type TypeWorker interface {
	GetProperties() *map[string] interface{}
}


//Start of ContinuousMove
type ProfileToken struct {
	ReferenceToken 	string
}

type Velocity struct {
	XMLName xml.Name
	Pan 	float64 `xml:"x,attr"`
	Tilt 	float64 `xml:"y,attr"`
	Zoom Zoom
}

type Zoom struct {
	XMLName xml.Name
	Zoom int `xml:"x,attr"`
}

type ContinuousMove struct {
	XMLName xml.Name `xml:"wsdl:ContinuousMove"`
	ProfileToken ProfileToken 	`xml:"wsdl:ProfileToken"`
	Velocity 		Velocity 	`xml:"wsdl:Velocity"`
	Timeout 		int 		`xml:"wsdl:Timeout"`   //In seconds. Maybe need set more possibilities
}

func (t *Velocity)GetProperties() *map[string] interface{} {
	return &map[string]interface{}{
		"Pan":t.Pan,
		"Tilt":t.Tilt,
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
		"ProfileToken": t.ProfileToken,
		"Velocity": t.Velocity,
		"Timeout": t.Timeout,
		}
}
//End of ContinuousMove