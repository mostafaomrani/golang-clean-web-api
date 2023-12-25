package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type TestHandler struct {
}

func NewTestHandler() *TestHandler {
	return &TestHandler{}
}

func (T *TestHandler) Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"result": "Test",
	})
}

func (T *TestHandler) Users(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"result": "Test",
	})
}

func (T *TestHandler) UserById(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"result": "Test",
		"id":     id,
	})
}

func (T *TestHandler) UserByUsername(c *gin.Context) {
	username := c.Param("username")
	c.JSON(http.StatusOK, gin.H{
		"result":   "Test",
		"username": username,
	})
}

func (T *TestHandler) Accounts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"result": "Test",
	})
}

func (T *TestHandler) AddUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"result": "Test",
	})
}
