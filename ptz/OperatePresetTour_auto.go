// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package ptz

import (
	"context"
	"github.com/juju/errors"
	"github.com/use-go/onvif/device"
	"github.com/use-go/onvif/networking"
)

// Call_OperatePresetTour forwards the call to dev.CallMethod() then parses the payload of the reply as a OperatePresetTourResponse.
func Call_OperatePresetTour(ctx context.Context, dev *device.Device, request OperatePresetTour) (OperatePresetTourResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			OperatePresetTourResponse OperatePresetTourResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.OperatePresetTourResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "OperatePresetTour")
		return reply.Body.OperatePresetTourResponse, errors.Annotate(err, "reply")
	}
}
