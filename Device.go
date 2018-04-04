package goonvif

import (
	"net"
	"encoding/xml"
	"log"
	"fmt"
	"github.com/beevik/etree"
	"github.com/yakovlevdmv/goonvif/Networking"
	"github.com/yakovlevdmv/gosoap"
	"reflect"
	"strings"
	"github.com/yakovlevdmv/goonvif/Device"
	"github.com/yakovlevdmv/goonvif/Imaging"
	"github.com/yakovlevdmv/goonvif/Media"
	"github.com/yakovlevdmv/goonvif/PTZ"
	"errors"
	"strconv"
)

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

func buildMethodSOAP(msg, space string) (gosoap.SoapMessage, error) {
	doc := etree.NewDocument()
	if err := doc.ReadFromString(msg); err != nil {
		log.Println("Got error")
		return "", err
	}
	element := doc.Root()


	soap := gosoap.NewEmptySOAP()
	soap.AddBodyContent(element)
	soap.AddRootNamespace("onvif", "http://www.onvif.org/ver10/device/wsdl")
	soap.AddRootNamespace("wsdl", space)

	return soap, nil
}

func setXMLNamespaces(strct interface{}) (string, error) {
	pkgName := reflect.TypeOf( strct ).PkgPath()
	pkgSplit := strings.Split(pkgName, "/")
	pkgName = pkgSplit[len(pkgSplit)-1]
	switch pkgName {
	case "Device":
		return Device.WSDL, nil
	case "Imaging":
		return Imaging.WSDL, nil
	case "Media":
		return Media.WSDL, nil
	case "PTZ":
		return PTZ.WSDL, nil
	default:
		return "", errors.New("Can't find Service")
	}
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

	wsdlSpaces, err := setXMLNamespaces(method)
	if err != nil {
		fmt.Println(err)
	}

	/*
	Build an SOAP request with <method>
	 */
	soap, err := buildMethodSOAP(string(output), wsdlSpaces)
	if err != nil {
		log.Printf("error: %v\n", err)
		return "", err
	}

	/*
	Sending request and returns the response
	 */
	return Networking.SendSoap(endpoint, soap.String()), nil
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

	wsdlSpaces, err := setXMLNamespaces(method)
	if err != nil {
		fmt.Println(err)
	}

	/*
	Build an SOAP request with <method>
	 */
	soap, err := buildMethodSOAP(string(output), wsdlSpaces)
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

	soapReq, err := xml.MarshalIndent(auth, "", "  ")
	if err != nil {
		log.Printf("error: %v\n", err.Error())
		return "", err
	}
	/*
	Adding WS-Security struct to SOAP header
	 */
	soap.AddStringHeaderContent(string(soapReq))

	/*
	Sending request and returns the response
	 */
	return Networking.SendSoap(endpoint, soap.String()), nil
}
