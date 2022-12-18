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

// Call_GetVideoSourceModes forwards the call to dev.CallMethod() then parses the payload of the reply as a GetVideoSourceModesResponse.
func Call_GetVideoSourceModes(ctx context.Context, dev *onvif.Device, request GetVideoSourceModes) (GetVideoSourceModesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetVideoSourceModesResponse GetVideoSourceModesResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetVideoSourceModesResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "GetVideoSourceModes")
		return reply.Body.GetVideoSourceModesResponse, errors.Annotate(err, "reply")
	}
}
