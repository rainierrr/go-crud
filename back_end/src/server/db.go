package server

import "github.com/jinzhu/gorm"

func DbConnect() *gorm.DB {
  DBMS     := "mysql"
  USER     := "go"
  PASS     := "go"
  PROTOCOL := "tcp(db:3306)"
  DBNAME   := "go_db"

  CONNECT := USER+":"+PASS+"@"+PROTOCOL+"/"+DBNAME
  db,err := gorm.Open(DBMS, CONNECT)

  if err != nil {
    panic(err.Error())
  }
  return db
}