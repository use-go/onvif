package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/use-go/onvif"
	"github.com/use-go/onvif/gosoap"
	"github.com/use-go/onvif/networking"
)

func main() {
	RunApi()
}

func RunApi() {
	router := gin.Default()

	router.POST("/:service/:method", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		//c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")

		//serviceName := c.Param("service")
		methodName := c.Param("method")
		username := c.GetHeader("username")
		pass := c.GetHeader("password")
		xaddr := c.GetHeader("xaddr")
		acceptedData, err := c.GetRawData()
		if err != nil {
			fmt.Println(err)
		}
		dev, err := onvif.NewDevice(onvif.DeviceParams{
			Xaddr:    xaddr,
			Username: username,
			Password: pass,
		})
		message, err := CallOnvifCommand(dev, methodName, string(acceptedData))
		//message, err := callNecessaryMethod(serviceName, methodName, string(acceptedData), username, pass, xaddr)
		if err != nil {
			c.XML(http.StatusBadRequest, err.Error())
		} else {
			c.String(http.StatusOK, message)
		}
	})

	router.Run()
}

func CallOnvifCommand(dev *onvif.Device, commandName, acceptedData string) (string, error) {
	requestStruct, err := GetRequestStructByCommand(commandName)
	if err != nil { //done
		return "", err
	}

	if len(acceptedData) > 0 {
		err = json.Unmarshal([]byte(acceptedData), requestStruct)
		if err != nil {
			return "", err
		}
	}

	requestBody, err := xml.Marshal(requestStruct)
	if err != nil {
		return "", err
	}

	endpoint, err := dev.GetEndpointByRequestStruct(requestStruct)
	if err != nil {
		return "", err
	}

	soap := gosoap.NewEmptySOAP()
	soap.AddStringBodyContent(string(requestBody))
	soap.AddRootNamespaces(onvif.Xlmns)
	soap.AddWSSecurity(dev.GetDeviceParams().Username, dev.GetDeviceParams().Password)

	servResp, err := networking.SendSoap(new(http.Client), endpoint, soap.String())
	if err != nil {
		return "", err
	}

	rsp, err := ioutil.ReadAll(servResp.Body)
	if err != nil {
		return "", err
	}

	responseStruct, err := GetResponseStructByCommand(commandName)
	if err != nil {
		return "", err
	}
	responseEnvelope := gosoap.NewSOAPEnvelope(responseStruct)
	err = xml.Unmarshal(rsp, responseEnvelope)
	if err != nil {
		return "", err
	}

	if responseEnvelope.Body.Fault != nil {
		jsonData, err := json.Marshal(responseEnvelope.Body.Fault)
		if err != nil {
			return "", err
		}
		return string(jsonData), nil
	} else {
		jsonData, err := json.Marshal(responseEnvelope.Body.Content)
		if err != nil {
			return "", err
		}
		return string(jsonData), nil
	}

}
