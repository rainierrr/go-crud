package models

import (
	"github.com/rainierrr/go-crud/db"
	"github.com/rainierrr/go-crud/entries"
)

type TaskModel struct{}

func (_ TaskModel) GetAll() ([]entries.Task, error) {
	db := db.GetDB()
	var u []entries.Task
	if err := db.Table("tasks").Select("name, category, created_at, updated_at, deleted_at").Scan(&u).Error; err != nil {
		return nil, err
	}
	return u, nil
}
