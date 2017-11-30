package goonvif

import (
	"encoding/xml"
	"time"
	"github.com/yakovlevdmv/goonvif/scheme"
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

type Body struct {
	XMLName xml.Name
	Content scheme.TypeWorker
}

type Security struct {
	XMLName xml.Name  `xml:"wsse:Security"`
	Auth WSAuth
}

type Password struct {
	XMLName xml.Name `xml:"wsse:Password"`
	Type string `xml:"Type,attr"`
}

type WSAuth struct {
	XMLName xml.Name  `xml:"wsse:UsernameToken"`
	Username string   `xml:"wsse:Username"`
	Password Password `xml:"wsse:Password"`
	Nonce string      `xml:"wsse:Nonce"`
	Created time.Time `xml:"wsse:Created"`
}