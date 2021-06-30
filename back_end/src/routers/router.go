package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/rainierrr/go-crud/controllers"
)

func Init() {
	r := router()
	r.Run(":8080")
}

func router() *gin.Engine {
	r := gin.Default()
	u := r.Group("/tasks")
	{
		ctrl := controllers.UserController{}
		u.GET("", ctrl.Index)
		u.POST("", ctrl.Create)
		u.PATCH("/:id", ctrl.Update)
		//u.DELETE("/:id", ctrl.Delete)
	}
	return r
}
