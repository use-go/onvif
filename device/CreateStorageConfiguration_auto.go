// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package device

import (
	"context"
	"github.com/juju/errors"
	"github.com/use-go/onvif/networking"
)

// Call_CreateStorageConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a CreateStorageConfigurationResponse.
func Call_CreateStorageConfiguration(ctx context.Context, dev *Device, request CreateStorageConfiguration) (CreateStorageConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			CreateStorageConfigurationResponse CreateStorageConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.CreateStorageConfigurationResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "CreateStorageConfiguration")
		return reply.Body.CreateStorageConfigurationResponse, errors.Annotate(err, "reply")
	}
}
