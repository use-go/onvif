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

// Call_GetDot1XConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a GetDot1XConfigurationResponse.
func Call_GetDot1XConfiguration(ctx context.Context, dev *onvif.Device, request GetDot1XConfiguration) (GetDot1XConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetDot1XConfigurationResponse GetDot1XConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetDot1XConfigurationResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "GetDot1XConfiguration")
		return reply.Body.GetDot1XConfigurationResponse, errors.Annotate(err, "reply")
	}
}
