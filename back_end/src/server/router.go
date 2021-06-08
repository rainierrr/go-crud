package server

import 	(
	"github.com/rainierrr/go-crud/controllers"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine{
	r := gin.Default()
	r.GET("/", controllers.Tasks)
	return r
}