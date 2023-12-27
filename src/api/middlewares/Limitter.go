package middlewares

import (
	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func LimittByRequest() gin.HandlerFunc {
	lmt := tollbooth.NewLimiter(1, &limiter.ExpirableOptions{DefaultExpirationTTL: time.Hour})
	lmt.SetMessage("You have reached maximum request limit.")

	return func(context *gin.Context) {
		err := tollbooth.LimitByRequest(lmt, context.Writer, context.Request)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": err.Error(),
			})
			return
		} else {
			context.Next()
		}
	}
}
