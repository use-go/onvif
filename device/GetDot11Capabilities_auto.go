// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package device

import (
	"context"
	"github.com/juju/errors"
	"github.com/use-go/onvif"
	"github.com/use-go/onvif/networking"
)

// Call_GetDot11Capabilities forwards the call to dev.CallMethod() then parses the payload of the reply as a GetDot11CapabilitiesResponse.
func Call_GetDot11Capabilities(ctx context.Context, dev *onvif.Device, request GetDot11Capabilities) (GetDot11CapabilitiesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetDot11CapabilitiesResponse GetDot11CapabilitiesResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetDot11CapabilitiesResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "GetDot11Capabilities")
		return reply.Body.GetDot11CapabilitiesResponse, errors.Annotate(err, "reply")
	}
}
