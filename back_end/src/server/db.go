package server

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

var (
	db  *gorm.DB
	err error
)

func DBConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "go"
	PASS := "go"
	PROTOCOL := "tcp(db:3306)"
	DBNAME := "go_db"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err = gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return db
}

func GetDB() *gorm.DB {
	return db
}

func DBClose() {
	if err := db.Close(); err != nil {
		fmt.Println("test")
		panic(err.Error())
	}
}
