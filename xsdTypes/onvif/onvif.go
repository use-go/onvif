package onvif

import "github.com/yakovlevdmv/goonvif/xsdTypes"

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
type DNSName xsdTypes.token

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
	X float64 `xml:"X,attr"`
	Y float64 `xml:"Y,attr"`
	Z float64 `xml:"Z,attr"`
	Colorspace xsdTypes.AnyURI `xml:"Colorspace,attr"`
}


type OSDTextConfigurationExtension struct {
	Any string
}

type OSDImgConfiguration struct {
	ImgPath xsdTypes.AnyURI
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
	BoundaryType string
	BoundaryOffset float64
	ResponseTime xsdTypes.Duration
	Extension IrCutFilterAutoAdjustmentExtension
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
	Encoding VideoEncoding
	Resolution VideoResolution
	Quality float64
	RateControl VideoRateControl
	MPEG4 Mpeg4Configuration
	H264 H264Configuration
	Multicast MulticastConfiguration
	SessionTimeout xsdTypes.Duration
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
	Address xsdTypes.token
}

type IPv6Address struct {
	Address xsdTypes.token
}

type AudioEncoderConfiguration struct {
	ConfigurationEntity
	Encoding AudioEncoding
	Bitrate int
	SampleRate int
	Multicast MulticastConfiguration
	SessionTimeout xsdTypes.Duration
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
	Name 		string 			`xml:"Name,attr"`
	Type 		xsdTypes.QName 	`xml:"Type,attr"`
	Parameters 	ItemList
}

type ItemList struct {
	SimpleItem
	ElementItem
	Extension ItemListExtension
}

type SimpleItem struct {
	Name 	string 			`xml:"Name,attr"`
	Value 	xsdTypes.anySimpleType 	`xml:"Value,attr"`
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
	MoveRamp 		int `xml:"MoveRamp,attr"`
	PresetRamp 		int `xml:"PresetRamp,attr"`
	PresetTourRamp 	int `xml:"PresetTourRamp,attr"`
	NodeToken ReferenceToken
	DefaultAbsolutePantTiltPositionSpace xsdTypes.AnyURI
	DefaultAbsoluteZoomPositionSpace xsdTypes.AnyURI
	DefaultRelativePanTiltTranslationSpace xsdTypes.AnyURI
	DefaultRelativeZoomTranslationSpace xsdTypes.AnyURI
	DefaultContinuousPanTiltVelocitySpace xsdTypes.AnyURI
	DefaultContinuousZoomVelocitySpace xsdTypes.AnyURI
	DefaultPTZSpeed PTZSpeed
	DefaultPTZTimeout xsdTypes.Duration
	PanTiltLimits PanTiltLimits
	ZoomLimits ZoomLimits
	Extension PTZConfigurationExtension
}

type PTZSpeed struct {
	PanTilt Vector2D
	Zoom Vector1D
}

type Vector2D struct {
	X 		float64 `xml:"x,attr"`
	Y 		float64 `xml:"y,attr"`
	Space 	xsdTypes.AnyURI 	`xml:"space,attr"`
}

type Vector1D struct {
	X 		float64 `xml:"x,attr"`
	Space 	xsdTypes.AnyURI 	`xml:"space,attr"`
}

type PanTiltLimits struct {
	Range Space2DDescription
}

type Space2DDescription struct {
	URI xsdTypes.AnyURI
	XRange FloatRange
	YRange FloatRange
}

type ZoomLimits struct {
	Range Space1DDescription
}

type Space1DDescription struct {
	URI xsdTypes.AnyURI
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
	CompressionType 				string `xml:"CompressionType,attr"`
	PTZStatus 						PTZFilter
	Events 							EventSubscription
	Analytics 						bool
	Multicast 						MulticastConfiguration
	SessionTimeout 					xsdTypes.Duration
	AnalyticsEngineConfiguration 	AnalyticsEngineConfiguration
	Extension 						MetadataConfigurationExtension
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
	SendPrimacy xsdTypes.AnyURI
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
	SendPrimacyOptions xsdTypes.AnyURI
	OutputLevelRange IntRange
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
	Uri xsdTypes.AnyURI
	InvalidAfterConnect bool
	InvalidAfterReboot bool
	Timeout xsdTypes.Duration
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
	X FloatRange
	Y FloatRange
	Z FloatRange
	Colorspace xsdTypes.AnyURI
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

	ImagePath xsdTypes.AnyURI
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
	FixedHomePosition 	xsdTypes.Boolean `xml:"FixedHomePosition,attr"`
	GeoMove 			xsdTypes.Boolean `xml:"GeoMove,attr"`
	Name Name
	SupportedPTZSpaces PTZSpaces
	MaximumNumberOfPresets int
	HomeSupported xsdTypes.Boolean
	AuxiliaryCommands AuxiliaryData
	Extension PTZNodeExtension
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
	Min xsdTypes.Duration
	Max xsdTypes.Duration
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
	Position PTZVector
	MoveStatus PTZMoveStatus
	Error string
	UtcTime xsdTypes.DateTime
}

type PTZMoveStatus struct {
	PanTilt MoveStatus
	Zoom MoveStatus
}

type MoveStatus struct {
	Status string
}

type GeoLocation struct {
	Lon 		xsdTypes.Double 	`xml:"lon,attr"`
	Lat 		xsdTypes.Double 	`xml:"lat,attr"`
	Elevation 	xsdTypes.Float 	`xml:"elevation,attr"`
}

type PresetTour struct {
	Token ReferenceToken `xml:"token,attr"`
	Name Name
	Status PTZPresetTourStatus
	AutoStart xsdTypes.Boolean
	StartingCondition PTZPresetTourStartingCondition
	TourSpot PTZPresetTourSpot
	Extension PTZPresetTourExtension
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
	Speed PTZSpeed
	StayTime xsdTypes.Duration
	Extension PTZPresetTourSpotExtension
}

type PTZPresetTourPresetDetail struct {
	PresetToken ReferenceToken
	Home xsdTypes.Boolean
	PTZPosition PTZVector
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
	RandomPresetOrder xsdTypes.Boolean `xml:"RandomPresetOrder,attr"`
	RecurringTime int
	RecurringDuration xsdTypes.Duration
	Direction PTZPresetTourDirection
	Extension PTZPresetTourStartingConditionExtension
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
	AutoStart xsdTypes.Boolean
	StartingCondition PTZPresetTourStartingConditionOptions
	TourSpot PTZPresetTourSpotOptions
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
	PresetToken ReferenceToken
	Home xsdTypes.Boolean
	PanTiltPositionSpace Space2DDescription
	ZoomPositionSpace Space1DDescription
	Extension PTZPresetTourPresetDetailOptionsExtension
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
	TZ xsdTypes.token
}

type SystemDateTime struct {
	DateTimeType SetDateTimeType
	DaylightSavings xsdTypes.Boolean
	TimeZone TimeZone
	UTCDateTime xsdTypes.DateTime
	LocalDateTime xsdTypes.DateTime
	Extension SystemDateTimeExtension
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
	Href xsdTypes.AnyURI `xml:"href,attr"`
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
	ScopeDef ScopeDefinition
	ScopeItem xsdTypes.AnyURI
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
	Username string
	Password string
	UseDerivedPassword xsdTypes.Boolean
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
	XAddr 					xsdTypes.AnyURI
	RuleSupport 			xsdTypes.Boolean
	AnalyticsModuleSupport 	xsdTypes.Boolean
}

type DeviceCapabilities struct {
	XAddr 		xsdTypes.AnyURI
	Network 	NetworkCapabilities
	System 		SystemCapabilities
	IO 			IOCapabilities
	Security 	SecurityCapabilities
	Extension 	DeviceCapabilitiesExtension
}

type NetworkCapabilities struct {
	IPFilter 			xsdTypes.Boolean
	ZeroConfiguration 	xsdTypes.Boolean
	IPVersion6 			xsdTypes.Boolean
	DynDNS 				xsdTypes.Boolean
	Extension 			NetworkCapabilitiesExtension
}

type NetworkCapabilitiesExtension struct {
	Dot11Configuration 	xsdTypes.Boolean
	Extension 			NetworkCapabilitiesExtension2
}

type NetworkCapabilitiesExtension2 struct {
	Any string
}

type SystemCapabilities struct {
	DiscoveryResolve 	xsdTypes.Boolean
	DiscoveryBye 		xsdTypes.Boolean
	RemoteDiscovery 	xsdTypes.Boolean
	SystemBackup 		xsdTypes.Boolean
	SystemLogging 		xsdTypes.Boolean
	FirmwareUpgrade 	xsdTypes.Boolean
	SupportedVersions 	OnvifVersion
	Extension 			SystemCapabilitiesExtension
}

type SystemCapabilitiesExtension struct {
	HttpFirmwareUpgrade 	xsdTypes.Boolean
	HttpSystemBackup 		xsdTypes.Boolean
	HttpSystemLogging 		xsdTypes.Boolean
	HttpSupportInformation 	xsdTypes.Boolean
	Extension 				SystemCapabilitiesExtension2
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
	Auxiliary 			xsdTypes.Boolean
	AuxiliaryCommands 	AuxiliaryData
	Extension 			IOCapabilitiesExtension2
}

type IOCapabilitiesExtension2 struct {
	Any string
}

type SecurityCapabilities struct {
	TLS1_1 					xsdTypes.Boolean
	TLS1_2 					xsdTypes.Boolean
	OnboardKeyGeneration 	xsdTypes.Boolean
	AccessPolicyConfig 		xsdTypes.Boolean
	X_509Token 				xsdTypes.Boolean
	SAMLToken 				xsdTypes.Boolean
	KerberosToken 			xsdTypes.Boolean
	RELToken 				xsdTypes.Boolean
	Extension 				SecurityCapabilitiesExtension
}

type SecurityCapabilitiesExtension struct {
	TLS1_0 		xsdTypes.Boolean
	Extension 	SecurityCapabilitiesExtension2
}

type SecurityCapabilitiesExtension2 struct {
	Dot1X 				xsdTypes.Boolean
	SupportedEAPMethod 	int
	RemoteUserHandling 	xsdTypes.Boolean
}

type DeviceCapabilitiesExtension struct {
	Any string
}

type EventCapabilities struct {
	XAddr 											xsdTypes.AnyURI
	WSSubscriptionPolicySupport 					xsdTypes.Boolean
	WSPullPointSupport 								xsdTypes.Boolean
	WSPausableSubscriptionManagerInterfaceSupport 	xsdTypes.Boolean
}

type ImagingCapabilities struct {
	XAddr xsdTypes.AnyURI
}

type MediaCapabilities struct {
	XAddr 					xsdTypes.AnyURI
	StreamingCapabilities 	RealTimeStreamingCapabilities
	Extension 				MediaCapabilitiesExtension
}

type RealTimeStreamingCapabilities struct {
	RTPMulticast 	xsdTypes.Boolean
	RTP_TCP 		xsdTypes.Boolean
	RTP_RTSP_TCP 	xsdTypes.Boolean
	Extension 		RealTimeStreamingCapabilitiesExtension
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
	XAddr xsdTypes.AnyURI
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
	XAddr 			xsdTypes.AnyURI
	VideoSources 	int
	VideoOutputs 	int
	AudioSources 	int
	AudioOutputs 	int
	RelayOutputs 	int
}

type DisplayCapabilities struct {
	XAddr 		xsdTypes.AnyURI
	FixedLayout xsdTypes.Boolean
}

type RecordingCapabilities struct {
	XAddr 				xsdTypes.AnyURI
	ReceiverSource 		xsdTypes.Boolean
	MediaProfileSource 	xsdTypes.Boolean
	DynamicRecordings 	xsdTypes.Boolean
	DynamicTracks 		xsdTypes.Boolean
	MaxStringLength 	int
}

type SearchCapabilities struct {
	XAddr 			xsdTypes.AnyURI
	MetadataSearch 	xsdTypes.Boolean
}

type ReplayCapabilities struct {
	XAddr xsdTypes.AnyURI
}

type ReceiverCapabilities struct {
	XAddr 					xsdTypes.AnyURI
	RTP_Multicast 			xsdTypes.Boolean
	RTP_TCP 				xsdTypes.Boolean
	RTP_RTSP_TCP 			xsdTypes.Boolean
	SupportedReceivers		int
	MaximumRTSPURILength 	int
}

type AnalyticsDeviceCapabilities struct {
	XAddr 		xsdTypes.AnyURI
	RuleSupport xsdTypes.Boolean
	Extension 	AnalyticsDeviceExtension
}

type AnalyticsDeviceExtension struct {
	Any string
}

type CapabilitiesExtension2 struct {
	Any string
}

type HostnameInformation struct {
	FromDHCP 	xsdTypes.Boolean
	Name 		xsdTypes.token
	Extension 	HostnameInformationExtension
}

type HostnameInformationExtension struct {
	Any string
}

type DNSInformation struct {
	FromDHCP 		xsdTypes.Boolean
	SearchDomain 	xsdTypes.token
	DNSFromDHCP 	IPAddress
	DNSManual 		IPAddress
	Extension 		DNSInformationExtension
}

type DNSInformationExtension struct {
	Any string
}