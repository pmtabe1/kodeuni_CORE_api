package soap_client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/paulmsegeya/pos/utils"
)

var getDefaultSoapTemplate = `<soapenv:Envelope
 xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/"
 xmlns:api="http://soapdummies.com/api">
 <soapenv:Header/>
 <soapenv:Body>
  <api:Command
   xmlns="http://soapdummies.com/api">
   <api:Credentials>
    <api:Username>{{.Username}}</api:Username>
    <api:Password>{{.Password}}</api:Password>
   </api:Credentials>
   <api:Body>
    <SOAPDummy schemaVersion="3.0"
     xmlns="http://soapdummies.com/products/request">
     <Identity>      
      <Title/>
      <FirstName>{{.FirstName}}</FirstName>
      <MiddleName>{{.MiddleName}}</MiddleName>
      <LastName>{{.LastName}}</LastName>
      <Suffix/>
      <DOB>{{.Dob}}</DOB>
      <Address>
       <Line1>{{.AddressLine1}}</Line1>
       <Line2>{{.AddressLine2}}</Line2>
      </Address>
      <City>{{.City}}</City>
      <State>{{.State}}</State>
      <Zip>{{.ZipCode}}</Zip>
      <MobilePhone>{{.MobilePhone}}</MobilePhone>
     </Identity>
    </SOAPDummy>
   </api:Body>
  </api:Command>
 </soapenv:Body>
</soapenv:Envelope>`

type SoapTemplateEntity struct {
	Protocol string `json:"protocol"`
	Template string `json:"template"`
}

func GetSoapRequestTemplate(soapRequestFile string) (soapRequestTemplate string) {

	if soapRequestFile == "" || len(soapRequestFile) == 0 {
		soapRequestTemplate = getDefaultSoapTemplate
	} else {
		soapRequestTemplate = LoadSoapRequestTemplateConfigFile("").Template
	}

	return soapRequestTemplate
}

func GetSoapResponseTemplate(soapResponseFile string) (soapResponseTemplate string) {

	if soapResponseFile == "" || len(soapResponseFile) == 0 {
		soapResponseTemplate = getDefaultSoapTemplate
	} else {
		soapResponseTemplate = LoadSoapRequestTemplateConfigFile("").Template
	}

	return soapResponseTemplate
}

func GetConfiguredSoapTemplatePath() string {
	log.Printf("Getting  Soap Template File Location \n")
	filename := []string{"templates/", "soap_template1", ".json"}
	_, dirname, _, _ := runtime.Caller(0)
	filePath := path.Join(filepath.Dir(dirname), strings.Join(filename, ""))

	fmt.Println(filePath)
	return filePath
}

func BuildJSON(soapTemplateXMLData string) string {

	var templateEntity SoapTemplateEntity
	templateEntity.Protocol = "soap"
	templateEntity.Template = utils.StringToBase64Encoded(soapTemplateXMLData)

	bytes, err := json.Marshal(&templateEntity)

	if err != nil {

		message := "Error while  Marshalling Struct to JSON Bytes :" + err.Error()
		log.Println(message)
	}

	message := "Successfull Marshered Struct to JSON Bytes  with string JSON :" + string(bytes)
	log.Println(message)

	return string(bytes)

}

func CreateSoapConfigurationTemplate(soapTemplateXMLData string) (status bool) {

	filename := GetConfiguredSoapTemplatePath()
	file, err := os.Create(filename)
	if err != nil {
		status = false
		message := "Error while  creating File :" + err.Error()
		log.Println(message)

	}
	status = true
	message := "Successfully created File :" + file.Name()
	log.Println(message)

	// Write content to the created File

	if utils.FileWriter(file, BuildJSON(soapTemplateXMLData)) {
		status = true
		message := "Successfully written data :" + file.Name()
		log.Println(message)
	} else {

		status = false
		message := "Failed to write data to the specified file due to error :" + err.Error()
		log.Println(message)
	}

	return status
}

//LoadConfigFile ...
func LoadSoapRequestTemplateConfigFile(file string) (templateEntity SoapTemplateEntity) {

	if file == "" {
		file = GetConfiguredSoapTemplatePath()
	}

	configMap := make(map[string]string)
	var configData string

	configFile, configError := os.Open(file)

	if configError != nil {
		fmt.Println(configError.Error())
	}

	bytes, bytesErr := ioutil.ReadFile(file)

	if bytesErr != nil {
		message := fmt.Sprintf("Could not load soap template configuration file due to error %v   ", bytesErr)
		log.Fatalf(message)

	}
	configData = string(bytes)

	// save config data to the environmental variable
	os.Setenv("SOAP_TEMPLATE", configData)
	configMap["SOAP_TEMPLATE"] = configData

	defer configFile.Close()

	//Read config values from map

	//configDataJSONString := configMap["configJSON"]
	//configJSON := make(map[string][]AppConfig)
	configDataJSONString := configData
	err := json.Unmarshal([]byte(configDataJSONString), &templateEntity)

	if err != nil {
		message := fmt.Sprintf("Could not Unmarshal JSON data to TemplateEntiy Struct due to error : %v   ", bytesErr)
		log.Fatalf(message)
		panic(err)
	}

	message := fmt.Sprintf("Successfully Unmarshared JSON data to TemplateEntity Struct  giving template data as %v :", utils.Base64EncodedToString(templateEntity.Template))
	log.Println(message)

	// Update template to Decoded String

	templateEntity.Template = utils.Base64EncodedToString(templateEntity.Template)

	fmt.Println(templateEntity.Template)

	return templateEntity
}
