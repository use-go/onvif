// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package main

import (
	"context"
	"github.com/use-go/onvif/device"
	"github.com/use-go/onvif/media"
)

type MediaOutput struct { 
	AudioDecoderConfigurations               *media.GetAudioDecoderConfigurationsResponse
	AudioEncoderConfigurations               *media.GetAudioEncoderConfigurationsResponse
	AudioOutputConfigurations                *media.GetAudioOutputConfigurationsResponse
	AudioSourceConfigurations                *media.GetAudioSourceConfigurationsResponse
	AudioSources                             *media.GetAudioSourcesResponse
	MetadataConfigurations                   *media.GetMetadataConfigurationsResponse
	OSDs                                     *media.GetOSDsResponse
	Profiles                                 *media.GetProfilesResponse
	ServiceCapabilities                      *media.GetServiceCapabilitiesResponse
	VideoAnalyticsConfigurations             *media.GetVideoAnalyticsConfigurationsResponse
	VideoEncoderConfigurations               *media.GetVideoEncoderConfigurationsResponse
	VideoSourceConfigurations                *media.GetVideoSourceConfigurationsResponse
	VideoSources                             *media.GetVideoSourcesResponse
}

func detailMedia(ctx context.Context, dev *device.Device) MediaOutput {
	var out MediaOutput
	calls := make([]func(c context.Context), 0)

	calls = append(calls, func(c context.Context) {
		if p, err := media.Call_GetAudioDecoderConfigurations(c, dev, media.GetAudioDecoderConfigurations {}); err == nil {
			out.AudioDecoderConfigurations = &p
		} else {
			Logger.Trace().Err(err).Str("rpc", "AudioDecoderConfigurations").Msg("media")
		}
	})

	calls = append(calls, func(c context.Context) {
		if p, err := media.Call_GetAudioEncoderConfigurations(c, dev, media.GetAudioEncoderConfigurations {}); err == nil {
			out.AudioEncoderConfigurations = &p
		} else {
			Logger.Trace().Err(err).Str("rpc", "AudioEncoderConfigurations").Msg("media")
		}
	})

	calls = append(calls, func(c context.Context) {
		if p, err := media.Call_GetAudioOutputConfigurations(c, dev, media.GetAudioOutputConfigurations {}); err == nil {
			out.AudioOutputConfigurations = &p
		} else {
			Logger.Trace().Err(err).Str("rpc", "AudioOutputConfigurations").Msg("media")
		}
	})

	calls = append(calls, func(c context.Context) {
		if p, err := media.Call_GetAudioSourceConfigurations(c, dev, media.GetAudioSourceConfigurations {}); err == nil {
			out.AudioSourceConfigurations = &p
		} else {
			Logger.Trace().Err(err).Str("rpc", "AudioSourceConfigurations").Msg("media")
		}
	})

	calls = append(calls, func(c context.Context) {
		if p, err := media.Call_GetAudioSources(c, dev, media.GetAudioSources {}); err == nil {
			out.AudioSources = &p
		} else {
			Logger.Trace().Err(err).Str("rpc", "AudioSources").Msg("media")
		}
	})

	calls = append(calls, func(c context.Context) {
		if p, err := media.Call_GetMetadataConfigurations(c, dev, media.GetMetadataConfigurations {}); err == nil {
			out.MetadataConfigurations = &p
		} else {
			Logger.Trace().Err(err).Str("rpc", "MetadataConfigurations").Msg("media")
		}
	})

	calls = append(calls, func(c context.Context) {
		if p, err := media.Call_GetOSDs(c, dev, media.GetOSDs {}); err == nil {
			out.OSDs = &p
		} else {
			Logger.Trace().Err(err).Str("rpc", "OSDs").Msg("media")
		}
	})

	calls = append(calls, func(c context.Context) {
		if p, err := media.Call_GetProfiles(c, dev, media.GetProfiles {}); err == nil {
			out.Profiles = &p
		} else {
			Logger.Trace().Err(err).Str("rpc", "Profiles").Msg("media")
		}
	})

	calls = append(calls, func(c context.Context) {
		if p, err := media.Call_GetServiceCapabilities(c, dev, media.GetServiceCapabilities {}); err == nil {
			out.ServiceCapabilities = &p
		} else {
			Logger.Trace().Err(err).Str("rpc", "ServiceCapabilities").Msg("media")
		}
	})

	calls = append(calls, func(c context.Context) {
		if p, err := media.Call_GetVideoAnalyticsConfigurations(c, dev, media.GetVideoAnalyticsConfigurations {}); err == nil {
			out.VideoAnalyticsConfigurations = &p
		} else {
			Logger.Trace().Err(err).Str("rpc", "VideoAnalyticsConfigurations").Msg("media")
		}
	})

	calls = append(calls, func(c context.Context) {
		if p, err := media.Call_GetVideoEncoderConfigurations(c, dev, media.GetVideoEncoderConfigurations {}); err == nil {
			out.VideoEncoderConfigurations = &p
		} else {
			Logger.Trace().Err(err).Str("rpc", "VideoEncoderConfigurations").Msg("media")
		}
	})

	calls = append(calls, func(c context.Context) {
		if p, err := media.Call_GetVideoSourceConfigurations(c, dev, media.GetVideoSourceConfigurations {}); err == nil {
			out.VideoSourceConfigurations = &p
		} else {
			Logger.Trace().Err(err).Str("rpc", "VideoSourceConfigurations").Msg("media")
		}
	})

	calls = append(calls, func(c context.Context) {
		if p, err := media.Call_GetVideoSources(c, dev, media.GetVideoSources {}); err == nil {
			out.VideoSources = &p
		} else {
			Logger.Trace().Err(err).Str("rpc", "VideoSources").Msg("media")
		}
	})

	runAll(ctx, calls...)
	return out
}
