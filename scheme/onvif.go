package scheme

//Types

type TypeWorker interface {
	GetProperties() *map[string] interface{}
}

type ContinuousMove struct {
	ReferenceToken 	string
	Pan, Tilt 		float64
	Zoom 			float64
	Timeout 		int                 /*****In seconds. Maybe need set more possibilities*****/
}


func (t *ContinuousMove) GetProperties() *map[string] interface{} {
	return &map[string]interface{}{"ReferenceToken": t.ReferenceToken, "Pan": t.Pan, "Tilt": t.Tilt,"Zoom": t.Zoom, "Timeout": t.Timeout}

}