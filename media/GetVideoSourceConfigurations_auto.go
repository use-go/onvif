// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package media

import (
	"context"
	"github.com/juju/errors"
	"github.com/use-go/onvif"
	"github.com/use-go/onvif/networking"
)

// Call_GetVideoSourceConfigurations forwards the call to dev.CallMethod() then parses the payload of the reply as a GetVideoSourceConfigurationsResponse.
func Call_GetVideoSourceConfigurations(ctx context.Context, dev *onvif.Device, request GetVideoSourceConfigurations) (GetVideoSourceConfigurationsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetVideoSourceConfigurationsResponse GetVideoSourceConfigurationsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetVideoSourceConfigurationsResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "GetVideoSourceConfigurations")
		return reply.Body.GetVideoSourceConfigurationsResponse, errors.Annotate(err, "reply")
	}
}
