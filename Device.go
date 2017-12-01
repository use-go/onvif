package goonvif

import "net"

type Device struct {

	xaddr net.IP
	login string
	password string

	token [64]uint8
}

type Service struct {
	name string

	device *Device
}

func NewService(name string, dev *Device) Service {
	return Service{name:name, device:dev}
}


