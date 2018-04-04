package goonvif

import (
	"net"
	"encoding/xml"
	"log"
	"fmt"
	"github.com/beevik/etree"
	"github.com/yakovlevdmv/goonvif/networking"
	"github.com/yakovlevdmv/gosoap"
	"strconv"
)

var xlmns = map[string]string {
	"onvif":"http://www.onvif.org/ver10/schema",
	"tds":"http://www.onvif.org/ver10/device/wsdl",
	"trt":"http://www.onvif.org/ver10/media/wsdl",
	"tev":"http://www.onvif.org/ver10/events/wsdl",
	"tpz":"http://www.onvif.org/ver20/ptz/wsdl",
	"timg":"http://www.onvif.org/ver20/imaging/wsdl",
	"xmime":"http://www.w3.org/2005/05/xmlmime",
	"wsnt":"http://docs.oasis-open.org/wsn/b-2",
	"xop":"http://www.w3.org/2004/08/xop/include",
	"wsa":"http://www.w3.org/2005/08/addressing",
	"wstop":"http://docs.oasis-open.org/wsn/t-1",
	"wsntw":"http://docs.oasis-open.org/wsn/bw-2",
	"wsrf-rw":"http://docs.oasis-open.org/wsrf/rw-2",
	"wsaw":"http://www.w3.org/2006/05/addressing/wsdl",
}

type DeviceType int

const (
	NVD DeviceType = iota
	NVS
	NVA
	NVT
)

func (devType DeviceType) String() string {
	stringRepresentation := []string {
		"NetworkVideoDisplay",
		"NetworkVideoStorage",
		"NetworkVideoAnalytics",
		"NetworkVideoTransmitter",
	}
	i := uint8(devType)
	switch {
	case i <= uint8(NVT):
		return stringRepresentation[i]
	default:
		return strconv.Itoa(int(i))
	}
}

//deviceInfo struct contains general information about ONVIF device
type deviceInfo struct {
	Manufacturer string
	Model string
	FirmwareVersion string
	SerialNumber string
	HardwareId string

}

//deviceInfo struct represents an abstract ONVIF device.
//It contains methods, which helps to communicate with ONVIF device
type device struct {

	xaddr net.IP
	login string
	password string

	token [64]uint8

	endpoints map[string]string
	info deviceInfo

}

func getAvailableDevicesAtEthernet(interfaceName string) {

}

//NewDevice function construct a ONVIF Device entity
func NewDevice() *device {
	return &device{}
}

//Authenticate function authenticate client in the ONVIF Device.
//Function takes <username> and <password> params.
//You should use this function to allow authorized requests to the ONVIF Device
//To change auth data call this function again.
func (dev *device) Authenticate(username, password string) {
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
	soap.AddRootNamespace("onvif", "http://www.onvif.org/ver10/device/wsdl")

	return soap, nil
}



//CallMethod functions call an method, defined <method> struct.
//You should use Authenticate method to call authorized requests.
func (dev device) CallMethod(endpoint string, method interface{}) (string, error) {
	//TODO: Get endpoint automatically
	if dev.login != "" && dev.password != "" {
		return dev.CallAuthorizedMethod(endpoint, method)
	} else {
		return dev.CallNonAuthorizedMethod(endpoint, method)
	}
}

//CallNonAuthorizedMethod functions call an method, defined <method> struct without authentication data
func (dev device) CallNonAuthorizedMethod(endpoint string, method interface{}) (string, error) {
	//TODO: Get endpoint automatically
	/*
	Converting <method> struct to xml string representation
	 */
	output, err := xml.MarshalIndent(method, "  ", "    ")
	if err != nil {
		log.Printf("error: %v\n", err.Error())
		return "", err
	}

	if err != nil {
		fmt.Println(err)
	}

	/*
	Build an SOAP request with <method>
	 */
	soap, err := buildMethodSOAP(string(output))
	if err != nil {
		log.Printf("error: %v\n", err)
		return "", err
	}

	soap.AddRootNamespaces(xlmns)

	fmt.Println(soap.String())

	/*
	Sending request and returns the response
	 */
	return networking.SendSoap(endpoint, soap.String()), nil
}

//CallMethod functions call an method, defined <method> struct with authentication data
func (dev device) CallAuthorizedMethod(endpoint string, method interface{}) (string, error) {
	/*
	Converting <method> struct to xml string representation
	 */
	output, err := xml.MarshalIndent(method, "  ", "    ")
	if err != nil {
		log.Printf("error: %v\n", err.Error())
		return "", err
	}

	if err != nil {
		fmt.Println(err)
	}


	/*
	Build an SOAP request with <method>
	 */
	soap, err := buildMethodSOAP(string(output))
	if err != nil {
		log.Fatal(err)
	}

	/*
	Getting an WS-Security struct representation
	 */
	auth := newSecurity(dev.login, dev.password)

	/*
	Adding WS-Security namespaces to root element of SOAP message
	 */
	soap.AddRootNamespace("wsse", "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext1.0.xsd")
	soap.AddRootNamespace("wsu", "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility1.0.xsd")

	soap.AddRootNamespaces(xlmns)

	soapReq, err := xml.MarshalIndent(auth, "", "  ")
	if err != nil {
		log.Printf("error: %v\n", err.Error())
		return "", err
	}

	/*
	Adding WS-Security struct to SOAP header
	 */
	soap.AddStringHeaderContent(string(soapReq))

	fmt.Println(soap.String())
	/*
	Sending request and returns the response
	 */
	return networking.SendSoap(endpoint, soap.String()), nil
}
