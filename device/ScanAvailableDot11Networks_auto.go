// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package device

import (
	"context"
	"github.com/juju/errors"
	"github.com/use-go/onvif/networking"
)

// Call_ScanAvailableDot11Networks forwards the call to dev.CallMethod() then parses the payload of the reply as a ScanAvailableDot11NetworksResponse.
func Call_ScanAvailableDot11Networks(ctx context.Context, dev *Device, request ScanAvailableDot11Networks) (ScanAvailableDot11NetworksResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			ScanAvailableDot11NetworksResponse ScanAvailableDot11NetworksResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.ScanAvailableDot11NetworksResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "ScanAvailableDot11Networks")
		return reply.Body.ScanAvailableDot11NetworksResponse, errors.Annotate(err, "reply")
	}
}
