// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package device

import (
	"context"
	"github.com/juju/errors"
	"github.com/use-go/onvif"
	"github.com/use-go/onvif/networking"
)

// Call_SetClientCertificateMode forwards the call to dev.CallMethod() then parses the payload of the reply as a SetClientCertificateModeResponse.
func Call_SetClientCertificateMode(ctx context.Context, dev *onvif.Device, request SetClientCertificateMode) (SetClientCertificateModeResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetClientCertificateModeResponse SetClientCertificateModeResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetClientCertificateModeResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "SetClientCertificateMode")
		return reply.Body.SetClientCertificateModeResponse, errors.Annotate(err, "reply")
	}
}
