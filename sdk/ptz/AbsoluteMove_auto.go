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

// Call_AbsoluteMove forwards the call to dev.CallMethod() then parses the payload of the reply as a AbsoluteMoveResponse.
func Call_AbsoluteMove(ctx context.Context, dev *onvif.Device, request ptz.AbsoluteMove) (ptz.AbsoluteMoveResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			AbsoluteMoveResponse ptz.AbsoluteMoveResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.AbsoluteMoveResponse, err
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "AbsoluteMove")
		return reply.Body.AbsoluteMoveResponse, err
	}
}
