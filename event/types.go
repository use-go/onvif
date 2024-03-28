package event

import (
	"github.com/use-go/onvif/xsd"
	"github.com/use-go/onvif/xsd/onvif"
)

//Address Alias
type Address xsd.String

//CurrentTime alias
type CurrentTime xsd.DateTime //wsnt http://docs.oasis-open.org/wsn/b-2.xsd

//TerminationTime alias
type TerminationTime xsd.DateTime //wsnt http://docs.oasis-open.org/wsn/b-2.xsd

//FixedTopicSet alias
type FixedTopicSet xsd.Boolean //wsnt http://docs.oasis-open.org/wsn/b-2.xsd

//Documentation alias
type Documentation xsd.AnyType //wstop http://docs.oasis-open.org/wsn/t-1.xsd

//TopicExpressionDialect alias
type TopicExpressionDialect xsd.AnyURI

//Message alias
type Message xsd.AnyType

//MessageNotification alias
type MessageNotification struct {
	Message MessageNotificationHolderType
}

type MessageNotificationHolderType struct {
	UtcTime           xsd.DateTime     `xml:",attr"`
	PropertyOperation xsd.String       `xml:",attr"`
	Source            onvif.SimpleItem `xml:"Source>SimpleItem"`
	Data              onvif.SimpleItem `xml:"Data>SimpleItem"`
}

//ActionType for AttributedURIType
type ActionType AttributedURIType

//AttributedURIType in ws-addr
type AttributedURIType xsd.AnyURI //wsa https://www.w3.org/2005/08/addressing/ws-addr.xsd

// AbsoluteOrRelativeTimeType <xsd:union memberTypes="xsd:dateTime xsd:duration"/>
type AbsoluteOrRelativeTimeType xsd.AnySimpleType //wsnt http://docs.oasis-open.org/wsn/b-2.xsd

//EndpointReferenceType in ws-addr
type EndpointReferenceType struct { //wsa http://www.w3.org/2005/08/addressing/ws-addr.xsd
	Address             AttributedURIType
	ReferenceParameters ReferenceParametersType
	Metadata            MetadataType
}

// FilterType struct
type FilterType struct {
	TopicExpression TopicExpressionType `xml:"wsnt:TopicExpression"`
	MessageContent  *QueryExpressionType `xml:"wsnt:MessageContent"`
}

//EndpointReference alais
type EndpointReference EndpointReferenceType

//ReferenceParametersType in ws-addr
type ReferenceParametersType struct { //wsa https://www.w3.org/2005/08/addressing/ws-addr.xsd
	//Here can be anyAttribute
}

//Metadata in ws-addr
type Metadata MetadataType //wsa https://www.w3.org/2005/08/addressing/ws-addr.xsd

//MetadataType in ws-addr
type MetadataType struct { //wsa https://www.w3.org/2005/08/addressing/ws-addr.xsd

	//Here can be anyAttribute
}

//TopicSet alias
type TopicSet TopicSetType //wstop http://docs.oasis-open.org/wsn/t-1.xsd

//TopicSetType alias
type TopicSetType struct { //wstop http://docs.oasis-open.org/wsn/t-1.xsd
	ExtensibleDocumented
	//here can be any element
}

//ExtensibleDocumented struct
type ExtensibleDocumented struct { //wstop http://docs.oasis-open.org/wsn/t-1.xsd
	Documentation Documentation //к xsd-документе documentation с маленькой буквы начинается
	//here can be anyAttribute
}

//ProducerReference Alias
type ProducerReference EndpointReferenceType

//SubscriptionReference Alias
type SubscriptionReference EndpointReferenceType

//NotificationMessageHolderType Alias
type NotificationMessageHolderType struct {
	SubscriptionReference SubscriptionReference //wsnt http://docs.oasis-open.org/wsn/b-2.xsd
	Topic                 Topic
	ProducerReference     ProducerReference
	Message               MessageNotification
}

//NotificationMessage Alias
type NotificationMessage NotificationMessageHolderType //wsnt http://docs.oasis-open.org/wsn/b-2.xsd

//QueryExpressionType struct for wsnt:MessageContent
type QueryExpressionType struct { //wsnt http://docs.oasis-open.org/wsn/b-2.xsd
	Dialect     xsd.AnyURI `xml:"Dialect,attr"`
	MessageKind xsd.String `xml:",chardata"` // boolean(ncex:Producer="15")
}

//MessageContentType Alias
type MessageContentType QueryExpressionType

//QueryExpression Alias
type QueryExpression QueryExpressionType

//TopicExpressionType struct for wsnt:TopicExpression
type TopicExpressionType struct { //wsnt http://docs.oasis-open.org/wsn/b-2.xsd
	Dialect    xsd.AnyURI `xml:"Dialect,attr"`
	TopicKinds xsd.String `xml:",chardata"`
}

//Topic Alias
type Topic TopicExpressionType

// Capabilities of event
type Capabilities struct { //tev
	WSSubscriptionPolicySupport                   xsd.Boolean `xml:"WSSubscriptionPolicySupport,attr"`
	WSPullPointSupport                            xsd.Boolean `xml:"WSPullPointSupport,attr"`
	WSPausableSubscriptionManagerInterfaceSupport xsd.Boolean `xml:"WSPausableSubscriptionManagerInterfaceSupport,attr"`
	MaxNotificationProducers                      xsd.Int     `xml:"MaxNotificationProducers,attr"`
	MaxPullPoints                                 xsd.Int     `xml:"MaxPullPoints,attr"`
	PersistentNotificationStorage                 xsd.Boolean `xml:"PersistentNotificationStorage,attr"`
}

//ResourceUnknownFault response type
type ResourceUnknownFault struct {
}

//InvalidFilterFault response type
type InvalidFilterFault struct {
}

//TopicExpressionDialectUnknownFault response type
type TopicExpressionDialectUnknownFault struct {
}

//InvalidTopicExpressionFault response type
type InvalidTopicExpressionFault struct {
}

//TopicNotSupportedFault response type
type TopicNotSupportedFault struct {
}

//InvalidProducerPropertiesExpressionFault response type
type InvalidProducerPropertiesExpressionFault struct {
}

//InvalidMessageContentExpressionFault response type
type InvalidMessageContentExpressionFault struct {
}

//UnacceptableInitialTerminationTimeFault response type
type UnacceptableInitialTerminationTimeFault struct {
}

//UnrecognizedPolicyRequestFault response type
type UnrecognizedPolicyRequestFault struct {
}

//UnsupportedPolicyRequestFault response type
type UnsupportedPolicyRequestFault struct {
}

//NotifyMessageNotSupportedFault response type
type NotifyMessageNotSupportedFault struct {
}

//SubscribeCreationFailedFault response type
type SubscribeCreationFailedFault struct {
}
