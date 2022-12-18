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

// Call_GetMetadataConfigurationOptions forwards the call to dev.CallMethod() then parses the payload of the reply as a GetMetadataConfigurationOptionsResponse.
func Call_GetMetadataConfigurationOptions(ctx context.Context, dev *device.Device, request GetMetadataConfigurationOptions) (GetMetadataConfigurationOptionsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetMetadataConfigurationOptionsResponse GetMetadataConfigurationOptionsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetMetadataConfigurationOptionsResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "GetMetadataConfigurationOptions")
		return reply.Body.GetMetadataConfigurationOptionsResponse, errors.Annotate(err, "reply")
	}
}
