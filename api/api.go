package api

import (
	"github.com/yakovlevdmv/goonvif"
	"github.com/yakovlevdmv/goonvif/Device"
	"fmt"
	"reflect"
)

//func RunApi ()  {
//	router := gin.Default()
//
//	router.POST("/:service/:method", func(c *gin.Context) {
//		serviceName := c.Param("service")
//		methodName := c.Param("method")
//		message := callNecessaryMethod(&serviceName, &methodName, "192.168.13.12")
//		c.XML(http.StatusOK, message)
//	})
//	router.Run()
//}

func soapHandling(tp interface{}, tags* map[string]string)  {
	ifaceValue := reflect.ValueOf(tp).Elem()
	typeOfStruct := ifaceValue.Type()
	if ifaceValue.Kind() != reflect.Struct {
		return
	}
	for i := 0; i < ifaceValue.NumField(); i++ {
		field := ifaceValue.Field(i)
		tg, err := typeOfStruct.FieldByName(typeOfStruct.Field(i).Name)
		if err == false {
			fmt.Println(err)
		}
		(*tags)[typeOfStruct.Field(i).Name] = string(tg.Tag)

		subStruct := reflect.New(reflect.TypeOf( field.Interface() ))
		soapHandling(subStruct.Interface(), tags)
	}
}


func callNecessaryMethod(serviceName* string, methodName* string, deviceXaddr* string) *string {
	switch *serviceName {
	case "Device", "device":
		return callDeviceMethods(methodName, deviceXaddr)
	}
	return nil
}

func callDeviceMethods(methodName* string, deviceXaddr* string) *string {
	switch *methodName {
	case "GetCapabilities":
		return GetCapabilities(deviceXaddr)
	default:
		return nil
	}
}

func GetCapabilities(deviceXaddr* string) *string {
	device := goonvif.NewDevice(*deviceXaddr)
	data, err := device.CallMethod(Device.GetCapabilities{Category:"All"})
	if err != nil {
		fmt.Println(err)
	}
	return &data
}