package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rainierrr/go-crud/db"
	"github.com/rainierrr/go-crud/routers"
)

func main() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	db.Init()
	routers.Init()
	db.Close()
}
