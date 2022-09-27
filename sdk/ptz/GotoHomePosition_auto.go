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

// Call_GotoHomePosition forwards the call to dev.CallMethod() then parses the payload of the reply as a GotoHomePositionResponse.
func Call_GotoHomePosition(ctx context.Context, dev *onvif.Device, request ptz.GotoHomePosition) (ptz.GotoHomePositionResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GotoHomePositionResponse ptz.GotoHomePositionResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GotoHomePositionResponse, err
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GotoHomePosition")
		return reply.Body.GotoHomePositionResponse, err
	}
}
