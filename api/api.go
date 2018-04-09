package api

import (
	"fmt"
	"reflect"
	"regexp"
	"errors"
	"strings"
	"github.com/beevik/etree"
	"github.com/gin-gonic/gin"
	"github.com/yakovlevdmv/gosoap"
	"github.com/yakovlevdmv/goonvif"
	"github.com/yakovlevdmv/goonvif/networking"
	"log"
	"net/http"
	"io/ioutil"
	"encoding/xml"
)

func RunApi ()  {
	router := gin.Default()

	router.POST("/:service/:method", func(c *gin.Context) {
		serviceName := c.Param("service")
		methodName := c.Param("method")
		acceptedData, err := c.GetRawData()
		if err != nil {
			fmt.Println(err)
		}
		message := callNecessaryMethod(serviceName, methodName, string(acceptedData), "192.168.13.12")
		c.XML(http.StatusOK, message)
	})
	router.Run()
}

//func soapHandling(tp interface{}, tags* map[string]string)  {
//	ifaceValue := reflect.ValueOf(tp).Elem()
//	typeOfStruct := ifaceValue.Type()
//	if ifaceValue.Kind() != reflect.Struct {
//		return
//	}
//	for i := 0; i < ifaceValue.NumField(); i++ {
//		field := ifaceValue.Field(i)
//		tg, err := typeOfStruct.FieldByName(typeOfStruct.Field(i).Name)
//		if err == false {
//			fmt.Println(err)
//		}
//		(*tags)[typeOfStruct.Field(i).Name] = string(tg.Tag)
//
//		subStruct := reflect.New(reflect.TypeOf( field.Interface() ))
//		soapHandling(subStruct.Interface(), tags)
//	}
//}


func callNecessaryMethod(serviceName string, methodName string, acceptedData string, deviceXaddr string) string {
	var methodStruct interface{}
	var err error
	switch serviceName {
	case "Device", "device":
		methodStruct, err = GetDeviceStructByName(methodName)
	case "PTZ", "ptz":
		methodStruct, err = GetPTZStructByName(methodName)
	}
	if err != nil {
		//todo: нормально написать
	}
	resp, err := xmlAnalize(methodStruct, &acceptedData)
	if err != nil {
		//todo: нормально написать
	}

	soap := gosoap.NewEmptySOAP()
	soap.AddStringBodyContent(*resp)
	soap.AddRootNamespaces(goonvif.Xlmns)

	/*
	Getting an WS-Security struct representation
	 */
	auth := goonvif.NewSecurity("admin", "Supervisor")

	/*
	Adding WS-Security namespaces to root element of SOAP message
	 */
	soap.AddRootNamespace("wsse", "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext1.0.xsd")
	soap.AddRootNamespace("wsu", "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility1.0.xsd")

	soapReq, err := xml.MarshalIndent(auth, "", "  ")

	/*
	Adding WS-Security struct to SOAP header
	 */
	soap.AddStringHeaderContent(string(soapReq))

	endpoint:= getEndpoint(serviceName, deviceXaddr)
	servResp, srvErr := networking.SendSoap(endpoint, soap.String())
	if srvErr != nil {
		panic(srvErr)
		//todo: нормально написать
	}

	rsp, err := ioutil.ReadAll(servResp.Body)
	if err != nil {
		log.Println(err)
		return ""
	}
	return string(rsp)
}

func getEndpoint(service, xaddr string)  string {
	dev, err := goonvif.NewDevice(xaddr)
	if err != nil {
		log.Println(err)
		return ""
	}
	pkg := strings.ToLower(service)

	var endpoint string
	switch pkg {
		case "device": endpoint = dev.GetEndpoint("Device")
		case "event": endpoint = dev.GetEndpoint("Event")
		case "imaging": endpoint = dev.GetEndpoint("Imaging")
		case "media": endpoint = dev.GetEndpoint("Media")
		case "ptz": endpoint = dev.GetEndpoint("PTZ")
	}
	return endpoint
}

/*
NEW
 */

func xmlAnalize(methodStruct interface{}, acceptedData* string) (*string, error) {
	test := make([]map[string]string, 0) //tags
	testunMarshal := make([][]interface{}, 0) //data
	//tmp, err := GetPTZStructByName("Stop")
	//if err != nil {
	//	fmt.Println(err)
	//}

	soapHandling(methodStruct, &test)
	test = mapProcessing(test)

	doc := etree.NewDocument()
	if err := doc.ReadFromString(*acceptedData); err != nil {
		fmt.Println(err)
		return nil, nil //todo: нормально написать
	}
	etr := doc.FindElements("./*")
	xmlUnmarshal(etr, &testunMarshal)

	etr_ := doc.FindElements("./*")
	var mas []string //idnt
	getPos(etr_, &mas)
	ident(&mas)

	//todo: может возникнуть проблема с типами данных без тегов. От таких типов надо избавляться
	//todo: попробовать всунуть все вызовы (soapHandling, xmlUnmarshal)сюда


	document:= etree.NewDocument()
	var el *etree.Element
	var idntIndex = 0

	for lstIndex := 0; lstIndex < len(testunMarshal); {
		lst := (testunMarshal)[lstIndex]
		elemName, attr, value:= xmlMaker(&lst, &test, lstIndex)

		if mas[lstIndex] == "Push" && lstIndex == 0 { //done
			el = document.CreateElement(elemName)
			el.SetText(value)
			if len(attr) != 0 {
				for key, value := range attr {
					el.CreateAttr(key, value)
				}
			}
		} else if mas[idntIndex] == "Push" {
			pushTmp := etree.NewElement(elemName)
			pushTmp.SetText(value)
			if len(attr) != 0 {
				for key, value := range attr {
					pushTmp.CreateAttr(key, value)
				}
			}
			el.AddChild(pushTmp)
			el = pushTmp
		} else if mas[idntIndex] == "PushPop" { //done
			popTmp := etree.NewElement(elemName)
			popTmp.SetText(value)
			if len(attr) != 0 {
				for key, value := range attr {
					popTmp.CreateAttr(key, value)
				}
			}
			el.AddChild(popTmp)
		} else if mas[idntIndex] == "Pop" {
			el = el.Parent()
			lstIndex  -= 1
		}
		idntIndex += 1
		lstIndex  += 1
	}
	//document.Indent(2)
	//fmt.Println(document.WriteToString())
	resp, err := document.WriteToString()
	if err != nil {
		//todo нормально расписать
	}
	return &resp, err
}

func xmlMaker(lst* []interface{}, tags* []map[string]string, lstIndex int) (string, map[string]string, string)  {
	var elemName, value string
	attr := make(map[string]string)
	for tgIndx, tg := range *tags {
		if tgIndx == lstIndex {
			for index, elem := range *lst {
				if reflect.TypeOf(elem).String() == "[]etree.Attr" {
					conversion := elem.([]etree.Attr)
					for _, i := range conversion {
						attr[i.Key] = i.Value
					}
				} else {
					conversion := elem.(string)
					if index == 0 && lstIndex == 0 {
						res, err := xmlProcessing(tg["XMLName"])
						if err != nil {
							fmt.Println(err)
						}
						elemName = res
					} else if index == 0 {
						res, err := xmlProcessing(tg[conversion])
						if err != nil {
							fmt.Println(err)
						}
						elemName = res
					} else {
						value = conversion
					}
				}
			}
		}
	}
	return elemName, attr, value
}

func xmlProcessing (tg string) (string, error) {
	r, _ := regexp.Compile(`\"(.*?)\"`)
	str := r.FindStringSubmatch(tg)
	if len(str) == 0 {
		return "", errors.New("out of range")
	}
	attr := strings.Index(str[1], ",attr")
	omit := strings.Index(str[1], ",omitempty")
	if attr == -1 && omit == -1 { //todo: проработать вариант, когдв omitempty и attr вместе
		return str[1], nil
	} else if omit > -1 {
		return str[1][0:omit], nil
	}else {
		return str[0:attr][0], nil
	}
	return "", errors.New("something went wrong")
}

func mapProcessing(mapVar []map[string]string) []map[string]string {
	for indx := 0; indx < len(mapVar); indx++ {
		element := mapVar[indx]
		for _, value := range element {
			if value == "" {
				mapVar = append(mapVar[:indx], mapVar[indx+1:]...)
				indx--
			}
			if strings.Index(value, ",attr") != -1 {
				mapVar = append(mapVar[:indx], mapVar[indx+1:]...)
				indx--
			}
		}
	}
	return mapVar
}

func soapHandling(tp interface{}, tags* []map[string]string)  {
	s := reflect.ValueOf(tp).Elem()
	typeOfT := s.Type()
	if s.Kind() != reflect.Struct {
		return
	}
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		tmp, err := typeOfT.FieldByName(typeOfT.Field(i).Name)
		if err == false {
			fmt.Println(err)
		}
		*tags = append(*tags, map[string]string{typeOfT.Field(i).Name : string(tmp.Tag)})
		subStruct := reflect.New(reflect.TypeOf( f.Interface() ))
		soapHandling(subStruct.Interface(), tags)
	}
}

//todo: передавать в него отступы
func xmlUnmarshal(elems []*etree.Element, data* [][]interface{}) {
	for _, elem := range elems {
		*data = append(*data, []interface{}{elem.Tag,elem.Attr,elem.Text()})
		xmlUnmarshal(elem.FindElements("./*"), data)
	}
}


//todo лишняя функция, засунуть в xmlUnmarshal
func getPos(elems []*etree.Element,  mas* []string)  {
	for _, elem := range elems {
		*mas = append(*mas, "Push")
		getPos(elem.FindElements("./*"), mas)
		*mas = append(*mas, "Pop")
	}
}

func ident(mas* []string)  {
	var buffer string
	for _, j := range *mas {
		buffer += j + " "
	}
	buffer = strings.Replace(buffer, "Push Pop ", "PushPop ", -1)
	buffer = strings.TrimSpace(buffer)
	*mas = strings.Split(buffer, " ")
}