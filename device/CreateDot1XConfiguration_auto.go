// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package device

import (
	"context"
	"github.com/juju/errors"
	"github.com/use-go/onvif/networking"
)

// Call_CreateDot1XConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a CreateDot1XConfigurationResponse.
func Call_CreateDot1XConfiguration(ctx context.Context, dev *Device, request CreateDot1XConfiguration) (CreateDot1XConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			CreateDot1XConfigurationResponse CreateDot1XConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.CreateDot1XConfigurationResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "CreateDot1XConfiguration")
		return reply.Body.CreateDot1XConfigurationResponse, errors.Annotate(err, "reply")
	}
}
