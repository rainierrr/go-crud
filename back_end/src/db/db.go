package db

import (
	"github.com/jmoiron/sqlx"
)

var (
	db  *sqlx.DB
	err error
)

func Init() {
	DBConnect()
}

func DBConnect() *sqlx.DB {
	DBMS := "mysql"
	USER := "go"
	PASS := "go"
	PROTOCOL := "tcp(db:3306)"
	DBNAME := "go_db"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
	db, err = sqlx.Connect(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return db
}

func GetDB() *sqlx.DB {
	return db
}

func Close() {
	if err := db.Close(); err != nil {
		panic(err.Error())
	}
}
