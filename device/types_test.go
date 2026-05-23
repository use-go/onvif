package device

import (
	"encoding/xml"
	"strings"
	"testing"
)

// TestGetServicesResponse_PreservesAllServices guards against a regression
// where GetServicesResponse.Service was a single value, silently dropping
// every <Service> element except the last during xml.Unmarshal.
func TestGetServicesResponse_PreservesAllServices(t *testing.T) {
	payload := []byte(`
<GetServicesResponse xmlns="http://www.onvif.org/ver10/device/wsdl">
  <Service>
    <Namespace>http://www.onvif.org/ver10/device/wsdl</Namespace>
    <XAddr>http://192.0.2.10/onvif/device_service</XAddr>
    <Version><Major>2</Major><Minor>40</Minor></Version>
  </Service>
  <Service>
    <Namespace>http://www.onvif.org/ver10/media/wsdl</Namespace>
    <XAddr>http://192.0.2.10/onvif/media_service</XAddr>
    <Version><Major>2</Major><Minor>40</Minor></Version>
  </Service>
  <Service>
    <Namespace>http://www.onvif.org/ver20/ptz/wsdl</Namespace>
    <XAddr>http://192.0.2.10/onvif/ptz_service</XAddr>
    <Version><Major>2</Major><Minor>40</Minor></Version>
  </Service>
  <Service>
    <Namespace>http://www.onvif.org/ver10/events/wsdl</Namespace>
    <XAddr>http://192.0.2.10/onvif/events_service</XAddr>
    <Version><Major>2</Major><Minor>40</Minor></Version>
  </Service>
</GetServicesResponse>`)

	var resp GetServicesResponse
	if err := xml.Unmarshal(payload, &resp); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}

	if got, want := len(resp.Service), 4; got != want {
		t.Fatalf("len(Service) = %d, want %d (all <Service> elements must be preserved)", got, want)
	}

	wantNamespaces := []string{
		"http://www.onvif.org/ver10/device/wsdl",
		"http://www.onvif.org/ver10/media/wsdl",
		"http://www.onvif.org/ver20/ptz/wsdl",
		"http://www.onvif.org/ver10/events/wsdl",
	}
	for i, ns := range wantNamespaces {
		if got := string(resp.Service[i].Namespace); got != ns {
			t.Errorf("Service[%d].Namespace = %q, want %q", i, got, ns)
		}
		if v := resp.Service[i].Version; v.Major != 2 || v.Minor != 40 {
			t.Errorf("Service[%d].Version = {%d,%d}, want {2,40}", i, v.Major, v.Minor)
		}
		if got := string(resp.Service[i].XAddr); got == "" {
			t.Errorf("Service[%d].XAddr is empty", i)
		}
	}
}

// TestServiceCapabilities_PreservesInnerXML guards against a regression where
// the per-service <Capabilities> payload — declared as xs:any in the WSDL —
// was discarded. The raw inner XML must be captured so callers can re-unmarshal
// it into the appropriate service-specific capability type based on Namespace.
func TestServiceCapabilities_PreservesInnerXML(t *testing.T) {
	payload := []byte(`
<GetServicesResponse xmlns="http://www.onvif.org/ver10/device/wsdl">
  <Service>
    <Namespace>http://www.onvif.org/ver10/media/wsdl</Namespace>
    <XAddr>http://192.0.2.10/onvif/media_service</XAddr>
    <Capabilities>
      <trt:Capabilities xmlns:trt="http://www.onvif.org/ver10/media/wsdl" SnapshotUri="true" Rotation="false">
        <trt:ProfileCapabilities MaximumNumberOfProfiles="10"/>
      </trt:Capabilities>
    </Capabilities>
    <Version><Major>2</Major><Minor>40</Minor></Version>
  </Service>
</GetServicesResponse>`)

	var resp GetServicesResponse
	if err := xml.Unmarshal(payload, &resp); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}

	if len(resp.Service) != 1 {
		t.Fatalf("len(Service) = %d, want 1", len(resp.Service))
	}

	raw := resp.Service[0].Capabilities.Any
	if raw == "" {
		t.Fatal("Capabilities.Any is empty; ServiceCapabilities must capture inner XML via xml:\",innerxml\"")
	}
	for _, want := range []string{
		"trt:Capabilities",
		`SnapshotUri="true"`,
		`MaximumNumberOfProfiles="10"`,
	} {
		if !strings.Contains(raw, want) {
			t.Errorf("Capabilities.Any missing %q\n  got: %s", want, raw)
		}
	}

	// Demonstrate the intended consumption pattern: the caller re-unmarshals
	// the raw payload into a service-specific capability type chosen by
	// Service.Namespace.
	var caps struct {
		XMLName     xml.Name `xml:"Capabilities"`
		SnapshotURI string   `xml:"SnapshotUri,attr"`
		Profile     struct {
			MaximumNumberOfProfiles int `xml:"MaximumNumberOfProfiles,attr"`
		} `xml:"ProfileCapabilities"`
	}
	if err := xml.Unmarshal([]byte(raw), &caps); err != nil {
		t.Fatalf("re-unmarshal of Capabilities.Any failed: %v", err)
	}
	if caps.SnapshotURI != "true" {
		t.Errorf("SnapshotUri = %q, want %q", caps.SnapshotURI, "true")
	}
	if caps.Profile.MaximumNumberOfProfiles != 10 {
		t.Errorf("MaximumNumberOfProfiles = %d, want 10", caps.Profile.MaximumNumberOfProfiles)
	}
}
