// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package device

import (
	"context"
	"github.com/juju/errors"
	"github.com/use-go/onvif/networking"
)

// Call_SetStorageConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a SetStorageConfigurationResponse.
func Call_SetStorageConfiguration(ctx context.Context, dev *Device, request SetStorageConfiguration) (SetStorageConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetStorageConfigurationResponse SetStorageConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetStorageConfigurationResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "SetStorageConfiguration")
		return reply.Body.SetStorageConfigurationResponse, errors.Annotate(err, "reply")
	}
}
