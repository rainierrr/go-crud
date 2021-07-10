package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/rainierrr/go-crud/models"
)

type UserController struct{}

func (_ UserController) Index(c *gin.Context) {
	var model = models.UserModel{}
	var tasks, err = model.GetAll()
	if err != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"tasks": tasks,
		})
	}

}

func (_ UserController) Create(c *gin.Context) {
	var model = models.UserModel{}
	task, err := model.CreateModel(c)

	if err != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(201, task)
	}
}

func (_ UserController) Update(c *gin.Context) {
	var model = models.UserModel{}
	task, err := model.UpdateModel(c)

	if err != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(201, task)
	}
}

func (_ UserController) Delete(c *gin.Context) {
	var model = models.UserModel{}

	if err := model.DeleteModel(c); err != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.Status(200)
	}
}
