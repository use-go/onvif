// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package device

import (
	"context"
	"github.com/juju/errors"
	"github.com/use-go/onvif/networking"
)

// Call_GetRelayOutputs forwards the call to dev.CallMethod() then parses the payload of the reply as a GetRelayOutputsResponse.
func Call_GetRelayOutputs(ctx context.Context, dev *Device, request GetRelayOutputs) (GetRelayOutputsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetRelayOutputsResponse GetRelayOutputsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetRelayOutputsResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "GetRelayOutputs")
		return reply.Body.GetRelayOutputsResponse, errors.Annotate(err, "reply")
	}
}
