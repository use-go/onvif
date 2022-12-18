// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package device

import (
	"context"
	"github.com/juju/errors"
	"github.com/use-go/onvif/networking"
)

// Call_GetRemoteDiscoveryMode forwards the call to dev.CallMethod() then parses the payload of the reply as a GetRemoteDiscoveryModeResponse.
func Call_GetRemoteDiscoveryMode(ctx context.Context, dev *Device, request GetRemoteDiscoveryMode) (GetRemoteDiscoveryModeResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetRemoteDiscoveryModeResponse GetRemoteDiscoveryModeResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetRemoteDiscoveryModeResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "GetRemoteDiscoveryMode")
		return reply.Body.GetRemoteDiscoveryModeResponse, errors.Annotate(err, "reply")
	}
}
