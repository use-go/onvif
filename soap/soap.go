package goonvif

import (
	"encoding/xml"
	"time"
	"github.com/elgs/gostrgen"
	"github.com/yakovlevdmv/goonvif/scheme"
	"encoding/base64"
	"crypto/sha1"
	"fmt"
)

var xlmns = map[string]string {
//	SOAP namespase
	"env":"http://www.w3.org/2003/05/soap-envelope",
	"xs":"http://www.w3.org/2001/XMLSchema",
//	ONVIF namespace
	"tt":"http://www.onvif.org/ver10/schema",
	"tds":"http://www.onvif.org/ver10/device/wsdl",
	"trt":"http://www.onvif.org/ver10/media/wsdl",
	"tev":"http://www.onvif.org/ver10/events/wsdl",
	"ter":"http://www.onvif.org/ver10/error",
	"dn":"http://www.onvif.org/ver10/network/wsdl",
	"tns1":"http://www.onvif.org/ver10/topics",

}

var passwordType = "https://www.oasis-open.org/committees/download.php/13392/wss-v1.1-spec-pr-UsernameTokenProfile-01.htm#PasswordDigest"

type Envelope struct {
	XMLName xml.Name `xml:"http://www.w3.org/2003/05/soap-envelope Envelope"`
	Header Header
	Body Body
}

type Header struct {
	XMLName xml.Name
	WSAuth  Security
}

//type BodyContent struct {
//	XMLName xml.Name
//	Content string `xml:",innerxml"`
//}

type Body struct {
	XMLName xml.Name
	Content string `xml:",innerxml"`
}

type Security struct {
	XMLName xml.Name  `xml:"wsse:Security"`
	Auth WSAuth
}

type Password struct {
	XMLName xml.Name `xml:"wsse:Password"`
	Type string `xml:"Type,attr"`
	Password string `xml:",chardata"`
}

type WSAuth struct {
	XMLName xml.Name  `xml:"wsse:UsernameToken"`
	Username string   `xml:"wsse:Username"`
	Password Password `xml:"wsse:Password"`
	Nonce string      `xml:"wsse:Nonce"`
	Created string    `xml:"wsse:Created"`
}

//Digest = B64ENCODE( SHA1( B64DECODE( Nonce ) + Date + Password ) )
func generateToken(Username string, Nonce string, Created time.Time, Password string) string {

	sDec, _ := base64.StdEncoding.DecodeString(Nonce)


	hasher := sha1.New()
	//hasher.Write([]byte((base64.StdEncoding.EncodeToString([]byte(Nonce)) + Created.Format(time.RFC3339) + Password)))
	hasher.Write([]byte(string(sDec) + Created.Format(time.RFC3339) + Password))

	return base64.StdEncoding.EncodeToString(hasher.Sum(nil))
}

func NewEnvelope( funcName string, data scheme.TypeWorker ) Envelope {

	/** Generating Nonce sequence **/
	charsToGenerate := 16
	charSet := gostrgen.Lower | gostrgen.Digit

	nonce, _ := gostrgen.RandGen(charsToGenerate, charSet, "", "")

	output, _ := xml.MarshalIndent(data, "", "  ")
	fmt.Println(string(output))

	var soapData = Envelope{
		Header:Header {
			WSAuth:Security {
				Auth:WSAuth {
					Username:"admin",
					Password:Password {
						Type:passwordType,
						Password:generateToken("admin", nonce, time.Now(), "Supervisor"),
					},
					Nonce: nonce,
					Created: time.Now().Format(time.RFC3339),
				},
			},
		},
		Body:Body{
			//xml.Name{Local: funcName},
			Content:"\n\t"+string(output)+"\n",
			//Content:BodyContent{
			//	XMLName: xml.Name { Local: funcName },
			//	Content:string(output),
			//},
		},
	}

	//output, err := xml.MarshalIndent(soapData, "", "  ")
	//if err != nil {
	//	fmt.Printf("error: %v\n", err)
	//}
	//fmt.Println("Old XML")
	//fmt.Println(string(output))
	//
	//output = []byte(strings.Replace(strings.Replace(string(output), "<Content>\n", "", 1), "</Content>\n","",1))
	//fmt.Println("New XML")
	//fmt.Println(string(output))

	return soapData

}