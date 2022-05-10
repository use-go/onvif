// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package ptz

import (
	"context"
	"github.com/juju/errors"
	"github.com/use-go/onvif"
	"github.com/use-go/onvif/sdk"
	"github.com/use-go/onvif/ptz"
)

// Call_GetPresets forwards the call to dev.CallMethod() then parses the payload of the reply as a GetPresetsResponse.
func Call_GetPresets(ctx context.Context, dev *onvif.Device, request ptz.GetPresets) (ptz.GetPresetsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetPresetsResponse ptz.GetPresetsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetPresetsResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetPresets")
		return reply.Body.GetPresetsResponse, errors.Annotate(err, "reply")
	}
}
