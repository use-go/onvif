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

// Call_RemoveVideoAnalyticsConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a RemoveVideoAnalyticsConfigurationResponse.
func Call_RemoveVideoAnalyticsConfiguration(ctx context.Context, dev *onvif.Device, request RemoveVideoAnalyticsConfiguration) (RemoveVideoAnalyticsConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			RemoveVideoAnalyticsConfigurationResponse RemoveVideoAnalyticsConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.RemoveVideoAnalyticsConfigurationResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "RemoveVideoAnalyticsConfiguration")
		return reply.Body.RemoveVideoAnalyticsConfigurationResponse, errors.Annotate(err, "reply")
	}
}
