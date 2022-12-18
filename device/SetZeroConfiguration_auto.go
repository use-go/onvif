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

// Call_SetZeroConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a SetZeroConfigurationResponse.
func Call_SetZeroConfiguration(ctx context.Context, dev *onvif.Device, request SetZeroConfiguration) (SetZeroConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetZeroConfigurationResponse SetZeroConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetZeroConfigurationResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "SetZeroConfiguration")
		return reply.Body.SetZeroConfigurationResponse, errors.Annotate(err, "reply")
	}
}
