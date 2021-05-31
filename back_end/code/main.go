package main

import (
	"net/http"
	_ "github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world test",
		})
	})
	engine.Run(":8080")
}
