package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type User struct {
	Id       int    `json:"id"`
	Fullname string `json:"fullname"`
	Age      int    `json:"age"`
}

var Users = map[string]User{}

type HealthHandler struct {
}

func NewHeathHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Health(ctx *gin.Context) {
	var user *User
	ctx.JSON(http.StatusOK, user)
}

func (h *HealthHandler) PostHealth(ctx *gin.Context) {
	var userData *User
	err := json.NewDecoder(ctx.Request.Body).Decode(&userData)
	if err != nil {
		panic(err)
	}
	Users[strconv.Itoa(userData.Id)] = *userData
	ctx.JSON(http.StatusOK, userData)
}

func (h *HealthHandler) PostHealthById(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	user, exists := Users[id]

	if exists {
		ctx.JSON(http.StatusOK, user)
	} else {
		ctx.JSON(http.StatusNotFound, struct {
			Status string `json:"status"`
		}{
			Status: "User not found",
		})
	}
}
