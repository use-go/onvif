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

// Call_SetHomePosition forwards the call to dev.CallMethod() then parses the payload of the reply as a SetHomePositionResponse.
func Call_SetHomePosition(ctx context.Context, dev *onvif.Device, request ptz.SetHomePosition) (ptz.SetHomePositionResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetHomePositionResponse ptz.SetHomePositionResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetHomePositionResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "SetHomePosition")
		return reply.Body.SetHomePositionResponse, errors.Annotate(err, "reply")
	}
}
