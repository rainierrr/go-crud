package main

import (

	_ "github.com/go-sql-driver/mysql"
	"github.com/rainierrr/go-crud/server"
)



func main() {
	db := server.DbConnect()
	defer db.Close()
	r := server.GetRouter()
	r.Run(":8080")
}
