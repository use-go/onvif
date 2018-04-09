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
	"net/http"
	"io/ioutil"
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

func (dev *device)GetServices() map[string]string {
	return dev.endpoints
}

func readResponse(resp *http.Response) string {
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func GetAvailableDevicesAtSpecificEthernetInterface(interfaceName string) []device {
	/*
	Call an WS-Discovery Probe Message to Discover NVT type Devices
	 */
	devices := WS_Discovery.SendProbe(interfaceName, nil, []string{"dn:"+NVT.String()}, map[string]string{"dn":"http://www.onvif.org/ver10/network/wsdl"})
	nvtDevices := make([]device, 0)
	////fmt.Println(devices)
	for _, j := range devices {
		doc := etree.NewDocument()
		if err := doc.ReadFromString(j); err != nil {
			fmt.Errorf("%s", err.Error())
			return nil
		}
		////fmt.Println(j)
		endpoints := doc.Root().FindElements("./Body/ProbeMatches/ProbeMatch/XAddrs")
		for _, xaddr := range endpoints {
			//fmt.Println(xaddr.Tag,strings.Split(strings.Split(xaddr.Text(), " ")[0], "/")[2] )
			xaddr := strings.Split(strings.Split(xaddr.Text(), " ")[0], "/")[2]
			fmt.Println(xaddr)
			c := 0
			for c = 0; c < len(nvtDevices); c++ {
				if nvtDevices[c].xaddr == xaddr {
					fmt.Println(nvtDevices[c].xaddr, "==", xaddr)
					break
				}
			}
			if c < len(nvtDevices) {
				continue
			}
			dev, err := NewDevice(strings.Split(xaddr, " ")[0])
			//fmt.Println(dev)
			if err != nil {
				fmt.Println("Error", xaddr)
				fmt.Println(err)
				continue
			} else {
				////fmt.Println(dev)
				nvtDevices = append(nvtDevices, *dev)
			}
		}
		////fmt.Println(j)
		//nvtDevices[i] = NewDevice()
	}
	return nvtDevices
}

func (dev *device) getSupportedServices(resp *http.Response) {
	//resp, err := dev.CallMethod(Device.GetCapabilities{Category:"All"})
	//if err != nil {
	//	log.Println(err.Error())
		//return
	//} else {
		doc := etree.NewDocument()

		data, _ := ioutil.ReadAll(resp.Body)

		if err := doc.ReadFromBytes(data); err != nil {
			//log.Println(err.Error())
			return
		}
		services := doc.FindElements("./Envelope/Body/GetCapabilitiesResponse/Capabilities/*/XAddr")
		for _, j := range services{
			////fmt.Println(j.Text())
			////fmt.Println(j.Parent().Tag)
			dev.addEndpoint(j.Parent().Tag, j.Text())
		}
	//}
}

//NewDevice function construct a ONVIF Device entity
func NewDevice(xaddr string) (*device, error) {
	dev := new(device)
	dev.xaddr = xaddr
	dev.endpoints = make(map[string]string)
	dev.addEndpoint("Device", "http://"+xaddr+"/onvif/device_service")

	getCapabilities := Device.GetCapabilities{Category: "All"}

	resp, err := dev.CallMethod(getCapabilities)
	//fmt.Println(resp.Request.Host)
	//fmt.Println(readResponse(resp))
	if err != nil || resp.StatusCode != http.StatusOK {
		//panic(errors.New("camera is not available at " + xaddr + " or it does not support ONVIF services"))
		return nil, errors.New("camera is not available at " + xaddr + " or it does not support ONVIF services")
	}

	dev.getSupportedServices(resp)
	return dev, nil
}

func (dev *device)addEndpoint(Key, Value string) {
	dev.endpoints[Key]=Value
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
func (dev device) CallMethod(method interface{}) (*http.Response, error) {
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

	////fmt.Println("endpoint", endpoint)

	//TODO: Get endpoint automatically
	if dev.login != "" && dev.password != "" {
		resp, err := dev.callAuthorizedMethod(endpoint, method)
		if resp.StatusCode == 401 {
			return resp, errors.New("device requires HTTP digest authentication, which is not implemented there")
		}
		return resp, err
		//return dev.callAuthorizedMethod(endpoint, method)
	} else {
		resp, err := dev.callNonAuthorizedMethod(endpoint, method)
		if resp.StatusCode == 401 {
			return resp, errors.New("device requires HTTP digest authentication, which is not implemented there")
		}
		return resp, err
		//return dev.callNonAuthorizedMethod(endpoint, method)

	}
}

//CallNonAuthorizedMethod functions call an method, defined <method> struct without authentication data
func (dev device) callNonAuthorizedMethod(endpoint string, method interface{}) (*http.Response, error) {
	//TODO: Get endpoint automatically
	/*
	Converting <method> struct to xml string representation
	 */
	output, err := xml.MarshalIndent(method, "  ", "    ")
	if err != nil {
		//log.Printf("error: %v\n", err.Error())
		return nil, err
	}

	/*
	Build an SOAP request with <method>
	 */
	soap, err := buildMethodSOAP(string(output))
	if err != nil {
		//log.Printf("error: %v\n", err)
		return nil, err
	}

	soap.AddRootNamespaces(xlmns)

	/*
	Sending request and returns the response
	 */
	return networking.SendSoap(endpoint, soap.String())
}

//CallMethod functions call an method, defined <method> struct with authentication data
func (dev device) callAuthorizedMethod(endpoint string, method interface{}) (*http.Response, error) {
	/*
	Converting <method> struct to xml string representation
	 */
	output, err := xml.MarshalIndent(method, "  ", "    ")
	if err != nil {
		//log.Printf("error: %v\n", err.Error())
		return nil, err
	}

	/*
	Build an SOAP request with <method>
	 */
	soap, err := buildMethodSOAP(string(output))
	if err != nil {
		//log.Printf("error: %v\n", err.Error())
		return nil, err
	}

	soap.AddWSSecurity(dev.login, dev.password)

	/*
	Sending request and returns the response
	 */
	return networking.SendSoap(endpoint, soap.String())
}
