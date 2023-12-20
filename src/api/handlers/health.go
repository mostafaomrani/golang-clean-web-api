package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
}

func NewHeathHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Get Working!")
}

func (h *HealthHandler) PostHealth(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Post Working!")
}

func (h *HealthHandler) PostHealthById(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	ctx.JSON(http.StatusOK, fmt.Sprintf("Post Working ID : %s!", id))
}
