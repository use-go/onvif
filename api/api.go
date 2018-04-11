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
	"net/http"
	"io/ioutil"
)

func RunApi ()  {
	router := gin.Default()

	router.POST("/:service/:method", func(c *gin.Context) {
		serviceName := c.Param("service")
		methodName := c.Param("method")
		username := c.GetHeader("username")
		pass := c.GetHeader("password")
		xaddr := c.GetHeader("xaddr")
		acceptedData, err := c.GetRawData()
		if err != nil {
			fmt.Println(err)
		}

		message, err := callNecessaryMethod(serviceName, methodName, string(acceptedData), username, pass, xaddr)
		if err != nil {
			c.XML(http.StatusBadRequest, err.Error())
		} else {
			c.XML(http.StatusOK, message)
		}
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


func callNecessaryMethod(serviceName, methodName, acceptedData, username, password, xaddr string) (string, error) {
	var methodStruct interface{}
	var err error

	switch strings.ToLower(serviceName) {
	case "device":
		methodStruct, err = getDeviceStructByName(methodName)
	case "ptz":
		methodStruct, err = getPTZStructByName(methodName)
	case "media":
		methodStruct, err = getMediaStructByName(methodName)
	default:
		return "", errors.New("there is no such service")
	}
	if err != nil { //done
		return "", err
	}

	resp, err := xmlAnalize(methodStruct, &acceptedData)
	if err != nil {
		return "", err
	}

	endpoint, err := getEndpoint(serviceName, xaddr)
	if err != nil {
		return "", err
	}

	soap := gosoap.NewEmptySOAP()
	soap.AddStringBodyContent(*resp)
	soap.AddRootNamespaces(goonvif.Xlmns)
	soap.AddWSSecurity(username, password)

	servResp, err := networking.SendSoap(endpoint, soap.String())
	if err != nil {
		return "", err
	}

	rsp, err := ioutil.ReadAll(servResp.Body)
	if err != nil {
		return "", err
	}

	return string(rsp), nil
}

func getEndpoint(service, xaddr string)  (string, error) {
	dev, err := goonvif.NewDevice(xaddr)
	if err != nil {
		return "", err
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
	return endpoint, nil
}

func xmlAnalize(methodStruct interface{}, acceptedData* string) (*string, error) {
	test := make([]map[string]string, 0) //tags
	testunMarshal := make([][]interface{}, 0) //data
	var mas []string //idnt

	soapHandling(methodStruct, &test)
	test = mapProcessing(test)

	doc := etree.NewDocument()
	if err := doc.ReadFromString(*acceptedData); err != nil {
		return nil, err
	}
	etr := doc.FindElements("./*")
	xmlUnmarshal(etr, &testunMarshal, &mas)
	ident(&mas)

	document:= etree.NewDocument()
	var el *etree.Element
	var idntIndex = 0

	for lstIndex := 0; lstIndex < len(testunMarshal); {
		lst := (testunMarshal)[lstIndex]
		elemName, attr, value, err := xmlMaker(&lst, &test, lstIndex)
		if err != nil {
			return nil, err
		}

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
		} else if mas[idntIndex] == "PushPop" {
			popTmp := etree.NewElement(elemName)
			popTmp.SetText(value)
			if len(attr) != 0 {
				for key, value := range attr {
					popTmp.CreateAttr(key, value)
				}
			}
			if el == nil {
				document.AddChild(popTmp)
			} else {
				el.AddChild(popTmp)
			}
		} else if mas[idntIndex] == "Pop" {
			el = el.Parent()
			lstIndex  -= 1
		}
		idntIndex += 1
		lstIndex  += 1
	}

	resp, err := document.WriteToString()
	if err != nil {
		return nil, err
	}

	return &resp, err
}

func xmlMaker(lst* []interface{}, tags* []map[string]string, lstIndex int) (string, map[string]string, string, error)  {
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
							return "", nil, "", err
						}
						elemName = res
					} else if index == 0 {
						res, err := xmlProcessing(tg[conversion])
						if err != nil {
							return "", nil, "", err
						}
						elemName = res
					} else {
						value = conversion
					}
				}
			}
		}
	}
	return elemName, attr, value, nil
}

func xmlProcessing (tg string) (string, error) {
	r, _ := regexp.Compile(`"(.*?)"`)
	str := r.FindStringSubmatch(tg)
	if len(str) == 0 {
		return "", errors.New("out of range")
	}
	attr := strings.Index(str[1], ",attr")
	omit := strings.Index(str[1], ",omitempty")
	attrOmit := strings.Index(str[1], ",attr,omitempty")
	omitAttr := strings.Index(str[1], ",omitempty,attr")

	if attr > -1 && attrOmit == -1 && omitAttr == -1 {
		return str[1][0:attr], nil
	} else if omit > -1 && attrOmit == -1 && omitAttr == -1 {
		return str[1][0:omit], nil
	} else if attr == -1 && omit == -1 {
		return str[1], nil
	} else if attrOmit > -1 {
		return str[1][0:attrOmit], nil
	} else {
		return str[1][0:omitAttr], nil
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


func xmlUnmarshal(elems []*etree.Element, data* [][]interface{}, mas* []string) {
	for _, elem := range elems {
		*data = append(*data, []interface{}{elem.Tag,elem.Attr,elem.Text()})
		*mas = append(*mas, "Push")
		xmlUnmarshal(elem.FindElements("./*"), data, mas)
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