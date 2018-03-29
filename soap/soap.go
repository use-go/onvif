package goonvif

import (
	"encoding/xml"
	"time"
	"github.com/elgs/gostrgen"
	"github.com/yakovlevdmv/goonvif/scheme"
	"encoding/base64"
	"crypto/sha1"
	"fmt"
	"log"
	"net/http"
	"bytes"
	"io/ioutil"
)

var xlmns = map[string]string {
//	SOAP namespase
	"soap":"http://www.w3.org/2003/05/soap-envelope",
	"xs":"http://www.w3.org/2001/XMLSchema",
//	ONVIF namespace
	"tt":"http://www.onvif.org/ver10/schema",
	"tds":"http://www.onvif.org/ver10/device/wsdl",
	"trt":"http://www.onvif.org/ver10/media/wsdl",
	"tev":"http://www.onvif.org/ver10/events/wsdl",
	"ter":"http://www.onvif.org/ver10/error",
	"dn":"http://www.onvif.org/ver10/network/wsdl",
	"tns1":"http://www.onvif.org/ver10/topics",
	"tpz":"http://www.onvif.org/ver20/ptz/wsdl",
	"dev":"https://www.onvif.org/ver10/device/wsdl/devicemgmt.wsdl",

	"xsi":"http://www.w3.org/2001/XMLSchema-instance",
	"xsd":"http://www.w3.org/2001/XMLSchema",
	"wsse":"http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd",
	"wsu":"http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd",

}

var passwordType = "https://www.oasis-open.org/committees/download.php/13392/wss-v1.1-spec-pr-UsernameTokenProfile-01.htm#PasswordDigest"

type Envelope struct {
	XMLName xml.Name `xml:"soap:Envelope"`
	XmlnsSoap string `xml:"xmlns:soap,attr,omitempty"`
	XmlnsXsi string `xml:"xmlns:xsi,attr,omitempty"`
	XmlnsXsd string `xml:"xmlns:xsd,attr,omitempty"`
	XmlnsTt string `xml:"xmlns:tt,attr,omitempty"`
	XmlnsWsse string `xml:"xmlns:wsse,attr,omitempty"`
	XmlnsWsu string `xml:"xmlns:wsu,attr,omitempty"`
	XmlnsWsdl string `xml:"xmlns:wsdl,attr,omitempty"`
	Header Header
	Body Body
}

type Header struct {
	XMLName xml.Name `xml:"soap:Header"`
	WSAuth  Security
}

type BodyContent struct {
	XMLName xml.Name
	Content string `xml:",innerxml"`
}

type Body struct {
	XMLName xml.Name `xml:"soap:Body"`
	Content BodyContent
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

//func (t *emptyInput)createSelfCloseTag() string {
//	t.EmptyInput
//}

func NewEnvelope( funcName string, data scheme.TypeWorker ) Envelope {

	/** Generating Nonce sequence **/
	charsToGenerate := 16
	charSet := gostrgen.Lower | gostrgen.Digit

	nonce, _ := gostrgen.RandGen(charsToGenerate, charSet, "", "")

	output, err := xml.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	var soapData = Envelope{
		XmlnsSoap:xlmns["soap"],
		XmlnsXsd:xlmns["xsd"],
		XmlnsXsi:xlmns["xsi"],
		XmlnsWsdl:xlmns["dev"],
		XmlnsTt:xlmns["tt"],
		XmlnsWsse:xlmns["wsse"],
		XmlnsWsu:xlmns["wsu"],
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
			Content:BodyContent{
				XMLName:xml.Name{Local:"wsdl:" + funcName},
				Content:"\n\t"+string(output)+"\n",},
		},
	}


	return soapData

}

func SendSoap() {
	env := NewEnvelope("GetDNS", nil)
	soapReq, _ := xml.MarshalIndent(env, "", "  ")

	httpClient := new(http.Client)
	resp, _ := httpClient.Post("http://192.168.13.53", "application/soap+xml; charset=utf-8", bytes.NewBufferString(string(soapReq)))

	b, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(b))
}