package main

import (
	"net/http"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gin-gonic/gin"
)
func gormConnect() *gorm.DB {
  DBMS     := "mysql"
  USER     := "go"
  PASS     := "go"
  PROTOCOL := "tcp(db:3306)"
  DBNAME   := "go_db"

  CONNECT := USER+":"+PASS+"@"+PROTOCOL+"/"+DBNAME
  db,err := gorm.Open(DBMS, CONNECT)

  if err != nil {
    panic(err.Error())
  }
  return db
}

func main() {
	db := gormConnect()
  defer db.Close()
	engine := gin.Default()
	engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world test",
		})
	})
	engine.Run(":8080")
}
