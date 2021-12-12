package soap_client

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"

	//"io/ioutil"
	"net/http"
	"text/template"

	"github.com/paulmsegeya/pos/cmd/clients/soap_client/soap_request"
	"github.com/paulmsegeya/pos/cmd/clients/soap_client/soap_response"
)

type ISoapClient interface {
	NewSoapClient()
}
type SoapClient struct {
	SoapRequest soap_request.SoapRequest
}

func NewSoapClient() *SoapClient {

	soapRequest := soap_request.NewSoapRequest()

	return &SoapClient{SoapRequest: soapRequest}
}

func (s *SoapClient) SoapCaller() {

	//s.SoapRequest.SoapCall(&s.SoapRequest)

	// req := populateRequest()

	// httpReq, err := generateSOAPRequest(req)
	// if err != nil {
	// 	fmt.Println("Some problem occurred in request generation")
	// }

	// response, err := soapCall(httpReq)
	// if err != nil {
	// 	fmt.Println("Problem occurred in making a SOAP call")
	// }
}

func (r *SoapClient) GenerateSOAPRequest(req *soap_request.Envelope, soapEndpointUrl string) (*http.Request, error) {
	// Using the var getTemplate to construct request
	template, err := template.New("InputRequest").Parse(getDefaultSoapTemplate)
	if err != nil {
		fmt.Printf("Error while marshling object. %s ", err.Error())
		return nil, err
	}

	doc := &bytes.Buffer{}
	// Replacing the doc from template with actual req values
	err = template.Execute(doc, req)
	if err != nil {
		fmt.Printf("template.Execute error. %s ", err.Error())
		return nil, err
	}

	buffer := &bytes.Buffer{}
	encoder := xml.NewEncoder(buffer)
	err = encoder.Encode(doc.String())
	if err != nil {
		fmt.Printf("encoder.Encode error. %s ", err.Error())
		return nil, err
	}

	response, err := http.NewRequest(http.MethodPost, soapEndpointUrl, bytes.NewBuffer([]byte(doc.String())))
	if err != nil {
		fmt.Printf("Error making a request. %s ", err.Error())
		return nil, err
	}

	return response, nil
}

func (r *SoapClient) SoapCall(req *http.Request) (*soap_response.Envelope, error) {
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	response := &soap_response.Envelope{}
	err = xml.Unmarshal(body, &r)

	if err != nil {
		return nil, err
	}

	if response.Body.Response.Body.Status != "200" {
		return nil, err
	}

	return response, nil
}
