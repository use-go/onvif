// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package ptz

import (
	"context"
	"github.com/juju/errors"
	"github.com/use-go/onvif"
	"github.com/use-go/onvif/networking"
)

// Call_GetConfigurations forwards the call to dev.CallMethod() then parses the payload of the reply as a GetConfigurationsResponse.
func Call_GetConfigurations(ctx context.Context, dev *onvif.Device, request GetConfigurations) (GetConfigurationsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetConfigurationsResponse GetConfigurationsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetConfigurationsResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "GetConfigurations")
		return reply.Body.GetConfigurationsResponse, errors.Annotate(err, "reply")
	}
}
