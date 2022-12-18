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

// Call_GetStorageConfigurations forwards the call to dev.CallMethod() then parses the payload of the reply as a GetStorageConfigurationsResponse.
func Call_GetStorageConfigurations(ctx context.Context, dev *onvif.Device, request GetStorageConfigurations) (GetStorageConfigurationsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetStorageConfigurationsResponse GetStorageConfigurationsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetStorageConfigurationsResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "GetStorageConfigurations")
		return reply.Body.GetStorageConfigurationsResponse, errors.Annotate(err, "reply")
	}
}
