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

// Call_DeleteStorageConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a DeleteStorageConfigurationResponse.
func Call_DeleteStorageConfiguration(ctx context.Context, dev *onvif.Device, request DeleteStorageConfiguration) (DeleteStorageConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			DeleteStorageConfigurationResponse DeleteStorageConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.DeleteStorageConfigurationResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "DeleteStorageConfiguration")
		return reply.Body.DeleteStorageConfigurationResponse, errors.Annotate(err, "reply")
	}
}
