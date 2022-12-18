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

// Call_AddVideoSourceConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a AddVideoSourceConfigurationResponse.
func Call_AddVideoSourceConfiguration(ctx context.Context, dev *onvif.Device, request AddVideoSourceConfiguration) (AddVideoSourceConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			AddVideoSourceConfigurationResponse AddVideoSourceConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.AddVideoSourceConfigurationResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "AddVideoSourceConfiguration")
		return reply.Body.AddVideoSourceConfigurationResponse, errors.Annotate(err, "reply")
	}
}
