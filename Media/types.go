package Media

import "github.com/yakovlevdmv/goonvif/xsd/onvif"

type Capabilities struct {
	SnapshotUri 		bool `xml:"SnapshotUri,attr"`
	Rotation 			bool `xml:"Rotation,attr"`
	VideoSourceMode 	bool `xml:"VideoSourceMode,attr"`
	OSD 				bool `xml:"OSD,attr"`
	TemporaryOSDText 	bool `xml:"TemporaryOSDText,attr"`
	EXICompression 		bool `xml:"EXICompression,attr"`
	ProfileCapabilities ProfileCapabilities
	StreamingCapabilities StreamingCapabilities
}

type ProfileCapabilities struct {
	MaximumNumberOfProfiles int `xml:"MaximumNumberOfProfiles,attr"`
}

type StreamingCapabilities struct {
	RTPMulticast 		bool `xml:"RTPMulticast,attr"`
	RTP_TCP 			bool `xml:"RTP_TCP,attr"`
	RTP_RTSP_TCP 		bool `xml:"RTP_RTSP_TCP,attr"`
	NonAggregateControl bool `xml:"NonAggregateControl,attr"`
	NoRTSPStreaming 	bool `xml:"NoRTSPStreaming,attr"`
}

//Media main types

type GetServiceCapabilities struct {

}


type GetServiceCapabilitiesResponse struct {
	Capabilities Capabilities

}


type GetVideoSources struct {

}


type GetVideoSourcesResponse struct {
	VideoSources onvif.VideoSource

}


type GetAudioSources struct {

}


type GetAudioSourcesResponse struct {
	AudioSources onvif.AudioSource

}


type GetAudioOutputs struct {

}


type GetAudioOutputsResponse struct {
	AudioOutputs onvif.AudioOutput

}


type CreateProfile struct {
	Name onvif.Name
	Token onvif.ReferenceToken

}


type CreateProfileResponse struct {
	Profile onvif.Profile

}


type GetProfile struct {
	ProfileToken onvif.ReferenceToken

}


type GetProfileResponse struct {
	Profile onvif.Profile

}


type GetProfiles struct {

}


type GetProfilesResponse struct {
	Profiles onvif.Profile

}


type AddVideoEncoderConfiguration struct {
	ProfileToken onvif.ReferenceToken
	ConfigurationToken onvif.ReferenceToken

}


type AddVideoEncoderConfigurationResponse struct {

}


type RemoveVideoEncoderConfiguration struct {
	ProfileToken onvif.ReferenceToken

}


type RemoveVideoEncoderConfigurationResponse struct {

}


type AddVideoSourceConfiguration struct {
	ProfileToken onvif.ReferenceToken
	ConfigurationToken onvif.ReferenceToken

}


type AddVideoSourceConfigurationResponse struct {

}


type RemoveVideoSourceConfiguration struct {
	ProfileToken onvif.ReferenceToken

}


type RemoveVideoSourceConfigurationResponse struct {

}


type AddAudioEncoderConfiguration struct {
	ProfileToken onvif.ReferenceToken
	ConfigurationToken onvif.ReferenceToken

}


type AddAudioEncoderConfigurationResponse struct {

}


type RemoveAudioEncoderConfiguration struct {
	ProfileToken onvif.ReferenceToken

}


type RemoveAudioEncoderConfigurationResponse struct {

}


type AddAudioSourceConfiguration struct {
	ProfileToken onvif.ReferenceToken
	ConfigurationToken onvif.ReferenceToken

}


type AddAudioSourceConfigurationResponse struct {

}


type RemoveAudioSourceConfiguration struct {
	ProfileToken onvif.ReferenceToken

}


type RemoveAudioSourceConfigurationResponse struct {

}


type AddPTZConfiguration struct {
	ProfileToken onvif.ReferenceToken
	ConfigurationToken onvif.ReferenceToken

}


type AddPTZConfigurationResponse struct {

}


type RemovePTZConfiguration struct {
	ProfileToken onvif.ReferenceToken

}


type RemovePTZConfigurationResponse struct {

}


type AddVideoAnalyticsConfiguration struct {
	ProfileToken onvif.ReferenceToken
	ConfigurationToken onvif.ReferenceToken

}


type AddVideoAnalyticsConfigurationResponse struct {

}


type RemoveVideoAnalyticsConfiguration struct {
	ProfileToken onvif.ReferenceToken

}


type RemoveVideoAnalyticsConfigurationResponse struct {

}


type AddMetadataConfiguration struct {
	ProfileToken onvif.ReferenceToken
	ConfigurationToken onvif.ReferenceToken

}


type AddMetadataConfigurationResponse struct {

}


type RemoveMetadataConfiguration struct {
	ProfileToken onvif.ReferenceToken

}


type RemoveMetadataConfigurationResponse struct {

}


type AddAudioOutputConfiguration struct {
	ProfileToken onvif.ReferenceToken
	ConfigurationToken onvif.ReferenceToken

}


type AddAudioOutputConfigurationResponse struct {

}


type RemoveAudioOutputConfiguration struct {
	ProfileToken onvif.ReferenceToken

}


type RemoveAudioOutputConfigurationResponse struct {

}


type AddAudioDecoderConfiguration struct {
	ProfileToken onvif.ReferenceToken
	ConfigurationToken onvif.ReferenceToken

}


type AddAudioDecoderConfigurationResponse struct {

}


type RemoveAudioDecoderConfiguration struct {
	ProfileToken onvif.ReferenceToken

}


type RemoveAudioDecoderConfigurationResponse struct {

}


type DeleteProfile struct {
	ProfileToken onvif.ReferenceToken

}


type DeleteProfileResponse struct {

}


type GetVideoSourceConfigurations struct {

}


type GetVideoSourceConfigurationsResponse struct {
	Configurations onvif.VideoSourceConfiguration

}


type GetVideoEncoderConfigurations struct {

}


type GetVideoEncoderConfigurationsResponse struct {
	Configurations onvif.VideoEncoderConfiguration

}


type GetAudioSourceConfigurations struct {

}


type GetAudioSourceConfigurationsResponse struct {
	Configurations onvif.AudioSourceConfiguration

}


type GetAudioEncoderConfigurations struct {

}


type GetAudioEncoderConfigurationsResponse struct {
	Configurations onvif.AudioEncoderConfiguration

}


type GetVideoAnalyticsConfigurations struct {

}


type GetVideoAnalyticsConfigurationsResponse struct {
	Configurations onvif.VideoAnalyticsConfiguration

}


type GetMetadataConfigurations struct {

}


type GetMetadataConfigurationsResponse struct {
	Configurations onvif.MetadataConfiguration

}


type GetAudioOutputConfigurations struct {

}


type GetAudioOutputConfigurationsResponse struct {
	Configurations onvif.AudioOutputConfiguration

}


type GetAudioDecoderConfigurations struct {

}


type GetAudioDecoderConfigurationsResponse struct {
	Configurations onvif.AudioDecoderConfiguration

}


type GetVideoSourceConfiguration struct {
	ConfigurationToken onvif.ReferenceToken

}


type GetVideoSourceConfigurationResponse struct {
	Configuration onvif.VideoSourceConfiguration

}


type GetVideoEncoderConfiguration struct {
	ConfigurationToken onvif.ReferenceToken

}


type GetVideoEncoderConfigurationResponse struct {
	Configuration onvif.VideoEncoderConfiguration

}


type GetAudioSourceConfiguration struct {
	ConfigurationToken onvif.ReferenceToken

}


type GetAudioSourceConfigurationResponse struct {
	Configuration onvif.AudioSourceConfiguration

}


type GetAudioEncoderConfiguration struct {
	ConfigurationToken onvif.ReferenceToken

}


type GetAudioEncoderConfigurationResponse struct {
	Configuration onvif.AudioEncoderConfiguration

}


type GetVideoAnalyticsConfiguration struct {
	ConfigurationToken onvif.ReferenceToken

}


type GetVideoAnalyticsConfigurationResponse struct {
	Configuration onvif.VideoAnalyticsConfiguration

}


type GetMetadataConfiguration struct {
	ConfigurationToken onvif.ReferenceToken

}


type GetMetadataConfigurationResponse struct {
	Configuration onvif.MetadataConfiguration

}


type GetAudioOutputConfiguration struct {
	ConfigurationToken onvif.ReferenceToken

}


type GetAudioOutputConfigurationResponse struct {
	Configuration onvif.AudioOutputConfiguration

}


type GetAudioDecoderConfiguration struct {
	ConfigurationToken onvif.ReferenceToken

}


type GetAudioDecoderConfigurationResponse struct {
	Configuration onvif.AudioDecoderConfiguration

}


type GetCompatibleVideoEncoderConfigurations struct {
	ProfileToken onvif.ReferenceToken

}


type GetCompatibleVideoEncoderConfigurationsResponse struct {
	Configurations onvif.VideoEncoderConfiguration

}


type GetCompatibleVideoSourceConfigurations struct {
	ProfileToken onvif.ReferenceToken

}


type GetCompatibleVideoSourceConfigurationsResponse struct {
	Configurations onvif.VideoSourceConfiguration

}


type GetCompatibleAudioEncoderConfigurations struct {
	ProfileToken onvif.ReferenceToken

}


type GetCompatibleAudioEncoderConfigurationsResponse struct {
	Configurations onvif.AudioEncoderConfiguration

}


type GetCompatibleAudioSourceConfigurations struct {
	ProfileToken onvif.ReferenceToken

}


type GetCompatibleAudioSourceConfigurationsResponse struct {
	Configurations onvif.AudioSourceConfiguration

}


type GetCompatibleVideoAnalyticsConfigurations struct {
	ProfileToken onvif.ReferenceToken

}


type GetCompatibleVideoAnalyticsConfigurationsResponse struct {
	Configurations onvif.VideoAnalyticsConfiguration

}


type GetCompatibleMetadataConfigurations struct {
	ProfileToken onvif.ReferenceToken

}


type GetCompatibleMetadataConfigurationsResponse struct {
	Configurations onvif.MetadataConfiguration

}


type GetCompatibleAudioOutputConfigurations struct {
	ProfileToken onvif.ReferenceToken

}


type GetCompatibleAudioOutputConfigurationsResponse struct {
	Configurations onvif.AudioOutputConfiguration

}


type GetCompatibleAudioDecoderConfigurations struct {
	ProfileToken onvif.ReferenceToken

}


type GetCompatibleAudioDecoderConfigurationsResponse struct {
	Configurations onvif.AudioDecoderConfiguration

}


type SetVideoSourceConfiguration struct {
	Configuration onvif.VideoSourceConfiguration
	ForcePersistence bool

}


type SetVideoSourceConfigurationResponse struct {

}


type SetVideoEncoderConfiguration struct {
	Configuration onvif.VideoEncoderConfiguration
	ForcePersistence bool

}


type SetVideoEncoderConfigurationResponse struct {

}


type SetAudioSourceConfiguration struct {
	Configuration onvif.AudioSourceConfiguration
	ForcePersistence bool

}


type SetAudioSourceConfigurationResponse struct {

}


type SetAudioEncoderConfiguration struct {
	Configuration onvif.AudioEncoderConfiguration
	ForcePersistence bool

}


type SetAudioEncoderConfigurationResponse struct {

}


type SetVideoAnalyticsConfiguration struct {
	Configuration onvif.VideoAnalyticsConfiguration
	ForcePersistence bool

}


type SetVideoAnalyticsConfigurationResponse struct {

}


type SetMetadataConfiguration struct {
	Configuration onvif.MetadataConfiguration
	ForcePersistence bool

}


type SetMetadataConfigurationResponse struct {

}


type SetAudioOutputConfiguration struct {
	Configuration onvif.AudioOutputConfiguration
	ForcePersistence bool

}


type SetAudioOutputConfigurationResponse struct {

}


type SetAudioDecoderConfiguration struct {
	Configuration onvif.AudioDecoderConfiguration
	ForcePersistence bool

}


type SetAudioDecoderConfigurationResponse struct {

}


type GetVideoSourceConfigurationOptions struct {
	ConfigurationToken onvif.ReferenceToken
	ProfileToken onvif.ReferenceToken

}


type GetVideoSourceConfigurationOptionsResponse struct {
	Options onvif.VideoSourceConfigurationOptions

}


type GetVideoEncoderConfigurationOptions struct {
	ConfigurationToken onvif.ReferenceToken
	ProfileToken onvif.ReferenceToken

}


type GetVideoEncoderConfigurationOptionsResponse struct {
	Options onvif.VideoEncoderConfigurationOptions

}


type GetAudioSourceConfigurationOptions struct {
	ConfigurationToken onvif.ReferenceToken
	ProfileToken onvif.ReferenceToken

}


type GetAudioSourceConfigurationOptionsResponse struct {
	Options onvif.AudioSourceConfigurationOptions

}


type GetAudioEncoderConfigurationOptions struct {
	ConfigurationToken onvif.ReferenceToken
	ProfileToken onvif.ReferenceToken

}


type GetAudioEncoderConfigurationOptionsResponse struct {
	Options onvif.AudioEncoderConfigurationOptions

}


type GetMetadataConfigurationOptions struct {
	ConfigurationToken onvif.ReferenceToken
	ProfileToken onvif.ReferenceToken

}


type GetMetadataConfigurationOptionsResponse struct {
	Options onvif.MetadataConfigurationOptions

}


type GetAudioOutputConfigurationOptions struct {
	ConfigurationToken onvif.ReferenceToken
	ProfileToken onvif.ReferenceToken

}


type GetAudioOutputConfigurationOptionsResponse struct {
	Options onvif.AudioOutputConfigurationOptions

}


type GetAudioDecoderConfigurationOptions struct {
	ConfigurationToken onvif.ReferenceToken
	ProfileToken onvif.ReferenceToken

}


type GetAudioDecoderConfigurationOptionsResponse struct {
	Options onvif.AudioDecoderConfigurationOptions

}


type GetGuaranteedNumberOfVideoEncoderInstances struct {
	ConfigurationToken onvif.ReferenceToken

}


type GetGuaranteedNumberOfVideoEncoderInstancesResponse struct {
	TotalNumber int
	JPEG int
	H264 int
	MPEG4 int

}


type GetStreamUri struct {
	StreamSetup onvif.StreamSetup
	ProfileToken onvif.ReferenceToken

}


type GetStreamUriResponse struct {
	MediaUri onvif.MediaUri

}


type StartMulticastStreaming struct {
	ProfileToken onvif.ReferenceToken

}


type StartMulticastStreamingResponse struct {

}


type StopMulticastStreaming struct {
	ProfileToken onvif.ReferenceToken

}


type StopMulticastStreamingResponse struct {

}


type SetSynchronizationPoint struct {
	ProfileToken onvif.ReferenceToken

}


type SetSynchronizationPointResponse struct {

}


type GetSnapshotUri struct {
	ProfileToken onvif.ReferenceToken

}


type GetSnapshotUriResponse struct {
	MediaUri onvif.MediaUri

}


type GetVideoSourceModes struct {
	VideoSourceToken onvif.ReferenceToken

}


type GetVideoSourceModesResponse struct {
	VideoSourceModes onvif.VideoSourceMode

}


type SetVideoSourceMode struct {
	VideoSourceToken onvif.ReferenceToken
	VideoSourceModeToken onvif.ReferenceToken

}


type SetVideoSourceModeResponse struct {
	Reboot bool

}


type GetOSDs struct {
	ConfigurationToken onvif.ReferenceToken

}


type GetOSDsResponse struct {
	OSDs onvif.OSDConfiguration

}


type GetOSD struct {
	OSDToken onvif.ReferenceToken

}


type GetOSDResponse struct {
	OSD onvif.OSDConfiguration

}


type GetOSDOptions struct {
	ConfigurationToken onvif.ReferenceToken

}


type GetOSDOptionsResponse struct {
	OSDOptions onvif.OSDConfigurationOptions

}


type SetOSD struct {
	OSD onvif.OSDConfiguration

}


type SetOSDResponse struct {

}


type CreateOSD struct {
	OSD onvif.OSDConfiguration

}


type CreateOSDResponse struct {
	OSDToken onvif.ReferenceToken

}


type DeleteOSD struct {
	OSDToken onvif.ReferenceToken

}


type DeleteOSDResponse struct {

}