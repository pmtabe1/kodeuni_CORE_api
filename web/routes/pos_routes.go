package routes

import (
	"errors"
	"log"
	"net/http"
	"os"
	"path/filepath"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"github.com/paulmsegeya/pos/core/models/auth_models"
	"github.com/paulmsegeya/pos/internal/handlers"
	"github.com/paulmsegeya/pos/internal/middlewares/jwt_auth_middleware"
	"github.com/paulmsegeya/pos/utils/httputil"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type IRoutes interface {
}

type Routes struct {
}

func New() *Routes {

	return &Routes{}
}

var identityKey = "id"

func StartIntergrator() {
	port := os.Getenv("PORT")
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	if port == "" {
		port = "8000"
	}

	// the jwt middleware
	initiateJWTMiddlewareInstance := jwt_auth_middleware.New()

	authMiddleware := initiateJWTMiddlewareInstance.GetJWTAuthMiddleware()
	// When you use jwt.New(), the function is already automatically called for checking,
	// which means you don't need to call it again.
	errInit := authMiddleware.MiddlewareInit()

	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}

	r.POST("/login", authMiddleware.LoginHandler)

	r.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
	// user, pass := environments.GetFineractCredentials()

	// headerEntity := entities_dto.FineractHeaderEnity{
	// 	ContentType: "application/json",
	// 	Authorization: entities_dto.FineractBasicAuthorizationEntity{
	// 		Username: user,
	// 		Password: pass,
	// 	},
	// 	FineractPlatformTenantId: "nexis",
	// }
	// log.Println(headerEntity)
	// integratonEntity := entities_dto.IntegrationEntity{
	// 	BankAPIEndpointURI:  environments.GetBankAPIEndpoint(),
	// 	ReceivedPayload:     "",
	// 	HttpAPIHeaders:      map[string]interface{}{},
	// 	FineractHookHeaders: map[string]string{},
	// }

	//rootHandlers := handlers.New()
	v1 := r.Group("/api/v1")
	auth := r.Group("/auth")
	// Refresh time can be longer than token timeout
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)

	middleware := v1.Group("/sage")
	{
		middleware.Use(authMiddleware.MiddlewareFunc())
		{
			auth.GET("/ping", pingClaimsHandlers)
			// middleware.POST("/bankCallbacks", rootHandlers.BankCallbacks)
			// middleware.POST("/bankRequests", rootHandlers.BankRequests)
			// middleware.POST("/channelsRequests", rootHandlers.ChannelsRequests)
			// middleware.POST("/hooks", rootHandlers.EventHooks)
			// middleware.POST("/bankSimulations", rootHandlers.BankSimulations)

		}

	}

	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}

func (s *Routes) StartPosServer() {
	log.Println("Starting SAGE MIDDLEWARE SERVER ....")

	// "integration_deployment") == "dev"  determines which db to pick Live OR uat

	if len(os.Getenv("DEPLOYMENT_STATUS")) == 0 {

		os.Setenv("DEPLOYMENT_STATUS", "dev")
	}

	// if len(os.Getenv("DEPLOYMENT")) == 0 {
	// 	os.Setenv("DBENGINE", "mysql")

	// } else {
	// 	os.Setenv("DBENGINE", "")
	// }

	if len(os.Getenv("POS_CONFIG")) == 0 {
		os.Setenv("POS_CONFIG", filepath.FromSlash("/etc/integrations/conf/pos/config.%v.json"))

	}
	port := os.Getenv("PORT")
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	if port == "" {
		port = "9001"
	}

	// user, pass := environments.GetFineractCredentials()

	// headerEntity := entities_dto.FineractHeaderEnity{
	// 	ContentType: "application/json",
	// 	Authorization: entities_dto.FineractBasicAuthorizationEntity{
	// 		Username: user,
	// 		Password: pass,
	// 	},
	// 	FineractPlatformTenantId: environments.GetPlatformTENANTID(),
	// }
	// log.Println(headerEntity)
	// integratonEntity := entities_dto.IntegrationEntity{
	// 	BankAPIEndpointURI:  environments.GetBankAPIEndpoint(),
	// 	ReceivedPayload:     "",
	// 	HttpAPIHeaders:      map[string]interface{}{},
	// 	FineractHookHeaders: map[string]string{},
	// }

	rootHandlers := handlers.New()
	v2 := r.Group("/api/v1/pos")

	middleware := v2.Group("/auth")
	{
		middleware.Use(auth())
		middleware.POST("/auth", rootHandlers.AuthHandlers.Auth)

	}

	middleware = v2.Group("/purchase")
	{

		middleware.POST("/add", rootHandlers.PurchaseHandlers.Add)
		middleware.POST("/addOrupdate/:id", rootHandlers.PurchaseHandlers.AddOrUpdate)
		middleware.PUT("/update/:id", rootHandlers.PurchaseHandlers.Update)
		middleware.DELETE("/delete/:id", rootHandlers.PurchaseHandlers.Delete)
		middleware.GET("/locale/:locale", rootHandlers.PurchaseHandlers.GetByLocate)
		middleware.GET("/id/:id", rootHandlers.PurchaseHandlers.GetByID)
		middleware.GET("/owner/:owner", rootHandlers.PurchaseHandlers.GetByOwnerRef)
		middleware.GET("/type/:type", rootHandlers.PurchaseHandlers.GetByType)
		middleware.GET("/stage/:stage", rootHandlers.PurchaseHandlers.GetByStage)
		middleware.GET("/enabled/:enabled", rootHandlers.PurchaseHandlers.GetByEnabled)
		middleware.GET("/all", rootHandlers.PurchaseHandlers.GetAll)
		middleware.GET("/date/:date", rootHandlers.PurchaseHandlers.GetByDate)

	}

	middleware = v2.Group("/invoice")
	{

		middleware.POST("/add", rootHandlers.InvoiceHandlers.Add)
		middleware.POST("/addOrupdate/:id", rootHandlers.InvoiceHandlers.AddOrUpdate)
		middleware.PUT("/update/:id", rootHandlers.InvoiceHandlers.Update)
		middleware.DELETE("/delete/:id", rootHandlers.InvoiceHandlers.Delete)
		middleware.GET("/locale/:locale", rootHandlers.InvoiceHandlers.GetByLocate)
		middleware.GET("/id/:id", rootHandlers.InvoiceHandlers.GetByID)
		middleware.GET("/owner/:owner", rootHandlers.InvoiceHandlers.GetByOwnerRef)
		middleware.GET("/type/:type", rootHandlers.InvoiceHandlers.GetByType)
		middleware.GET("/stage/:stage", rootHandlers.InvoiceHandlers.GetByStage)
		middleware.GET("/enabled/:enabled", rootHandlers.InvoiceHandlers.GetByEnabled)
		middleware.GET("/all", rootHandlers.InvoiceHandlers.GetAll)
		middleware.GET("/date/:date", rootHandlers.InvoiceHandlers.GetByDate)

	}

	middleware = v2.Group("/paymentmethod")
	{

		middleware.POST("/add", rootHandlers.PaymentMethodHandlers.Add)
		middleware.POST("/addOrupdate/:id", rootHandlers.PaymentMethodHandlers.AddOrUpdate)
		middleware.PUT("/update/:id", rootHandlers.PaymentMethodHandlers.Update)
		middleware.DELETE("/delete/:id", rootHandlers.PaymentMethodHandlers.Delete)
		middleware.GET("/locale/:locale", rootHandlers.PaymentMethodHandlers.GetByLocate)
		middleware.GET("/id/:id", rootHandlers.PaymentMethodHandlers.GetByID)
		middleware.GET("/owner/:owner", rootHandlers.PaymentMethodHandlers.GetByOwnerRef)
		middleware.GET("/type/:type", rootHandlers.PaymentMethodHandlers.GetByType)
		middleware.GET("/stage/:stage", rootHandlers.PaymentMethodHandlers.GetByStage)
		middleware.GET("/enabled/:enabled", rootHandlers.PaymentMethodHandlers.GetByEnabled)
		middleware.GET("/all", rootHandlers.PaymentMethodHandlers.GetAll)
		middleware.GET("/date/:date", rootHandlers.PaymentMethodHandlers.GetByDate)

	}

	middleware = v2.Group("/datalog")
	{

		middleware.POST("/add", rootHandlers.DatalogHandlers.Add)
		middleware.POST("/addOrupdate/:id", rootHandlers.DatalogHandlers.AddOrUpdate)
		middleware.PUT("/update/:id", rootHandlers.DatalogHandlers.Update)
		middleware.DELETE("/delete/:id", rootHandlers.DatalogHandlers.Delete)
		middleware.GET("/locale/:locale", rootHandlers.DatalogHandlers.GetByLocate)
		middleware.GET("/id/:id", rootHandlers.DatalogHandlers.GetByID)
		middleware.GET("/owner/:owner", rootHandlers.DatalogHandlers.GetByOwnerRef)
		middleware.GET("/type/:type", rootHandlers.DatalogHandlers.GetByType)
		middleware.GET("/stage/:stage", rootHandlers.DatalogHandlers.GetByStage)
		middleware.GET("/enabled/:enabled", rootHandlers.DatalogHandlers.GetByEnabled)
		middleware.GET("/all", rootHandlers.DatalogHandlers.GetAll)
		middleware.GET("/date/:date", rootHandlers.DatalogHandlers.GetByDate)

	}

	middleware = v2.Group("/reference")
	{

		middleware.POST("/add", rootHandlers.ReferenceHandlers.Add)
		middleware.POST("/addOrupdate/:id", rootHandlers.ReferenceHandlers.AddOrUpdate)
		middleware.PUT("/update/:id", rootHandlers.ReferenceHandlers.Update)
		middleware.DELETE("/delete/:id", rootHandlers.ReferenceHandlers.Delete)
		middleware.GET("/locale/:locale", rootHandlers.ReferenceHandlers.GetByLocate)
		middleware.GET("/id/:id", rootHandlers.ReferenceHandlers.GetByID)
		middleware.GET("/owner/:owner", rootHandlers.ReferenceHandlers.GetByOwnerRef)
		middleware.GET("/type/:type", rootHandlers.ReferenceHandlers.GetByType)
		middleware.GET("/stage/:stage", rootHandlers.ReferenceHandlers.GetByStage)
		middleware.GET("/enabled/:enabled", rootHandlers.ReferenceHandlers.GetByEnabled)
		middleware.GET("/all", rootHandlers.ReferenceHandlers.GetAll)
		middleware.GET("/date/:date", rootHandlers.ReferenceHandlers.GetByDate)

	}

	middleware = v2.Group("/inventory")
	{

		middleware.POST("/add", rootHandlers.InventoryHandlers.Add)
		middleware.POST("/addOrupdate/:id", rootHandlers.InventoryHandlers.AddOrUpdate)
		middleware.PUT("/update/:id", rootHandlers.InventoryHandlers.Update)
		middleware.DELETE("/delete/:id", rootHandlers.InventoryHandlers.Delete)
		middleware.GET("/locale/:locale", rootHandlers.InventoryHandlers.GetByLocate)
		middleware.GET("/id/:id", rootHandlers.InventoryHandlers.GetByID)
		middleware.GET("/owner/:owner", rootHandlers.InventoryHandlers.GetByOwnerRef)
		middleware.GET("/type/:type", rootHandlers.InventoryHandlers.GetByType)
		middleware.GET("/stage/:stage", rootHandlers.InventoryHandlers.GetByStage)
		middleware.GET("/enabled/:enabled", rootHandlers.InventoryHandlers.GetByEnabled)
		middleware.GET("/all", rootHandlers.InventoryHandlers.GetAll)
		middleware.GET("/date/:date", rootHandlers.InventoryHandlers.GetByDate)

	}
	middleware = v2.Group("/tax")
	{

		middleware.POST("/add", rootHandlers.TaxHandlers.Add)
		middleware.POST("/addOrupdate/:id", rootHandlers.TaxHandlers.AddOrUpdate)
		middleware.PUT("/update/:id", rootHandlers.TaxHandlers.Update)
		middleware.DELETE("/delete/:id", rootHandlers.TaxHandlers.Delete)
		middleware.GET("/locale/:locale", rootHandlers.TaxHandlers.GetByLocate)
		middleware.GET("/id/:id", rootHandlers.TaxHandlers.GetByID)
		middleware.GET("/owner/:owner", rootHandlers.TaxHandlers.GetByOwnerRef)
		middleware.GET("/type/:type", rootHandlers.TaxHandlers.GetByType)
		middleware.GET("/stage/:stage", rootHandlers.TaxHandlers.GetByStage)
		middleware.GET("/enabled/:enabled", rootHandlers.TaxHandlers.GetByEnabled)
		middleware.GET("/all", rootHandlers.TaxHandlers.GetAll)
		middleware.GET("/date/:date", rootHandlers.TaxHandlers.GetByDate)

	}
	middleware = v2.Group("/workflow")
	{

		middleware.POST("/add", rootHandlers.WorkflowHandlers.Add)
		middleware.POST("/addOrupdate/:id", rootHandlers.WorkflowHandlers.AddOrUpdate)
		middleware.PUT("/update/:id", rootHandlers.WorkflowHandlers.Update)
		middleware.DELETE("/delete/:id", rootHandlers.WorkflowHandlers.Delete)
		middleware.GET("/locale/:locale", rootHandlers.WorkflowHandlers.GetByLocate)
		middleware.GET("/id/:id", rootHandlers.WorkflowHandlers.GetByID)
		middleware.GET("/owner/:owner", rootHandlers.WorkflowHandlers.GetByOwnerRef)
		middleware.GET("/type/:type", rootHandlers.WorkflowHandlers.GetByType)
		middleware.GET("/stage/:stage", rootHandlers.WorkflowHandlers.GetByStage)
		middleware.GET("/enabled/:enabled", rootHandlers.WorkflowHandlers.GetByEnabled)
		middleware.GET("/all", rootHandlers.WorkflowHandlers.GetAll)
		middleware.GET("/date/:date", rootHandlers.WorkflowHandlers.GetByDate)

	}

	middleware = v2.Group("/order")
	{

		middleware.POST("/add", rootHandlers.OrderHandlers.Add)
		middleware.POST("/addOrupdate/:id", rootHandlers.OrderHandlers.AddOrUpdate)
		middleware.PUT("/update/:id", rootHandlers.OrderHandlers.Update)
		middleware.DELETE("/delete/:id", rootHandlers.OrderHandlers.Delete)
		middleware.GET("/locale/:locale", rootHandlers.OrderHandlers.GetByLocate)
		middleware.GET("/id/:id", rootHandlers.OrderHandlers.GetByID)
		middleware.GET("/owner/:owner", rootHandlers.OrderHandlers.GetByOwnerRef)
		middleware.GET("/type/:type", rootHandlers.OrderHandlers.GetByType)
		middleware.GET("/stage/:stage", rootHandlers.OrderHandlers.GetByStage)
		middleware.GET("/enabled/:enabled", rootHandlers.OrderHandlers.GetByEnabled)
		middleware.GET("/all", rootHandlers.OrderHandlers.GetAll)
		middleware.GET("/date/:date", rootHandlers.OrderHandlers.GetByDate)

	}

	middleware = v2.Group("/product")
	{

		middleware.POST("/add", rootHandlers.ProductHandlers.Add)
		middleware.POST("/addOrupdate/:id", rootHandlers.ProductHandlers.AddOrUpdate)
		middleware.PUT("/update/:id", rootHandlers.ProductHandlers.Update)
		middleware.DELETE("/delete/:id", rootHandlers.ProductHandlers.Delete)
		middleware.GET("/locale/:locale", rootHandlers.ProductHandlers.GetByLocate)
		middleware.GET("/id/:id", rootHandlers.ProductHandlers.GetByID)
		middleware.GET("/owner/:owner", rootHandlers.ProductHandlers.GetByOwnerRef)
		middleware.GET("/type/:type", rootHandlers.ProductHandlers.GetByType)
		middleware.GET("/stage/:stage", rootHandlers.ProductHandlers.GetByStage)
		middleware.GET("/enabled/:enabled", rootHandlers.ProductHandlers.GetByEnabled)
		middleware.GET("/all", rootHandlers.ProductHandlers.GetAll)
		middleware.GET("/date/:date", rootHandlers.ProductHandlers.GetByDate)

	}
	middleware = v2.Group("/marketing")
	{

		middleware.POST("/add", rootHandlers.MarketingHandlers.Add)
		middleware.POST("/addOrupdate/:id", rootHandlers.MarketingHandlers.AddOrUpdate)
		middleware.PUT("/update/:id", rootHandlers.MarketingHandlers.Update)
		middleware.DELETE("/delete/:id", rootHandlers.MarketingHandlers.Delete)
		middleware.GET("/locale/:locale", rootHandlers.MarketingHandlers.GetByLocate)
		middleware.GET("/id/:id", rootHandlers.MarketingHandlers.GetByID)
		middleware.GET("/owner/:owner", rootHandlers.MarketingHandlers.GetByOwnerRef)
		middleware.GET("/type/:type", rootHandlers.MarketingHandlers.GetByType)
		middleware.GET("/stage/:stage", rootHandlers.MarketingHandlers.GetByStage)
		middleware.GET("/enabled/:enabled", rootHandlers.MarketingHandlers.GetByEnabled)
		middleware.GET("/all", rootHandlers.MarketingHandlers.GetAll)
		middleware.GET("/date/:date", rootHandlers.MarketingHandlers.GetByDate)

	}

	middleware = v2.Group("/employee")
	{

		middleware.POST("/add", rootHandlers.EmployeeHandlers.Add)
		middleware.POST("/addOrupdate/:id", rootHandlers.EmployeeHandlers.AddOrUpdate)
		middleware.PUT("/update/:id", rootHandlers.EmployeeHandlers.Update)
		middleware.DELETE("/delete/:id", rootHandlers.EmployeeHandlers.Delete)
		middleware.GET("/locale/:locale", rootHandlers.EmployeeHandlers.GetByLocate)
		middleware.GET("/id/:id", rootHandlers.EmployeeHandlers.GetByID)
		middleware.GET("/owner/:owner", rootHandlers.EmployeeHandlers.GetByOwnerRef)
		middleware.GET("/type/:type", rootHandlers.EmployeeHandlers.GetByType)
		middleware.GET("/stage/:stage", rootHandlers.EmployeeHandlers.GetByStage)
		middleware.GET("/enabled/:enabled", rootHandlers.EmployeeHandlers.GetByEnabled)
		middleware.GET("/all", rootHandlers.EmployeeHandlers.GetAll)
		middleware.GET("/date/:date", rootHandlers.EmployeeHandlers.GetByDate)

	}

	middleware = v2.Group("/change")
	{

		middleware.POST("/add", rootHandlers.ChangeHandlers.Add)
		middleware.POST("/addOrupdate/:id", rootHandlers.ChangeHandlers.AddOrUpdate)
		middleware.PUT("/update/:id", rootHandlers.ChangeHandlers.Update)
		middleware.DELETE("/delete/:id", rootHandlers.ChangeHandlers.Delete)
		middleware.GET("/locale/:locale", rootHandlers.ChangeHandlers.GetByLocate)
		middleware.GET("/id/:id", rootHandlers.ChangeHandlers.GetByID)
		middleware.GET("/owner/:owner", rootHandlers.ChangeHandlers.GetByOwnerRef)
		middleware.GET("/type/:type", rootHandlers.ChangeHandlers.GetByType)
		middleware.GET("/stage/:stage", rootHandlers.ChangeHandlers.GetByStage)
		middleware.GET("/enabled/:enabled", rootHandlers.ChangeHandlers.GetByEnabled)
		middleware.GET("/all", rootHandlers.ChangeHandlers.GetAll)
		middleware.GET("/date/:date", rootHandlers.ChangeHandlers.GetByDate)

	}

	middleware = v2.Group("/contact")
	{

		middleware.POST("/add", rootHandlers.ContactHandlers.Add)
		middleware.POST("/addOrupdate/:id", rootHandlers.ContactHandlers.AddOrUpdate)
		middleware.PUT("/update/:id", rootHandlers.ContactHandlers.Update)
		middleware.DELETE("/delete/:id", rootHandlers.ContactHandlers.Delete)
		middleware.GET("/locale/:locale", rootHandlers.ContactHandlers.GetByLocate)
		middleware.GET("/id/:id", rootHandlers.ContactHandlers.GetByID)
		middleware.GET("/owner/:owner", rootHandlers.ContactHandlers.GetByOwnerRef)
		middleware.GET("/type/:type", rootHandlers.ContactHandlers.GetByType)
		middleware.GET("/stage/:stage", rootHandlers.ContactHandlers.GetByStage)
		middleware.GET("/enabled/:enabled", rootHandlers.ContactHandlers.GetByEnabled)
		middleware.GET("/all", rootHandlers.ContactHandlers.GetAll)
		middleware.GET("/date/:date", rootHandlers.ContactHandlers.GetByDate)

	}

	middleware = v2.Group("/catalogue")
	{

		middleware.POST("/add", rootHandlers.CatalogueHandlers.Add)
		middleware.POST("/addOrupdate/:id", rootHandlers.CatalogueHandlers.AddOrUpdate)
		middleware.PUT("/update/:id", rootHandlers.CatalogueHandlers.Update)
		middleware.DELETE("/delete/:id", rootHandlers.CatalogueHandlers.Delete)
		middleware.GET("/locale/:locale", rootHandlers.CatalogueHandlers.GetByLocate)
		middleware.GET("/id/:id", rootHandlers.CatalogueHandlers.GetByID)
		middleware.GET("/owner/:owner", rootHandlers.CatalogueHandlers.GetByOwnerRef)
		middleware.GET("/type/:type", rootHandlers.CatalogueHandlers.GetByType)
		middleware.GET("/stage/:stage", rootHandlers.CatalogueHandlers.GetByStage)
		middleware.GET("/enabled/:enabled", rootHandlers.CatalogueHandlers.GetByEnabled)
		middleware.GET("/all", rootHandlers.CatalogueHandlers.GetAll)
		middleware.GET("/date/:date", rootHandlers.CatalogueHandlers.GetByDate)

	}

	middleware = v2.Group("/customer")
	{

		middleware.POST("/add", rootHandlers.CustomerHandlers.Add)
		middleware.POST("/addOrupdate/:id", rootHandlers.CustomerHandlers.AddOrUpdate)
		middleware.PUT("/update/:id", rootHandlers.CustomerHandlers.Update)
		middleware.DELETE("/delete/:id", rootHandlers.CustomerHandlers.Delete)
		middleware.GET("/locale/:locale", rootHandlers.CustomerHandlers.GetByLocate)
		middleware.GET("/id/:id", rootHandlers.CustomerHandlers.GetByID)
		middleware.GET("/owner/:owner", rootHandlers.CustomerHandlers.GetByOwnerRef)
		middleware.GET("/type/:type", rootHandlers.CustomerHandlers.GetByType)
		middleware.GET("/stage/:stage", rootHandlers.CustomerHandlers.GetByStage)
		middleware.GET("/enabled/:enabled", rootHandlers.CustomerHandlers.GetByEnabled)
		middleware.GET("/all", rootHandlers.CustomerHandlers.GetAll)
		middleware.GET("/date/:date", rootHandlers.CustomerHandlers.GetByDate)

	}

	middleware = v2.Group("/department")
	{

		middleware.POST("/add", rootHandlers.DepartmentHandlers.Add)
		middleware.POST("/addOrupdate/:id", rootHandlers.DepartmentHandlers.AddOrUpdate)
		middleware.PUT("/update/:id", rootHandlers.DepartmentHandlers.Update)
		middleware.DELETE("/delete/:id", rootHandlers.DepartmentHandlers.Delete)
		middleware.GET("/locale/:locale", rootHandlers.DepartmentHandlers.GetByLocate)
		middleware.GET("/id/:id", rootHandlers.DepartmentHandlers.GetByID)
		middleware.GET("/owner/:owner", rootHandlers.DepartmentHandlers.GetByOwnerRef)
		middleware.GET("/type/:type", rootHandlers.DepartmentHandlers.GetByType)
		middleware.GET("/stage/:stage", rootHandlers.DepartmentHandlers.GetByStage)
		middleware.GET("/enabled/:enabled", rootHandlers.DepartmentHandlers.GetByEnabled)
		middleware.GET("/all", rootHandlers.DepartmentHandlers.GetAll)
		middleware.GET("/date/:date", rootHandlers.DepartmentHandlers.GetByDate)

	}

	middleware = v2.Group("/eod")
	{

		middleware.POST("/add", rootHandlers.EodHandlers.Add)
		middleware.POST("/addOrupdate/:id", rootHandlers.EodHandlers.AddOrUpdate)
		middleware.PUT("/update/:id", rootHandlers.EodHandlers.Update)
		middleware.DELETE("/delete/:id", rootHandlers.EodHandlers.Delete)
		middleware.GET("/locale/:locale", rootHandlers.EodHandlers.GetByLocate)
		middleware.GET("/id/:id", rootHandlers.EodHandlers.GetByID)
		middleware.GET("/owner/:owner", rootHandlers.EodHandlers.GetByOwnerRef)
		middleware.GET("/type/:type", rootHandlers.EodHandlers.GetByType)
		middleware.GET("/stage/:stage", rootHandlers.EodHandlers.GetByStage)
		middleware.GET("/enabled/:enabled", rootHandlers.EodHandlers.GetByEnabled)
		middleware.GET("/all", rootHandlers.EodHandlers.GetAll)
		middleware.GET("/date/:date", rootHandlers.EodHandlers.GetByDate)

	}

	middleware = v2.Group("/appstore")
	{

		middleware.POST("/add", rootHandlers.AppStoreHandlers.Add)
		middleware.POST("/addOrupdate/:id", rootHandlers.AppStoreHandlers.AddOrUpdate)
		middleware.PUT("/update/:id", rootHandlers.AppStoreHandlers.Update)
		middleware.DELETE("/delete/:id", rootHandlers.AppStoreHandlers.Delete)
		middleware.GET("/locale/:locale", rootHandlers.AppStoreHandlers.GetByLocate)
		middleware.GET("/id/:id", rootHandlers.AppStoreHandlers.GetByID)
		middleware.GET("/owner/:owner", rootHandlers.AppStoreHandlers.GetByOwnerRef)
		middleware.GET("/type/:type", rootHandlers.AppStoreHandlers.GetByType)
		middleware.GET("/stage/:stage", rootHandlers.AppStoreHandlers.GetByStage)
		middleware.GET("/enabled/:enabled", rootHandlers.AppStoreHandlers.GetByEnabled)
		middleware.GET("/all", rootHandlers.AppStoreHandlers.GetAll)
		middleware.GET("/date/:date", rootHandlers.AppStoreHandlers.GetByDate)

	}

	middleware = v2.Group("/sale")
	{

		middleware.POST("/add", rootHandlers.SaleHandlers.Add)
		middleware.POST("/addOrupdate/:id", rootHandlers.SaleHandlers.AddOrUpdate)
		middleware.PUT("/update/:id", rootHandlers.SaleHandlers.Update)
		middleware.DELETE("/delete/:id", rootHandlers.SaleHandlers.Delete)
		middleware.GET("/locale/:locale", rootHandlers.SaleHandlers.GetByLocate)
		middleware.GET("/id/:id", rootHandlers.SaleHandlers.GetByID)
		middleware.GET("/owner/:owner", rootHandlers.SaleHandlers.GetByOwnerRef)
		middleware.GET("/type/:type", rootHandlers.SaleHandlers.GetByType)
		middleware.GET("/stage/:stage", rootHandlers.SaleHandlers.GetByStage)
		middleware.GET("/enabled/:enabled", rootHandlers.SaleHandlers.GetByEnabled)
		middleware.GET("/all", rootHandlers.SaleHandlers.GetAll)
		middleware.GET("/date/:date", rootHandlers.SaleHandlers.GetByDate)

	}

	middleware = v2.Group("/promotion")
	{

		middleware.POST("/add", rootHandlers.PromotionHandlers.Add)
		middleware.POST("/addOrupdate/:id", rootHandlers.PromotionHandlers.AddOrUpdate)
		middleware.PUT("/update/:id", rootHandlers.PromotionHandlers.Update)
		middleware.DELETE("/delete/:id", rootHandlers.PromotionHandlers.Delete)
		middleware.GET("/locale/:locale", rootHandlers.PromotionHandlers.GetByLocate)
		middleware.GET("/id/:id", rootHandlers.PromotionHandlers.GetByID)
		middleware.GET("/owner/:owner", rootHandlers.PromotionHandlers.GetByOwnerRef)
		middleware.GET("/type/:type", rootHandlers.PromotionHandlers.GetByType)
		middleware.GET("/stage/:stage", rootHandlers.PromotionHandlers.GetByStage)
		middleware.GET("/enabled/:enabled", rootHandlers.PromotionHandlers.GetByEnabled)
		middleware.GET("/all", rootHandlers.PromotionHandlers.GetAll)
		middleware.GET("/date/:date", rootHandlers.PromotionHandlers.GetByDate)

	}

	middleware = v2.Group("/sell")
	{

		middleware.POST("/add", rootHandlers.SellHandlers.Add)
		middleware.POST("/addOrupdate/:id", rootHandlers.SellHandlers.AddOrUpdate)
		middleware.PUT("/update/:id", rootHandlers.SellHandlers.Update)
		middleware.DELETE("/delete/:id", rootHandlers.SellHandlers.Delete)
		middleware.GET("/locale/:locale", rootHandlers.SellHandlers.GetByLocate)
		middleware.GET("/id/:id", rootHandlers.SellHandlers.GetByID)
		middleware.GET("/owner/:owner", rootHandlers.SellHandlers.GetByOwnerRef)
		middleware.GET("/type/:type", rootHandlers.SellHandlers.GetByType)
		middleware.GET("/stage/:stage", rootHandlers.SellHandlers.GetByStage)
		middleware.GET("/enabled/:enabled", rootHandlers.SellHandlers.GetByEnabled)
		middleware.GET("/all", rootHandlers.SellHandlers.GetAll)
		middleware.GET("/date/:date", rootHandlers.SellHandlers.GetByDate)

	}

	middleware = v2.Group("/app")
	{

		middleware.POST("/add", rootHandlers.AppHandlers.Add)
		middleware.POST("/addOrupdate/:id", rootHandlers.AppHandlers.AddOrUpdate)
		middleware.PUT("/update/:id", rootHandlers.AppHandlers.Update)
		middleware.DELETE("/delete/:id", rootHandlers.AppHandlers.Delete)
		middleware.GET("/locale/:locale", rootHandlers.AppHandlers.GetByLocate)
		middleware.GET("/id/:id", rootHandlers.AppHandlers.GetByID)
		middleware.GET("/owner/:owner", rootHandlers.AppHandlers.GetByOwnerRef)
		middleware.GET("/type/:type", rootHandlers.AppHandlers.GetByType)
		middleware.GET("/stage/:stage", rootHandlers.AppHandlers.GetByStage)
		middleware.GET("/enabled/:enabled", rootHandlers.AppHandlers.GetByEnabled)
		middleware.GET("/all", rootHandlers.AppHandlers.GetAll)
		middleware.GET("/date/:date", rootHandlers.AppHandlers.GetByDate)

	}
	middleware = v2.Group("/shipping")
	{

		middleware.POST("/add", rootHandlers.ShippingHandlers.Add)
		middleware.POST("/addOrupdate/:id", rootHandlers.ShippingHandlers.AddOrUpdate)
		middleware.PUT("/update/:id", rootHandlers.ShippingHandlers.Update)
		middleware.DELETE("/delete/:id", rootHandlers.ShippingHandlers.Delete)
		middleware.GET("/locale/:locale", rootHandlers.ShippingHandlers.GetByLocate)
		middleware.GET("/id/:id", rootHandlers.ShippingHandlers.GetByID)
		middleware.GET("/owner/:owner", rootHandlers.ShippingHandlers.GetByOwnerRef)
		middleware.GET("/type/:type", rootHandlers.ShippingHandlers.GetByType)
		middleware.GET("/stage/:stage", rootHandlers.ShippingHandlers.GetByStage)
		middleware.GET("/enabled/:enabled", rootHandlers.ShippingHandlers.GetByEnabled)
		middleware.GET("/all", rootHandlers.ShippingHandlers.GetAll)
		middleware.GET("/date/:date", rootHandlers.ShippingHandlers.GetByDate)

	}

	middleware = v2.Group("/shipper")
	{

		middleware.POST("/add", rootHandlers.ShipperHandlers.Add)
		middleware.POST("/addOrupdate/:id", rootHandlers.ShipperHandlers.AddOrUpdate)
		middleware.PUT("/update/:id", rootHandlers.ShipperHandlers.Update)
		middleware.DELETE("/delete/:id", rootHandlers.ShipperHandlers.Delete)
		middleware.GET("/locale/:locale", rootHandlers.ShipperHandlers.GetByLocate)
		middleware.GET("/id/:id", rootHandlers.ShipperHandlers.GetByID)
		middleware.GET("/owner/:owner", rootHandlers.ShipperHandlers.GetByOwnerRef)
		middleware.GET("/type/:type", rootHandlers.ShipperHandlers.GetByType)
		middleware.GET("/stage/:stage", rootHandlers.ShipperHandlers.GetByStage)
		middleware.GET("/enabled/:enabled", rootHandlers.ShipperHandlers.GetByEnabled)
		middleware.GET("/all", rootHandlers.ShipperHandlers.GetAll)
		middleware.GET("/date/:date", rootHandlers.ShipperHandlers.GetByDate)

	}

	middleware = v2.Group("/till")
	{

		middleware.POST("/add", rootHandlers.TillHandlers.Add)
		middleware.POST("/addOrupdate/:id", rootHandlers.TillHandlers.AddOrUpdate)
		middleware.PUT("/update/:id", rootHandlers.TillHandlers.Update)
		middleware.DELETE("/delete/:id", rootHandlers.TillHandlers.Delete)
		middleware.GET("/locale/:locale", rootHandlers.TillHandlers.GetByLocate)
		middleware.GET("/id/:id", rootHandlers.TillHandlers.GetByID)
		middleware.GET("/owner/:owner", rootHandlers.TillHandlers.GetByOwnerRef)
		middleware.GET("/type/:type", rootHandlers.TillHandlers.GetByType)
		middleware.GET("/stage/:stage", rootHandlers.TillHandlers.GetByStage)
		middleware.GET("/enabled/:enabled", rootHandlers.TillHandlers.GetByEnabled)
		middleware.GET("/all", rootHandlers.TillHandlers.GetAll)
		middleware.GET("/date/:date", rootHandlers.TillHandlers.GetByDate)

	}

	middleware = v2.Group("/payment")
	{

		middleware.POST("/add", rootHandlers.PaymentHandlers.Add)
		middleware.POST("/addOrupdate/:id", rootHandlers.PaymentHandlers.AddOrUpdate)
		middleware.PUT("/update/:id", rootHandlers.PaymentHandlers.Update)
		middleware.DELETE("/delete/:id", rootHandlers.PaymentHandlers.Delete)
		middleware.GET("/locale/:locale", rootHandlers.PaymentHandlers.GetByLocate)
		middleware.GET("/id/:id", rootHandlers.PaymentHandlers.GetByID)
		middleware.GET("/owner/:owner", rootHandlers.PaymentHandlers.GetByOwnerRef)
		middleware.GET("/type/:type", rootHandlers.PaymentHandlers.GetByType)
		middleware.GET("/stage/:stage", rootHandlers.PaymentHandlers.GetByStage)
		middleware.GET("/enabled/:enabled", rootHandlers.PaymentHandlers.GetByEnabled)
		middleware.GET("/all", rootHandlers.PaymentHandlers.GetAll)
		middleware.GET("/date/:date", rootHandlers.PaymentHandlers.GetByDate)

	}

	middleware = v2.Group("/transaction")
	{

		middleware.POST("/add", rootHandlers.TransactionHandlers.Add)
		middleware.POST("/addOrupdate/:id", rootHandlers.TransactionHandlers.AddOrUpdate)
		middleware.PUT("/update/:id", rootHandlers.TransactionHandlers.Update)
		middleware.DELETE("/delete/:id", rootHandlers.TransactionHandlers.Delete)
		middleware.GET("/locale/:locale", rootHandlers.TransactionHandlers.GetByLocate)
		middleware.GET("/id/:id", rootHandlers.TransactionHandlers.GetByID)
		middleware.GET("/owner/:owner", rootHandlers.TransactionHandlers.GetByOwnerRef)
		middleware.GET("/type/:type", rootHandlers.TransactionHandlers.GetByType)
		middleware.GET("/stage/:stage", rootHandlers.TransactionHandlers.GetByStage)
		middleware.GET("/enabled/:enabled", rootHandlers.TransactionHandlers.GetByEnabled)
		middleware.GET("/all", rootHandlers.TransactionHandlers.GetAll)
		middleware.GET("/date/:date", rootHandlers.TransactionHandlers.GetByDate)

	}

	middleware = v2.Group("/register")
	{

		middleware.POST("/add", rootHandlers.RegisterHandlers.Add)
		middleware.POST("/addOrupdate/:id", rootHandlers.RegisterHandlers.AddOrUpdate)
		middleware.PUT("/update/:id", rootHandlers.RegisterHandlers.Update)
		middleware.DELETE("/delete/:id", rootHandlers.RegisterHandlers.Delete)
		middleware.GET("/locale/:locale", rootHandlers.RegisterHandlers.GetByLocate)
		middleware.GET("/id/:id", rootHandlers.RegisterHandlers.GetByID)
		middleware.GET("/owner/:owner", rootHandlers.RegisterHandlers.GetByOwnerRef)
		middleware.GET("/type/:type", rootHandlers.RegisterHandlers.GetByType)
		middleware.GET("/stage/:stage", rootHandlers.RegisterHandlers.GetByStage)
		middleware.GET("/enabled/:enabled", rootHandlers.RegisterHandlers.GetByEnabled)
		middleware.GET("/all", rootHandlers.RegisterHandlers.GetAll)
		middleware.GET("/date/:date", rootHandlers.RegisterHandlers.GetByDate)

	}

	middleware = v2.Group("/wallet")
	{

		middleware.POST("/add", rootHandlers.WalletHandlers.Add)
		middleware.POST("/addOrupdate/:id", rootHandlers.WalletHandlers.AddOrUpdate)
		middleware.PUT("/update/:id", rootHandlers.WalletHandlers.Update)
		middleware.DELETE("/delete/:id", rootHandlers.WalletHandlers.Delete)
		middleware.GET("/locale/:locale", rootHandlers.WalletHandlers.GetByLocate)
		middleware.GET("/id/:id", rootHandlers.WalletHandlers.GetByID)
		middleware.GET("/owner/:owner", rootHandlers.WalletHandlers.GetByOwnerRef)
		middleware.GET("/type/:type", rootHandlers.WalletHandlers.GetByType)
		middleware.GET("/stage/:stage", rootHandlers.WalletHandlers.GetByStage)
		middleware.GET("/enabled/:enabled", rootHandlers.WalletHandlers.GetByEnabled)
		middleware.GET("/all", rootHandlers.WalletHandlers.GetAll)
		middleware.GET("/date/:date", rootHandlers.WalletHandlers.GetByDate)

	}

	middleware = v2.Group("/channel")
	{

		middleware.POST("/add", rootHandlers.ChannelHandlers.Add)
		middleware.POST("/addOrupdate/:id", rootHandlers.ChannelHandlers.AddOrUpdate)
		middleware.PUT("/update/:id", rootHandlers.ChannelHandlers.Update)
		middleware.DELETE("/delete/:id", rootHandlers.ChannelHandlers.Delete)
		middleware.GET("/locale/:locale", rootHandlers.ChannelHandlers.GetByLocate)
		middleware.GET("/id/:id", rootHandlers.ChannelHandlers.GetByID)
		middleware.GET("/owner/:owner", rootHandlers.ChannelHandlers.GetByOwnerRef)
		middleware.GET("/type/:type", rootHandlers.ChannelHandlers.GetByType)
		middleware.GET("/stage/:stage", rootHandlers.ChannelHandlers.GetByStage)
		middleware.GET("/enabled/:enabled", rootHandlers.ChannelHandlers.GetByEnabled)
		middleware.GET("/all", rootHandlers.ChannelHandlers.GetAll)
		middleware.GET("/date/:date", rootHandlers.ChannelHandlers.GetByDate)

	}

	middleware = v2.Group("/consumption")
	{

		middleware.POST("/add", rootHandlers.ConsumptionHandlers.Add)
		middleware.POST("/addOrupdate/:id", rootHandlers.ConsumptionHandlers.AddOrUpdate)
		middleware.PUT("/update/:id", rootHandlers.ConsumptionHandlers.Update)
		middleware.DELETE("/delete/:id", rootHandlers.ConsumptionHandlers.Delete)
		middleware.GET("/locale/:locale", rootHandlers.ConsumptionHandlers.GetByLocate)
		middleware.GET("/id/:id", rootHandlers.ConsumptionHandlers.GetByID)
		middleware.GET("/owner/:owner", rootHandlers.ConsumptionHandlers.GetByOwnerRef)
		middleware.GET("/type/:type", rootHandlers.ConsumptionHandlers.GetByType)
		middleware.GET("/stage/:stage", rootHandlers.ConsumptionHandlers.GetByStage)
		middleware.GET("/enabled/:enabled", rootHandlers.ConsumptionHandlers.GetByEnabled)
		middleware.GET("/all", rootHandlers.ConsumptionHandlers.GetAll)
		middleware.GET("/date/:date", rootHandlers.ConsumptionHandlers.GetByDate)

	}

	middleware = v2.Group("/virtualcard")
	{

		middleware.POST("/add", rootHandlers.VirtualCardHandlers.Add)
		middleware.POST("/addOrupdate/:id", rootHandlers.VirtualCardHandlers.AddOrUpdate)
		middleware.PUT("/update/:id", rootHandlers.VirtualCardHandlers.Update)
		middleware.DELETE("/delete/:id", rootHandlers.VirtualCardHandlers.Delete)
		middleware.GET("/locale/:locale", rootHandlers.VirtualCardHandlers.GetByLocate)
		middleware.GET("/id/:id", rootHandlers.VirtualCardHandlers.GetByID)
		middleware.GET("/owner/:owner", rootHandlers.VirtualCardHandlers.GetByOwnerRef)
		middleware.GET("/type/:type", rootHandlers.VirtualCardHandlers.GetByType)
		middleware.GET("/stage/:stage", rootHandlers.VirtualCardHandlers.GetByStage)
		middleware.GET("/enabled/:enabled", rootHandlers.VirtualCardHandlers.GetByEnabled)
		middleware.GET("/all", rootHandlers.VirtualCardHandlers.GetAll)
		middleware.GET("/date/:date", rootHandlers.VirtualCardHandlers.GetByDate)

	}

	middleware = v2.Group("/supplier")
	{

		middleware.POST("/add", rootHandlers.SupplierHandlers.Add)
		middleware.POST("/addOrupdate/:id", rootHandlers.SupplierHandlers.AddOrUpdate)
		middleware.PUT("/update/:id", rootHandlers.SupplierHandlers.Update)
		middleware.DELETE("/delete/:id", rootHandlers.SupplierHandlers.Delete)
		middleware.GET("/locale/:locale", rootHandlers.SupplierHandlers.GetByLocate)
		middleware.GET("/id/:id", rootHandlers.SupplierHandlers.GetByID)
		middleware.GET("/owner/:owner", rootHandlers.SupplierHandlers.GetByOwnerRef)
		middleware.GET("/type/:type", rootHandlers.SupplierHandlers.GetByType)
		middleware.GET("/stage/:stage", rootHandlers.SupplierHandlers.GetByStage)
		middleware.GET("/enabled/:enabled", rootHandlers.SupplierHandlers.GetByEnabled)
		middleware.GET("/all", rootHandlers.SupplierHandlers.GetAll)
		middleware.GET("/date/:date", rootHandlers.SupplierHandlers.GetByDate)

	}

	middleware = v2.Group("/report")
	{

		middleware.POST("/add", rootHandlers.ReportHandlers.Add)
		middleware.POST("/addOrupdate/:id", rootHandlers.ReportHandlers.AddOrUpdate)
		middleware.PUT("/update/:id", rootHandlers.ReportHandlers.Update)
		middleware.DELETE("/delete/:id", rootHandlers.ReportHandlers.Delete)
		middleware.GET("/locale/:locale", rootHandlers.ReportHandlers.GetByLocate)
		middleware.GET("/id/:id", rootHandlers.ReportHandlers.GetByID)
		middleware.GET("/owner/:owner", rootHandlers.ReportHandlers.GetByOwnerRef)
		middleware.GET("/type/:type", rootHandlers.ReportHandlers.GetByType)
		middleware.GET("/stage/:stage", rootHandlers.ReportHandlers.GetByStage)
		middleware.GET("/enabled/:enabled", rootHandlers.ReportHandlers.GetByEnabled)
		middleware.GET("/all", rootHandlers.ReportHandlers.GetAll)
		middleware.GET("/date/:date", rootHandlers.ReportHandlers.GetByDate)

	}

	middleware = v2.Group("/dependant")
	{

		middleware.POST("/add", rootHandlers.DependantHandlers.Add)
		middleware.POST("/addOrupdate/:id", rootHandlers.DependantHandlers.AddOrUpdate)
		middleware.PUT("/update/:id", rootHandlers.DependantHandlers.Update)
		middleware.DELETE("/delete/:id", rootHandlers.DependantHandlers.Delete)
		middleware.GET("/locale/:locale", rootHandlers.DependantHandlers.GetByLocate)
		middleware.GET("/id/:id", rootHandlers.DependantHandlers.GetByID)
		middleware.GET("/owner/:owner", rootHandlers.DependantHandlers.GetByOwnerRef)
		middleware.GET("/type/:type", rootHandlers.DependantHandlers.GetByType)
		middleware.GET("/stage/:stage", rootHandlers.DependantHandlers.GetByStage)
		middleware.GET("/enabled/:enabled", rootHandlers.DependantHandlers.GetByEnabled)
		middleware.GET("/all", rootHandlers.DependantHandlers.GetAll)
		middleware.GET("/date/:date", rootHandlers.DependantHandlers.GetByDate)

	}

	middleware = v2.Group("/member")
	{

		middleware.POST("/add", rootHandlers.MemberHandlers.Add)
		middleware.POST("/addOrupdate/:id", rootHandlers.MemberHandlers.AddOrUpdate)
		middleware.PUT("/update/:id", rootHandlers.MemberHandlers.Update)
		middleware.DELETE("/delete/:id", rootHandlers.MemberHandlers.Delete)
		middleware.GET("/locale/:locale", rootHandlers.MemberHandlers.GetByLocate)
		middleware.GET("/id/:id", rootHandlers.MemberHandlers.GetByID)
		middleware.GET("/owner/:owner", rootHandlers.MemberHandlers.GetByOwnerRef)
		middleware.GET("/type/:type", rootHandlers.MemberHandlers.GetByType)
		middleware.GET("/stage/:stage", rootHandlers.MemberHandlers.GetByStage)
		middleware.GET("/enabled/:enabled", rootHandlers.MemberHandlers.GetByEnabled)
		middleware.GET("/all", rootHandlers.MemberHandlers.GetAll)
		middleware.GET("/date/:date", rootHandlers.MemberHandlers.GetByDate)

	}

	middleware = v2.Group("/utilization")
	{

		middleware.POST("/add", rootHandlers.UtilizationHandlers.Add)
		middleware.POST("/addOrupdate/:id", rootHandlers.UtilizationHandlers.AddOrUpdate)
		middleware.PUT("/update/:id", rootHandlers.UtilizationHandlers.Update)
		middleware.DELETE("/delete/:id", rootHandlers.UtilizationHandlers.Delete)
		middleware.GET("/locale/:locale", rootHandlers.UtilizationHandlers.GetByLocate)
		middleware.GET("/id/:id", rootHandlers.UtilizationHandlers.GetByID)
		middleware.GET("/owner/:owner", rootHandlers.UtilizationHandlers.GetByOwnerRef)
		middleware.GET("/type/:type", rootHandlers.UtilizationHandlers.GetByType)
		middleware.GET("/stage/:stage", rootHandlers.UtilizationHandlers.GetByStage)
		middleware.GET("/enabled/:enabled", rootHandlers.UtilizationHandlers.GetByEnabled)
		middleware.GET("/all", rootHandlers.UtilizationHandlers.GetAll)
		middleware.GET("/date/:date", rootHandlers.UtilizationHandlers.GetByDate)

	}

	middleware = v2.Group("/store")
	{

		middleware.POST("/add", rootHandlers.StoreHandlers.Add)
		middleware.POST("/addOrupdate/:id", rootHandlers.StoreHandlers.AddOrUpdate)
		middleware.PUT("/update/:id", rootHandlers.StoreHandlers.Update)
		middleware.DELETE("/delete/:id", rootHandlers.StoreHandlers.Delete)
		middleware.GET("/locale/:locale", rootHandlers.StoreHandlers.GetByLocate)
		middleware.GET("/id/:id", rootHandlers.StoreHandlers.GetByID)
		middleware.GET("/owner/:owner", rootHandlers.StoreHandlers.GetByOwnerRef)
		middleware.GET("/type/:type", rootHandlers.StoreHandlers.GetByType)
		middleware.GET("/stage/:stage", rootHandlers.StoreHandlers.GetByStage)
		middleware.GET("/enabled/:enabled", rootHandlers.StoreHandlers.GetByEnabled)
		middleware.GET("/all", rootHandlers.StoreHandlers.GetAll)
		middleware.GET("/date/:date", rootHandlers.StoreHandlers.GetByDate)

	}
	v2.GET("header", rootHandlers.AuthHandlers.AuthorizationHeader)
	v2.GET("securities", rootHandlers.AuthHandlers.SecuritiesAuthorization)
	v2.GET("attribute", rootHandlers.AuthHandlers.Attribute)
	v2.GET("signup", rootHandlers.AuthHandlers.Signup)
	v2.GET("login", rootHandlers.AuthHandlers.Login)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}

}

func pingClaimsHandlers(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	user, _ := c.Get(identityKey)
	c.JSON(200, gin.H{
		"userID":   claims[identityKey],
		"userName": user.(*auth_models.User).Username,
		"text":     "Cool Claims Pinged....",
	})
}

func auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.GetHeader("Authorization")) == 0 {
			httputil.NewError(c, http.StatusUnauthorized, errors.New("authorization is required Header"))
			c.Abort()
		}
		c.Next()
	}
}
