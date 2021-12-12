package soap_client

import (
	"bytes"
	"encoding/xml"
	"fmt"

	//"io/ioutil"
	"net/http"
	"text/template"

	"github.com/paulmsegeya/pos/cmd/clients/soap_client/soap_request"
)

type ISoapRequest interface {
}

type SoapRequest struct {
	//Values are set in below fields as per the request
	FirstName    string
	LastName     string
	MiddleName   string
	AddressLine1 string
	AddressLine2 string
	City         string
	State        string
	ZipCode      string
	MobilePhone  string
	SSN          string
	Dob          string
	Username     string
	Password     string
}

func NewSoapRequest() SoapRequest {
	req := SoapRequest{}
	req.FirstName = "Tony"
	req.MiddleName = ""
	req.LastName = "Blaire"
	req.Dob = "1946-08-08"
	req.AddressLine1 = "866 Atlas Dr"
	req.AddressLine2 = "Apt 999"
	req.City = "London"
	req.State = "England"
	req.ZipCode = "SW15 5PU"
	req.MobilePhone = "9876543210"
	req.Username = "tony1"
	req.Password = "password1"

	return req

}

func (r *SoapRequest) PopulateRequest(soapRequest SoapRequest) (populatedSoapRequest *SoapRequest) {

	req := SoapRequest{}

	detectedRequest := soapRequest

	req.FirstName = detectedRequest.FirstName
	req.MiddleName = detectedRequest.MiddleName
	req.LastName = detectedRequest.LastName
	req.Dob = detectedRequest.Dob
	req.AddressLine1 = detectedRequest.AddressLine1
	req.AddressLine2 = detectedRequest.AddressLine2
	req.City = detectedRequest.City
	req.State = detectedRequest.State
	req.ZipCode = detectedRequest.ZipCode
	req.MobilePhone = detectedRequest.MobilePhone
	req.Username = detectedRequest.Username
	req.Password = detectedRequest.Password

	populatedSoapRequest = &req

	return populatedSoapRequest
}

func (r *SoapRequest) GenerateSOAPRequest(req *soap_request.Envelope, soapEndpointUrl string) (*http.Request, error) {
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

// func (r *SoapRequest) SoapCall(req *http.Request) (*soap_response.Envelope, error) {
// 	client := &http.Client{}
// 	resp, err := client.Do(req)

// 	if err != nil {
// 		return nil, err
// 	}

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()

// 	response := &soap_response.Envelope{}
// 	err = xml.Unmarshal(body, &r)

// 	if err != nil {
// 		return nil, err
// 	}

// 	if response.Body.Response.Body.Status != "200" {
// 		return nil, err
// 	}

// 	return response, nil
// }
