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
	"github.com/paulmsegeya/pos/internal/middlewares/jwt_auth_middleware"
	"github.com/paulmsegeya/pos/utils/httputil"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var identityKey_ = "id"

func StartSageIntergrator() {
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

	auth := r.Group("/auth")
	// Refresh time can be longer than token timeout
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/hello", helloHandler)
	}

	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}

func StartSage300Server2() {

	// "integration_deployment") == "dev"  determines which db to pick Live OR uat

	if len(os.Getenv("integration_deployment")) == 0 {

		os.Setenv("integration_deployment", "dev")
	}

	if len(os.Getenv("DEPLOYMENT")) == 0 {
		os.Setenv("DBENGINE", "mysql")

	} else {
		os.Setenv("DBENGINE", "")
	}

	if len(os.Getenv("POS_CONFIG")) == 0 {
		os.Setenv("POS_CONFIG", filepath.FromSlash("/etc/integrations/conf/sage/config.%v.json"))

	}
	port := os.Getenv("PORT")
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	if port == "" {
		port = "9001"
	}

	// sage300Handlers := sage_300_OE_handlers.New()
	// sageAUTHHandlers := sage_AUTH_handlers.New()
	//v1 := r.Group("/api/v1")

	// sage := v1.Group("/sage")
	// {
	// 	sage.Use(auth0())
	// 	sage.POST("/auth", sageAUTHHandlers.Auth)
	// 	sage.POST("/purchaseItem/add", sage300Handlers.LogOrderDetail)
	// 	sage.POST("/purchase/add", sage300Handlers.LogOrderHeader)
	// 	sage.POST("/orderEntry", sage300Handlers.OrderEntry)

	// }
	// // sage.POST("/orderEntry", sage300Handlers.OrderEntry)
	// v1.GET("header", sageAUTHHandlers.AuthorizationHeader)
	// v1.GET("securities", sageAUTHHandlers.SecuritiesAuthorization)
	// v1.GET("attribute", sageAUTHHandlers.Attribute)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}

}

func helloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	user, _ := c.Get(identityKey)
	c.JSON(200, gin.H{
		"userID":   claims[identityKey],
		"userName": user.(auth_models.User).Username,
		"text":     "Hello World.",
	})
}

func auth0() gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.GetHeader("Authorization")) == 0 {
			httputil.NewError(c, http.StatusUnauthorized, errors.New("authorization is required Header"))
			c.Abort()
		}
		c.Next()
	}
}
