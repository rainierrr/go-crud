package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/rainierrr/go-crud/server"
)

func main() {
	server.Init()
}
