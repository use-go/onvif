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

// Call_SetMetadataConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a SetMetadataConfigurationResponse.
func Call_SetMetadataConfiguration(ctx context.Context, dev *onvif.Device, request SetMetadataConfiguration) (SetMetadataConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetMetadataConfigurationResponse SetMetadataConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetMetadataConfigurationResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "SetMetadataConfiguration")
		return reply.Body.SetMetadataConfigurationResponse, errors.Annotate(err, "reply")
	}
}
