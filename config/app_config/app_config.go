package app_config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	// "path/filepath"
	"strings"

	"github.com/paulmsegeya/subscription/cmd/environments"
	"github.com/paulmsegeya/subscription/constants/app_constants"
	"github.com/paulmsegeya/subscription/constants/defaults_constants"
)

type AppConfig struct {
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

type IAppConfig interface {
}

func New() *AppConfig {

	return &AppConfig{}
}

func (r *AppConfig) LoadConfiguration() *AppConfig {
	configurationNew := New()
	configurationLocal := configurationNew.FromJSON(configurationNew.LoadJSONConfig(""))
	return &configurationLocal
}

func (config *AppConfig) FromJSON(jsonStringData string) AppConfig {

	if len(jsonStringData) == 0 {
		log.Panicln("JSON STRING PROVIDED IS EMPTY OR ITS POINTER IS NIL")

		jsonStringData = environments.GetConfiguredConfigurationFile()
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
func (config *AppConfig) ToJSON() string {
	jsonBytes, _ := json.Marshal(config)
	x := fmt.Sprintf("%v", string(jsonBytes))
	return x
}

func GetFilepathTemplate() string {

	return fmt.Sprintf(defaults_constants.CONFIG_FILE_LOCATION, app_constants.AppName, "%v")
}

func (config *AppConfig) LoadJSONConfig(filePathLocation string) string {

	var jsonFileLocation string

	var configFileLocation string

	if len(filePathLocation) == 0 {
		// DEFAULT TO DEV MODE CONFIGURATION
		configFileLocation = os.Getenv(app_constants.AppConfigLocationEnvName)

		// Read from Environmental Config

		deployment := os.Getenv("DEPLOYMENT_STATUS")

		var jsonFileLocationTemplate string

		if len(deployment) == 0 || deployment == "dev" {
			log.Printf("SET %v default Dev path", app_constants.AppConfigLocationEnvName)
			deployment = "dev"

			if len(configFileLocation) == 0 {
				jsonFileLocationTemplate = fmt.Sprintf(defaults_constants.CONFIG_FILE_LOCATION, app_constants.AppName, "%v")
				jsonFileLocation = fmt.Sprintf(jsonFileLocationTemplate, deployment)
			} else {
				jsonFileLocationTemplate = configFileLocation
				jsonFileLocation = fmt.Sprintf(jsonFileLocationTemplate, deployment)
			}

		}
		if deployment == "prod" {
			deployment = "prod"
			if len(configFileLocation) == 0 {
				jsonFileLocationTemplate = fmt.Sprintf(defaults_constants.CONFIG_FILE_LOCATION, app_constants.AppName, "%v")
				jsonFileLocation = fmt.Sprintf(jsonFileLocationTemplate, deployment)
			} else {
				jsonFileLocationTemplate = configFileLocation
				jsonFileLocation = fmt.Sprintf(jsonFileLocationTemplate, deployment)
			}

		} else if deployment == "stagging" {
			deployment = "stagging"
			if len(configFileLocation) == 0 {
				jsonFileLocationTemplate = fmt.Sprintf(defaults_constants.CONFIG_FILE_LOCATION, app_constants.AppName, "%v")
				jsonFileLocation = fmt.Sprintf(jsonFileLocationTemplate, deployment)
			} else {
				jsonFileLocationTemplate = configFileLocation
				jsonFileLocation = fmt.Sprintf(jsonFileLocationTemplate, deployment)
			}
		} else if deployment == "uat" {
			deployment = "uat"
			if len(configFileLocation) == 0 {
				jsonFileLocationTemplate = fmt.Sprintf(defaults_constants.CONFIG_FILE_LOCATION, app_constants.AppName, "%v")
				jsonFileLocation = fmt.Sprintf(jsonFileLocationTemplate, deployment)
			} else {
				jsonFileLocationTemplate = configFileLocation
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
func (config *AppConfig) ReadConfiguration() AppConfig {
	jsonConfiguration := config.LoadJSONConfig("")
	log.Println(jsonConfiguration + "  Reading configuration" + jsonConfiguration)
	configuration := (config.FromJSON(jsonConfiguration))
	log.Println(configuration)

	return configuration
}
