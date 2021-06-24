package db

import (
	"github.com/jinzhu/gorm"
	"github.com/rainierrr/go-crud/entries"
)

var (
	db  *gorm.DB
	err error
)

func Init() {
	DBConnect()
	entries.AutoMigrate(db)
}

func DBConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "go"
	PASS := "go"
	PROTOCOL := "tcp(db:3306)"
	DBNAME := "go_db"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
	db, err = gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return db
}

func GetDB() *gorm.DB {
	return db
}

func Close() {
	if err := db.Close(); err != nil {
		panic(err.Error())
	}
}
