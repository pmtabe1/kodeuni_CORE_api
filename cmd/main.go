package main

/**
 * @Author Paul Msegeya
 */
import (
	"log"
	"os"

	"github.com/paulmsegeya/pos/databases/pos_databases"
	"github.com/paulmsegeya/pos/databases/pos_databases_migrations"
	_ "github.com/paulmsegeya/pos/docs"
	"github.com/paulmsegeya/pos/web/routes"
)

// @title   Fineract Middleware API
// @version 1.0
// @description This is a finaract middleware API  documentation .
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:9001
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information
func main() {
	//os.Setenv("GORDB","aaa")
	//os.Setenv("GORDB", "GORM")
	//os.Setenv("DEPLOYMENT_TARGET", "local")
	os.Setenv("DEPLOYMENT_STATUS", "dev")
	//os.Setenv("POS_CONFIG", "./conf/config.%v.json")
	os.Setenv("POS_CONFIG", "/etc/integrations/conf/pos/config.%v.json")
	os.Getenv("CONFIGURED PATH >>>" + os.Getenv("POS_CONFIG"))
	os.Setenv("DEPLOYMENT_TARGET", "local")
	os.Setenv("DEV_BASE_URL", "https://fineract.nexiss.cloud/fineract-provider/api/v1")
	os.Setenv("PROD_BASE_URL", "") // Set it to actual production FINERACT BASE URL
	os.Setenv("UAT_BASE_URL", "")  // Set it to actual production FINERACT BASE URL
	os.Setenv("GORMDB", "GORMDB")

	if len(os.Getenv("REFERENCE_SIZE")) == 0 {
		os.Setenv("REFERENCE_SIZE", "1000")
	}
	os.Setenv("FINERACT_PROD_DOMAIN", "fineract.nexiss.cloud/fineract-provider")
	os.Setenv("FINERACT_UAT_DOMAIN", "fineract.nexiss.cloud/fineract-provider")
	os.Setenv("FINERACT_DEV_DOMAIN", "fineract.nexiss.cloud/fineract-provider")
	os.Setenv("BANK_PROD_ENDPOINT_URL", "https:prod.bankabc.com/api/whatever")
	os.Setenv("BANK_DEV_ENDPOINT_URL", "https:dev.bankabc.com/api/whatever")
	os.Setenv("BANK_PIPEDRIVE_SIMULATION_API", "")
	os.Setenv("FINERACT_PLATFORM_TENANTID", "nexis")
	os.Setenv("FINERACT_USER", "mifos")
	os.Setenv("FINERACT_PASSWORD", "password")

	err := pos_databases_migrations.Migrate(pos_databases.New().DBConnection())

	if err != nil {
		log.Panicln("MIGRATION ERROR .....")
	}
	log.Println("configuration....")
	routes.New().StartPosServer()

}
