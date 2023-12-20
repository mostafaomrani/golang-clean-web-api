package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mostafaomrani/golang-clean-web-api/config"

	"github.com/mostafaomrani/golang-clean-web-api/api/routers"
)

func InitServer() {

	cfg := config.GetConfig()

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	v1 := r.Group("/api/v1")

	{
		health := v1.Group("/health")
		routers.Health(health)
	}

	// v1.GET("/health", func(ctx *gin.Context) {
	// })

	r.Run(fmt.Sprintf(":%s", cfg.Server.Port))

}
