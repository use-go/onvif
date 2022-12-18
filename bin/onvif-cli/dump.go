// Copyright (c) 2022-2022 Jean-Francois SMIGIELSKI

package main

import (
	"context"
	"encoding/json"
	"os"
	"sync"

	"github.com/juju/errors"
	"github.com/use-go/onvif/device"
)

// Used once to generate the first skeleton of code
//#go:generate go run github.com/use-go/onvif/bin/onvif-codegen cli ptz    ../../ptz/calls.txt
//#go:generate go run github.com/use-go/onvif/bin/onvif-codegen cli device ../../device/calls.txt
//#go:generate go run github.com/use-go/onvif/bin/onvif-codegen cli media  ../../media/calls.txt
//#go:generate go run github.com/use-go/onvif/bin/onvif-codegen cli event  ../../event/calls.txt

type OnvifOutput struct {
	Endpoint string
	Device   DeviceOutput
	Media    MediaOutput
	Ptz      PtzOutput
}

func details(ctx context.Context, endpoint string) error {

	dev, err := device.NewDevice(device.DeviceParams{
		Xaddr:    endpoint,
		Username: "admin",
		Password: "ollyhgqo",
	})
	if err != nil {
		return errors.Trace(err)
	}

	out := OnvifOutput{}
	out.Endpoint = endpoint

	wg := sync.WaitGroup{}
	wg.Add(3)
	doSync := func(wg *sync.WaitGroup, fn func()) {
		defer wg.Done()
		fn()
	}
	go doSync(&wg, func() { out.Device = detailDevice(ctx, dev) })
	go doSync(&wg, func() { out.Media = detailMedia(ctx, dev) })
	go doSync(&wg, func() { out.Ptz = detailPtz(ctx, dev) })
	wg.Wait()

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(out)
}
