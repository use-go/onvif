package onvif

import "github.com/yakovlevdmv/goonvif/xsd"

//TODO: rename package xsdTypes
//TODO: enumerations
//TODO: type <typeName> struct {Any string} convert to type <typeName> AnyType
//TODO: process restrictions

//todo посмотреть все Extensions (Any string)
//todo что делать с xs:any = Any
//todo IntList и ему подобные. Проверить нужен ли слайс. Изменить на slice
//todo посмотреть можно ли заменить StreamType и ему подобные типы на вмтроенные типы
//todo оттестировать тип VideoSourceMode из-за Description-а

//todo в документации описать, что Capabilities повторяеся у каждого сервиса, поэтому у каждого свой Capabilities (MediaCapabilities)
//todo AuxiliaryData и другие simpleTypes, как реализовать рестрикшн
//todo Name и ему подобные необходимо изучить на наличие "List of..." ошибок


//todo Add in buit in
type ContentType string // minLength value="3"
type DNSName xsd.Token

type DeviceEntity struct {
	Token ReferenceToken `xml:"token,attr"`
}

type ReferenceToken struct {
	Token string
}

type Name struct {
	Name string
}

type IntRectangle struct {
	X int `xml:"x,attr"`
	Y int `xml:"y,attr"`
	Width int `xml:"width,attr"`
	Height int `xml:"height,attr"`
}

type IntRectangleRange struct {
	XRange IntRange
	YRange IntRange
	WidthRange IntRange
	HeightRange IntRange
}

type IntRange struct {
	Min int
	Max int
}

type FloatRange struct {
	Min float64
	Max float64
}

type OSDConfiguration struct {
	DeviceEntity
	VideoSourceConfigurationToken OSDReference
	Type OSDType
	Position OSDPosConfiguration
	TextString OSDTextConfiguration
	Image OSDImgConfiguration
	Extension OSDConfigurationExtension
}

type OSDType struct {
	Type string
}

type OSDPosConfiguration struct {
	Type string
	Pos Vector
	Extension OSDPosConfigurationExtension
}

type Vector struct {
	X float64 `xml:"x,attr"`
	Y float64 `xml:"y,attr"`
}

type OSDPosConfigurationExtension struct {
	Any string //while
}

type OSDReference struct {
	ReferenceToken
}

type OSDTextConfiguration struct {
	Type string
	DateFormat string
	TimeFormat string
	FontSize int
	FontColor OSDColor
	BackgroundColor OSDColor
	PlainText string
	Extension OSDTextConfigurationExtension
}

type OSDColor struct {
	Color Color
	Transparent int
}

type Color struct {
	X          float64    `xml:"X,attr"`
	Y          float64    `xml:"Y,attr"`
	Z          float64    `xml:"Z,attr"`
	Colorspace xsd.AnyURI `xml:"Colorspace,attr"`
}


type OSDTextConfigurationExtension struct {
	Any string
}

type OSDImgConfiguration struct {
	ImgPath   xsd.AnyURI
	Extension OSDImgConfigurationExtension
}

type OSDImgConfigurationExtension struct {
	Any string
}

type OSDConfigurationExtension struct {
	Any string //While
}

type VideoSource struct {
	DeviceEntity
	Framerate 	float64
	Resolution 	VideoResolution
	Imaging 	ImagingSettings
	Extension 	VideoSourceExtension
}

type VideoResolution struct {
	Width 	int
	Height 	int
}

type ImagingSettings struct {
	BacklightCompensation 	BacklightCompensation
	Brightness 				float64
	ColorSaturation 		float64
	Contrast 				float64
	Exposure 				Exposure
	Focus 					FocusConfiguration
	IrCutFilter 			IrCutFilterMode
	Sharpness 				float64
	WideDynamicRange 		WideDynamicRange
	WhiteBalance 			WhiteBalance
	Extension 				ImagingSettingsExtension
}

type BacklightCompensation struct {
	Mode 	BacklightCompensationMode
	Level 	float64
}

type BacklightCompensationMode struct {
	BacklightCompensation string
}

type Exposure struct {
	Mode 			ExposureMode
	Priority 		ExposurePriority
	Window 			Rectangle
	MinExposureTime float64
	MaxExposureTime float64
	MinGain 		float64
	MaxGain 		float64
	MinIris 		float64
	MaxIris 		float64
	ExposureTime 	float64
	Gain 			float64
	Iris 			float64
}

type ExposureMode struct {
	Mode string
}

type ExposurePriority struct {
	Priority string
}

type Rectangle struct {
	Bottom 	float64 `xml:"bottom,attr"`
	Top 	float64 `xml:"top,attr"`
	Right 	float64 `xml:"right,attr"`
	Left 	float64 `xml:"left,attr"`
}

type FocusConfiguration struct {
	AutoFocusMode AutoFocusMode
	DefaultSpeed float64
	NearLimit float64
	FarLimit float64
}

type AutoFocusMode struct {
	FocusMode string
}

type IrCutFilterMode struct {
	FilterMode string
}

type WideDynamicRange struct {
	Mode 	WideDynamicMode
	Level 	float64
}

type WideDynamicMode struct {
	DynamicMode string
}

type WhiteBalance struct {
	Mode WhiteBalanceMode
	CrGain float64
	CbGain float64
}

type WhiteBalanceMode struct {
	BalanceMode string
}

type ImagingSettingsExtension struct {
	Any string
}

type VideoSourceExtension struct {
	Imaging ImagingSettings20
	Extension VideoSourceExtension2
}

type ImagingSettings20 struct {
	BacklightCompensation BacklightCompensation20
	Brightness float64
	ColorSaturation float64
	Contrast float64
	Exposure Exposure20
	Focus FocusConfiguration20
	IrCutFilter IrCutFilterMode
	Sharpness float64
	WideDynamicRange WideDynamicRange20
	WhiteBalance WhiteBalance20
	Extension ImagingSettingsExtension20
}

type BacklightCompensation20 struct {
	Mode BacklightCompensationMode
	Level float64
}

type Exposure20 struct {
	Mode ExposureMode
	Priority ExposurePriority
	Window Rectangle
	MinExposureTime float64
	MaxExposureTime float64
	MinGain float64
	MaxGain float64
	MinIris float64
	MaxIris float64
	ExposureTime float64
	Gain float64
	Iris float64
}

type FocusConfiguration20 struct {
	AutoFocusMode AutoFocusMode
	DefaultSpeed float64
	NearLimit float64
	FarLimit float64
	Extension FocusConfiguration20Extension
}

type FocusConfiguration20Extension struct {
	Any string
}

type WideDynamicRange20 struct {
	Mode WideDynamicMode
	Level float64
}

type WhiteBalance20 struct {
	Mode WhiteBalanceMode
	CrGain float64
	CbGain float64
	Extension WhiteBalance20Extension
}

type WhiteBalance20Extension struct {
	Any string
}

type ImagingSettingsExtension20 struct {
	ImageStabilization ImageStabilization
	Extension ImagingSettingsExtension202
}

type ImageStabilization struct {
	Mode ImageStabilizationMode
	Level float64
	Extension ImageStabilizationExtension
}

type ImageStabilizationMode struct {
	StabilizationMode string
}

type ImageStabilizationExtension struct {
	Any string
}

type ImagingSettingsExtension202 struct {
	IrCutFilterAutoAdjustment IrCutFilterAutoAdjustment
	Extension ImagingSettingsExtension203
}

type IrCutFilterAutoAdjustment struct {
	BoundaryType   string
	BoundaryOffset float64
	ResponseTime   xsd.Duration
	Extension      IrCutFilterAutoAdjustmentExtension
}

type IrCutFilterAutoAdjustmentExtension struct {
	Any string
}

type ImagingSettingsExtension203 struct {
	ToneCompensation ToneCompensation
	Defogging Defogging
	NoiseReduction NoiseReduction
	Extension ImagingSettingsExtension204
}

type ToneCompensation struct {
	Mode string
	Level float64
	Extension ToneCompensationExtension
}

type ToneCompensationExtension struct {
	Any string
}

type Defogging struct {
	Mode string
	Level float64
	Extension DefoggingExtension
}

type DefoggingExtension struct {
	Any string
}

type NoiseReduction struct {
	Level float64
}

type ImagingSettingsExtension204 struct {
	Any string
}

type VideoSourceExtension2 struct {
	Any string
}

type AudioSource struct {
	DeviceEntity
	Channels int
}

type AudioOutput struct {
	DeviceEntity
}

type Profile struct {
	Token ReferenceToken 	`xml:"token,attr"`
	Fixed bool 				`xml:"fixed,attr"`
	Name Name
	VideoSourceConfiguration VideoSourceConfiguration
	AudioSourceConfiguration AudioSourceConfiguration
	VideoEncoderConfiguration VideoEncoderConfiguration
	AudioEncoderConfiguration AudioEncoderConfiguration
	VideoAnalyticsConfiguration VideoAnalyticsConfiguration
	PTZConfiguration PTZConfiguration
	MetadataConfiguration MetadataConfiguration
	Extension ProfileExtension
}

type VideoSourceConfiguration struct {
	ConfigurationEntity
	ViewMode string `xml:"ViewMode,attr"`
	SourceToken ReferenceToken
	Bounds IntRectangle
	Extension VideoSourceConfigurationExtension
}

type ConfigurationEntity struct {
	Token ReferenceToken `xml:"token,attr"`
	Name Name
	UseCount int
}

type VideoSourceConfigurationExtension struct {
	Rotate Rotate
	Extension VideoSourceConfigurationExtension2
}

type Rotate struct {
	Mode RotateMode
	Degree int
	Extension RotateExtension
}

type RotateMode struct {
	RotateMode string
}

type RotateExtension struct {
	Any string
}

type VideoSourceConfigurationExtension2 struct {
	LensDescription LensDescription
	SceneOrientation SceneOrientation
}

type LensDescription struct {
	FocalLength float64 `xml:"FocalLength,attr"`
	Offset LensOffset
	Projection LensProjection
	XFactor float64
}

type LensOffset struct {
	X float64 `xml:"x,attr"`
	Y float64 `xml:"y,attr"`
}

type LensProjection struct {
	Angle float64
	Radius float64
	Transmittance float64
}

type SceneOrientation struct {
	Mode SceneOrientationMode
	Orientation string
}

type SceneOrientationMode struct {
	OrientationMode string
}

type AudioSourceConfiguration struct {
	ConfigurationEntity
	SourceToken ReferenceToken
}

type VideoEncoderConfiguration struct {
	ConfigurationEntity
	Encoding       VideoEncoding
	Resolution     VideoResolution
	Quality        float64
	RateControl    VideoRateControl
	MPEG4          Mpeg4Configuration
	H264           H264Configuration
	Multicast      MulticastConfiguration
	SessionTimeout xsd.Duration
}

type VideoEncoding struct {
	VideoEncoding string
}

type VideoRateControl struct {
	FrameRateLimit int
	EncodingInterval int
	BitrateLimit int
}

type Mpeg4Configuration struct {
	GovLength int
	Mpeg4Profile Mpeg4Profile
}

type Mpeg4Profile struct {
	Profile string
}

type H264Configuration struct {
	GovLength int
	H264Profile H264Profile
}

type H264Profile struct {
	Profile string
}

type MulticastConfiguration struct {
	Address IPAddress
	Port int
	TTL int
	AutoStart bool
}

type IPAddress struct {
	Type IPType
	IPv4Address IPv4Address
	IPv6Address IPv6Address
}

type IPType struct {
	Type string
}

type IPv4Address struct {
	Address xsd.Token
}

type IPv6Address struct {
	Address xsd.Token
}

type AudioEncoderConfiguration struct {
	ConfigurationEntity
	Encoding       AudioEncoding
	Bitrate        int
	SampleRate     int
	Multicast      MulticastConfiguration
	SessionTimeout xsd.Duration
}

type AudioEncoding struct {
	Encoding string
}

type VideoAnalyticsConfiguration struct {
	ConfigurationEntity
	AnalyticsEngineConfiguration AnalyticsEngineConfiguration
	RuleEngineConfiguration RuleEngineConfiguration
}

type AnalyticsEngineConfiguration struct {
	AnalyticsModule Config
	Extension AnalyticsEngineConfigurationExtension
}

type Config struct {
	Name       string    `xml:"Name,attr"`
	Type       xsd.QName `xml:"Type,attr"`
	Parameters ItemList
}

type ItemList struct {
	SimpleItem
	ElementItem
	Extension ItemListExtension
}

type SimpleItem struct {
	Name  string            `xml:"Name,attr"`
	Value xsd.AnySimpleType `xml:"Value,attr"`
}

type ElementItem struct {
	Name string `xml:"Name,attr"`
}

type ItemListExtension struct {
	Any string
}

type AnalyticsEngineConfigurationExtension struct {
	Any string
}

type RuleEngineConfiguration struct {
	Rule Config
	Extension RuleEngineConfigurationExtension
}

type RuleEngineConfigurationExtension struct {
	Any string
}

type PTZConfiguration struct {
	ConfigurationEntity
	MoveRamp                               int `xml:"MoveRamp,attr"`
	PresetRamp                             int `xml:"PresetRamp,attr"`
	PresetTourRamp                         int `xml:"PresetTourRamp,attr"`
	NodeToken                              ReferenceToken
	DefaultAbsolutePantTiltPositionSpace   xsd.AnyURI
	DefaultAbsoluteZoomPositionSpace       xsd.AnyURI
	DefaultRelativePanTiltTranslationSpace xsd.AnyURI
	DefaultRelativeZoomTranslationSpace    xsd.AnyURI
	DefaultContinuousPanTiltVelocitySpace  xsd.AnyURI
	DefaultContinuousZoomVelocitySpace     xsd.AnyURI
	DefaultPTZSpeed                        PTZSpeed
	DefaultPTZTimeout                      xsd.Duration
	PanTiltLimits                          PanTiltLimits
	ZoomLimits                             ZoomLimits
	Extension                              PTZConfigurationExtension
}

type PTZSpeed struct {
	PanTilt Vector2D
	Zoom Vector1D
}

type Vector2D struct {
	X     float64    `xml:"x,attr"`
	Y     float64    `xml:"y,attr"`
	Space xsd.AnyURI `xml:"space,attr"`
}

type Vector1D struct {
	X     float64    `xml:"x,attr"`
	Space xsd.AnyURI `xml:"space,attr"`
}

type PanTiltLimits struct {
	Range Space2DDescription
}

type Space2DDescription struct {
	URI    xsd.AnyURI
	XRange FloatRange
	YRange FloatRange
}

type ZoomLimits struct {
	Range Space1DDescription
}

type Space1DDescription struct {
	URI    xsd.AnyURI
	XRange FloatRange
}

type PTZConfigurationExtension struct {
	PTControlDirection PTControlDirection
	Extension PTZConfigurationExtension2
}

type PTControlDirection struct {
	EFlip EFlip
	Reverse Reverse
	Extension PTControlDirectionExtension
}

type EFlip struct {
	Mode EFlipMode
}

type EFlipMode struct {
	EFlipMode string
}

type Reverse struct {
	Mode ReverseMode
}

type ReverseMode struct {
	ReverseMode string
}

type PTControlDirectionExtension struct {
	Any string
}

type PTZConfigurationExtension2 struct {
	Any string
}

type MetadataConfiguration struct {
	ConfigurationEntity
	CompressionType              string `xml:"CompressionType,attr"`
	PTZStatus                    PTZFilter
	Events                       EventSubscription
	Analytics                    bool
	Multicast                    MulticastConfiguration
	SessionTimeout               xsd.Duration
	AnalyticsEngineConfiguration AnalyticsEngineConfiguration
	Extension                    MetadataConfigurationExtension
}

type PTZFilter struct {
	Status bool
	Position bool
}

type EventSubscription struct {
	Filter FilterType
	SubscriptionPolicy
}

type FilterType struct {
	Any string
}

type SubscriptionPolicy struct {
	Any string
}

type MetadataConfigurationExtension struct {
	Any string
}

type ProfileExtension struct {
	AudioOutputConfiguration AudioOutputConfiguration
	AudioDecoderConfiguration AudioDecoderConfiguration
	Extension ProfileExtension2
}

type AudioOutputConfiguration struct {
	ConfigurationEntity
	OutputToken ReferenceToken
	SendPrimacy xsd.AnyURI
	OutputLevel int
}

type AudioDecoderConfiguration struct {
	ConfigurationEntity
}

type ProfileExtension2 struct {
	Any string
}

type VideoSourceConfigurationOptions struct {
	MaximumNumberOfProfiles int `xml:"MaximumNumberOfProfiles,attr"`
	BoundsRange IntRectangleRange
	VideoSourceTokensAvailable ReferenceToken
	Extension VideoSourceConfigurationOptionsExtension
}

type VideoSourceConfigurationOptionsExtension struct {
	Rotate RotateOptions
	Extension VideoSourceConfigurationOptionsExtension2
}

type RotateOptions struct {
	Mode RotateMode
	DegreeList IntList
	Extension RotateOptionsExtension
}

type IntList struct {
	Items []int
}

type RotateOptionsExtension struct {
	Any string
}

type VideoSourceConfigurationOptionsExtension2 struct {
	SceneOrientationMode SceneOrientationMode
}

type VideoEncoderConfigurationOptions struct {
	QualityRange IntRange
	JPEG JpegOptions
	MPEG4 Mpeg4Options
	H264 H264Options
	Extension VideoEncoderOptionsExtension
}

type JpegOptions struct {
	ResolutionsAvailable VideoResolution
	FrameRateRange IntRange
	EncodingIntervalRange IntRange
}

type Mpeg4Options struct {
	ResolutionsAvailable VideoResolution
	GovLengthRange IntRange
	FrameRateRange IntRange
	EncodingIntervalRange IntRange
	Mpeg4ProfilesSupported Mpeg4Profile
}

type H264Options struct {
	ResolutionsAvailable VideoResolution
	GovLengthRange IntRange
	FrameRateRange IntRange
	EncodingIntervalRange IntRange
	H264ProfilesSupported H264Profile
}

type VideoEncoderOptionsExtension struct {
	JPEG JpegOptions2
	MPEG4 Mpeg4Options2
	H264 H264Options2
	Extension VideoEncoderOptionsExtension2
}

type JpegOptions2 struct {
	JpegOptions
	BitrateRange IntRange
}

type Mpeg4Options2 struct {
	Mpeg4Options
	BitrateRange IntRange
}

type H264Options2 struct {
	H264Options
	BitrateRange IntRange
}

type VideoEncoderOptionsExtension2 struct {
	Any string
}

type AudioSourceConfigurationOptions struct {
	InputTokensAvailable ReferenceToken
	Extension AudioSourceOptionsExtension
}

type AudioSourceOptionsExtension struct {
	Any string
}

type AudioEncoderConfigurationOptions struct {
	Options AudioEncoderConfigurationOption
}

type AudioEncoderConfigurationOption struct {
	Encoding AudioEncoding
	BitrateList IntList
	SampleRateList IntList
}

type MetadataConfigurationOptions struct {
	PTZStatusFilterOptions PTZStatusFilterOptions
	Extension MetadataConfigurationOptionsExtension
}

type PTZStatusFilterOptions struct {
	PanTiltStatusSupported bool
	ZoomStatusSupported bool
	PanTiltPositionSupported bool
	ZoomPositionSupported bool
	Extension PTZStatusFilterOptionsExtension
}

type PTZStatusFilterOptionsExtension struct {
	Any string
}

type MetadataConfigurationOptionsExtension struct {
	CompressionType string
	Extension MetadataConfigurationOptionsExtension2
}

type MetadataConfigurationOptionsExtension2 struct {
	Any string
}

type AudioOutputConfigurationOptions struct {
	OutputTokensAvailable ReferenceToken
	SendPrimacyOptions    xsd.AnyURI
	OutputLevelRange      IntRange
}

type AudioDecoderConfigurationOptions struct {
	AACDecOptions AACDecOptions
	G711DecOptions G711DecOptions
	G726DecOptions G726DecOptions
	Extension AudioDecoderConfigurationOptionsExtension
}

type AACDecOptions struct {
	Bitrate IntList
	SampleRateRange IntList
}

type G711DecOptions struct {
	Bitrate IntList
	SampleRateRange IntList
}

type G726DecOptions struct {
	Bitrate IntList
	SampleRateRange IntList
}

type AudioDecoderConfigurationOptionsExtension struct {
	Any string
}

type StreamSetup struct {
	Stream StreamType
	Transport Transport
}

type StreamType struct {
	Type string
}

type Transport struct {
	Protocol TransportProtocol
	Tunnel Transport
}

//enum
type TransportProtocol struct {
	TransportProtocol string
}

type MediaUri struct {
	Uri                 xsd.AnyURI
	InvalidAfterConnect bool
	InvalidAfterReboot  bool
	Timeout             xsd.Duration
}

type VideoSourceMode struct {
	Token 	ReferenceToken 	`xml:"token,attr"`
	Enabled bool 			`xml:"Enabled,attr"`
	MaxFramerate float64
	MaxResolution VideoResolution
	Encodings EncodingTypes
	Reboot bool
	Description Description
	Extension VideoSourceModeExtension
}

type EncodingTypes struct {
	EncodingTypes []string
}

type Description struct {
	Description string
}

type VideoSourceModeExtension struct {
	Any string
}

type OSDConfigurationOptions struct {
	MaximumNumberOfOSDs MaximumNumberOfOSDs
	Type OSDType
	PositionOption string
	TextOption OSDTextOptions
	ImageOption OSDImgOptions
	Extension OSDConfigurationOptionsExtension
}

type MaximumNumberOfOSDs struct {
	Total 		int `xml:"Total,attr"`
	Image 		int `xml:"Image,attr"`
	PlainText 	int `xml:"PlainText,attr"`
	Date 		int `xml:"Date,attr"`
	Time 		int `xml:"Time,attr"`
	DateAndTime int `xml:"DateAndTime,attr"`
}

type OSDTextOptions struct {
	Type string
	FontSizeRange IntRange
	DateFormat string
	TimeFormat string
	FontColor OSDColorOptions
	BackgroundColor OSDColorOptions
	Extension OSDTextOptionsExtension
}

type OSDColorOptions struct {
	Color ColorOptions
	Transparent IntRange
	Extension OSDColorOptionsExtension
}

type ColorOptions struct {
	ColorList Color
	ColorspaceRange ColorspaceRange
}

type ColorspaceRange struct {
	X          FloatRange
	Y          FloatRange
	Z          FloatRange
	Colorspace xsd.AnyURI
}

type OSDColorOptionsExtension struct {
	Any string
}

type OSDTextOptionsExtension struct {
	Any string
}

type OSDImgOptions struct {
	FormatsSupported 	StringAttrList 	`xml:"FormatsSupported,attr"`
	MaxSize 			int 			`xml:"MaxSize,attr"`
	MaxWidth 			int 			`xml:"MaxWidth,attr"`
	MaxHeight 			int 			`xml:"MaxHeight,attr"`

	ImagePath xsd.AnyURI
	Extension OSDImgOptionsExtension
}

type StringAttrList struct {
	AttrList []string
}

type OSDImgOptionsExtension struct {
	Any string
}

type OSDConfigurationOptionsExtension struct {
	Any string
}

//PTZ

type PTZNode struct {
	DeviceEntity
	FixedHomePosition      xsd.Boolean `xml:"FixedHomePosition,attr"`
	GeoMove                xsd.Boolean `xml:"GeoMove,attr"`
	Name                   Name
	SupportedPTZSpaces     PTZSpaces
	MaximumNumberOfPresets int
	HomeSupported          xsd.Boolean
	AuxiliaryCommands      AuxiliaryData
	Extension              PTZNodeExtension
}

type PTZSpaces struct {
	AbsolutePanTiltPositionSpace Space2DDescription
	AbsoluteZoomPositionSpace Space1DDescription
	RelativePanTiltTranslationSpace Space2DDescription
	RelativeZoomTranslationSpace Space1DDescription
	ContinuousPanTiltVelocitySpace Space2DDescription
	ContinuousZoomVelocitySpace Space1DDescription
	PanTiltSpeedSpace Space1DDescription
	ZoomSpeedSpace Space1DDescription
	Extension PTZSpacesExtension
}

type PTZSpacesExtension struct {
	Any string
}

//TODO: restriction
type AuxiliaryData struct {
	Data string
}

type PTZNodeExtension struct {
	SupportedPresetTour PTZPresetTourSupported
	Extension PTZNodeExtension2
}

type PTZPresetTourSupported struct {
	MaximumNumberOfPresetTours int
	PTZPresetTourOperation PTZPresetTourOperation
	Extension PTZPresetTourSupportedExtension
}

type PTZPresetTourOperation struct {
	Operation string
}

type PTZPresetTourSupportedExtension struct {
	Any string
}

type PTZNodeExtension2 struct {
	Any string
}

type PTZConfigurationOptions struct {
	PTZRamps IntAttrList `xml:"PTZRamps,attr"`
	Spaces PTZSpaces
	PTZTimeout DurationRange
	PTControlDirection PTControlDirectionOptions
	Extension PTZConfigurationOptions2
}

type IntAttrList struct {
	IntAttrList []int
}

type DurationRange struct {
	Min xsd.Duration
	Max xsd.Duration
}

type PTControlDirectionOptions struct {
	EFlip EFlipOptions
	Reverse ReverseOptions
	Extension PTControlDirectionOptionsExtension
}

type EFlipOptions struct {
	Mode EFlipMode
	Extension EFlipOptionsExtension
}

type EFlipOptionsExtension struct {
	Any string
}

type ReverseOptions struct {
	Mode ReverseMode
	Extension ReverseOptionsExtension
}

type ReverseOptionsExtension struct {
	Any string
}

type PTControlDirectionOptionsExtension struct {
	Any string
}

type PTZConfigurationOptions2 struct {
	Any string
}

type PTZPreset struct {
	Token ReferenceToken `xml:"token,attr"`
	Name Name
	PTZPosition PTZVector
}

type PTZVector struct {
	PanTilt Vector2D
	Zoom Vector1D
}

type PTZStatus struct {
	Position   PTZVector
	MoveStatus PTZMoveStatus
	Error      string
	UtcTime    xsd.DateTime
}

type PTZMoveStatus struct {
	PanTilt MoveStatus
	Zoom MoveStatus
}

type MoveStatus struct {
	Status string
}

type GeoLocation struct {
	Lon       xsd.Double `xml:"lon,attr"`
	Lat       xsd.Double `xml:"lat,attr"`
	Elevation xsd.Float  `xml:"elevation,attr"`
}

type PresetTour struct {
	Token             ReferenceToken `xml:"token,attr"`
	Name              Name
	Status            PTZPresetTourStatus
	AutoStart         xsd.Boolean
	StartingCondition PTZPresetTourStartingCondition
	TourSpot          PTZPresetTourSpot
	Extension         PTZPresetTourExtension
}

type PTZPresetTourStatus struct {
	State PTZPresetTourState
	CurrentTourSpot PTZPresetTourSpot
	Extension PTZPresetTourStatusExtension
}

type PTZPresetTourState struct {
	State string
}

type PTZPresetTourSpot struct {
	PresetDetail PTZPresetTourPresetDetail
	Speed        PTZSpeed
	StayTime     xsd.Duration
	Extension    PTZPresetTourSpotExtension
}

type PTZPresetTourPresetDetail struct {
	PresetToken   ReferenceToken
	Home          xsd.Boolean
	PTZPosition   PTZVector
	TypeExtension PTZPresetTourTypeExtension
}

type PTZPresetTourTypeExtension struct {
	Any string
}

type PTZPresetTourSpotExtension struct {
	Any string
}

type PTZPresetTourStatusExtension struct {
	Any string
}

type PTZPresetTourStartingCondition struct {
	RandomPresetOrder xsd.Boolean `xml:"RandomPresetOrder,attr"`
	RecurringTime     int
	RecurringDuration xsd.Duration
	Direction         PTZPresetTourDirection
	Extension         PTZPresetTourStartingConditionExtension
}

type PTZPresetTourDirection struct {
	Direction string
}

type PTZPresetTourStartingConditionExtension struct {
	Any string
}

type PTZPresetTourExtension struct {
	Any string
}

type PTZPresetTourOptions struct {
	AutoStart         xsd.Boolean
	StartingCondition PTZPresetTourStartingConditionOptions
	TourSpot          PTZPresetTourSpotOptions
}

type PTZPresetTourStartingConditionOptions struct {
	RecurringTime IntRange
	RecurringDuration DurationRange
	Direction PTZPresetTourDirection
	Extension PTZPresetTourStartingConditionOptionsExtension
}

type PTZPresetTourStartingConditionOptionsExtension struct {
	Any string
}

type PTZPresetTourSpotOptions struct {
	PresetDetail PTZPresetTourPresetDetailOptions
	StayTime DurationRange
}

type PTZPresetTourPresetDetailOptions struct {
	PresetToken          ReferenceToken
	Home                 xsd.Boolean
	PanTiltPositionSpace Space2DDescription
	ZoomPositionSpace    Space1DDescription
	Extension            PTZPresetTourPresetDetailOptionsExtension
}

type PTZPresetTourPresetDetailOptionsExtension struct {
	Any string
}

//Device

type OnvifVersion struct {
	Major int
	Minor int
}

type SetDateTimeType struct {
	Type string
}

type TimeZone struct {
	TZ xsd.Token
}

type SystemDateTime struct {
	DateTimeType    SetDateTimeType
	DaylightSavings xsd.Boolean
	TimeZone        TimeZone
	UTCDateTime     xsd.DateTime
	LocalDateTime   xsd.DateTime
	Extension       SystemDateTimeExtension
}

type SystemDateTimeExtension struct {
	Any string
}

type FactoryDefaultType struct {
	Type string
}

type AttachmentData struct {
	ContentType ContentType `xml:"contentType,attr"`
	Include Include
}

type Include struct {
	Href xsd.AnyURI `xml:"href,attr"`
}

type BackupFile struct {
	Name string
	Data AttachmentData
}

type SystemLogType struct {
	Type string
}

type SystemLog struct {
	Binary AttachmentData
	String string
}

type SupportInformation struct {
	Binary AttachmentData
	String string
}

type Scope struct {
	ScopeDef  ScopeDefinition
	ScopeItem xsd.AnyURI
}

type ScopeDefinition struct {
	Definition string
}

type DiscoveryMode struct {
	Mode string
}

type NetworkHost struct {
	Type 		NetworkHostType
	IPv4Address IPv4Address
	IPv6Address IPv6Address
	DNSname 	DNSName
	Extension 	NetworkHostExtension
}

type NetworkHostType struct {
	Type string
}

type NetworkHostExtension struct {
	Any string
}

type RemoteUser struct {
	Username           string
	Password           string
	UseDerivedPassword xsd.Boolean
}

type User struct {
	Username string
	Password string
	UserLevel UserLevel
	Extension UserExtension
}

type UserLevel struct {
	Level string
}

type UserExtension struct {
	Any string
}

type CapabilityCategory struct {
	Category string
}

type Capabilities struct {
	Analytics 	AnalyticsCapabilities
	Device 		DeviceCapabilities
	Events 		EventCapabilities
	Imaging 	ImagingCapabilities
	Media 		MediaCapabilities
	PTZ 		PTZCapabilities
	Extension 	CapabilitiesExtension
}

type AnalyticsCapabilities struct {
	XAddr                  xsd.AnyURI
	RuleSupport            xsd.Boolean
	AnalyticsModuleSupport xsd.Boolean
}

type DeviceCapabilities struct {
	XAddr     xsd.AnyURI
	Network   NetworkCapabilities
	System    SystemCapabilities
	IO        IOCapabilities
	Security  SecurityCapabilities
	Extension DeviceCapabilitiesExtension
}

type NetworkCapabilities struct {
	IPFilter          xsd.Boolean
	ZeroConfiguration xsd.Boolean
	IPVersion6        xsd.Boolean
	DynDNS            xsd.Boolean
	Extension         NetworkCapabilitiesExtension
}

type NetworkCapabilitiesExtension struct {
	Dot11Configuration xsd.Boolean
	Extension          NetworkCapabilitiesExtension2
}

type NetworkCapabilitiesExtension2 struct {
	Any string
}

type SystemCapabilities struct {
	DiscoveryResolve  xsd.Boolean
	DiscoveryBye      xsd.Boolean
	RemoteDiscovery   xsd.Boolean
	SystemBackup      xsd.Boolean
	SystemLogging     xsd.Boolean
	FirmwareUpgrade   xsd.Boolean
	SupportedVersions OnvifVersion
	Extension         SystemCapabilitiesExtension
}

type SystemCapabilitiesExtension struct {
	HttpFirmwareUpgrade    xsd.Boolean
	HttpSystemBackup       xsd.Boolean
	HttpSystemLogging      xsd.Boolean
	HttpSupportInformation xsd.Boolean
	Extension              SystemCapabilitiesExtension2
}

type SystemCapabilitiesExtension2 struct {
	Any string
}

type IOCapabilities struct {
	InputConnectors int
	RelayOutputs 	int
	Extension 		IOCapabilitiesExtension
}

type IOCapabilitiesExtension struct {
	Auxiliary         xsd.Boolean
	AuxiliaryCommands AuxiliaryData
	Extension         IOCapabilitiesExtension2
}

type IOCapabilitiesExtension2 struct {
	Any string
}

type SecurityCapabilities struct {
	TLS1_1               xsd.Boolean
	TLS1_2               xsd.Boolean
	OnboardKeyGeneration xsd.Boolean
	AccessPolicyConfig   xsd.Boolean
	X_509Token           xsd.Boolean
	SAMLToken            xsd.Boolean
	KerberosToken        xsd.Boolean
	RELToken             xsd.Boolean
	Extension            SecurityCapabilitiesExtension
}

type SecurityCapabilitiesExtension struct {
	TLS1_0    xsd.Boolean
	Extension SecurityCapabilitiesExtension2
}

type SecurityCapabilitiesExtension2 struct {
	Dot1X              xsd.Boolean
	SupportedEAPMethod int
	RemoteUserHandling xsd.Boolean
}

type DeviceCapabilitiesExtension struct {
	Any string
}

type EventCapabilities struct {
	XAddr                                         xsd.AnyURI
	WSSubscriptionPolicySupport                   xsd.Boolean
	WSPullPointSupport                            xsd.Boolean
	WSPausableSubscriptionManagerInterfaceSupport xsd.Boolean
}

type ImagingCapabilities struct {
	XAddr xsd.AnyURI
}

type MediaCapabilities struct {
	XAddr                 xsd.AnyURI
	StreamingCapabilities RealTimeStreamingCapabilities
	Extension             MediaCapabilitiesExtension
}

type RealTimeStreamingCapabilities struct {
	RTPMulticast xsd.Boolean
	RTP_TCP      xsd.Boolean
	RTP_RTSP_TCP xsd.Boolean
	Extension    RealTimeStreamingCapabilitiesExtension
}

type RealTimeStreamingCapabilitiesExtension struct {
	Any string
}

type MediaCapabilitiesExtension struct {
	ProfileCapabilities ProfileCapabilities
}

type ProfileCapabilities struct {
	MaximumNumberOfProfiles int
}

type PTZCapabilities struct {
	XAddr xsd.AnyURI
}

type CapabilitiesExtension struct {
	DeviceIO 		DeviceIOCapabilities
	Display 		DisplayCapabilities
	Recording 		RecordingCapabilities
	Search 			SearchCapabilities
	Replay 			ReplayCapabilities
	Receiver 		ReceiverCapabilities
	AnalyticsDevice AnalyticsDeviceCapabilities
	Extensions 		CapabilitiesExtension2
}

type DeviceIOCapabilities struct {
	XAddr        xsd.AnyURI
	VideoSources int
	VideoOutputs int
	AudioSources int
	AudioOutputs int
	RelayOutputs int
}

type DisplayCapabilities struct {
	XAddr       xsd.AnyURI
	FixedLayout xsd.Boolean
}

type RecordingCapabilities struct {
	XAddr              xsd.AnyURI
	ReceiverSource     xsd.Boolean
	MediaProfileSource xsd.Boolean
	DynamicRecordings  xsd.Boolean
	DynamicTracks      xsd.Boolean
	MaxStringLength    int
}

type SearchCapabilities struct {
	XAddr          xsd.AnyURI
	MetadataSearch xsd.Boolean
}

type ReplayCapabilities struct {
	XAddr xsd.AnyURI
}

type ReceiverCapabilities struct {
	XAddr                xsd.AnyURI
	RTP_Multicast        xsd.Boolean
	RTP_TCP              xsd.Boolean
	RTP_RTSP_TCP         xsd.Boolean
	SupportedReceivers   int
	MaximumRTSPURILength int
}

type AnalyticsDeviceCapabilities struct {
	XAddr       xsd.AnyURI
	RuleSupport xsd.Boolean
	Extension   AnalyticsDeviceExtension
}

type AnalyticsDeviceExtension struct {
	Any string
}

type CapabilitiesExtension2 struct {
	Any string
}

type HostnameInformation struct {
	FromDHCP  xsd.Boolean
	Name      xsd.Token
	Extension HostnameInformationExtension
}

type HostnameInformationExtension struct {
	Any string
}

type DNSInformation struct {
	FromDHCP     xsd.Boolean
	SearchDomain xsd.Token
	DNSFromDHCP  IPAddress
	DNSManual    IPAddress
	Extension    DNSInformationExtension
}

type DNSInformationExtension struct {
	Any string
}

type NTPInformation struct {
	FromDHCP    xsd.Boolean
	NTPFromDHCP NetworkHost
	NTPManual   NetworkHost
	Extension   NTPInformationExtension
}

type NTPInformationExtension struct {
	Any string
}

type DynamicDNSInformation struct {
	Type      DynamicDNSType
	Name      DNSName
	TTL       xsd.Duration
	Extension DynamicDNSInformationExtension
}

//TODO: enumeration
type DynamicDNSType struct {
	Type string
}

type DynamicDNSInformationExtension struct {
	Any string
}

type NetworkInterface struct {
	DeviceEntity
	Enabled   xsd.Boolean
	Info      NetworkInterfaceInfo
	Link      NetworkInterfaceLink
	IPv4      IPv4NetworkInterface
	IPv6      IPv6NetworkInterface
	Extension NetworkInterfaceExtension
}

type NetworkInterfaceInfo struct {
	Name      xsd.String
	HwAddress HwAddress
	MTU       xsd.Int
}

type HwAddress xsd.Token

type NetworkInterfaceLink struct {
	AdminSettings NetworkInterfaceConnectionSetting
	OperSettings NetworkInterfaceConnectionSetting
	InterfaceType IANA_IfTypes `xml:"IANA-IfTypes"`
}

type IANA_IfTypes xsd.Int

type NetworkInterfaceConnectionSetting struct {
	AutoNegotiation xsd.Boolean
	Speed           xsd.Int
	Duplex          Duplex
}

//TODO: enum
type Duplex struct {

}

type NetworkInterfaceExtension struct {
	InterfaceType IANA_IfTypes
	Dot3 		  Dot3Configuration
	Dot11         Dot11Configuration
	Extension     NetworkInterfaceExtension2
}

type NetworkInterfaceExtension2 xsd.AnyType

type Dot11Configuration struct {
	SSID Dot11SSIDType
	Mode Dot11StationMode
	Alias Name
	Priority NetworkInterfaceConfigPriority
	Security Dot11SecurityConfiguration
}

type Dot11SecurityConfiguration struct {
	Mode Dot11SecurityMode
	Algorithm Dot11Cipher
	PSK Dot11PSKSet
	Dot1X ReferenceToken
	Extension Dot11SecurityConfigurationExtension
}

type Dot11SecurityConfigurationExtension xsd.AnyType

type Dot11PSKSet struct {
	Key Dot11PSK
	Passphrase Dot11PSKPassphrase
	Extension Dot11PSKSetExtension
}

type Dot11PSKSetExtension xsd.AnyType

type Dot11PSKPassphrase xsd.String

type Dot11PSK xsd.HexBinary

//TODO: enumeration
type Dot11Cipher struct {

}

//TODO: enumeration
type Dot11SecurityMode struct {

}

//TODO: restrictions
type NetworkInterfaceConfigPriority xsd.Integer

//TODO: enumeration
type Dot11StationMode struct {

}

//TODO: restrictions
type Dot11SSIDType xsd.HexBinary

type Dot3Configuration xsd.AnyType

type IPv6NetworkInterface struct {
	Enabled xsd.Boolean
	Config IPv6Configuration
}

type IPv6Configuration struct {
	AcceptRouterAdvert xsd.Boolean
	DHCP IPv6DHCPConfiguration
	Manual PrefixedIPv6Address
	LinkLocal PrefixedIPv6Address
	FromDHCP PrefixedIPv6Address
	FromRA PrefixedIPv6Address
	Extension IPv6ConfigurationExtension
}

type IPv6ConfigurationExtension xsd.AnyType

type PrefixedIPv6Address struct {
	Address IPv6Address
	PrefixLength xsd.Int
}

//TODO: enumeration
type IPv6DHCPConfiguration struct {

}

type IPv4NetworkInterface struct {
	Enabled xsd.Boolean
	Config IPv4Configuration
}

type IPv4Configuration struct {
	Manual PrefixedIPv4Address
	LinkLocal PrefixedIPv4Address
	FromDHCP PrefixedIPv4Address
	DHCP xsd.Boolean
}

type PrefixedIPv4Address struct {
	Address IPv4Address
	PrefixLength xsd.Int
}

type NetworkInterfaceSetConfiguration struct {
	Enabled xsd.Boolean
	Link NetworkInterfaceConnectionSetting
	MTU xsd.Int
	IPv4 IPv4NetworkInterfaceSetConfiguration
	IPv6 IPv6NetworkInterfaceSetConfiguration
	Extension NetworkInterfaceSetConfigurationExtension
}

type NetworkInterfaceSetConfigurationExtension struct {
	Dot3 Dot3Configuration
	Dot11 Dot11Configuration
	Extension NetworkInterfaceSetConfigurationExtension2
}

type NetworkInterfaceSetConfigurationExtension2 xsd.AnyType

type IPv6NetworkInterfaceSetConfiguration struct {
	Enabled xsd.Boolean
	AcceptRouterAdvert xsd.Boolean
	Manual PrefixedIPv6Address
	DHCP IPv6DHCPConfiguration
}

type IPv4NetworkInterfaceSetConfiguration struct {
	Enabled xsd.Boolean
	Manual PrefixedIPv4Address
	DHCP xsd.Boolean
}

type NetworkProtocol struct {
	Name NetworkProtocolType
	Enabled xsd.Boolean
	Port xsd.Int
	Extension NetworkProtocolExtension
}

type NetworkProtocolExtension xsd.AnyType

//TODO: enumeration
type NetworkProtocolType struct {

}

type NetworkGateway struct {
	IPv4Address IPv4Address
	IPv6Address IPv6Address
}

type NetworkZeroConfiguration struct {
	InterfaceToken ReferenceToken
	Enabled xsd.Boolean
	Addresses IPv4Address
	Extension NetworkZeroConfigurationExtension
}

type NetworkZeroConfigurationExtension struct {
	Additional NetworkZeroConfiguration
	Extension NetworkZeroConfigurationExtension2
}

type NetworkZeroConfigurationExtension2 xsd.AnyType


type IPAddressFilter struct {
	Type IPAddressFilterType
	IPv4Address PrefixedIPv4Address
	IPv6Address PrefixedIPv6Address
	Extension IPAddressFilterExtension
}

type IPAddressFilterExtension xsd.AnyType

//TODO: enumeration
type IPAddressFilterType struct {

}

//TODO: attribite <xs:attribute ref="xmime:contentType" use="optional"/>
type BinaryData struct {
	X ContentType `xml:"xmime:contentType,attr"`
	Data xsd.Base64Binary
}

type Certificate struct {
	CertificateID xsd.Token
	Certificate   BinaryData
}

type CertificateStatus struct {
	CertificateID xsd.Token
	Status xsd.Boolean
}

type RelayOutput struct {
	DeviceEntity
	Properties RelayOutputSettings
}

type RelayOutputSettings struct {
	Mode RelayMode
	DelayTime xsd.Duration
	IdleState RelayIdleState
}

//TODO:enumeration
type RelayIdleState struct {

}

//TODO: enumeration
type RelayMode struct {

}

//TODO: enumeration
type RelayLogicalState struct {

}

type CertificateWithPrivateKey struct {
	CertificateID xsd.Token
	Certificate BinaryData
	PrivateKey BinaryData
}

type CertificateInformation struct {
	CertificateID xsd.Token
	IssuerDN xsd.String
	SubjectDN xsd.String
	KeyUsage CertificateUsage
	ExtendedKeyUsage CertificateUsage
	KeyLength xsd.Int
	Version xsd.String
	SerialNum xsd.String
	SignatureAlgorithm xsd.String
	Validity DateTimeRange
	Extension CertificateInformationExtension
}

type CertificateInformationExtension xsd.AnyType

type DateTimeRange struct {
	From xsd.DateTime
	Until xsd.DateTime
}

type CertificateUsage struct {
	Critical xsd.Boolean `xml:"Critical,attr"`
	CertificateUsage xsd.String
}

type Dot1XConfiguration struct {
	Dot1XConfigurationToken ReferenceToken
	Identity xsd.String
	AnonymousID xsd.String
	EAPMethod xsd.Int
	CACertificateID xsd.Token
	EAPMethodConfiguration EAPMethodConfiguration
	Extension Dot1XConfigurationExtension
}

type Dot1XConfigurationExtension xsd.AnyType

type EAPMethodConfiguration struct {
	TLSConfiguration TLSConfiguration
	Password xsd.String
	Extension EapMethodExtension
}

type EapMethodExtension xsd.AnyType

type TLSConfiguration struct {
	CertificateID xsd.Token
}

type Dot11Capabilities struct {
	TKIP xsd.Boolean
	ScanAvailableNetworks xsd.Boolean
	MultipleConfiguration xsd.Boolean
	AdHocStationMode xsd.Boolean
	WEP xsd.Boolean
}

type Dot11Status struct {
	SSID Dot11SSIDType
	BSSID xsd.String
	PairCipher Dot11Cipher
	GroupCipher Dot11Cipher
	SignalStrength Dot11SignalStrength
	ActiveConfigAlias ReferenceToken
}

//TODO: enumeration
type Dot11SignalStrength struct {

}

type Dot11AvailableNetworks struct {
	SSID Dot11SSIDType
	BSSID xsd.String
	AuthAndMangementSuite Dot11AuthAndMangementSuite
	PairCipher Dot11Cipher
	GroupCipher Dot11Cipher
	SignalStrength Dot11SignalStrength
	Extension Dot11AvailableNetworksExtension
}

type Dot11AvailableNetworksExtension xsd.AnyType

//TODO: enumeration
type Dot11AuthAndMangementSuite struct {

}

type SystemLogUriList struct {
	SystemLog SystemLogUri
}

type SystemLogUri struct {
	Type SystemLogType
	Uri xsd.AnyURI
}

type LocationEntity struct {
	Entity xsd.String `xml:"Entity,attr"`
	Token ReferenceToken `xml:"ReferenceToken,attr"`
	Fixed xsd.Boolean `xml:"Fixed,attr"`
	GeoSource xsd.AnyURI `xml:"GeoSource,attr"`
	AutoGeo xsd.Boolean `xml:"AutoGeo,attr"`

	GeoLocation GeoLocation
	GeoOrientation GeoOrientation
	LocalLocation LocalLocation
	LocalOrientation LocalOrientation
}

type LocalOrientation struct {
	Lon xsd.Double `xml:"lon,attr"`
	Lat xsd.Double `xml:"lat,attr"`
	Elevation xsd.Float `xml:"elevation,attr"`
}

type LocalLocation struct {
	X xsd.Float `xml:"x,attr"`
	Y xsd.Float `xml:"y,attr"`
	Z xsd.Float `xml:"z,attr"`
}

type GeoOrientation struct {
	Roll xsd.Float `xml:"roll,attr"`
	Pitch xsd.Float `xml:"pitch,attr"`
	Yaw xsd.Float `xml:"yaw,attr"`
}

type FocusMove struct {
	Absolute AbsoluteFocus
	Relative RelativeFocus
	Continuous ContinuousFocus
}

type ContinuousFocus struct {
	Speed xsd.Float
}

type RelativeFocus struct {
	Distance xsd.Float
	Speed xsd.Float
}

type AbsoluteFocus struct {
	Position xsd.Float
	Speed xsd.Float
}


























