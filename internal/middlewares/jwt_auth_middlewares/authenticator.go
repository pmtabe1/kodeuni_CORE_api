package jwt_auth_middlewares

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/paulmsegeya/pos/core/models/auth_models"
)

func GetAuthenticator(c *gin.Context) (interface{}, error) {
	var loginVals auth_models.Login
	if err := c.ShouldBind(&loginVals); err != nil {
		return "", jwt.ErrMissingLoginValues
	}
	userID := loginVals.Username
	password := loginVals.Password

	if (userID == "admin" && password == "admin") || (userID == "test" && password == "test") {
		return &auth_models.User{
			Username:  userID,
			Lastname:  "Supper",
			Firstname: "Geeky",
		}, nil
	}

	return nil, jwt.ErrFailedAuthentication
}
