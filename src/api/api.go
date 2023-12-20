package api

import (
	"github.com/gin-gonic/gin"

	"github.com/mostafaomrani/golang-clean-web-api/api/routers"
)

func InitServer() {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	v1 := r.Group("/api/v1")

	{
		health := v1.Group("/health")
		routers.Health(health)
	}

	// v1.GET("/health", func(ctx *gin.Context) {
	// })

	r.Run(":5005")
}
