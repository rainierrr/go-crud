package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/rainierrr/go-crud/db"
	"github.com/rainierrr/go-crud/routers"
)

func main() {
	db.Init()
	routers.Init()
	db.Close()
}
