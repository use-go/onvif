package goonvif

import (
	"net"
	"encoding/xml"
	"log"
	"fmt"
	"github.com/beevik/etree"
	"github.com/yakovlevdmv/gosoap"
	"github.com/yakovlevdmv/goonvif/Networking"
	"time"
	"encoding/base64"
	"crypto/sha1"
	"github.com/elgs/gostrgen"
)

type DeviceInfo struct {
	Manufacturer string
	Model string
	FirmwareVersion string
	SerialNumber string
	HardwareId string

}

type device struct {

	xaddr net.IP
	login string
	password string

	token [64]uint8

	endpoints map[string]string
	info DeviceInfo

}

func NewDevice() *device {
	return &device{}
}

func (dev *device) AddAuthentification(username, password string) {
	dev.login = username
	dev.password = password
}

func buildMethodSOAP(msg string) (gosoap.SoapMessage, error) {
	doc := etree.NewDocument()
	if err := doc.ReadFromString(msg); err != nil {
		log.Println("Got error")
		return "", err
	}
	element := doc.Root()

	soap := gosoap.NewEmptySOAP()
	soap.AddBodyContent(element)
	soap.AddRootNamespace("wsdl", "http://www.onvif.org/ver10/device/wsdl")

	return soap, nil
}

//TODO: Get endpoint automatically
func (dev device) CallMethod(endpoint string, method interface{}) {

	output, err := xml.MarshalIndent(method, "  ", "    ")
	if err != nil {
		log.Printf("error: %v\n", err)
	} else {
		log.Println("Marshalled struct: ", string(output))
	}

	fmt.Println(string(output))
	soap, err := buildMethodSOAP(string(output))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Send soap\n")
	fmt.Println(soap.String())

	Networking.SendSoap(endpoint, soap.String())
}

/*************************
	WS-Security types
*************************/
const (passwordType = "https://www.oasis-open.org/committees/download.php/13392/wss-v1.1-spec-pr-UsernameTokenProfile-01.htm#PasswordDigest")

/*
xmlns:wsse="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext1.0.xsd"

xmlns:wsu="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility1.0.xsd"

 */
type security struct {
	XMLName xml.Name  `xml:"wsse:Security"`
	Auth wsAuth
}

type password struct {
	XMLName xml.Name `xml:"wsse:Password"`
	Type string `xml:"Type,attr"`
	Password string `xml:",chardata"`
}

type wsAuth struct {
	XMLName xml.Name  `xml:"wsse:UsernameToken"`
	Username string   `xml:"wsse:Username"`
	Password password `xml:"wsse:Password"`
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


func (dev device) CallAuthorizedMethod(endpoint string, method interface{}) {
	output, err := xml.MarshalIndent(method, "  ", "    ")
	if err != nil {
		log.Printf("error: %v\n", err)
	} else {
		log.Println("Marshalled struct: ", string(output))
	}

	soap, err := buildMethodSOAP(string(output))
	if err != nil {
		log.Fatal(err)
	}

	/**soap := gosoap.NewEmptySOAP()

	soap.AddStringBodyContent(
		//`<tds:SetHostname xmlns:tds="http://www.onvif.org/ver10/device/wsdl"><tds:Name>Test</tds:Name></tds:SetHostname>`,
		`<tds:GetHostname xmlns:tds="http://www.onvif.org/ver10/device/wsdl" />`,
	)*/

	/** Generating Nonce sequence **/
	charsToGenerate := 16
	charSet := gostrgen.Lower | gostrgen.Digit

	nonce, _ := gostrgen.RandGen(charsToGenerate, charSet, "", "")

	auth := security{
		Auth:wsAuth{
			Username:dev.login,
			Password:password {
				Type:passwordType,
				Password:generateToken(dev.login, nonce, time.Now(), dev.password),
			},
			Nonce: nonce,
			Created: time.Now().Format(time.RFC3339),
		},
	}

	soap.AddRootNamespace("wsse", "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext1.0.xsd")
	soap.AddRootNamespace("wsu", "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility1.0.xsd")

	soapReq, err := xml.MarshalIndent(auth, "", "  ")
	if err != nil {
		log.Panic(err.Error())
		return
	} else {
		soap.AddStringHeaderContent(string(soapReq))

		log.Println("Soap to send")
		log.Println(soap.String())

	}

	Networking.SendSoap(endpoint, soap.String())
	//fmt.Println(string(output))
}
