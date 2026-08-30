package event

import (
	"encoding/xml"
	"strings"
	"testing"
)

func TestSubscribeMarshalsChangedOnlyAsChildElement(t *testing.T) {
	req := Subscribe{
		SubscriptionPolicy: &SubscriptionPolicy{ChangedOnly: &ChangedOnly{}},
	}

	got, err := xml.Marshal(req)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	const want = "<wsnt:SubscriptionPolicy><tev:ChangedOnly></tev:ChangedOnly></wsnt:SubscriptionPolicy>"
	if !strings.Contains(string(got), want) {
		t.Errorf("policy not serialized as ONVIF declares\n got: %s\nwant substring: %s", got, want)
	}
}

func TestSubscribeOmitsUnsetSubscriptionPolicy(t *testing.T) {
	got, err := xml.Marshal(Subscribe{})
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	if strings.Contains(string(got), "SubscriptionPolicy") {
		t.Errorf("unset optional policy must not be sent, got: %s", got)
	}
}

func TestCreatePullPointSubscriptionMarshalsAllChildrenInTevNamespace(t *testing.T) {
	terminationTime := AbsoluteOrRelativeTimeType("PT60S")
	req := CreatePullPointSubscription{
		Filter: &FilterType{
			TopicExpression: TopicExpressionType{
				Dialect:    "http://www.onvif.org/ver10/tev/topicExpression/ConcreteSet",
				TopicKinds: "tns1:RuleEngine//.",
			},
		},
		InitialTerminationTime: &terminationTime,
		SubscriptionPolicy:     &SubscriptionPolicy{ChangedOnly: &ChangedOnly{}},
	}

	got, err := xml.Marshal(req)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	const want = "<tev:CreatePullPointSubscription>" +
		"<tev:Filter>" +
		`<wsnt:TopicExpression Dialect="http://www.onvif.org/ver10/tev/topicExpression/ConcreteSet">` +
		"tns1:RuleEngine//." +
		"</wsnt:TopicExpression>" +
		"</tev:Filter>" +
		"<tev:InitialTerminationTime>PT60S</tev:InitialTerminationTime>" +
		"<tev:SubscriptionPolicy><tev:ChangedOnly></tev:ChangedOnly></tev:SubscriptionPolicy>" +
		"</tev:CreatePullPointSubscription>"
	if string(got) != want {
		t.Errorf("payload does not match the Event WSDL\n got: %s\nwant: %s", got, want)
	}
}

func TestCreatePullPointSubscriptionOmitsUnsetOptionalElements(t *testing.T) {
	got, err := xml.Marshal(CreatePullPointSubscription{})
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	const want = "<tev:CreatePullPointSubscription></tev:CreatePullPointSubscription>"
	if string(got) != want {
		t.Errorf("every child of CreatePullPointSubscription is minOccurs=0\n got: %s\nwant: %s", got, want)
	}
}
