package models

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rainierrr/go-crud/db"
	"github.com/rainierrr/go-crud/entries"
)

type UserModel struct{}

func (_ UserModel) GetAll() ([]entries.Task, error) {
	db := db.GetDB()
	var task []entries.Task
	if err := db.Model(&task).Scan(&task).Error; err != nil {
		return nil, err
	}
	return task, nil
}

func (_ UserModel) CreateModel(c *gin.Context) (entries.Task, error) {
	db := db.GetDB()
	task := entries.Task{}
	now := time.Now()
	task.CreatedAt = now
	task.UpdatedAt = now

	if err := c.BindJSON(&task); err != nil {
		return task, err
	}
	if err := db.Create(&task).Error; err != nil {
		return task, err
	}
	return task, nil
}
