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

	api := r.Group("/api")

	v1 := api.Group("/v1")
	{
		health := v1.Group("/health")
		routers.Health(health)

		testRouter := v1.Group("/test")
		routers.TestRouter(testRouter)
	}

	v2 := api.Group("/v2")
	{
		testRouter := v2.Group("/test")
		routers.TestRouter(testRouter)
	}

	// v1.GET("/health", func(ctx *gin.Context) {
	// })

	r.Run(fmt.Sprintf(":%s", cfg.Server.Port))

}
