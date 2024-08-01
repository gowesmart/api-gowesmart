package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gowesmart/api-gowesmart/model/web"
	"github.com/gowesmart/api-gowesmart/utils"
)

func JwtAuthMiddleware(c *gin.Context) {
	err := utils.TokenValid(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, &web.WebError{Code: http.StatusUnauthorized, Errors: err.Error()})
		return
	}
	c.Next()
}
