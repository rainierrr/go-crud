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
		if err := rows.StructScan(&nullAbleTask); err != nil {
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

func (_ UserModel) UpdateModel(c *gin.Context) (entries.Task, error) {
	db := db.GetDB()
	task := entries.Task{}
	id := c.Param("id")

	if err := c.BindJSON(&task); err != nil {
		return task, err
	}

	query := "UPDATE tasks SET name=?, category=?, explanation=? WHERE id = ?"

	if _, err := db.Exec(query, task.Name, task.Category, task.Explanation, id); err != nil {
		return task, err
	}

	if err := db.Get(&task, "SELECT * FROM tasks where id = ?", id); err != nil {
		return task, err
	}

	return task, nil
}

func (_ UserModel) DeleteModel(c *gin.Context) error {
	db := db.GetDB()
	id := c.Param("id")

	query := "DELETE FROM tasks WHERE id = ?"
	if _, err := db.Exec(query, id); err != nil {
		return err
	}

	return nil
}
