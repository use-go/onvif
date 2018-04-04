package Imaging

import (
	"github.com/yakovlevdmv/goonvif/xsd/onvif"
	"github.com/yakovlevdmv/goonvif/xsd"
)

type GetServiceCapabilities struct {
	XMLName string `xml:"wsdl:GetServiceCapabilities"`
}


type GetImagingSettings struct {
	XMLName string `xml:"wsdl:GetImagingSettings"`
	VideoSourceToken onvif.ReferenceToken `xml:"wsdl:VideoSourceToken"`

}


type SetImagingSettings struct {
	XMLName string `xml:"wsdl:SetImagingSettings"`
	VideoSourceToken onvif.ReferenceToken `xml:"wsdl:VideoSourceToken"`
	ImagingSettings onvif.ImagingSettings20 `xml:"wsdl:ImagingSettings"`
	ForcePersistence xsd.Boolean `xml:"wsdl:ForcePersistence"`

}


type GetOptions struct {
	XMLName string `xml:"wsdl:GetOptions"`
	VideoSourceToken onvif.ReferenceToken `xml:"wsdl:VideoSourceToken"`

}


type Move struct {
	XMLName string `xml:"wsdl:Move"`
	VideoSourceToken onvif.ReferenceToken `xml:"wsdl:VideoSourceToken"`
	Focus onvif.FocusMove `xml:"wsdl:Focus"`

}


type GetMoveOptions struct {
	XMLName string `xml:"wsdl:GetMoveOptions"`
	VideoSourceToken onvif.ReferenceToken `xml:"wsdl:VideoSourceToken"`

}


type Stop struct {
	XMLName string `xml:"wsdl:Stop"`
	VideoSourceToken onvif.ReferenceToken `xml:"wsdl:VideoSourceToken"`

}


type GetStatus struct {
	XMLName string `xml:"wsdl:GetStatus"`
	VideoSourceToken onvif.ReferenceToken `xml:"wsdl:VideoSourceToken"`

}


type GetPresets struct {
	XMLName string `xml:"wsdl:GetPresets"`
	VideoSourceToken onvif.ReferenceToken `xml:"wsdl:VideoSourceToken"`

}


type GetCurrentPreset struct {
	XMLName string `xml:"wsdl:GetCurrentPreset"`
	VideoSourceToken onvif.ReferenceToken `xml:"wsdl:VideoSourceToken"`

}


type SetCurrentPreset struct {
	XMLName string `xml:"wsdl:SetCurrentPreset"`
	VideoSourceToken onvif.ReferenceToken `xml:"wsdl:VideoSourceToken"`
	PresetToken onvif.ReferenceToken `xml:"wsdl:PresetToken"`

}