package PTZ

import (
	"github.com/yakovlevdmv/goonvif/xsd"
	"github.com/yakovlevdmv/goonvif/xsd/onvif"
)

type Capabilities struct {
	EFlip 						xsd.Boolean `xml:"EFlip,attr"`
	Reverse 					xsd.Boolean `xml:"Reverse,attr"`
	GetCompatibleConfigurations xsd.Boolean `xml:"GetCompatibleConfigurations,attr"`
	MoveStatus 					xsd.Boolean `xml:"MoveStatus,attr"`
	StatusPosition 				xsd.Boolean `xml:"StatusPosition,attr"`
}

//PTZ main types

type GetServiceCapabilities struct {
	XMLName string `xml:"tptz:GetServiceCapabilities"`
}


type GetServiceCapabilitiesResponse struct {
	Capabilities Capabilities

}


type GetNodes struct {
	XMLName string `xml:"tptz:GetNodes"`
}


type GetNodesResponse struct {
	PTZNode onvif.PTZNode

}


type GetNode struct {
	XMLName string `xml:"tptz:GetNode"`
	NodeToken onvif.ReferenceToken `xml:"tptz:NodeToken"`

}


type GetNodeResponse struct {
	PTZNode onvif.PTZNode

}


type GetConfiguration struct {
	XMLName string `xml:"tptz:GetConfiguration"`
	ProfileToken onvif.ReferenceToken `xml:"tptz:ProfileToken"`
}


type GetConfigurationResponse struct {
	PTZConfiguration onvif.PTZConfiguration

}


type GetConfigurations struct {
	XMLName string `xml:"tptz:GetConfigurations"`
}


type GetConfigurationsResponse struct {
	PTZConfiguration onvif.PTZConfiguration

}


type SetConfiguration struct {
	XMLName string `xml:"tptz:SetConfiguration"`
	PTZConfiguration onvif.PTZConfiguration `xml:"tptz:PTZConfiguration"`
	ForcePersistence xsd.Boolean `xml:"tptz:ForcePersistence"`

}


type SetConfigurationResponse struct {

}


type GetConfigurationOptions struct {
	XMLName string `xml:"tptz:GetConfigurationOptions"`
	ProfileToken onvif.ReferenceToken `xml:"tptz:ProfileToken"`

}


type GetConfigurationOptionsResponse struct {
	PTZConfigurationOptions onvif.PTZConfigurationOptions

}


type SendAuxiliaryCommand struct {
	XMLName string `xml:"tptz:SendAuxiliaryCommand"`
	ProfileToken onvif.ReferenceToken `xml:"tptz:ProfileToken"`
	AuxiliaryData onvif.AuxiliaryData `xml:"tptz:AuxiliaryData"`

}


type SendAuxiliaryCommandResponse struct {
	AuxiliaryResponse onvif.AuxiliaryData

}


type GetPresets struct {
	XMLName string `xml:"tptz:GetPresets"`
	ProfileToken onvif.ReferenceToken `xml:"tptz:ProfileToken"`
}


type GetPresetsResponse struct {
	Preset onvif.PTZPreset

}


type SetPreset struct {
	XMLName string `xml:"tptz:SetPreset"`
	ProfileToken onvif.ReferenceToken `xml:"tptz:ProfileToken"`
	PresetName xsd.String `xml:"tptz:PresetName"`
	PresetToken onvif.ReferenceToken `xml:"tptz:PresetToken"`
}


type SetPresetResponse struct {
	PresetToken onvif.ReferenceToken

}


type RemovePreset struct {
	XMLName string `xml:"tptz:RemovePreset"`
	ProfileToken onvif.ReferenceToken `xml:"tptz:ProfileToken"`
	PresetToken onvif.ReferenceToken `xml:"tptz:PresetToken"`

}


type RemovePresetResponse struct {

}


type GotoPreset struct {
	XMLName string `xml:"tptz:GotoPreset"`
	ProfileToken onvif.ReferenceToken `xml:"tptz:ProfileToken"`
	PresetToken onvif.ReferenceToken `xml:"tptz:PresetToken"`
	Speed onvif.PTZSpeed `xml:"tptz:Speed"`

}


type GotoPresetResponse struct {

}


type GotoHomePosition struct {
	XMLName string `xml:"tptz:GotoHomePosition"`
	ProfileToken onvif.ReferenceToken `xml:"tptz:ProfileToken"`
	Speed onvif.PTZSpeed `xml:"tptz:Speed"`

}


type GotoHomePositionResponse struct {

}


type SetHomePosition struct {
	XMLName string `xml:"tptz:SetHomePosition"`
	ProfileToken onvif.ReferenceToken `xml:"tptz:ProfileToken"`
}


type SetHomePositionResponse struct {

}


type ContinuousMove struct {
	XMLName string `xml:"tptz:ContinuousMove"`
	ProfileToken onvif.ReferenceToken `xml:"tptz:ProfileToken"`
	Velocity onvif.PTZSpeed `xml:"tptz:Velocity"`
	Timeout xsd.Duration `xml:"tptz:Timeout"`

}


type ContinuousMoveResponse struct {

}


type RelativeMove struct {
	XMLName string `xml:"tptz:RelativeMove"`
	ProfileToken onvif.ReferenceToken `xml:"tptz:ProfileToken"`
	Translation onvif.PTZVector `xml:"tptz:Translation"`
	Speed onvif.PTZSpeed `xml:"tptz:Speed"`

}


type RelativeMoveResponse struct {

}


type GetStatus struct {
	XMLName string `xml:"tptz:GetStatus"`
	ProfileToken onvif.ReferenceToken `xml:"tptz:ProfileToken"`
}


type GetStatusResponse struct {
	PTZStatus onvif.PTZStatus

}


type AbsoluteMove struct {
	XMLName string `xml:"tptz:AbsoluteMove"`
	ProfileToken onvif.ReferenceToken `xml:"tptz:ProfileToken"`
	Position onvif.PTZVector `xml:"tptz:Position"`
	Speed onvif.PTZSpeed `xml:"tptz:Speed"`

}


type AbsoluteMoveResponse struct {

}


type GeoMove struct {
	XMLName string `xml:"tptz:GeoMove"`
	ProfileToken onvif.ReferenceToken `xml:"tptz:ProfileToken"`
	Target onvif.GeoLocation `xml:"tptz:Target"`
	Speed onvif.PTZSpeed `xml:"tptz:Speed"`
	AreaHeight xsd.Float `xml:"tptz:AreaHeight"`
	AreaWidth xsd.Float `xml:"tptz:AreaWidth"`

}


type GeoMoveResponse struct {

}


type Stop struct {
	XMLName string `xml:"tptz:Stop"`
	ProfileToken onvif.ReferenceToken `xml:"tptz:ProfileToken"`
	PanTilt xsd.Boolean `xml:"tptz:PanTilt"`
	Zoom xsd.Boolean `xml:"tptz:Zoom"`

}


type StopResponse struct {

}


type GetPresetTours struct {
	XMLName string `xml:"tptz:GetPresetTours"`
	ProfileToken onvif.ReferenceToken `xml:"tptz:ProfileToken"`
}


type GetPresetToursResponse struct {
	PresetTour onvif.PresetTour

}


type GetPresetTour struct {
	XMLName string `xml:"tptz:GetPresetTour"`
	ProfileToken onvif.ReferenceToken `xml:"tptz:ProfileToken"`
	PresetTourToken onvif.ReferenceToken `xml:"tptz:PresetTourToken"`

}


type GetPresetTourResponse struct {
	PresetTour onvif.PresetTour

}


type GetPresetTourOptions struct {
	XMLName string `xml:"tptz:GetPresetTourOptions"`
	ProfileToken onvif.ReferenceToken `xml:"tptz:ProfileToken"`
	PresetTourToken onvif.ReferenceToken `xml:"tptz:PresetTourToken"`

}


type GetPresetTourOptionsResponse struct {
	Options onvif.PTZPresetTourOptions

}


type CreatePresetTour struct {
	XMLName string `xml:"tptz:CreatePresetTour"`
	ProfileToken onvif.ReferenceToken `xml:"tptz:ProfileToken"`

}


type CreatePresetTourResponse struct {
	PresetTourToken onvif.ReferenceToken

}


type ModifyPresetTour struct {
	XMLName string `xml:"tptz:ModifyPresetTour"`
	ProfileToken onvif.ReferenceToken `xml:"tptz:ProfileToken"`
	PresetTour onvif.PresetTour `xml:"tptz:PresetTour"`

}


type ModifyPresetTourResponse struct {

}


type OperatePresetTour struct {
	XMLName string `xml:"tptz:OperatePresetTour"`
	ProfileToken onvif.ReferenceToken `xml:"tptz:ProfileToken"`
	PresetTourToken onvif.ReferenceToken `xml:"onvif:PresetTourToken"`
	Operation onvif.PTZPresetTourOperation `xml:"onvif:Operation"`

}


type OperatePresetTourResponse struct {

}


type RemovePresetTour struct {
	XMLName string `xml:"tptz:RemovePresetTour"`
	ProfileToken onvif.ReferenceToken `xml:"tptz:ProfileToken"`
	PresetTourToken onvif.ReferenceToken `xml:"tptz:PresetTourToken"`

}


type RemovePresetTourResponse struct {

}


type GetCompatibleConfigurations struct {
	XMLName string `xml:"tptz:GetCompatibleConfigurations"`
	ProfileToken onvif.ReferenceToken `xml:"tptz:ProfileToken"`

}


type GetCompatibleConfigurationsResponse struct {
	PTZConfiguration onvif.PTZConfiguration

}
