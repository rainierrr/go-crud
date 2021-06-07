package server

import 	(
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine{
	engine := gin.Default()
	engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world testtrt",
		})
	})
	return engine
}