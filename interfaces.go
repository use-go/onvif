package onvif

type Function interface {
	Request() interface{}
	Response() interface{}
}
