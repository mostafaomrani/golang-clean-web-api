package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func TestMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		api := c.GetHeader("x-api-key")
		if api == "1" {
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"result": "Api key is required",
			})
			return
			
		}

	}
}
