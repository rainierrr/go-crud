package models

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rainierrr/go-crud/db"
	"github.com/rainierrr/go-crud/entries"
)

type UserModel struct{}

func (_ UserModel) GetAll() ([]entries.Task, error) {
	db := db.GetDB()
	var task []entries.Task
	if err := db.Table("tasks").Select("name, category, created_at, updated_at, deleted_at").Scan(&task).Error; err != nil {
		return nil, err
	}

	//log.Printf("get!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	return task, nil
}

func (_ UserModel) CreateModel(c *gin.Context) (entries.Task, error) {
	db := db.GetDB()
	task := entries.Task{}
	now := time.Now()
	task.CreatedAt = now
	task.UpdatedAt = now

	log.Println("c.Query:" + c.Query("ID"))
	if err := c.Bind(&task); err != nil {
		log.Printf("%#v", task)

		return task, err
	}
	log.Printf("%#v", task)
	if err := db.Create(&task).Error; err != nil {
		return task, err
	}
	return task, nil
}
