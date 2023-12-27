package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mostafaomrani/golang-clean-web-api/api/handlers"
)

func TestRouter(r *gin.RouterGroup) {
	h := handlers.NewTestHandler()

	r.GET("/", h.Test)
	r.GET("/users", h.Users)
	r.GET("/user/:id", h.UserById)
	r.GET("/user/get-user-by-username/:username", h.UserByUsername)
	r.GET("/user/:id/accounts", h.Accounts)
	r.POST("/add-user/", h.AddUser)
	r.GET("/get-header/", h.Headerbinder2)
	r.GET("/QueryBinder1/", h.QueryBinder1)
	r.GET("/UriBinder/:uri", h.UriBinder)
	r.GET("/body", h.BodyBinder)
	r.GET("/form", h.FormBinder)
	r.GET("/formfile", h.FormFile)
}
