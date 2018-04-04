package Event

import (
	"github.com/yakovlevdmv/goonvif/xsd"
)

//const WSDL = "http://www.onvif.org/ver10/events/wsdl"

type FilterType struct { //wsnt http://docs.oasis-open.org/wsn/b-2.xsd
	Any string
}

//<xsd:union memberTypes="xsd:dateTime xsd:duration"/>
type AbsoluteOrRelativeTimeType struct { //wsnt http://docs.oasis-open.org/wsn/b-2.xsd
	xsd.DateTime
	xsd.Duration
}

type SubscriptionPolicy struct { //tev http://www.onvif.org/ver10/events/wsdl
	ChangedOnly xsd.Boolean `xml:"ChangedOnly,attr"`
}

type Capabilities struct { //tev
	WSSubscriptionPolicySupport xsd.Boolean `xml:"WSSubscriptionPolicySupport,attr"`
	WSPullPointSupport xsd.Boolean `xml:"WSPullPointSupport,attr"`
	WSPausableSubscriptionManagerInterfaceSupport xsd.Boolean `xml:"WSPausableSubscriptionManagerInterfaceSupport,attr"`
	MaxNotificationProducers xsd.Int `xml:"MaxNotificationProducers,attr"`
	MaxPullPoints xsd.Int `xml:"MaxPullPoints,attr"`
	PersistentNotificationStorage xsd.Boolean `xml:"PersistentNotificationStorage,attr"`
}

type EndpointReferenceType struct { //wsa http://www.w3.org/2005/08/addressing/ws-addr.xsd
	Address AttributedURIType
	ReferenceParameters ReferenceParametersType
	Metadata // todo:разобраться с этим: понять, на какой тип ссылается
}

type AttributedURIType struct { //wsa https://www.w3.org/2005/08/addressing/ws-addr.xsd
	Any xsd.AnyURI //extension
	//Here can be anyAttribute
}

type ReferenceParametersType struct { //wsa https://www.w3.org/2005/08/addressing/ws-addr.xsd
	Any string
	//Here can be anyAttribute
}

type Metadata MetadataType //wsa https://www.w3.org/2005/08/addressing/ws-addr.xsd

type MetadataType struct { //wsa https://www.w3.org/2005/08/addressing/ws-addr.xsd
	Any string
	//Here can be anyAttribute
}

type CurrentTime xsd.DateTime //wsnt http://docs.oasis-open.org/wsn/b-2.xsd
type TerminationTime xsd.DateTime  //wsnt http://docs.oasis-open.org/wsn/b-2.xsd
type FixedTopicSet xsd.Boolean //wsnt http://docs.oasis-open.org/wsn/b-2.xsd

type TopicSet TopicSetType //wstop http://docs.oasis-open.org/wsn/t-1.xsd

type TopicSetType struct { //wstop http://docs.oasis-open.org/wsn/t-1.xsd
	ExtensibleDocumented
	//here can be any element
}

type ExtensibleDocumented struct { //wstop http://docs.oasis-open.org/wsn/t-1.xsd
	Documentation Documentation //к xsd-документе documentation с маленькой буквы начинается
	//here can be anyAttribute
}

type Documentation xsd.AnyType //wstop http://docs.oasis-open.org/wsn/t-1.xsd

type TopicExpressionDialect xsd.AnyURI

type NotificationMessage NotificationMessageHolderType //wsnt http://docs.oasis-open.org/wsn/b-2.xsd

type NotificationMessageHolderType struct {
	SubscriptionReference SubscriptionReference //wsnt http://docs.oasis-open.org/wsn/b-2.xsd
	Topic Topic
	ProducerReference ProducerReference
	Message Message
}

type SubscriptionReference EndpointReferenceType
type Topic TopicExpressionType
type ProducerReference EndpointReferenceType
type Message xsd.AnyType

type TopicExpressionType struct { //wsnt http://docs.oasis-open.org/wsn/b-2.xsd
	Dialect xsd.AnyURI `xml:"Dialect,attr"`
} 


//Event main types

type GetServiceCapabilities struct {

}


type GetServiceCapabilitiesResponse struct {
	Capabilities Capabilities

}


type CreatePullPointSubscription struct {
	Filter FilterType
	InitialTerminationTime AbsoluteOrRelativeTimeType
	SubscriptionPolicy SubscriptionPolicy

}


type CreatePullPointSubscriptionResponse struct {
	SubscriptionReference EndpointReferenceType
	CurrentTime CurrentTime
	TerminationTime TerminationTime

}

type ResourceUnknownFault struct {

}

type InvalidFilterFault struct {

}

type TopicExpressionDialectUnknownFault struct {

}

type InvalidTopicExpressionFault struct {

}

type TopicNotSupportedFault struct {

}

type InvalidProducerPropertiesExpressionFault struct {

}

type InvalidMessageContentExpressionFault struct {

}

type UnacceptableInitialTerminationTimeFault struct {

}

type UnrecognizedPolicyRequestFault struct {

}

type UnsupportedPolicyRequestFault struct {

}

type NotifyMessageNotSupportedFault struct {

}

type SubscribeCreationFailedFault struct {

}


type GetEventProperties struct {

}


type GetEventPropertiesResponse struct {
	TopicNamespaceLocation xsd.AnyURI
	FixedTopicSet FixedTopicSet
	TopicSet TopicSet
	TopicExpressionDialect TopicExpressionDialect
	MessageContentFilterDialect xsd.AnyURI
	ProducerPropertiesFilterDialect xsd.AnyURI
	MessageContentSchemaLocation xsd.AnyURI

}

//Port type PullPointSubscription

type PullMessages struct {
	Timeout xsd.Duration
	MessageLimit xsd.Int
}

type PullMessagesResponse struct {
	CurrentTime CurrentTime
	TerminationTime TerminationTime
	NotificationMessage NotificationMessage
}

type PullMessagesFaultResponse struct {
	MaxTimeout xsd.Duration
	MaxMessageLimit xsd.Int
}

type Seek struct {
	UtcTime xsd.DateTime
	Reverse xsd.Boolean
}

type SeekResponse struct {

}

type SetSynchronizationPoint struct {

}

type SetSynchronizationPointResponse struct {

}