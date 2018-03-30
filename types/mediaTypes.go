package types

type GetMediaServiceCapabilities struct {

}


type GetMediaServiceCapabilitiesResponse struct {
	Capabilities MediaCapabilities

}


type GetVideoSources struct {

}


type GetVideoSourcesResponse struct {
	VideoSources VideoSource

}


type GetAudioSources struct {

}


type GetAudioSourcesResponse struct {
	AudioSources AudioSource

}


type GetAudioOutputs struct {

}


type GetAudioOutputsResponse struct {
	AudioOutputs AudioOutput

}


type CreateProfile struct {
	Name Name
	Token ReferenceToken

}


type CreateProfileResponse struct {
	Profile Profile

}


type GetProfile struct {
	ProfileToken ReferenceToken

}


type GetProfileResponse struct {
	Profile Profile

}


type GetProfiles struct {

}


type GetProfilesResponse struct {
	Profiles Profile

}


type AddVideoEncoderConfiguration struct {
	ProfileToken ReferenceToken
	ConfigurationToken ReferenceToken

}


type AddVideoEncoderConfigurationResponse struct {

}


type RemoveVideoEncoderConfiguration struct {
	ProfileToken ReferenceToken

}


type RemoveVideoEncoderConfigurationResponse struct {

}


type AddVideoSourceConfiguration struct {
	ProfileToken ReferenceToken
	ConfigurationToken ReferenceToken

}


type AddVideoSourceConfigurationResponse struct {

}


type RemoveVideoSourceConfiguration struct {
	ProfileToken ReferenceToken

}


type RemoveVideoSourceConfigurationResponse struct {

}


type AddAudioEncoderConfiguration struct {
	ProfileToken ReferenceToken
	ConfigurationToken ReferenceToken

}


type AddAudioEncoderConfigurationResponse struct {

}


type RemoveAudioEncoderConfiguration struct {
	ProfileToken ReferenceToken

}


type RemoveAudioEncoderConfigurationResponse struct {

}


type AddAudioSourceConfiguration struct {
	ProfileToken ReferenceToken
	ConfigurationToken ReferenceToken

}


type AddAudioSourceConfigurationResponse struct {

}


type RemoveAudioSourceConfiguration struct {
	ProfileToken ReferenceToken

}


type RemoveAudioSourceConfigurationResponse struct {

}


type AddPTZConfiguration struct {
	ProfileToken ReferenceToken
	ConfigurationToken ReferenceToken

}


type AddPTZConfigurationResponse struct {

}


type RemovePTZConfiguration struct {
	ProfileToken ReferenceToken

}


type RemovePTZConfigurationResponse struct {

}


type AddVideoAnalyticsConfiguration struct {
	ProfileToken ReferenceToken
	ConfigurationToken ReferenceToken

}


type AddVideoAnalyticsConfigurationResponse struct {

}


type RemoveVideoAnalyticsConfiguration struct {
	ProfileToken ReferenceToken

}


type RemoveVideoAnalyticsConfigurationResponse struct {

}


type AddMetadataConfiguration struct {
	ProfileToken ReferenceToken
	ConfigurationToken ReferenceToken

}


type AddMetadataConfigurationResponse struct {

}


type RemoveMetadataConfiguration struct {
	ProfileToken ReferenceToken

}


type RemoveMetadataConfigurationResponse struct {

}


type AddAudioOutputConfiguration struct {
	ProfileToken ReferenceToken
	ConfigurationToken ReferenceToken

}


type AddAudioOutputConfigurationResponse struct {

}


type RemoveAudioOutputConfiguration struct {
	ProfileToken ReferenceToken

}


type RemoveAudioOutputConfigurationResponse struct {

}


type AddAudioDecoderConfiguration struct {
	ProfileToken ReferenceToken
	ConfigurationToken ReferenceToken

}


type AddAudioDecoderConfigurationResponse struct {

}


type RemoveAudioDecoderConfiguration struct {
	ProfileToken ReferenceToken

}


type RemoveAudioDecoderConfigurationResponse struct {

}


type DeleteProfile struct {
	ProfileToken ReferenceToken

}


type DeleteProfileResponse struct {

}


type GetVideoSourceConfigurations struct {

}


type GetVideoSourceConfigurationsResponse struct {
	Configurations VideoSourceConfiguration

}


type GetVideoEncoderConfigurations struct {

}


type GetVideoEncoderConfigurationsResponse struct {
	Configurations VideoEncoderConfiguration

}


type GetAudioSourceConfigurations struct {

}


type GetAudioSourceConfigurationsResponse struct {
	Configurations AudioSourceConfiguration

}


type GetAudioEncoderConfigurations struct {

}


type GetAudioEncoderConfigurationsResponse struct {
	Configurations AudioEncoderConfiguration

}


type GetVideoAnalyticsConfigurations struct {

}


type GetVideoAnalyticsConfigurationsResponse struct {
	Configurations VideoAnalyticsConfiguration

}


type GetMetadataConfigurations struct {

}


type GetMetadataConfigurationsResponse struct {
	Configurations MetadataConfiguration

}


type GetAudioOutputConfigurations struct {

}


type GetAudioOutputConfigurationsResponse struct {
	Configurations AudioOutputConfiguration

}


type GetAudioDecoderConfigurations struct {

}


type GetAudioDecoderConfigurationsResponse struct {
	Configurations AudioDecoderConfiguration

}


type GetVideoSourceConfiguration struct {
	ConfigurationToken ReferenceToken

}


type GetVideoSourceConfigurationResponse struct {
	Configuration VideoSourceConfiguration

}


type GetVideoEncoderConfiguration struct {
	ConfigurationToken ReferenceToken

}


type GetVideoEncoderConfigurationResponse struct {
	Configuration VideoEncoderConfiguration

}


type GetAudioSourceConfiguration struct {
	ConfigurationToken ReferenceToken

}


type GetAudioSourceConfigurationResponse struct {
	Configuration AudioSourceConfiguration

}


type GetAudioEncoderConfiguration struct {
	ConfigurationToken ReferenceToken

}


type GetAudioEncoderConfigurationResponse struct {
	Configuration AudioEncoderConfiguration

}


type GetVideoAnalyticsConfiguration struct {
	ConfigurationToken ReferenceToken

}


type GetVideoAnalyticsConfigurationResponse struct {
	Configuration VideoAnalyticsConfiguration

}


type GetMetadataConfiguration struct {
	ConfigurationToken ReferenceToken

}


type GetMetadataConfigurationResponse struct {
	Configuration MetadataConfiguration

}


type GetAudioOutputConfiguration struct {
	ConfigurationToken ReferenceToken

}


type GetAudioOutputConfigurationResponse struct {
	Configuration AudioOutputConfiguration

}


type GetAudioDecoderConfiguration struct {
	ConfigurationToken ReferenceToken

}


type GetAudioDecoderConfigurationResponse struct {
	Configuration AudioDecoderConfiguration

}


type GetCompatibleVideoEncoderConfigurations struct {
	ProfileToken ReferenceToken

}


type GetCompatibleVideoEncoderConfigurationsResponse struct {
	Configurations VideoEncoderConfiguration

}


type GetCompatibleVideoSourceConfigurations struct {
	ProfileToken ReferenceToken

}


type GetCompatibleVideoSourceConfigurationsResponse struct {
	Configurations VideoSourceConfiguration

}


type GetCompatibleAudioEncoderConfigurations struct {
	ProfileToken ReferenceToken

}


type GetCompatibleAudioEncoderConfigurationsResponse struct {
	Configurations AudioEncoderConfiguration

}


type GetCompatibleAudioSourceConfigurations struct {
	ProfileToken ReferenceToken

}


type GetCompatibleAudioSourceConfigurationsResponse struct {
	Configurations AudioSourceConfiguration

}


type GetCompatibleVideoAnalyticsConfigurations struct {
	ProfileToken ReferenceToken

}


type GetCompatibleVideoAnalyticsConfigurationsResponse struct {
	Configurations VideoAnalyticsConfiguration

}


type GetCompatibleMetadataConfigurations struct {
	ProfileToken ReferenceToken

}


type GetCompatibleMetadataConfigurationsResponse struct {
	Configurations MetadataConfiguration

}


type GetCompatibleAudioOutputConfigurations struct {
	ProfileToken ReferenceToken

}


type GetCompatibleAudioOutputConfigurationsResponse struct {
	Configurations AudioOutputConfiguration

}


type GetCompatibleAudioDecoderConfigurations struct {
	ProfileToken ReferenceToken

}


type GetCompatibleAudioDecoderConfigurationsResponse struct {
	Configurations AudioDecoderConfiguration

}


type SetVideoSourceConfiguration struct {
	Configuration VideoSourceConfiguration
	ForcePersistence bool

}


type SetVideoSourceConfigurationResponse struct {

}


type SetVideoEncoderConfiguration struct {
	Configuration VideoEncoderConfiguration
	ForcePersistence bool

}


type SetVideoEncoderConfigurationResponse struct {

}


type SetAudioSourceConfiguration struct {
	Configuration AudioSourceConfiguration
	ForcePersistence bool

}


type SetAudioSourceConfigurationResponse struct {

}


type SetAudioEncoderConfiguration struct {
	Configuration AudioEncoderConfiguration
	ForcePersistence bool

}


type SetAudioEncoderConfigurationResponse struct {

}


type SetVideoAnalyticsConfiguration struct {
	Configuration VideoAnalyticsConfiguration
	ForcePersistence bool

}


type SetVideoAnalyticsConfigurationResponse struct {

}


type SetMetadataConfiguration struct {
	Configuration MetadataConfiguration
	ForcePersistence bool

}


type SetMetadataConfigurationResponse struct {

}


type SetAudioOutputConfiguration struct {
	Configuration AudioOutputConfiguration
	ForcePersistence bool

}


type SetAudioOutputConfigurationResponse struct {

}


type SetAudioDecoderConfiguration struct {
	Configuration AudioDecoderConfiguration
	ForcePersistence bool

}


type SetAudioDecoderConfigurationResponse struct {

}


type GetVideoSourceConfigurationOptions struct {
	ConfigurationToken ReferenceToken
	ProfileToken ReferenceToken

}


type GetVideoSourceConfigurationOptionsResponse struct {
	Options VideoSourceConfigurationOptions

}


type GetVideoEncoderConfigurationOptions struct {
	ConfigurationToken ReferenceToken
	ProfileToken ReferenceToken

}


type GetVideoEncoderConfigurationOptionsResponse struct {
	Options VideoEncoderConfigurationOptions

}


type GetAudioSourceConfigurationOptions struct {
	ConfigurationToken ReferenceToken
	ProfileToken ReferenceToken

}


type GetAudioSourceConfigurationOptionsResponse struct {
	Options AudioSourceConfigurationOptions

}


type GetAudioEncoderConfigurationOptions struct {
	ConfigurationToken ReferenceToken
	ProfileToken ReferenceToken

}


type GetAudioEncoderConfigurationOptionsResponse struct {
	Options AudioEncoderConfigurationOptions

}


type GetMetadataConfigurationOptions struct {
	ConfigurationToken ReferenceToken
	ProfileToken ReferenceToken

}


type GetMetadataConfigurationOptionsResponse struct {
	Options MetadataConfigurationOptions

}


type GetAudioOutputConfigurationOptions struct {
	ConfigurationToken ReferenceToken
	ProfileToken ReferenceToken

}


type GetAudioOutputConfigurationOptionsResponse struct {
	Options AudioOutputConfigurationOptions

}


type GetAudioDecoderConfigurationOptions struct {
	ConfigurationToken ReferenceToken
	ProfileToken ReferenceToken

}


type GetAudioDecoderConfigurationOptionsResponse struct {
	Options AudioDecoderConfigurationOptions

}


type GetGuaranteedNumberOfVideoEncoderInstances struct {
	ConfigurationToken ReferenceToken

}


type GetGuaranteedNumberOfVideoEncoderInstancesResponse struct {
	TotalNumber int
	JPEG int
	H264 int
	MPEG4 int

}


type GetStreamUri struct {
	StreamSetup StreamSetup
	ProfileToken ReferenceToken

}


type GetStreamUriResponse struct {
	MediaUri MediaUri

}


type StartMulticastStreaming struct {
	ProfileToken ReferenceToken

}


type StartMulticastStreamingResponse struct {

}


type StopMulticastStreaming struct {
	ProfileToken ReferenceToken

}


type StopMulticastStreamingResponse struct {

}


type SetSynchronizationPoint struct {
	ProfileToken ReferenceToken

}


type SetSynchronizationPointResponse struct {

}


type GetSnapshotUri struct {
	ProfileToken ReferenceToken

}


type GetSnapshotUriResponse struct {
	MediaUri MediaUri

}


type GetVideoSourceModes struct {
	VideoSourceToken ReferenceToken

}


type GetVideoSourceModesResponse struct {
	VideoSourceModes VideoSourceMode

}


type SetVideoSourceMode struct {
	VideoSourceToken ReferenceToken
	VideoSourceModeToken ReferenceToken

}


type SetVideoSourceModeResponse struct {
	Reboot bool

}


type GetOSDs struct {
	ConfigurationToken ReferenceToken

}


type GetOSDsResponse struct {
	OSDs OSDConfiguration

}


type GetOSD struct {
	OSDToken ReferenceToken

}


type GetOSDResponse struct {
	OSD OSDConfiguration

}


type GetOSDOptions struct {
	ConfigurationToken ReferenceToken

}


type GetOSDOptionsResponse struct {
	OSDOptions OSDConfigurationOptions

}


type SetOSD struct {
	OSD OSDConfiguration

}


type SetOSDResponse struct {

}


type CreateOSD struct {
	OSD OSDConfiguration

}


type CreateOSDResponse struct {
	OSDToken ReferenceToken

}


type DeleteOSD struct {
	OSDToken ReferenceToken

}


type DeleteOSDResponse struct {

}

func FnGetMediaServiceCapabilities (arg GetServiceCapabilities) {

}

func FnGetMediaServiceCapabilitiesResponse (arg GetServiceCapabilitiesResponse) {

}

func FnGetVideoSources (arg GetVideoSources) {

}

func FnGetVideoSourcesResponse (arg GetVideoSourcesResponse) {

}

func FnGetAudioSources (arg GetAudioSources) {

}

func FnGetAudioSourcesResponse (arg GetAudioSourcesResponse) {

}

func FnGetAudioOutputs (arg GetAudioOutputs) {

}

func FnGetAudioOutputsResponse (arg GetAudioOutputsResponse) {

}

func FnCreateProfile (arg CreateProfile) {

}

func FnCreateProfileResponse (arg CreateProfileResponse) {

}

func FnGetProfile (arg GetProfile) {

}

func FnGetProfileResponse (arg GetProfileResponse) {

}

func FnGetProfiles (arg GetProfiles) {

}

func FnGetProfilesResponse (arg GetProfilesResponse) {

}

func FnAddVideoEncoderConfiguration (arg AddVideoEncoderConfiguration) {

}

func FnAddVideoEncoderConfigurationResponse (arg AddVideoEncoderConfigurationResponse) {

}

func FnRemoveVideoEncoderConfiguration (arg RemoveVideoEncoderConfiguration) {

}

func FnRemoveVideoEncoderConfigurationResponse (arg RemoveVideoEncoderConfigurationResponse) {

}

func FnAddVideoSourceConfiguration (arg AddVideoSourceConfiguration) {

}

func FnAddVideoSourceConfigurationResponse (arg AddVideoSourceConfigurationResponse) {

}

func FnRemoveVideoSourceConfiguration (arg RemoveVideoSourceConfiguration) {

}

func FnRemoveVideoSourceConfigurationResponse (arg RemoveVideoSourceConfigurationResponse) {

}

func FnAddAudioEncoderConfiguration (arg AddAudioEncoderConfiguration) {

}

func FnAddAudioEncoderConfigurationResponse (arg AddAudioEncoderConfigurationResponse) {

}

func FnRemoveAudioEncoderConfiguration (arg RemoveAudioEncoderConfiguration) {

}

func FnRemoveAudioEncoderConfigurationResponse (arg RemoveAudioEncoderConfigurationResponse) {

}

func FnAddAudioSourceConfiguration (arg AddAudioSourceConfiguration) {

}

func FnAddAudioSourceConfigurationResponse (arg AddAudioSourceConfigurationResponse) {

}

func FnRemoveAudioSourceConfiguration (arg RemoveAudioSourceConfiguration) {

}

func FnRemoveAudioSourceConfigurationResponse (arg RemoveAudioSourceConfigurationResponse) {

}

func FnAddPTZConfiguration (arg AddPTZConfiguration) {

}

func FnAddPTZConfigurationResponse (arg AddPTZConfigurationResponse) {

}

func FnRemovePTZConfiguration (arg RemovePTZConfiguration) {

}

func FnRemovePTZConfigurationResponse (arg RemovePTZConfigurationResponse) {

}

func FnAddVideoAnalyticsConfiguration (arg AddVideoAnalyticsConfiguration) {

}

func FnAddVideoAnalyticsConfigurationResponse (arg AddVideoAnalyticsConfigurationResponse) {

}

func FnRemoveVideoAnalyticsConfiguration (arg RemoveVideoAnalyticsConfiguration) {

}

func FnRemoveVideoAnalyticsConfigurationResponse (arg RemoveVideoAnalyticsConfigurationResponse) {

}

func FnAddMetadataConfiguration (arg AddMetadataConfiguration) {

}

func FnAddMetadataConfigurationResponse (arg AddMetadataConfigurationResponse) {

}

func FnRemoveMetadataConfiguration (arg RemoveMetadataConfiguration) {

}

func FnRemoveMetadataConfigurationResponse (arg RemoveMetadataConfigurationResponse) {

}

func FnAddAudioOutputConfiguration (arg AddAudioOutputConfiguration) {

}

func FnAddAudioOutputConfigurationResponse (arg AddAudioOutputConfigurationResponse) {

}

func FnRemoveAudioOutputConfiguration (arg RemoveAudioOutputConfiguration) {

}

func FnRemoveAudioOutputConfigurationResponse (arg RemoveAudioOutputConfigurationResponse) {

}

func FnAddAudioDecoderConfiguration (arg AddAudioDecoderConfiguration) {

}

func FnAddAudioDecoderConfigurationResponse (arg AddAudioDecoderConfigurationResponse) {

}

func FnRemoveAudioDecoderConfiguration (arg RemoveAudioDecoderConfiguration) {

}

func FnRemoveAudioDecoderConfigurationResponse (arg RemoveAudioDecoderConfigurationResponse) {

}

func FnDeleteProfile (arg DeleteProfile) {

}

func FnDeleteProfileResponse (arg DeleteProfileResponse) {

}

func FnGetVideoSourceConfigurations (arg GetVideoSourceConfigurations) {

}

func FnGetVideoSourceConfigurationsResponse (arg GetVideoSourceConfigurationsResponse) {

}

func FnGetVideoEncoderConfigurations (arg GetVideoEncoderConfigurations) {

}

func FnGetVideoEncoderConfigurationsResponse (arg GetVideoEncoderConfigurationsResponse) {

}

func FnGetAudioSourceConfigurations (arg GetAudioSourceConfigurations) {

}

func FnGetAudioSourceConfigurationsResponse (arg GetAudioSourceConfigurationsResponse) {

}

func FnGetAudioEncoderConfigurations (arg GetAudioEncoderConfigurations) {

}

func FnGetAudioEncoderConfigurationsResponse (arg GetAudioEncoderConfigurationsResponse) {

}

func FnGetVideoAnalyticsConfigurations (arg GetVideoAnalyticsConfigurations) {

}

func FnGetVideoAnalyticsConfigurationsResponse (arg GetVideoAnalyticsConfigurationsResponse) {

}

func FnGetMetadataConfigurations (arg GetMetadataConfigurations) {

}

func FnGetMetadataConfigurationsResponse (arg GetMetadataConfigurationsResponse) {

}

func FnGetAudioOutputConfigurations (arg GetAudioOutputConfigurations) {

}

func FnGetAudioOutputConfigurationsResponse (arg GetAudioOutputConfigurationsResponse) {

}

func FnGetAudioDecoderConfigurations (arg GetAudioDecoderConfigurations) {

}

func FnGetAudioDecoderConfigurationsResponse (arg GetAudioDecoderConfigurationsResponse) {

}

func FnGetVideoSourceConfiguration (arg GetVideoSourceConfiguration) {

}

func FnGetVideoSourceConfigurationResponse (arg GetVideoSourceConfigurationResponse) {

}

func FnGetVideoEncoderConfiguration (arg GetVideoEncoderConfiguration) {

}

func FnGetVideoEncoderConfigurationResponse (arg GetVideoEncoderConfigurationResponse) {

}

func FnGetAudioSourceConfiguration (arg GetAudioSourceConfiguration) {

}

func FnGetAudioSourceConfigurationResponse (arg GetAudioSourceConfigurationResponse) {

}

func FnGetAudioEncoderConfiguration (arg GetAudioEncoderConfiguration) {

}

func FnGetAudioEncoderConfigurationResponse (arg GetAudioEncoderConfigurationResponse) {

}

func FnGetVideoAnalyticsConfiguration (arg GetVideoAnalyticsConfiguration) {

}

func FnGetVideoAnalyticsConfigurationResponse (arg GetVideoAnalyticsConfigurationResponse) {

}

func FnGetMetadataConfiguration (arg GetMetadataConfiguration) {

}

func FnGetMetadataConfigurationResponse (arg GetMetadataConfigurationResponse) {

}

func FnGetAudioOutputConfiguration (arg GetAudioOutputConfiguration) {

}

func FnGetAudioOutputConfigurationResponse (arg GetAudioOutputConfigurationResponse) {

}

func FnGetAudioDecoderConfiguration (arg GetAudioDecoderConfiguration) {

}

func FnGetAudioDecoderConfigurationResponse (arg GetAudioDecoderConfigurationResponse) {

}

func FnGetCompatibleVideoEncoderConfigurations (arg GetCompatibleVideoEncoderConfigurations) {

}

func FnGetCompatibleVideoEncoderConfigurationsResponse (arg GetCompatibleVideoEncoderConfigurationsResponse) {

}

func FnGetCompatibleVideoSourceConfigurations (arg GetCompatibleVideoSourceConfigurations) {

}

func FnGetCompatibleVideoSourceConfigurationsResponse (arg GetCompatibleVideoSourceConfigurationsResponse) {

}

func FnGetCompatibleAudioEncoderConfigurations (arg GetCompatibleAudioEncoderConfigurations) {

}

func FnGetCompatibleAudioEncoderConfigurationsResponse (arg GetCompatibleAudioEncoderConfigurationsResponse) {

}

func FnGetCompatibleAudioSourceConfigurations (arg GetCompatibleAudioSourceConfigurations) {

}

func FnGetCompatibleAudioSourceConfigurationsResponse (arg GetCompatibleAudioSourceConfigurationsResponse) {

}

func FnGetCompatibleVideoAnalyticsConfigurations (arg GetCompatibleVideoAnalyticsConfigurations) {

}

func FnGetCompatibleVideoAnalyticsConfigurationsResponse (arg GetCompatibleVideoAnalyticsConfigurationsResponse) {

}

func FnGetCompatibleMetadataConfigurations (arg GetCompatibleMetadataConfigurations) {

}

func FnGetCompatibleMetadataConfigurationsResponse (arg GetCompatibleMetadataConfigurationsResponse) {

}

func FnGetCompatibleAudioOutputConfigurations (arg GetCompatibleAudioOutputConfigurations) {

}

func FnGetCompatibleAudioOutputConfigurationsResponse (arg GetCompatibleAudioOutputConfigurationsResponse) {

}

func FnGetCompatibleAudioDecoderConfigurations (arg GetCompatibleAudioDecoderConfigurations) {

}

func FnGetCompatibleAudioDecoderConfigurationsResponse (arg GetCompatibleAudioDecoderConfigurationsResponse) {

}

func FnSetVideoSourceConfiguration (arg SetVideoSourceConfiguration) {

}

func FnSetVideoSourceConfigurationResponse (arg SetVideoSourceConfigurationResponse) {

}

func FnSetVideoEncoderConfiguration (arg SetVideoEncoderConfiguration) {

}

func FnSetVideoEncoderConfigurationResponse (arg SetVideoEncoderConfigurationResponse) {

}

func FnSetAudioSourceConfiguration (arg SetAudioSourceConfiguration) {

}

func FnSetAudioSourceConfigurationResponse (arg SetAudioSourceConfigurationResponse) {

}

func FnSetAudioEncoderConfiguration (arg SetAudioEncoderConfiguration) {

}

func FnSetAudioEncoderConfigurationResponse (arg SetAudioEncoderConfigurationResponse) {

}

func FnSetVideoAnalyticsConfiguration (arg SetVideoAnalyticsConfiguration) {

}

func FnSetVideoAnalyticsConfigurationResponse (arg SetVideoAnalyticsConfigurationResponse) {

}

func FnSetMetadataConfiguration (arg SetMetadataConfiguration) {

}

func FnSetMetadataConfigurationResponse (arg SetMetadataConfigurationResponse) {

}

func FnSetAudioOutputConfiguration (arg SetAudioOutputConfiguration) {

}

func FnSetAudioOutputConfigurationResponse (arg SetAudioOutputConfigurationResponse) {

}

func FnSetAudioDecoderConfiguration (arg SetAudioDecoderConfiguration) {

}

func FnSetAudioDecoderConfigurationResponse (arg SetAudioDecoderConfigurationResponse) {

}

func FnGetVideoSourceConfigurationOptions (arg GetVideoSourceConfigurationOptions) {

}

func FnGetVideoSourceConfigurationOptionsResponse (arg GetVideoSourceConfigurationOptionsResponse) {

}

func FnGetVideoEncoderConfigurationOptions (arg GetVideoEncoderConfigurationOptions) {

}

func FnGetVideoEncoderConfigurationOptionsResponse (arg GetVideoEncoderConfigurationOptionsResponse) {

}

func FnGetAudioSourceConfigurationOptions (arg GetAudioSourceConfigurationOptions) {

}

func FnGetAudioSourceConfigurationOptionsResponse (arg GetAudioSourceConfigurationOptionsResponse) {

}

func FnGetAudioEncoderConfigurationOptions (arg GetAudioEncoderConfigurationOptions) {

}

func FnGetAudioEncoderConfigurationOptionsResponse (arg GetAudioEncoderConfigurationOptionsResponse) {

}

func FnGetMetadataConfigurationOptions (arg GetMetadataConfigurationOptions) {

}

func FnGetMetadataConfigurationOptionsResponse (arg GetMetadataConfigurationOptionsResponse) {

}

func FnGetAudioOutputConfigurationOptions (arg GetAudioOutputConfigurationOptions) {

}

func FnGetAudioOutputConfigurationOptionsResponse (arg GetAudioOutputConfigurationOptionsResponse) {

}

func FnGetAudioDecoderConfigurationOptions (arg GetAudioDecoderConfigurationOptions) {

}

func FnGetAudioDecoderConfigurationOptionsResponse (arg GetAudioDecoderConfigurationOptionsResponse) {

}

func FnGetGuaranteedNumberOfVideoEncoderInstances (arg GetGuaranteedNumberOfVideoEncoderInstances) {

}

func FnGetGuaranteedNumberOfVideoEncoderInstancesResponse (arg GetGuaranteedNumberOfVideoEncoderInstancesResponse) {

}

func FnGetStreamUri (arg GetStreamUri) {

}

func FnGetStreamUriResponse (arg GetStreamUriResponse) {

}

func FnStartMulticastStreaming (arg StartMulticastStreaming) {

}

func FnStartMulticastStreamingResponse (arg StartMulticastStreamingResponse) {

}

func FnStopMulticastStreaming (arg StopMulticastStreaming) {

}

func FnStopMulticastStreamingResponse (arg StopMulticastStreamingResponse) {

}

func FnSetSynchronizationPoint (arg SetSynchronizationPoint) {

}

func FnSetSynchronizationPointResponse (arg SetSynchronizationPointResponse) {

}

func FnGetSnapshotUri (arg GetSnapshotUri) {

}

func FnGetSnapshotUriResponse (arg GetSnapshotUriResponse) {

}

func FnGetVideoSourceModes (arg GetVideoSourceModes) {

}

func FnGetVideoSourceModesResponse (arg GetVideoSourceModesResponse) {

}

func FnSetVideoSourceMode (arg SetVideoSourceMode) {

}

func FnSetVideoSourceModeResponse (arg SetVideoSourceModeResponse) {

}

func FnGetOSDs (arg GetOSDs) {

}

func FnGetOSDsResponse (arg GetOSDsResponse) {

}

func FnGetOSD (arg GetOSD) {

}

func FnGetOSDResponse (arg GetOSDResponse) {

}

func FnGetOSDOptions (arg GetOSDOptions) {

}

func FnGetOSDOptionsResponse (arg GetOSDOptionsResponse) {

}

func FnSetOSD (arg SetOSD) {

}

func FnSetOSDResponse (arg SetOSDResponse) {

}

func FnCreateOSD (arg CreateOSD) {

}

func FnCreateOSDResponse (arg CreateOSDResponse) {

}

func FnDeleteOSD (arg DeleteOSD) {

}

func FnDeleteOSDResponse (arg DeleteOSDResponse) {

}