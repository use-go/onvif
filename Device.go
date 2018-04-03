package goonvif

import (
	"net"
	"encoding/xml"
	"log"
	"fmt"
	"github.com/beevik/etree"
	"github.com/yakovlevdmv/gosoap"
	"github.com/yakovlevdmv/goonvif/Networking"
)

type DeviceInfo struct {
	Manufacturer string
	Model string
	FirmwareVersion string
	SerialNumber string
	HardwareId string

}

type Device struct {

	xaddr net.IP
	login string
	password string

	token [64]uint8

	endpoints map[string]string
	info DeviceInfo

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
func (dev Device) CallMethod(endpoint string, method interface{}) {

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


