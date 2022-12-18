// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package device

import (
	"context"
	"github.com/juju/errors"
	"github.com/use-go/onvif/networking"
)

// Call_StartFirmwareUpgrade forwards the call to dev.CallMethod() then parses the payload of the reply as a StartFirmwareUpgradeResponse.
func Call_StartFirmwareUpgrade(ctx context.Context, dev *Device, request StartFirmwareUpgrade) (StartFirmwareUpgradeResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			StartFirmwareUpgradeResponse StartFirmwareUpgradeResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.StartFirmwareUpgradeResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "StartFirmwareUpgrade")
		return reply.Body.StartFirmwareUpgradeResponse, errors.Annotate(err, "reply")
	}
}
