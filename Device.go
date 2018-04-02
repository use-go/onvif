package goonvif

import (
	"net"
	"github.com/yakovlevdmv/goonvif/SOAP"
	"encoding/xml"
	"log"
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

//TODO: Get endpoint automatically
func (dev Device) CallMethod(endpoint string, method interface{}) {

	output, err := xml.MarshalIndent(method, "  ", "    ")
	if err != nil {
		log.Printf("error: %v\n", err)
	} else {
		log.Println("Marshalled struct: ", string(output))
	}
	soap := SOAP.BuildMethodSOAP(string(output))

	Networking.SendSoap(endpoint, soap.String())
}


