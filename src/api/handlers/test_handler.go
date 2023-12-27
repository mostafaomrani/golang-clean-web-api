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

func (h *TestHandler) Headerbinder2(c *gin.Context) {
	header := struct {
		UserId  string
		Browser string
	}{}
	c.BindHeader(&header)
	c.JSON(http.StatusOK, gin.H{
		"result": "HeaderBinder2",
		"header": header,
	})

}

func (h *TestHandler) QueryBinder1(c *gin.Context) {
	id := c.Query("id")
	//ids := c.QueryArray("id")
	c.JSON(http.StatusOK, gin.H{
		"result": "QueryBinder1",
		"header": id,
	})

}

func (h *TestHandler) UriBinder(c *gin.Context) {
	uri := c.Param("uri")
	//ids := c.QueryArray("id")
	c.JSON(http.StatusOK, gin.H{
		"result": "UriBinder",
		"uri":    uri,
	})

}

func (h *TestHandler) BodyBinder(c *gin.Context) {
	body := struct {
		UserId  string
		Browser string
	}{}
	c.BindJSON(&body)
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"result": "BodyBinder",
		"body":   body,
	})

}

func (h *TestHandler) FormBinder(c *gin.Context) {
	form := struct {
		FirstName string
		LastName  string
	}{}
	c.Bind(&form)
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"result": "FormBinder",
		"form":   form,
	})

}

func (h *TestHandler) FormFile(c *gin.Context) {
	file, _ := c.FormFile("file")
	err := c.SaveUploadedFile(file, "files/")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"result": "FormFile",
		"form":   file,
	})

}
