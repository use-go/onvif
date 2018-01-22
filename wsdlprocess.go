package goonvif

import (
	"github.com/beevik/etree"
	"fmt"
	"strings"
)

const PREFIX = "wsdl:"
const DIVIDER = "**********************************"

const SCHEME = "src/github.com/yakovlevdmv/goonvif/scheme/onvif.go"

func ReadWSDL(path string) {
	doc := etree.NewDocument()
	if err := doc.ReadFromFile(path); err != nil {
		panic(err)
	}
	fmt.Println(DIVIDER)

	fmt.Println("Operations")

	root := doc.SelectElement(PREFIX + "definitions")
	//fmt.Println("ROOT element:", root.Tag)
	//for _, attr := range root.Attr {
	//	fmt.Printf("  ATTR: \n\tprefix:%s\n\tname:%s\n\tnamespace:%s\n", attr.Space, attr.Key, attr.Value)
	//}

	fmt.Println(DIVIDER)

	portType := root.SelectElement(PREFIX + "portType")
	fmt.Println("Port Type: ", portType.SelectAttr("name").Value)

	fmt.Println(DIVIDER)

	counter := 1
	for _, operation := range portType.SelectElements("wsdl:operation") {
		operationName :=operation.SelectAttr("name").Value
		fmt.Println(fmt.Sprintf("%d. %s", counter, operationName))
		counter++

		/**
		Getting Soap Action endpoint
		 */
		fmt.Println("\tSOAP action:")
		fmt.Println("\t" + root.FindElement("//wsdl:binding/wsdl:operation[@name='" + operationName + "']/soap:operation").SelectAttr("soapAction").Value)

		/**
		Обработка вложенных тегов
		 */
		for _, element := range operation.ChildElements() {
			if element.Tag == "documentation" {
				//fmt.Println("\tDescription: " + element.Text())
				fmt.Println("\tDescription: " + element.Text())
			}
			if element.Tag == "input" {
				fmt.Println("\tInput:")
				name := strings.Split(element.SelectAttr("message").Value, ":")[1]
				//fmt.Println("\t[" + name + "]")

				/**
				Get type name from message attribute
				 */
				 typeName := strings.Split(root.FindElement("//wsdl:message[@name='" + name + "']/wsdl:part").SelectAttr("element").Value, ":")[1]
				 fmt.Println("\t[" + typeName + "]")


				 parseType(typeName, path)
			}
			if element.Tag == "output" {
				fmt.Println("\tOutput:")
				name := strings.Split(element.SelectAttr("message").Value, ":")[1]
				//fmt.Println("\t[" + name + "]")

				/**
				Get type name from message attribute
				 */
				typeName := strings.Split(root.FindElement("//wsdl:message[@name='" + name + "']/wsdl:part").SelectAttr("element").Value, ":")[1]
				fmt.Println("\t[" + typeName + "]")
				parseType(typeName, path)
			}
		}
	}
}

func parseType(typeName string, wsdl_path string) {
	doc := etree.NewDocument()
	if err := doc.ReadFromFile(wsdl_path); err != nil {
		panic(err)
	}

	root := doc.SelectElement(PREFIX + "definitions")

	t := root.FindElement("//wsdl:types//xs:element[@name='" + typeName + "']")
	//t := types.FindElement("//xs:element[@name='" + "DeleteGeoLocationRequest" + "']")
	//fmt.Println("//xs:element[@name='" + typeName + "']")
	_type := t.FindElement("//xs:element")
	if _type != nil {
		t_name := strings.Split(_type.SelectAttr("type").Value + "]", ":")[1]
		t_namespace := strings.Split(_type.SelectAttr("type").Value + "]", ":")[0]
		fmt.Println("\t\t" + _type.SelectAttr("name").Value + " [" + t_name)
		fmt.Println("\t\tNamespace of type: " + t_namespace)
		//fmt.Println("\t\t" + _type.SelectAttr("minOccurs").Value)
		//fmt.Println("\t\t" + _type.SelectAttr("maxOccurs").Value)

	}


}
