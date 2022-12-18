// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package device

import (
	"context"
	"github.com/juju/errors"
	"github.com/use-go/onvif/networking"
)

// Call_SetRemoteDiscoveryMode forwards the call to dev.CallMethod() then parses the payload of the reply as a SetRemoteDiscoveryModeResponse.
func Call_SetRemoteDiscoveryMode(ctx context.Context, dev *Device, request SetRemoteDiscoveryMode) (SetRemoteDiscoveryModeResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetRemoteDiscoveryModeResponse SetRemoteDiscoveryModeResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetRemoteDiscoveryModeResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "SetRemoteDiscoveryMode")
		return reply.Body.SetRemoteDiscoveryModeResponse, errors.Annotate(err, "reply")
	}
}
