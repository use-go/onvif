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

// Call_GetCompatibleMetadataConfigurations forwards the call to dev.CallMethod() then parses the payload of the reply as a GetCompatibleMetadataConfigurationsResponse.
func Call_GetCompatibleMetadataConfigurations(ctx context.Context, dev *onvif.Device, request GetCompatibleMetadataConfigurations) (GetCompatibleMetadataConfigurationsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetCompatibleMetadataConfigurationsResponse GetCompatibleMetadataConfigurationsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetCompatibleMetadataConfigurationsResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "GetCompatibleMetadataConfigurations")
		return reply.Body.GetCompatibleMetadataConfigurationsResponse, errors.Annotate(err, "reply")
	}
}
