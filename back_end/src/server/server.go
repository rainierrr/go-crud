package server

import "github.com/rainierrr/go-crud/models"

func Init() {
	DBConnect()
	models.AutoMigrate(db)
	r := router()
	r.Run(":8080")
}
