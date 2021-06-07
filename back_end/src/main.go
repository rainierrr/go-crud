package main

import (

	_ "github.com/go-sql-driver/mysql"
	"github.com/rainierrr/go-crud/server"
)



func main() {
	db := server.DbConnect()
	defer db.Close()
	engine := server.GetRouter()
	engine.Run(":8080")
}
