package scheme

//Types

type TypeWorker interface {
	GetProperties() *[]interface{}
}

type ContinuousMove struct {
	ReferenceToken 	string
	Pan, Tilt 		float64
	Zoom 			float64
	Timeout 		int //In seconds. Maybe need set more possibilities
}


func (t *ContinuousMove) GetProperties() *[]interface{} {
	return &[]interface{}{t.ReferenceToken, t.Pan, t.Tilt, t.Zoom, t.Timeout}
}