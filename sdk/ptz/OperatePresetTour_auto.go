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

// Call_OperatePresetTour forwards the call to dev.CallMethod() then parses the payload of the reply as a OperatePresetTourResponse.
func Call_OperatePresetTour(ctx context.Context, dev *onvif.Device, request ptz.OperatePresetTour) (ptz.OperatePresetTourResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			OperatePresetTourResponse ptz.OperatePresetTourResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.OperatePresetTourResponse, err
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "OperatePresetTour")
		return reply.Body.OperatePresetTourResponse, err
	}
}
