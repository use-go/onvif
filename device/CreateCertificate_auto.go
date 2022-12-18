// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package device

import (
	"context"
	"github.com/juju/errors"
	"github.com/use-go/onvif/networking"
)

// Call_CreateCertificate forwards the call to dev.CallMethod() then parses the payload of the reply as a CreateCertificateResponse.
func Call_CreateCertificate(ctx context.Context, dev *Device, request CreateCertificate) (CreateCertificateResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			CreateCertificateResponse CreateCertificateResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.CreateCertificateResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "CreateCertificate")
		return reply.Body.CreateCertificateResponse, errors.Annotate(err, "reply")
	}
}
