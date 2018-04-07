package goonvif

import (
	"encoding/xml"
	"fmt"
	"github.com/beevik/etree"
	"github.com/yakovlevdmv/gosoap"
	"strconv"
	"github.com/yakovlevdmv/WS-Discovery"
	"github.com/yakovlevdmv/goonvif/networking"
	"reflect"
	"strings"
	"github.com/yakovlevdmv/goonvif/Device"
	"errors"
)

var xlmns = map[string]string {
	"onvif":"http://www.onvif.org/ver10/schema",
	"tds":"http://www.onvif.org/ver10/device/wsdl",
	"trt":"http://www.onvif.org/ver10/media/wsdl",
	"tev":"http://www.onvif.org/ver10/events/wsdl",
	"tptz":"http://www.onvif.org/ver20/ptz/wsdl",
	"timg":"http://www.onvif.org/ver20/imaging/wsdl",
	"tan":"http://www.onvif.org/ver20/analytics/wsdl",
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

	xaddr string
	login string
	password string

	endpoints map[string]string
	info deviceInfo

}

func GetAvailableDevicesAtSpecificEthernetInterface(interfaceName string) []device {
	/*
	Call an WS-Discovery Probe Message to Discover NVT type Devices
	 */
	devices := WS_Discovery.SendProbe(interfaceName, nil, []string{"dn:"+NVT.String()}, map[string]string{"dn":"http://www.onvif.org/ver10/network/wsdl"})
	//nvtDevices := make([]device, len(devices))
	for _, j := range devices {
		fmt.Println(j)
		//nvtDevices[i] = NewDevice()
	}
	return nil
}

func (dev *device) getSupportedServices() {
	resp, err := dev.CallMethod(Device.GetCapabilities{})
	if err != nil {
		//log.Println(err.Error())
		return
	} else {
		doc := etree.NewDocument()
		if err := doc.ReadFromString(resp); err != nil {
			//log.Println(err.Error())
			return
		}
		services := doc.FindElements("./Envelope/Body/GetCapabilitiesResponse/Capabilities/*/XAddr")
		for _, j := range services{
			//fmt.Println(j.Text())
			//fmt.Println(j.Parent().Tag)
			dev.addEndpoint(j.Parent().Tag, j.Text())
		}
	}
}

//NewDevice function construct a ONVIF Device entity
func NewDevice(xaddr string) (*device, error) {
	dev := new(device)
	dev.xaddr = xaddr
	dev.endpoints = make(map[string]string)
	dev.addEndpoint("Device", "http://"+xaddr+"/onvif/device_service")

	systemDateTime := Device.GetDeviceInformation{}
	_, err := dev.CallMethod(systemDateTime)
	if err != nil {
		panic(errors.New("camera is not available at " + xaddr + " or it does not support ONVIF services"))
		return nil, err
	}

	dev.getSupportedServices()
	return dev, nil
}

func (dev *device)addEndpoint(Key, Value string) {
	dev.endpoints[Key]=Value
}

func newDeviceEntity() *device {
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
		//log.Println("Got error")

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
func (dev device) CallMethod(method interface{}) (string, error) {
	pkgPath := strings.Split(reflect.TypeOf(method).PkgPath(),"/")
	pkg := pkgPath[len(pkgPath)-1]

	var endpoint string
	switch pkg {
		case "Device": endpoint = dev.endpoints["Device"]
		case "Event": endpoint = dev.endpoints["Event"]
		case "Imaging": endpoint = dev.endpoints["Imaging"]
		case "Media": endpoint = dev.endpoints["Media"]
		case "PTZ": endpoint = dev.endpoints["PTZ"]
	}

	//TODO: Get endpoint automatically
	if dev.login != "" && dev.password != "" {
		/*resp, err := dev.сallAuthorizedMethod(endpoint, method)
		if err != nil {
			panic(err)
			return resp, err
		}

		return resp, err*/
		return dev.сallAuthorizedMethod(endpoint, method)
	} else {
		/*resp, err := dev.сallAuthorizedMethod(endpoint, method)
		if err != nil {
			panic(err)
			return resp, err
		}
		return resp, err*/
		return dev.сallNonAuthorizedMethod(endpoint, method)

	}
}

//CallNonAuthorizedMethod functions call an method, defined <method> struct without authentication data
func (dev device) сallNonAuthorizedMethod(endpoint string, method interface{}) (string, error) {
	//TODO: Get endpoint automatically
	/*
	Converting <method> struct to xml string representation
	 */
	output, err := xml.MarshalIndent(method, "  ", "    ")
	if err != nil {
		//log.Printf("error: %v\n", err.Error())
		return "", err
	}

	/*
	Build an SOAP request with <method>
	 */
	soap, err := buildMethodSOAP(string(output))
	if err != nil {
		//log.Printf("error: %v\n", err)
		return "", err
	}

	soap.AddRootNamespaces(xlmns)

	/*
	Sending request and returns the response
	 */
	return networking.SendSoap(endpoint, soap.String())
}

//CallMethod functions call an method, defined <method> struct with authentication data
func (dev device) сallAuthorizedMethod(endpoint string, method interface{}) (string, error) {
	/*
	Converting <method> struct to xml string representation
	 */
	output, err := xml.MarshalIndent(method, "  ", "    ")
	if err != nil {
		//log.Printf("error: %v\n", err.Error())
		return "", err
	}

	/*
	Build an SOAP request with <method>
	 */
	soap, err := buildMethodSOAP(string(output))
	if err != nil {
		//log.Printf("error: %v\n", err.Error())
		return "", err
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
		//log.Printf("error: %v\n", err.Error())
		return "", err
	}

	/*
	Adding WS-Security struct to SOAP header
	 */
	soap.AddStringHeaderContent(string(soapReq))

	/*
	Sending request and returns the response
	 */
	return networking.SendSoap(endpoint, soap.String())
}
