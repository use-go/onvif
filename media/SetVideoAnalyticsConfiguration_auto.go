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

// Call_SetVideoAnalyticsConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a SetVideoAnalyticsConfigurationResponse.
func Call_SetVideoAnalyticsConfiguration(ctx context.Context, dev *device.Device, request SetVideoAnalyticsConfiguration) (SetVideoAnalyticsConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetVideoAnalyticsConfigurationResponse SetVideoAnalyticsConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetVideoAnalyticsConfigurationResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "SetVideoAnalyticsConfiguration")
		return reply.Body.SetVideoAnalyticsConfigurationResponse, errors.Annotate(err, "reply")
	}
}
