package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/rainierrr/go-crud/models"
)

type UserController struct{}

func (_ UserController) Index(c *gin.Context) {
	var model = models.UserModel{}
	var p, _ = model.GetAll()
	c.JSON(http.StatusOK, gin.H{
		"tasks": p,
	})
}

func (_ UserController) Create(c *gin.Context) {
	var model = models.UserModel{}
	p, err := model.CreateModel(c)

	if err != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(201, p)
	}
}
