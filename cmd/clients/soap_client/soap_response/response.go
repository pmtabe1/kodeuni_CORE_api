package soap_response

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/paulmsegeya/pos/cmd/clients/soap_client/soap_request"
)

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	Soap    string   `xml:"soap,attr"`
	Xsd     string   `xml:"xsd,attr"`
	Xsi     string   `xml:"xsi,attr"`
	Body    struct {
		Text     string `xml:",chardata"`
		Response struct {
			Text      string `xml:",chardata"`
			Api       string `xml:"api,attr"`
			Solution  string `xml:"Solution"`
			RequestID string `xml:"RequestID"`
			Body      struct {
				Text          string `xml:",chardata"`
				Status        string `xml:"Status"`
				Salary        string `xml:"Salary"`
				Designation   string `xml:"Designation"`
				Manager       string `xml:"Manager"`
				Company       string `xml:"Company"`
				EmployedSince string `xml:"EmployedSince"`
			} `xml:"Body"`
		} `xml:"Response"`
	} `xml:"Body"`
}

type ISoapResponse interface {
}

type SoapResponse struct {
	*http.Request
}

func (r *SoapResponse) NewSoapResponse() *SoapResponse {

	var targetStruct interface{}

	requestBodyBytes, err := ioutil.ReadAll(r.Response.Body)

	if err != nil {

		message := fmt.Sprintf("Errors while reading request body ERROR:%v", err.Error())
		log.Panicln(message)

	}

	// now you have the bytes turn them to strings

	requestPayload := string(requestBodyBytes)

	if strings.Contains(requestPayload, "SOAPDummy") {

		// trun the matching soap template to STRUCT

		targetStruct = targetStruct.(soap_request.Envelope)

		err = xml.Unmarshal(requestBodyBytes, &targetStruct)

		if err != nil {
			log.Panicln("Error occurred when CONVERTONG SOAP XML bYTES TO STRUCT : ERROR " + err.Error())

		}
	}

	return &SoapResponse{}
}
