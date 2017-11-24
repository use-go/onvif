package goonvif

import "net"

type Device struct {

	xaddr net.IP
	login string
	password string

	token [64]uint8
}

