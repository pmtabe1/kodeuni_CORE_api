package environments

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type IEnvironment interface {
}

type Environment struct {
}

func New() *Environment {

	return &Environment{}
}

func GetConfiguredConfigurationFile() (fileLocation string) {

	templated := `"/etc/integrations/conf/pos/config.%v.json"`

	deployment := GetEnvironmentDeploymentStatus()

	if deployment == "dev" {

		fileLocation = os.Getenv("POS_CONFIG")

		if len(fileLocation) == 0 {

			os.Setenv("POS_CONFIG", filepath.FromSlash(fmt.Sprintf(templated, deployment)))
			deployment = GetConfiguredConfigurationFile()
		}

	} else if deployment == "prod" {
		fileLocation = os.Getenv("POS_CONFIG")

		if len(fileLocation) == 0 {

			os.Setenv("POS_CONFIG", filepath.FromSlash(fmt.Sprintf(templated, deployment)))
			deployment = GetConfiguredConfigurationFile()
		}

	}

	return fileLocation
}

func ActivateEnvironment(environment string) {

	if environment == "dev" {
		os.Setenv("DEPLOYMENT_STATUS", "dev")
		os.Setenv("SAGE_CONFIG", "./conf/config.%v.json")
		os.Setenv("GORMDB", "GORMDB")
		os.Setenv("SAGE_BASE_URL", "https://fineract.nexiss.cloud/fineract-provider/api/v1")
		os.Setenv("BASE_URL", "http://localhost")
		os.Setenv("FINERACT_SUPER_USERNAME", "mifos")
		os.Setenv("FINERACT_SUPER_PASSWORD", "password")
		os.Setenv("FINERACT_PLATFORM_TENANTID", "nexis")
		os.Setenv("DEPLOYMENT_TARGET", "local")
		os.Setenv("DEV_BASE_URL", "https://fineract.nexiss.cloud/fineract-provider/api/v1")
		os.Setenv("PROD_BASE_URL", "") // Set it to actual production FINERACT BASE URL
		os.Setenv("UAT_BASE_URL", "")  // Set it to actual production FINERACT BASE URL
		os.Setenv("GORMDB", "GORMDB")
		os.Setenv("REFERENCE_SIZE", "1000")
		os.Setenv("FINERACT_PROD_DOMAIN", "fineract.nexiss.cloud/fineract-provider")
		os.Setenv("FINERACT_UAT_DOMAIN", "fineract.nexiss.cloud/fineract-provider")
		os.Setenv("FINERACT_DEV_DOMAIN", "fineract.nexiss.cloud/fineract-provider")
		os.Setenv("BANK_PROD_ENDPOINT_URL", "https:prod.bankabc.com/api/whatever")
		os.Setenv("BANK_DEV_ENDPOINT_URL", "https:dev.bankabc.com/api/whatever")
		os.Setenv("FINERACT_USER", "mifos")
		os.Setenv("FINERACT_PASSWORD", "password")
		os.Setenv("FINERACT_PLATFORM_TENANTID", "nexis")
		log.Println("configuration....")
	} else if environment == "prod" {
		os.Setenv("DEPLOYMENT_STATUS", "prod")
		os.Setenv("FINERACT_CONFIG", "/etc/integrations/fineractmiddleware/conf/config.%v.json")
		os.Setenv("GORMDB", "GORMDB")
		os.Setenv("FINERACT_BASE_URL", "https://fineract.nexiss.cloud/fineract-provider/api/v1")
		os.Setenv("BASE_URL", "https://fineract.nexiss.cloud/fineract-provider/api/v1")
		os.Setenv("FINERACT_SUPER_USERNAME", "mifos")
		os.Setenv("FINERACT_SUPER_PASSWORD", "password")
		os.Setenv("FINERACT_PLATFORM_TENANTID", "nexis")
		os.Setenv("DEPLOYMENT_TARGET", "local")
		os.Setenv("DEV_BASE_URL", "https://fineract.nexiss.cloud/fineract-provider/api/v1")
		os.Setenv("PROD_BASE_URL", "") // Set it to actual production FINERACT BASE URL
		os.Setenv("UAT_BASE_URL", "")  // Set it to actual production FINERACT BASE URL
		os.Setenv("GORMDB", "GORMDB")
		os.Setenv("REFERENCE_SIZE", "1000")
		os.Setenv("FINERACT_PROD_DOMAIN", "fineract.nexiss.cloud/fineract-provider")
		os.Setenv("FINERACT_UAT_DOMAIN", "fineract.nexiss.cloud/fineract-provider")
		os.Setenv("FINERACT_DEV_DOMAIN", "fineract.nexiss.cloud/fineract-provider")
		os.Setenv("BANK_PROD_ENDPOINT_URL", "https:prod.bankabc.com/api/whatever")
		os.Setenv("BANK_DEV_ENDPOINT_URL", "https:dev.bankabc.com/api/whatever")
		os.Setenv("FINERACT_USER", "mifos")
		os.Setenv("FINERACT_PASSWORD", "password")
		os.Setenv("FINERACT_PLATFORM_TENANTID", "nexis")
		log.Println("configuration....")

	}
}

func GetDefaultCompanyOrDatabase() (databaseOrCompany string) {

	deployment := GetEnvironmentDeploymentStatus()

	if deployment == "dev" || len(deployment) == 0 {

		databaseOrCompany = os.Getenv("SAGE_DEVDB")

		if len(databaseOrCompany) == 0 {

			os.Setenv("SAGE_DEVDB", "TSTDAT")
			GetBankAPIEndpoint()
		}
	} else if deployment == "prod" {

		databaseOrCompany = os.Getenv("SAGE_PRODDB")

		if len(databaseOrCompany) == 0 {

			log.Panicln("SAGE_PRODDB    is not set on your environmental variable")
		}

	}

	return databaseOrCompany
}

func GetFineractDOMAIN(lookupStatus string) string {

	var deployment string

	if len(lookupStatus) == 0 {
		deployment = os.Getenv("DEPLOYMENT_STATUS")

	} else {
		deployment = lookupStatus
	}

	var domain string

	if len(deployment) == 0 || deployment == "dev" {

		domain = os.Getenv("FINERACT_DEV_DOMAIN")

		if len(domain) == 0 {

			domain = "fineract.nexiss.cloud/fineract-provider"
		}
	} else if deployment == "uat" {

		domain = os.Getenv("FINERACT_UAT_DOMAIN")
	} else if deployment == "prod" {

		domain = os.Getenv("FINERACT_PROD_DOMAIN")
	}

	return domain
}
func GetFineractBASEURL(lookupStatus string) string {

	var deployment string

	if len(lookupStatus) == 0 {
		deployment = os.Getenv("DEPLOYMENT_STATUS")

	} else {
		deployment = lookupStatus
	}

	var value string

	if len(deployment) == 0 || deployment == "dev" {

		value = os.Getenv("DEV_BASE_URL")

		if len(value) == 0 {

			value = "https://fineract.nexiss.cloud/fineract-provider/api/v1"
		}
	} else if deployment == "uat" {

		value = os.Getenv("UAT_BASE_URL")
	} else if deployment == "prod" {

		value = os.Getenv("PROD_BASE_URL")
	}

	return value
}

func GetSageBASEURL(lookupStatus string) string {

	var deployment string

	if len(lookupStatus) == 0 {
		deployment = os.Getenv("DEPLOYMENT_STATUS")

	} else {
		deployment = lookupStatus
	}

	var value string

	if len(deployment) == 0 || deployment == "dev" {

		value = os.Getenv("DEV_BASE_URL")

		if len(value) == 0 {

			value = "http://localhost/Sage300WebApi"
		}
	} else if deployment == "uat" {

		value = os.Getenv("UAT_BASE_URL")
	} else if deployment == "prod" {

		value = os.Getenv("PROD_BASE_URL")
	}
	//http://localhost/Sage300WebApi/v1.0/-/TSTDAT/OE/OECreditDebitNotes

	return value
}

func GetDefaultBaseURL() string {
	var baseURL string
	baseURL = GetSageBASEURL(GetEnvironmentDeploymentStatus())
	return baseURL
}

func SetDefaultEnvironmentalVariables() {
	os.Setenv("DEPLOYMENT_STATUS", "dev")
	os.Setenv("FINERACT_CONFIG", "./conf/config.%v.json")
	os.Setenv("DEPLOYMENT_TARGET", "local")
	os.Setenv("DEV_BASE_URL", "https://fineract.nexiss.cloud/fineract-provider/api/v1")
	os.Setenv("PROD_BASE_URL", "") // Set it to actual production FINERACT BASE URL
	os.Setenv("UAT_BASE_URL", "")  // Set it to actual production FINERACT BASE URL
	os.Setenv("GORMDB", "GORMDB")
	os.Setenv("REFERENCE_SIZE", "1000")
	os.Setenv("FINERACT_PROD_DOMAIN", "fineract.nexiss.cloud/fineract-provider")
	os.Setenv("FINERACT_UAT_DOMAIN", "fineract.nexiss.cloud/fineract-provider")
	os.Setenv("FINERACT_DEV_DOMAIN", "fineract.nexiss.cloud/fineract-provider")
	os.Setenv("BANK_PROD_ENDPOINT_URL", "https:prod.bankabc.com/api/whatever")
	os.Setenv("BANK_DEV_ENDPOINT_URL", "https:dev.bankabc.com/api/whatever")
	os.Setenv("FINERACT_USER", "mifos")
	os.Setenv("FINERACT_PASSWORD", "password")
	os.Setenv("FINERACT_PLATFORM_TENANTID", "nexis")
	log.Println("configuration....")
	log.Println(os.Getenv("SAGE_CONFIG"))

}

func GetBankAPIEndpoint() string {
	var bankEndpointUrl string
	deployment := GetEnvironmentDeploymentStatus()

	if deployment == "dev" || len(deployment) == 0 {

		bankEndpointUrl = os.Getenv("BANK_DEV_ENDPOINT_URL")

		if len(bankEndpointUrl) == 0 {

			os.Setenv("BANK_DEV_ENDPOINT_URL", "")
			GetBankAPIEndpoint()
		}
	} else if deployment == "prod" {

		bankEndpointUrl = os.Getenv("BANK_PROD_ENDPOINT_URL")

		if len(bankEndpointUrl) == 0 {

			log.Panicln("BANK_PROD_ENDPOINT_URL    is not set on your environmental variable")
		}

	}

	return bankEndpointUrl
}

func GetEnvironmentDeploymentStatus() string {

	deployment := os.Getenv("DEPLOYMENT_STATUS")

	if len(deployment) == 0 {
		deployment = "dev"
	}

	return deployment
}

func GetDefaultContentType() string {

	return "application/json"
}

func GetPlatformTENANTID() string {

	var tenantID string

	deployment := GetEnvironmentDeploymentStatus()

	if len(deployment) == 0 || deployment == "dev" {

		tenantID = os.Getenv("FINERACT_PLATFORM_TENANTID")

		if len(tenantID) == 0 {

			tenantID = "nexis"
		}
	} else {

		tenantID = os.Getenv("FINERACT_PLATFORM_TENANTID")
	}

	return tenantID
}

func GetSageCredentials() (user string, password string) {

	user = os.Getenv("SAGE_USER")
	password = os.Getenv("SAGE_PASSWORD")

	if len(user) == 0 {
		user = "API"
		os.Setenv("SAGE_USER", user)
		GetSageCredentials()

	}

	if len(password) == 0 {
		password = "A1"
		os.Setenv("SAGE_PASSWORD", password)
		GetSageCredentials()
	}

	return user, password
}

func GetFineractCredentials() (user string, password string) {

	user = os.Getenv("FINERACT_USER")
	password = os.Getenv("FINERACT_PASSWORD")

	if len(user) == 0 {
		user = "mifos"
		os.Setenv("FINERACT_USER", user)
		GetFineractCredentials()

	}

	if len(password) == 0 {
		password = "password"
		os.Setenv("FINERACT_PASSWORD", password)
		GetFineractCredentials()
	}

	return user, password
}
