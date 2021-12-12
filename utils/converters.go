package utils

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func HttpResponseToStruct(response *http.Response) (xmlStructWithData interface{}) {

	byteValue, err := ioutil.ReadAll(response.Body)

	if err != nil {

		log.Panicln("Error reading HTTP Response")
	}

	xmlStructWithData = string(byteValue)

	return xmlStructWithData
}

func XmlToStruct(xmlData string, targetStruct interface{}) (xmlStructWithData interface{}) {

	byteValue, err := ioutil.ReadAll(bytes.NewReader([]byte(xmlData)))

	if err != nil {

		log.Panicln("Error occurred when reading XML bYTES : ERROR " + err.Error())
	}

	err = xml.Unmarshal(byteValue, &targetStruct)

	if err != nil {
		log.Panicln("Error occurred when CONVERTONG XML bYTES TO STRUCT : ERROR " + err.Error())

	}
 
	return targetStruct
}

func StructToXml(xmlStructWithData interface{}) (xmlString string) {

	file, _ := xml.MarshalIndent(xmlStructWithData, "", " ")

	_ = ioutil.WriteFile("data.xml", file, 0644)

	log.Println("========")
	log.Println(string(file))
	log.Println("========")

	return xmlString

}

func StringToBase64Encoded(data string) (result string) {

	//base64.StdEncoding.Encode([]byte(result), []byte(data))
	result = string(base64.StdEncoding.EncodeToString([]byte(data)))
	fmt.Println(result)

	return result
}

func Base64EncodedToString(data string) (result string) {
	resultbytes, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		log.Fatal("error:", err)
	}
	//base64.StdEncoding.Decode([]byte(data), []byte(result))
	result = string(resultbytes)
	return result

}
