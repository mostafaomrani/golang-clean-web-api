package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/mostafaomrani/golang-clean-web-api/api/middlewares"
	"github.com/mostafaomrani/golang-clean-web-api/api/validations"
	"github.com/mostafaomrani/golang-clean-web-api/config"

	"github.com/mostafaomrani/golang-clean-web-api/api/routers"
)

func InitServer() {

	cfg := config.GetConfig()

	r := gin.New()

	r.Use(middlewares.Cors(cfg))
	r.Use(gin.Logger(), gin.Recovery(), middlewares.TestMiddleware(), middlewares.LimittByRequest())

	// Register validator
	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		err := val.RegisterValidation("mobile", validations.IranianMobileNumberValidatopr, true)
		if err != nil {
			return
		}
	}

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
