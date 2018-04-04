package PTZ

import "github.com/yakovlevdmv/goonvif/xsd"
import "github.com/yakovlevdmv/goonvif/xsd/onvif"

type Capabilities struct {
	EFlip 						xsd.Boolean `xml:"EFlip,attr"`
	Reverse 					xsd.Boolean `xml:"Reverse,attr"`
	GetCompatibleConfigurations xsd.Boolean `xml:"GetCompatibleConfigurations,attr"`
	MoveStatus 					xsd.Boolean `xml:"MoveStatus,attr"`
	StatusPosition 				xsd.Boolean `xml:"StatusPosition,attr"`
}

//PTZ main types

type GetServiceCapabilities struct {
	XMLName string `xml:"wsdl:GetServiceCapabilities"`
}


type GetServiceCapabilitiesResponse struct {
	Capabilities Capabilities

}


type GetNodes struct {
	XMLName string `xml:"wsdl:GetNodes"`
}


type GetNodesResponse struct {
	PTZNode onvif.PTZNode

}


type GetNode struct {
	XMLName string `xml:"wsdl:GetNode"`
	NodeToken onvif.ReferenceToken `xml:"wsdl:NodeToken"`

}


type GetNodeResponse struct {
	PTZNode onvif.PTZNode

}


type GetConfiguration struct {
	XMLName string `xml:"wsdl:GetConfiguration"`
	ProfileToken onvif.ReferenceToken `xml:"wsdl:ProfileToken"`
}


type GetConfigurationResponse struct {
	PTZConfiguration onvif.PTZConfiguration

}


type GetConfigurations struct {
	XMLName string `xml:"wsdl:GetConfigurations"`
}


type GetConfigurationsResponse struct {
	PTZConfiguration onvif.PTZConfiguration

}


type SetConfiguration struct {
	XMLName string `xml:"wsdl:SetConfiguration"`
	PTZConfiguration onvif.PTZConfiguration `xml:"wsdl:PTZConfiguration"`
	ForcePersistence xsd.Boolean `xml:"wsdl:ForcePersistence"`

}


type SetConfigurationResponse struct {

}


type GetConfigurationOptions struct {
	XMLName string `xml:"wsdl:GetConfigurationOptions"`
	ProfileToken onvif.ReferenceToken `xml:"wsdl:ProfileToken"`

}


type GetConfigurationOptionsResponse struct {
	PTZConfigurationOptions onvif.PTZConfigurationOptions

}


type SendAuxiliaryCommand struct {
	XMLName string `xml:"wsdl:SendAuxiliaryCommand"`
	ProfileToken onvif.ReferenceToken `xml:"wsdl:ProfileToken"`
	AuxiliaryData onvif.AuxiliaryData `xml:"wsdl:AuxiliaryData"`

}


type SendAuxiliaryCommandResponse struct {
	AuxiliaryResponse onvif.AuxiliaryData

}


type GetPresets struct {
	XMLName string `xml:"wsdl:GetPresets"`
	ProfileToken onvif.ReferenceToken `xml:"wsdl:ProfileToken"`
}


type GetPresetsResponse struct {
	Preset onvif.PTZPreset

}


type SetPreset struct {
	XMLName string `xml:"wsdl:SetPreset"`
	ProfileToken onvif.ReferenceToken `xml:"wsdl:ProfileToken"`
	PresetName xsd.String `xml:"wsdl:PresetName"`
	PresetToken onvif.ReferenceToken `xml:"wsdl:PresetToken"`
}


type SetPresetResponse struct {
	PresetToken onvif.ReferenceToken

}


type RemovePreset struct {
	XMLName string `xml:"wsdl:RemovePreset"`
	ProfileToken onvif.ReferenceToken `xml:"wsdl:ProfileToken"`
	PresetToken onvif.ReferenceToken `xml:"wsdl:PresetToken"`

}


type RemovePresetResponse struct {

}


type GotoPreset struct {
	XMLName string `xml:"wsdl:GotoPreset"`
	ProfileToken onvif.ReferenceToken `xml:"wsdl:ProfileToken"`
	PresetToken onvif.ReferenceToken `xml:"wsdl:PresetToken"`
	Speed onvif.PTZSpeed `xml:"wsdl:Speed"`

}


type GotoPresetResponse struct {

}


type GotoHomePosition struct {
	XMLName string `xml:"wsdl:GotoHomePosition"`
	ProfileToken onvif.ReferenceToken `xml:"wsdl:ProfileToken"`
	Speed onvif.PTZSpeed `xml:"wsdl:Speed"`

}


type GotoHomePositionResponse struct {

}


type SetHomePosition struct {
	XMLName string `xml:"wsdl:SetHomePosition"`
	ProfileToken onvif.ReferenceToken `xml:"wsdl:ProfileToken"`
}


type SetHomePositionResponse struct {

}


type ContinuousMove struct {
	XMLName string `xml:"wsdl:ContinuousMove"`
	ProfileToken onvif.ReferenceToken `xml:"wsdl:ProfileToken"`
	Velocity onvif.PTZSpeed `xml:"wsdl:Velocity"`
	Timeout xsd.Duration `xml:"wsdl:Timeout"`

}


type ContinuousMoveResponse struct {

}


type RelativeMove struct {
	XMLName string `xml:"wsdl:RelativeMove"`
	ProfileToken onvif.ReferenceToken `xml:"wsdl:ProfileToken"`
	Translation onvif.PTZVector `xml:"wsdl:Translation"`
	Speed onvif.PTZSpeed `xml:"wsdl:Speed"`

}


type RelativeMoveResponse struct {

}


type GetStatus struct {
	XMLName string `xml:"wsdl:GetStatus"`
	ProfileToken onvif.ReferenceToken `xml:"wsdl:ProfileToken"`
}


type GetStatusResponse struct {
	PTZStatus onvif.PTZStatus

}


type AbsoluteMove struct {
	XMLName string `xml:"wsdl:AbsoluteMove"`
	ProfileToken onvif.ReferenceToken `xml:"wsdl:ProfileToken"`
	Position onvif.PTZVector `xml:"wsdl:Position"`
	Speed onvif.PTZSpeed `xml:"wsdl:Speed"`

}


type AbsoluteMoveResponse struct {

}


type GeoMove struct {
	XMLName string `xml:"wsdl:GeoMove"`
	ProfileToken onvif.ReferenceToken `xml:"wsdl:ProfileToken"`
	Target onvif.GeoLocation `xml:"wsdl:Target"`
	Speed onvif.PTZSpeed `xml:"wsdl:Speed"`
	AreaHeight xsd.Float `xml:"wsdl:AreaHeight"`
	AreaWidth xsd.Float `xml:"wsdl:AreaWidth"`

}


type GeoMoveResponse struct {

}


type Stop struct {
	XMLName string `xml:"wsdl:Stop"`
	ProfileToken onvif.ReferenceToken `xml:"wsdl:ProfileToken"`
	PanTilt xsd.Boolean `xml:"wsdl:PanTilt"`
	Zoom xsd.Boolean `xml:"wsdl:Zoom"`

}


type StopResponse struct {

}


type GetPresetTours struct {
	XMLName string `xml:"wsdl:GetPresetTours"`
	ProfileToken onvif.ReferenceToken `xml:"wsdl:ProfileToken"`
}


type GetPresetToursResponse struct {
	PresetTour onvif.PresetTour

}


type GetPresetTour struct {
	XMLName string `xml:"wsdl:GetPresetTour"`
	ProfileToken onvif.ReferenceToken `xml:"wsdl:ProfileToken"`
	PresetTourToken onvif.ReferenceToken `xml:"wsdl:PresetTourToken"`

}


type GetPresetTourResponse struct {
	PresetTour onvif.PresetTour

}


type GetPresetTourOptions struct {
	XMLName string `xml:"wsdl:GetPresetTourOptions"`
	ProfileToken onvif.ReferenceToken `xml:"wsdl:ProfileToken"`
	PresetTourToken onvif.ReferenceToken `xml:"wsdl:PresetTourToken"`

}


type GetPresetTourOptionsResponse struct {
	Options onvif.PTZPresetTourOptions

}


type CreatePresetTour struct {
	XMLName string `xml:"wsdl:CreatePresetTour"`
	ProfileToken onvif.ReferenceToken `xml:"wsdl:ProfileToken"`

}


type CreatePresetTourResponse struct {
	PresetTourToken onvif.ReferenceToken

}


type ModifyPresetTour struct {
	XMLName string `xml:"wsdl:ModifyPresetTour"`
	ProfileToken onvif.ReferenceToken `xml:"wsdl:ProfileToken"`
	PresetTour onvif.PresetTour `xml:"wsdl:PresetTour"`

}


type ModifyPresetTourResponse struct {

}


type OperatePresetTour struct {
	XMLName string `xml:"wsdl:OperatePresetTour"`
	ProfileToken onvif.ReferenceToken `xml:"wsdl:ProfileToken"`
	PresetTourToken onvif.ReferenceToken `xml:"onvif:PresetTourToken"`
	Operation onvif.PTZPresetTourOperation `xml:"onvif:Operation"`

}


type OperatePresetTourResponse struct {

}


type RemovePresetTour struct {
	XMLName string `xml:"wsdl:RemovePresetTour"`
	ProfileToken onvif.ReferenceToken `xml:"wsdl:ProfileToken"`
	PresetTourToken onvif.ReferenceToken `xml:"wsdl:PresetTourToken"`

}


type RemovePresetTourResponse struct {

}


type GetCompatibleConfigurations struct {
	XMLName string `xml:"wsdl:GetCompatibleConfigurations"`
	ProfileToken onvif.ReferenceToken `xml:"wsdl:ProfileToken"`

}


type GetCompatibleConfigurationsResponse struct {
	PTZConfiguration onvif.PTZConfiguration

}
