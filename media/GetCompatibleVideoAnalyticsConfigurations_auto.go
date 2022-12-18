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

// Call_GetCompatibleVideoAnalyticsConfigurations forwards the call to dev.CallMethod() then parses the payload of the reply as a GetCompatibleVideoAnalyticsConfigurationsResponse.
func Call_GetCompatibleVideoAnalyticsConfigurations(ctx context.Context, dev *onvif.Device, request GetCompatibleVideoAnalyticsConfigurations) (GetCompatibleVideoAnalyticsConfigurationsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetCompatibleVideoAnalyticsConfigurationsResponse GetCompatibleVideoAnalyticsConfigurationsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetCompatibleVideoAnalyticsConfigurationsResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "GetCompatibleVideoAnalyticsConfigurations")
		return reply.Body.GetCompatibleVideoAnalyticsConfigurationsResponse, errors.Annotate(err, "reply")
	}
}
