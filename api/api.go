package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/yakovlevdmv/goonvif"
	"github.com/yakovlevdmv/goonvif/Device"
	"fmt"
)

func RunApi ()  {
	router := gin.Default()

	device := goonvif.NewDevice()
	data, err := device.CallMethod("http://192.168.13.12/onvif/device_service", Device.GetCapabilities{Category:"All"})
	//_=data
	if err != nil {
		fmt.Println(err)
	}

	router.GET("/:dima", func(c *gin.Context) {
		//name := c.Param("name")
		c.XML(http.StatusOK, data)
	})
	router.Run()
}

func callNecessaryMethod(serviceName, methodName string)  {
	switch serviceName {
	case "Device", "device":
		callDeviceMethonds(methodName)
	}
}

func callDeviceMethonds(methodName string)  {
	
}