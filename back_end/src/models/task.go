package models

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rainierrr/go-crud/db"
	"github.com/rainierrr/go-crud/entries"
)

type UserModel struct{}

func (_ UserModel) GetAll() ([]entries.Task, error) {
	db := db.GetDB()
	var task entries.Task
	var tasks []entries.Task
	var nullAbleTask entries.NullableTask
	rows, err := db.Queryx("SELECT * FROM tasks")

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err := rows.StructScan(&nullAbleTask)
		if err != nil {
			return nil, err
		}
		task = entries.Task{
			ID:          nullAbleTask.ID,
			Name:        nullAbleTask.Name,
			Category:    nullAbleTask.Category.String,
			Explanation: nullAbleTask.Explanation.String,
			CreatedAt:   nullAbleTask.CreatedAt,
			UpdatedAt:   nullAbleTask.UpdatedAt,
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (_ UserModel) CreateModel(c *gin.Context) (entries.Task, error) {
	db := db.GetDB()
	task := entries.Task{}
	uuid, err := uuid.NewRandom()
	if err != nil {
		return task, err
	}
	task.ID = uuid

	if err := c.BindJSON(&task); err != nil {
		return task, err
	}

	if _, err := db.Exec("insert into tasks (id, name, category, explanation) values(?,?,?,?)", task.ID, task.Name, task.Category, task.Explanation); err != nil {
		return task, err
	}
	return task, nil
}
