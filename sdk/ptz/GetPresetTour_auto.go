// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package ptz

import (
	"context"
	"github.com/use-go/onvif"
	"github.com/use-go/onvif/sdk"
	"github.com/use-go/onvif/ptz"
)

// Call_GetPresetTour forwards the call to dev.CallMethod() then parses the payload of the reply as a GetPresetTourResponse.
func Call_GetPresetTour(ctx context.Context, dev *onvif.Device, request ptz.GetPresetTour) (ptz.GetPresetTourResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetPresetTourResponse ptz.GetPresetTourResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetPresetTourResponse, err
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetPresetTour")
		return reply.Body.GetPresetTourResponse, err
	}
}
