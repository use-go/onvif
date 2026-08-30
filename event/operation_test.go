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
