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

// Call_GetProfiles forwards the call to dev.CallMethod() then parses the payload of the reply as a GetProfilesResponse.
func Call_GetProfiles(ctx context.Context, dev *onvif.Device, request GetProfiles) (GetProfilesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetProfilesResponse GetProfilesResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetProfilesResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "GetProfiles")
		return reply.Body.GetProfilesResponse, errors.Annotate(err, "reply")
	}
}
