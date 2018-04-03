package Imaging

import (
	"github.com/yakovlevdmv/goonvif/xsd/onvif"
	"github.com/yakovlevdmv/goonvif/xsd"
)

const WSDL  = "http://www.onvif.org/ver20/imaging/wsdl"

type GetServiceCapabilities struct {

}


type GetImagingSettings struct {
	VideoSourceToken onvif.ReferenceToken

}


type SetImagingSettings struct {
	VideoSourceToken onvif.ReferenceToken
	ImagingSettings onvif.ImagingSettings20
	ForcePersistence xsd.Boolean

}


type GetOptions struct {
	VideoSourceToken onvif.ReferenceToken

}


type Move struct {
	VideoSourceToken onvif.ReferenceToken
	Focus onvif.FocusMove

}


type GetMoveOptions struct {
	VideoSourceToken onvif.ReferenceToken

}


type Stop struct {
	VideoSourceToken onvif.ReferenceToken

}


type GetStatus struct {
	VideoSourceToken onvif.ReferenceToken

}


type GetPresets struct {
	VideoSourceToken onvif.ReferenceToken

}


type GetCurrentPreset struct {
	VideoSourceToken onvif.ReferenceToken

}


type SetCurrentPreset struct {
	VideoSourceToken onvif.ReferenceToken
	PresetToken onvif.ReferenceToken

}