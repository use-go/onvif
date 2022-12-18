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

// Call_GetOSDs forwards the call to dev.CallMethod() then parses the payload of the reply as a GetOSDsResponse.
func Call_GetOSDs(ctx context.Context, dev *onvif.Device, request GetOSDs) (GetOSDsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetOSDsResponse GetOSDsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetOSDsResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "GetOSDs")
		return reply.Body.GetOSDsResponse, errors.Annotate(err, "reply")
	}
}
