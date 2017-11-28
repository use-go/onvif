package goonvif

//import "encoding/xml"

//type Imd struct {
//	Anim string
//}

type Envelope struct {
	//Name 		xml.Name 	`xml:"notice"`
	Id 			int			`xml:"id,attr"`
	FirstName 	string   	`xml:"name>first"`
	LastName  	string   	`xml:"name>last"`
	NameId		string		`xml:"jorj, mas"`
	Age       	int      	`xml:"age"`
	Comment 	string 		`xml:",comment"`
}