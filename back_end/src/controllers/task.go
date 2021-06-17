package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (_ UserController) Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"tasks": "test1",
	})
}
