package jwt_auth_middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/paulmsegeya/pos/core/models/auth_models"
)



func GetAuthorizer(data interface{}, c *gin.Context) bool {
	if v, ok := data.(*auth_models.User); ok && v.Username == "admin" {
		return true
	}

	return false

}
