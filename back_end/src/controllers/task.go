package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/rainierrr/go-crud/models"
)

type UserController struct{}

func (_ UserController) Index(c *gin.Context) {
	var model = models.TaskModel{}
	var p, _ = model.GetAll()
	c.JSON(http.StatusOK, gin.H{
		"tasks": p,
	})
}
