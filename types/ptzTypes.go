package types


type GetServiceCapabilities struct {

}


type GetServiceCapabilitiesResponse struct {
	Capabilities PTZCapabilities

}


type GetNodes struct {

}


type GetNodesResponse struct {
	PTZNode PTZNode

}


type GetNode struct {
	NodeToken ReferenceToken

}


type GetNodeResponse struct {
	PTZNode PTZNode

}


type GetConfiguration struct {
	PTZConfigurationToken ReferenceToken

}


type GetConfigurationResponse struct {
	PTZConfiguration PTZConfiguration

}


type GetConfigurations struct {

}


type GetConfigurationsResponse struct {
	PTZConfiguration PTZConfiguration

}


type SetConfiguration struct {
	PTZConfiguration PTZConfiguration
	ForcePersistence Boolean

}


type SetConfigurationResponse struct {

}


type GetConfigurationOptions struct {
	ConfigurationToken ReferenceToken

}


type GetConfigurationOptionsResponse struct {
	PTZConfigurationOptions PTZConfigurationOptions

}


type SendAuxiliaryCommand struct {
	ProfileToken ReferenceToken
	AuxiliaryData AuxiliaryData

}


type SendAuxiliaryCommandResponse struct {
	AuxiliaryResponse AuxiliaryData

}


type GetPresets struct {
	ProfileToken ReferenceToken

}


type GetPresetsResponse struct {
	Preset PTZPreset

}


type SetPreset struct {
	ProfileToken ReferenceToken
	PresetName string
	PresetToken ReferenceToken

}


type SetPresetResponse struct {
	PresetToken ReferenceToken

}


type RemovePreset struct {
	ProfileToken ReferenceToken
	PresetToken ReferenceToken

}


type RemovePresetResponse struct {

}


type GotoPreset struct {
	ProfileToken ReferenceToken
	PresetToken ReferenceToken
	Speed PTZSpeed

}


type GotoPresetResponse struct {

}


type GotoHomePosition struct {
	ProfileToken ReferenceToken
	Speed PTZSpeed

}


type GotoHomePositionResponse struct {

}


type SetHomePosition struct {
	ProfileToken ReferenceToken

}


type SetHomePositionResponse struct {

}


type ContinuousMove struct {
	ProfileToken ReferenceToken
	Velocity PTZSpeed
	Timeout Duration

}


type ContinuousMoveResponse struct {

}


type RelativeMove struct {
	ProfileToken ReferenceToken
	Translation PTZVector
	Speed PTZSpeed

}


type RelativeMoveResponse struct {

}


type GetStatus struct {
	ProfileToken ReferenceToken

}


type GetStatusResponse struct {
	PTZStatus PTZStatus

}


type AbsoluteMove struct {
	ProfileToken ReferenceToken
	Position PTZVector
	Speed PTZSpeed

}


type AbsoluteMoveResponse struct {

}


type GeoMove struct {
	ProfileToken ReferenceToken
	Target GeoLocation
	Speed PTZSpeed
	AreaHeight float
	AreaWidth float

}


type GeoMoveResponse struct {

}


type Stop struct {
	ProfileToken ReferenceToken
	PanTilt Boolean
	Zoom Boolean

}


type StopResponse struct {

}


type GetPresetTours struct {
	ProfileToken ReferenceToken

}


type GetPresetToursResponse struct {
	PresetTour PresetTour

}


type GetPresetTour struct {
	ProfileToken ReferenceToken
	PresetTourToken ReferenceToken

}


type GetPresetTourResponse struct {
	PresetTour PresetTour

}


type GetPresetTourOptions struct {
	ProfileToken ReferenceToken
	PresetTourToken ReferenceToken

}


type GetPresetTourOptionsResponse struct {
	Options PTZPresetTourOptions

}


type CreatePresetTour struct {
	ProfileToken ReferenceToken

}


type CreatePresetTourResponse struct {
	PresetTourToken ReferenceToken

}


type ModifyPresetTour struct {
	ProfileToken ReferenceToken
	PresetTour PresetTour

}


type ModifyPresetTourResponse struct {

}


type OperatePresetTour struct {
	ProfileToken ReferenceToken
	PresetTourToken ReferenceToken
	Operation PTZPresetTourOperation

}


type OperatePresetTourResponse struct {

}


type RemovePresetTour struct {
	ProfileToken ReferenceToken
	PresetTourToken ReferenceToken

}


type RemovePresetTourResponse struct {

}


type GetCompatibleConfigurations struct {
	ProfileToken ReferenceToken

}


type GetCompatibleConfigurationsResponse struct {
	PTZConfiguration PTZConfiguration

}

func FnGetServiceCapabilities (arg GetServiceCapabilities) {

}

func FnGetServiceCapabilitiesResponse (arg GetServiceCapabilitiesResponse) {

}

func FnGetNodes (arg GetNodes) {

}

func FnGetNodesResponse (arg GetNodesResponse) {

}

func FnGetNode (arg GetNode) {

}

func FnGetNodeResponse (arg GetNodeResponse) {

}

func FnGetConfiguration (arg GetConfiguration) {

}

func FnGetConfigurationResponse (arg GetConfigurationResponse) {

}

func FnGetConfigurations (arg GetConfigurations) {

}

func FnGetConfigurationsResponse (arg GetConfigurationsResponse) {

}

func FnSetConfiguration (arg SetConfiguration) {

}

func FnSetConfigurationResponse (arg SetConfigurationResponse) {

}

func FnGetConfigurationOptions (arg GetConfigurationOptions) {

}

func FnGetConfigurationOptionsResponse (arg GetConfigurationOptionsResponse) {

}

func FnSendAuxiliaryCommand (arg SendAuxiliaryCommand) {

}

func FnSendAuxiliaryCommandResponse (arg SendAuxiliaryCommandResponse) {

}

func FnGetPresets (arg GetPresets) {

}

func FnGetPresetsResponse (arg GetPresetsResponse) {

}

func FnSetPreset (arg SetPreset) {

}

func FnSetPresetResponse (arg SetPresetResponse) {

}

func FnRemovePreset (arg RemovePreset) {

}

func FnRemovePresetResponse (arg RemovePresetResponse) {

}

func FnGotoPreset (arg GotoPreset) {

}

func FnGotoPresetResponse (arg GotoPresetResponse) {

}

func FnGotoHomePosition (arg GotoHomePosition) {

}

func FnGotoHomePositionResponse (arg GotoHomePositionResponse) {

}

func FnSetHomePosition (arg SetHomePosition) {

}

func FnSetHomePositionResponse (arg SetHomePositionResponse) {

}

func FnContinuousMove (arg ContinuousMove) {

}

func FnContinuousMoveResponse (arg ContinuousMoveResponse) {

}

func FnRelativeMove (arg RelativeMove) {

}

func FnRelativeMoveResponse (arg RelativeMoveResponse) {

}

func FnGetStatus (arg GetStatus) {

}

func FnGetStatusResponse (arg GetStatusResponse) {

}

func FnAbsoluteMove (arg AbsoluteMove) {

}

func FnAbsoluteMoveResponse (arg AbsoluteMoveResponse) {

}

func FnGeoMove (arg GeoMove) {

}

func FnGeoMoveResponse (arg GeoMoveResponse) {

}

func FnStop (arg Stop) {

}

func FnStopResponse (arg StopResponse) {

}

func FnGetPresetTours (arg GetPresetTours) {

}

func FnGetPresetToursResponse (arg GetPresetToursResponse) {

}

func FnGetPresetTour (arg GetPresetTour) {

}

func FnGetPresetTourResponse (arg GetPresetTourResponse) {

}

func FnGetPresetTourOptions (arg GetPresetTourOptions) {

}

func FnGetPresetTourOptionsResponse (arg GetPresetTourOptionsResponse) {

}

func FnCreatePresetTour (arg CreatePresetTour) {

}

func FnCreatePresetTourResponse (arg CreatePresetTourResponse) {

}

func FnModifyPresetTour (arg ModifyPresetTour) {

}

func FnModifyPresetTourResponse (arg ModifyPresetTourResponse) {

}

func FnOperatePresetTour (arg OperatePresetTour) {

}

func FnOperatePresetTourResponse (arg OperatePresetTourResponse) {

}

func FnRemovePresetTour (arg RemovePresetTour) {

}

func FnRemovePresetTourResponse (arg RemovePresetTourResponse) {

}

func FnGetCompatibleConfigurations (arg GetCompatibleConfigurations) {

}

func FnGetCompatibleConfigurationsResponse (arg GetCompatibleConfigurationsResponse) {

}