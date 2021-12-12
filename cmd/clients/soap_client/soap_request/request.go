package soap_request

import "encoding/xml"

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	Soapenv string   `xml:"soapenv,attr"`
	Api     string   `xml:"api,attr"`
	Header  string   `xml:"Header"`
	Body    struct {
		Text    string `xml:",chardata"`
		Command struct {
			Text        string `xml:",chardata"`
			Xmlns       string `xml:"xmlns,attr"`
			Credentials struct {
				Text     string `xml:",chardata"`
				Username string `xml:"Username"`
				Password string `xml:"Password"`
			} `xml:"Credentials"`
			Body struct {
				Text      string `xml:",chardata"`
				SOAPDummy struct {
					Text          string `xml:",chardata"`
					SchemaVersion string `xml:"schemaVersion,attr"`
					Xmlns         string `xml:"xmlns,attr"`
					Identity      struct {
						Text       string `xml:",chardata"`
						Title      string `xml:"Title"`
						FirstName  string `xml:"FirstName"`
						MiddleName string `xml:"MiddleName"`
						LastName   string `xml:"LastName"`
						Suffix     string `xml:"Suffix"`
						DOB        string `xml:"DOB"`
						Address    struct {
							Text  string `xml:",chardata"`
							Line1 string `xml:"Line1"`
							Line2 string `xml:"Line2"`
						} `xml:"Address"`
						City        string `xml:"City"`
						State       string `xml:"State"`
						Zip         string `xml:"Zip"`
						MobilePhone string `xml:"MobilePhone"`
					} `xml:"Identity"`
				} `xml:"SOAPDummy"`
			} `xml:"Body"`
		} `xml:"Command"`
	} `xml:"Body"`
}

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
