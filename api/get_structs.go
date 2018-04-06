package api

import (
	"github.com/yakovlevdmv/goonvif/PTZ"
	"errors"
)

func GetPTZStructByName(name string) (interface{}, error) {
	switch name {
	case "GetServiceCapabilities":
		return &PTZ.GetServiceCapabilities{}, nil
	case "GetNodes":
		return &PTZ.GetNodes{}, nil
	case "GetNode":
		return &PTZ.GetNode{}, nil
	case "GetConfiguration":
		return &PTZ.GetConfiguration{}, nil
	case "GetConfigurations":
		return &PTZ.GetConfigurations{}, nil
	case "SetConfiguration":
		return &PTZ.SetConfiguration{}, nil
	case "GetConfigurationOptions":
		return &PTZ.GetConfigurationOptions{}, nil
	case "SendAuxiliaryCommand":
		return &PTZ.SendAuxiliaryCommand{}, nil
	case "GetPresets":
		return &PTZ.GetPresets{}, nil
	case "SetPreset":
		return &PTZ.SetPreset{}, nil
	case "RemovePreset":
		return &PTZ.RemovePreset{}, nil
	case "GotoPreset":
		return &PTZ.GotoPreset{}, nil
	case "GotoHomePosition":
		return &PTZ.GotoHomePosition{}, nil
	case "SetHomePosition":
		return &PTZ.SetHomePosition{}, nil
	case "ContinuousMove":
		return &PTZ.ContinuousMove{}, nil
	case "RelativeMove":
		return &PTZ.RelativeMove{}, nil
	case "GetStatus":
		return &PTZ.GetStatus{}, nil
	case "AbsoluteMove":
		return &PTZ.AbsoluteMove{}, nil
	case "GeoMove":
		return &PTZ.GeoMove{}, nil
	case "Stop":
		return &PTZ.Stop{}, nil
	case "GetPresetTours":
		return &PTZ.GetPresetTours{}, nil
	case "GetPresetTour":
		return &PTZ.GetPresetTour{}, nil
	case "GetPresetTourOptions":
		return &PTZ.GetPresetTourOptions{}, nil
	case "CreatePresetTour":
		return &PTZ.CreatePresetTour{}, nil
	case "ModifyPresetTour":
		return &PTZ.ModifyPresetTour{}, nil
	case "OperatePresetTour":
		return &PTZ.OperatePresetTour{}, nil
	case "RemovePresetTour":
		return &PTZ.RemovePresetTour{}, nil
	case "GetCompatibleConfigurations":
		return &PTZ.GetCompatibleConfigurations{}, nil
	default:
		return nil, errors.New("invalid structure")
	}
}