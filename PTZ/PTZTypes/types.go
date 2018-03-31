package PTZTypes

import "github.com/yakovlevdmv/goonvif/xsdTypes"
import "github.com/yakovlevdmv/goonvif/xsdTypes/onvif"

type Capabilities struct {
	EFlip 						xsdTypes.Boolean `xml:"EFlip,attr"`
	Reverse 					xsdTypes.Boolean `xml:"Reverse,attr"`
	GetCompatibleConfigurations xsdTypes.Boolean `xml:"GetCompatibleConfigurations,attr"`
	MoveStatus 					xsdTypes.Boolean `xml:"MoveStatus,attr"`
	StatusPosition 				xsdTypes.Boolean `xml:"StatusPosition,attr"`
}


//PTZ main types

type GetServiceCapabilities struct {

}


type GetServiceCapabilitiesResponse struct {
	Capabilities Capabilities

}


type GetNodes struct {

}


type GetNodesResponse struct {
	PTZNode onvif.PTZNode

}


type GetNode struct {
	NodeToken onvif.ReferenceToken

}


type GetNodeResponse struct {
	PTZNode onvif.PTZNode

}


type GetConfiguration struct {
	PTZConfigurationToken onvif.ReferenceToken

}


type GetConfigurationResponse struct {
	PTZConfiguration onvif.PTZConfiguration

}


type GetConfigurations struct {

}


type GetConfigurationsResponse struct {
	PTZConfiguration onvif.PTZConfiguration

}


type SetConfiguration struct {
	PTZConfiguration onvif.PTZConfiguration
	ForcePersistence xsdTypes.Boolean

}


type SetConfigurationResponse struct {

}


type GetConfigurationOptions struct {
	ConfigurationToken onvif.ReferenceToken

}


type GetConfigurationOptionsResponse struct {
	PTZConfigurationOptions onvif.PTZConfigurationOptions

}


type SendAuxiliaryCommand struct {
	ProfileToken onvif.ReferenceToken
	AuxiliaryData onvif.AuxiliaryData

}


type SendAuxiliaryCommandResponse struct {
	AuxiliaryResponse onvif.AuxiliaryData

}


type GetPresets struct {
	ProfileToken onvif.ReferenceToken

}


type GetPresetsResponse struct {
	Preset onvif.PTZPreset

}


type SetPreset struct {
	ProfileToken onvif.ReferenceToken
	PresetName string
	PresetToken onvif.ReferenceToken

}


type SetPresetResponse struct {
	PresetToken onvif.ReferenceToken

}


type RemovePreset struct {
	ProfileToken onvif.ReferenceToken
	PresetToken onvif.ReferenceToken

}


type RemovePresetResponse struct {

}


type GotoPreset struct {
	ProfileToken onvif.ReferenceToken
	PresetToken onvif.ReferenceToken
	Speed onvif.PTZSpeed

}


type GotoPresetResponse struct {

}


type GotoHomePosition struct {
	ProfileToken onvif.ReferenceToken
	Speed onvif.PTZSpeed

}


type GotoHomePositionResponse struct {

}


type SetHomePosition struct {
	ProfileToken onvif.ReferenceToken

}


type SetHomePositionResponse struct {

}


type ContinuousMove struct {
	ProfileToken onvif.ReferenceToken
	Velocity onvif.PTZSpeed
	Timeout xsdTypes.Duration

}


type ContinuousMoveResponse struct {

}


type RelativeMove struct {
	ProfileToken onvif.ReferenceToken
	Translation onvif.PTZVector
	Speed onvif.PTZSpeed

}


type RelativeMoveResponse struct {

}


type GetStatus struct {
	ProfileToken onvif.ReferenceToken

}


type GetStatusResponse struct {
	PTZStatus onvif.PTZStatus

}


type AbsoluteMove struct {
	ProfileToken onvif.ReferenceToken
	Position onvif.PTZVector
	Speed onvif.PTZSpeed

}


type AbsoluteMoveResponse struct {

}


type GeoMove struct {
	ProfileToken onvif.ReferenceToken
	Target onvif.GeoLocation
	Speed onvif.PTZSpeed
	AreaHeight xsdTypes.Float
	AreaWidth xsdTypes.Float

}


type GeoMoveResponse struct {

}


type Stop struct {
	ProfileToken onvif.ReferenceToken
	PanTilt xsdTypes.Boolean
	Zoom xsdTypes.Boolean

}


type StopResponse struct {

}


type GetPresetTours struct {
	ProfileToken onvif.ReferenceToken

}


type GetPresetToursResponse struct {
	PresetTour onvif.PresetTour

}


type GetPresetTour struct {
	ProfileToken onvif.ReferenceToken
	PresetTourToken onvif.ReferenceToken

}


type GetPresetTourResponse struct {
	PresetTour onvif.PresetTour

}


type GetPresetTourOptions struct {
	ProfileToken onvif.ReferenceToken
	PresetTourToken onvif.ReferenceToken

}


type GetPresetTourOptionsResponse struct {
	Options onvif.PTZPresetTourOptions

}


type CreatePresetTour struct {
	ProfileToken onvif.ReferenceToken

}


type CreatePresetTourResponse struct {
	PresetTourToken onvif.ReferenceToken

}


type ModifyPresetTour struct {
	ProfileToken onvif.ReferenceToken
	PresetTour onvif.PresetTour

}


type ModifyPresetTourResponse struct {

}


type OperatePresetTour struct {
	ProfileToken onvif.ReferenceToken
	PresetTourToken onvif.ReferenceToken
	Operation onvif.PTZPresetTourOperation

}


type OperatePresetTourResponse struct {

}


type RemovePresetTour struct {
	ProfileToken onvif.ReferenceToken
	PresetTourToken onvif.ReferenceToken

}


type RemovePresetTourResponse struct {

}


type GetCompatibleConfigurations struct {
	ProfileToken onvif.ReferenceToken

}


type GetCompatibleConfigurationsResponse struct {
	PTZConfiguration onvif.PTZConfiguration

}
