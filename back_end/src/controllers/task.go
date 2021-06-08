package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Tasks (c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"tasks": "test1",
	})
}