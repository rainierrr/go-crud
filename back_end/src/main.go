package main

import (

	_ "github.com/go-sql-driver/mysql"
	"github.com/rainierrr/go-crud/server"
	"github.com/rainierrr/go-crud/entries"
)


func main() {
	db := server.DbConnect()
	entries.AutoMigrate(db)
	defer db.Close()
	r := server.GetRouter()
	r.Run(":8080")
}
