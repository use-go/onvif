package onvif

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/beevik/etree"
	"github.com/use-go/onvif/device"
	"github.com/use-go/onvif/gosoap"
	"github.com/use-go/onvif/networking"
	wsdiscovery "github.com/use-go/onvif/ws-discovery"
)

//Xlmns XML Scheam
var Xlmns = map[string]string{
	"onvif":   "http://www.onvif.org/ver10/schema",
	"tds":     "http://www.onvif.org/ver10/device/wsdl",
	"trt":     "http://www.onvif.org/ver10/media/wsdl",
	"tev":     "http://www.onvif.org/ver10/events/wsdl",
	"tptz":    "http://www.onvif.org/ver20/ptz/wsdl",
	"timg":    "http://www.onvif.org/ver20/imaging/wsdl",
	"tan":     "http://www.onvif.org/ver20/analytics/wsdl",
	"xmime":   "http://www.w3.org/2005/05/xmlmime",
	"wsnt":    "http://docs.oasis-open.org/wsn/b-2",
	"xop":     "http://www.w3.org/2004/08/xop/include",
	"wsa":     "http://www.w3.org/2005/08/addressing",
	"wstop":   "http://docs.oasis-open.org/wsn/t-1",
	"wsntw":   "http://docs.oasis-open.org/wsn/bw-2",
	"wsrf-rw": "http://docs.oasis-open.org/wsrf/rw-2",
	"wsaw":    "http://www.w3.org/2006/05/addressing/wsdl",
}

//DeviceType alias for int
type DeviceType int

// Onvif Device Tyoe
const (
	NVD DeviceType = iota
	NVS
	NVA
	NVT
)

func (devType DeviceType) String() string {
	stringRepresentation := []string{
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
	Manufacturer    string
	Model           string
	FirmwareVersion string
	SerialNumber    string
	HardwareId      string
}

//Device for a new device of onvif and deviceInfo
//struct represents an abstract ONVIF device.
//It contains methods, which helps to communicate with ONVIF device
type Device struct {
	xaddr     string
	login     string
	password  string
	endpoints map[string]string
	info      deviceInfo
}

//GetServices return available endpoints
func (dev *Device) GetServices() map[string]string {
	return dev.endpoints
}

func readResponse(resp *http.Response) string {
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return string(b)
}

//GetAvailableDevicesAtSpecificEthernetInterface ...
func GetAvailableDevicesAtSpecificEthernetInterface(interfaceName string) []Device {
	/*
		Call an ws-discovery Probe Message to Discover NVT type Devices
	*/
	devices := wsdiscovery.SendProbe(interfaceName, nil, []string{"dn:" + NVT.String()}, map[string]string{"dn": "http://www.onvif.org/ver10/network/wsdl"})
	nvtDevices := make([]Device, 0)
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

func (dev *Device) getSupportedServices(resp *http.Response) {
	//resp, err := dev.CallMethod(device.GetCapabilities{Category:"All"})
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
	for _, j := range services {
		////fmt.Println(j.Text())
		////fmt.Println(j.Parent().Tag)
		dev.addEndpoint(j.Parent().Tag, j.Text())
	}
	//}
}

//NewDevice function construct a ONVIF Device entity
func NewDevice(xaddr string) (*Device, error) {
	dev := new(Device)
	dev.xaddr = xaddr
	dev.endpoints = make(map[string]string)
	dev.addEndpoint("Device", "http://"+xaddr+"/onvif/device_service")

	getCapabilities := device.GetCapabilities{Category: "All"}

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

func (dev *Device) addEndpoint(Key, Value string) {

	//use lowCaseKey
	//make key having ability to handle Mixed Case for Different vendor devcie (e.g. Events EVENTS, events)
	lowCaseKey := strings.ToLower(Key)
	dev.endpoints[lowCaseKey] = Value
}

//Authenticate function authenticate client in the ONVIF device.
//Function takes <username> and <password> params.
//You should use this function to allow authorized requests to the ONVIF Device
//To change auth data call this function again.
func (dev *Device) Authenticate(username, password string) {
	dev.login = username
	dev.password = password
}

//GetEndpoint returns specific ONVIF service endpoint address
func (dev *Device) GetEndpoint(name string) string {
	return dev.endpoints[name]
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
	//soap.AddRootNamespace("onvif", "http://www.onvif.org/ver10/device/wsdl")

	return soap, nil
}

//getEndpoint functions get the target service endpoint in a better way
func (dev Device) getEndpoint(endpoint string) (string, error) {

	// common condition, endpointMark in map we use this.
	if endpointURL, bFound := dev.endpoints[endpoint]; bFound {
		return endpointURL, nil
	}

	//but ,if we have endpoint like event、analytic
	//and sametime the Targetkey like : events、analytics
	//we use fuzzy way to find the best match url
	var endpointURL string
	for targetKey := range dev.endpoints {
		if strings.Contains(targetKey, endpoint) {
			endpointURL = dev.endpoints[targetKey]
			return endpointURL, nil
		}
	}
	return endpointURL, errors.New("target endpoint service not found")
}

//CallMethod functions call an method, defined <method> struct.
//You should use Authenticate method to call authorized requests.
func (dev Device) CallMethod(method interface{}) (*http.Response, error) {
	pkgPath := strings.Split(reflect.TypeOf(method).PkgPath(), "/")
	pkg := strings.ToLower(pkgPath[len(pkgPath)-1])

	endpoint, err := dev.getEndpoint(pkg)
	if err != nil {
		return nil, err
	}
	return dev.callMethodDo(endpoint, method)
}

//CallMethod functions call an method, defined <method> struct with authentication data
func (dev Device) callMethodDo(endpoint string, method interface{}) (*http.Response, error) {
	/*
		Converting <method> struct to xml string representation
	*/
	output, err := xml.MarshalIndent(method, "  ", "    ")
	if err != nil {
		//log.Printf("error: %v\n", err.Error())
		return nil, err
	}
	//fmt.Println(gosoap.SoapMessage(string(output)).StringIndent())
	/*
		Build an SOAP request with <method>
	*/
	soap, err := buildMethodSOAP(string(output))
	if err != nil {
		//log.Printf("error: %v\n", err.Error())
		return nil, err
	}

	//fmt.Println(soap.StringIndent())
	/*
		Adding namespaces and WS-Security headers
	*/
	soap.AddRootNamespaces(Xlmns)

	//fmt.Println(soap.StringIndent())
	//Header handling
	soap.AddAction()

	//Auth Handling
	if dev.login != "" && dev.password != "" {
		soap.AddWSSecurity(dev.login, dev.password)
	}
	//fmt.Println(soap.StringIndent())
	/*
		Sending request and returns the response
	*/
	return networking.SendSoap(endpoint, soap.String())
}
