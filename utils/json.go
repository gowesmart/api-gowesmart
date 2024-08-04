package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gowesmart/api-gowesmart/model/web"
)

func ToResponseJSON[T any](c *gin.Context, code int, data T, pagination *web.Metadata) {
	if pagination == nil {
		c.JSON(code, web.WebSuccess[T]{
			Code:     code,
			Message:  http.StatusText(code),
			Payload:  data,
			Metadata: nil,
		})

		return
	}

	c.JSON(code, web.WebSuccess[T]{
		Code:     code,
		Message:  http.StatusText(code),
		Payload:  data,
		Metadata: pagination,
	})
}
