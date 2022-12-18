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

// Call_CreateProfile forwards the call to dev.CallMethod() then parses the payload of the reply as a CreateProfileResponse.
func Call_CreateProfile(ctx context.Context, dev *device.Device, request CreateProfile) (CreateProfileResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			CreateProfileResponse CreateProfileResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.CreateProfileResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "CreateProfile")
		return reply.Body.CreateProfileResponse, errors.Annotate(err, "reply")
	}
}
