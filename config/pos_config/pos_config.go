package pos_config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	// "path/filepath"
	"strings"

	"github.com/paulmsegeya/pos/cmd/environments"
	"github.com/paulmsegeya/pos/constants/defaults_constants"
)

type PosConfig struct {
	Integration struct {
		RouteSource string `json:"RouteSource"`
		From        struct {
			HostName string `json:"HostName"`
			Server   string `json:"Server"`
			Database struct {
				Source       string `json:"Source"`
				Host         string `json:"Host"`
				Engine       string `json:"Engine"`
				IP           string `json:"IP"`
				Port         string `json:"Port"`
				Username     string `json:"Username"`
				Password     string `json:"Password"`
				SslEnabled   bool   `json:"SslEnabled"`
				SslCertsPath string `json:"SslCertsPath"`
				Schema       string `json:"Schema"`
				DBName       string `json:"DBName"`
			} `json:"Database"`
		} `json:"From"`
		To struct {
			HostName string `json:"HostName"`
			Server   string `json:"Server"`
			Database struct {
				Source       string `json:"Source"`
				Host         string `json:"Host"`
				Engine       string `json:"Engine"`
				IP           string `json:"IP"`
				Port         string `json:"Port"`
				Username     string `json:"Username"`
				Password     string `json:"Password"`
				SslEnabled   bool   `json:"SslEnabled"`
				SslCertsPath string `json:"SslCertsPath"`
				Schema       string `json:"Schema"`
				DBName       string `json:"DBName"`
			} `json:"Database"`
		} `json:"To"`
	} `json:"Integration"`
}

type IPosConfig interface {
}

func New() *PosConfig {

	return &PosConfig{}
}

func (r *PosConfig) LoadConfiguration() *PosConfig {
	configurationNew := New()
	configurationLocal := configurationNew.FromJSON(configurationNew.LoadJSONConfig(""))
	return &configurationLocal
}

func (config *PosConfig) FromJSON(jsonStringData string) PosConfig {

	if len(jsonStringData) == 0 {
		log.Panicln("JSON STRING PROVIDED IS EMPTY OR ITS POINTER IS NIL")
		
		jsonStringData=environments.GetConfiguredConfigurationFile()
	}

	log.Println("Received JSONString :" + jsonStringData)

	err := json.Unmarshal([]byte(jsonStringData), &config)

	if err != nil {

		if strings.Contains(err.Error(), "unexpected end of JSON input") {

			log.Println("Failed to get configuration file")

		}
		log.Panicln(err.Error())

		log.Panicln("Error converting json string to struct")
	}
	log.Printf("%v", config)

	return *config
}
func (config *PosConfig) ToJSON() string {
	jsonBytes, _ := json.Marshal(config)
	x := fmt.Sprintf("%v", string(jsonBytes))
	return x
}

func GetFilepathTemplate() string  {
	
	return fmt.Sprintf(defaults_constants.CONFIG_FILE_LOCATION, "sage", "%v")
}

func (config *PosConfig) LoadJSONConfig(filePathLocation string) string {

	var jsonFileLocation string

	var fineractConfigFileLocation string

	if len(filePathLocation) == 0 {
		// DEFAULT TO DEV MODE CONFIGURATION
		fineractConfigFileLocation = os.Getenv("FINERACT_CONFIG")

		// Read from Environmental Config

		deployment := os.Getenv("DEPLOYMENT_STATUS")

		var jsonFileLocationTemplate string

		if len(deployment) == 0 || deployment == "dev" {
			log.Println("SET FINERACT_CONFIG default Dev path")
			deployment = "dev"

			if len(fineractConfigFileLocation) == 0 {
				jsonFileLocationTemplate = fmt.Sprintf(defaults_constants.CONFIG_FILE_LOCATION, "pos", "%v")
				jsonFileLocation = fmt.Sprintf(jsonFileLocationTemplate, deployment)
			} else {
				jsonFileLocationTemplate = fineractConfigFileLocation
				jsonFileLocation = fmt.Sprintf(jsonFileLocationTemplate, deployment)
			}

		}
		if deployment == "prod" {
			deployment = "prod"
			if len(fineractConfigFileLocation) == 0 {
				jsonFileLocationTemplate = fmt.Sprintf(defaults_constants.CONFIG_FILE_LOCATION, "sage", "%v")
				jsonFileLocation = fmt.Sprintf(jsonFileLocationTemplate, deployment)
			} else {
				jsonFileLocationTemplate = fineractConfigFileLocation
				jsonFileLocation = fmt.Sprintf(jsonFileLocationTemplate, deployment)
			}

		} else if deployment == "stagging" {
			deployment = "stagging"
			if len(fineractConfigFileLocation) == 0 {
				jsonFileLocationTemplate = fmt.Sprintf(defaults_constants.CONFIG_FILE_LOCATION, "sage", "%v")
				jsonFileLocation = fmt.Sprintf(jsonFileLocationTemplate, deployment)
			} else {
				jsonFileLocationTemplate = fineractConfigFileLocation
				jsonFileLocation = fmt.Sprintf(jsonFileLocationTemplate, deployment)
			}
		} else if deployment == "uat" {
			deployment = "uat"
			if len(fineractConfigFileLocation) == 0 {
				jsonFileLocationTemplate = fmt.Sprintf(defaults_constants.CONFIG_FILE_LOCATION, "sage", "%v")
				jsonFileLocation = fmt.Sprintf(jsonFileLocationTemplate, deployment)
			} else {
				jsonFileLocationTemplate = fineractConfigFileLocation
				jsonFileLocation = fmt.Sprintf(jsonFileLocationTemplate, deployment)
			}
		}

	} else {
		jsonFileLocation = filePathLocation
	}

	var JsonFileContents string

	if len(jsonFileLocation) == 0 {
		log.Panicln("JSONFILE LOCATION IS EMPTY | FIX IT")
	} else {
		log.Println("ONPENT JSON FILE FROM LOCATION >>" + jsonFileLocation)

		// Open our jsonFile
		jsonFile, err := os.Open(jsonFileLocation)
		// if we os.Open returns an error then handle it
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println("Successfully Opened >> " + jsonFileLocation)

		bytez, _ := ioutil.ReadFile(jsonFileLocation)

		JsonFileContents = fmt.Sprintf("%v", string(bytez))

		// defer the closing of our jsonFile so that we can parse it later on
		defer jsonFile.Close()
	}

	return JsonFileContents

}
func (config *PosConfig) ReadConfiguration() PosConfig {
	jsonConfiguration := config.LoadJSONConfig("")
	log.Println(jsonConfiguration + "  Reading configuration" + jsonConfiguration)
	configuration := (config.FromJSON(jsonConfiguration))
	log.Println(configuration)

	return configuration
}
