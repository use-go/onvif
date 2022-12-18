// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package media

import (
	"context"
	"github.com/juju/errors"
	"github.com/use-go/onvif/device"
	"github.com/use-go/onvif/networking"
)

// Call_GetCompatibleVideoSourceConfigurations forwards the call to dev.CallMethod() then parses the payload of the reply as a GetCompatibleVideoSourceConfigurationsResponse.
func Call_GetCompatibleVideoSourceConfigurations(ctx context.Context, dev *device.Device, request GetCompatibleVideoSourceConfigurations) (GetCompatibleVideoSourceConfigurationsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetCompatibleVideoSourceConfigurationsResponse GetCompatibleVideoSourceConfigurationsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetCompatibleVideoSourceConfigurationsResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "GetCompatibleVideoSourceConfigurations")
		return reply.Body.GetCompatibleVideoSourceConfigurationsResponse, errors.Annotate(err, "reply")
	}
}
