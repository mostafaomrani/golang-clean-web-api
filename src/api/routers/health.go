package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/mostafaomrani/golang-clean-web-api/api/handlers"
)

func Health(r *gin.RouterGroup) {
	handler := handlers.NewHeathHandler()
	r.GET("/", handler.Health)
	r.POST("/", handler.PostHealth)
	r.POST("/:id", handler.PostHealthById)
}
