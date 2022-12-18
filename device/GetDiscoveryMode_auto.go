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

// Call_GetDiscoveryMode forwards the call to dev.CallMethod() then parses the payload of the reply as a GetDiscoveryModeResponse.
func Call_GetDiscoveryMode(ctx context.Context, dev *onvif.Device, request GetDiscoveryMode) (GetDiscoveryModeResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetDiscoveryModeResponse GetDiscoveryModeResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetDiscoveryModeResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "GetDiscoveryMode")
		return reply.Body.GetDiscoveryModeResponse, errors.Annotate(err, "reply")
	}
}
