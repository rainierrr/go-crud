package models

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rainierrr/go-crud/db"
	"github.com/rainierrr/go-crud/entries"
)

type UserModel struct{}

func (model UserModel) GetAll() ([]entries.Task, error) {
	db := db.GetDB()
	var task entries.Task
	var tasks []entries.Task
	var nullAbleTask entries.NullAbleTask
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
			Title:       nullAbleTask.Title,
			Complete:    nullAbleTask.Complete,
			Explanation: nullAbleTask.Explanation.String,
			Due_date:    nullAbleTask.Due_date.Time,
			CreatedAt:   nullAbleTask.CreatedAt,
			UpdatedAt:   nullAbleTask.UpdatedAt,
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (model UserModel) CreateModel(c *gin.Context) (entries.Task, error) {
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

	if _, err := db.Exec("insert into tasks (id, title, explanation) values(?,?,?)", task.ID, task.Title, task.Explanation); err != nil {
		return task, err
	}
	return task, nil
}

func (model UserModel) UpdateModel(c *gin.Context) (entries.Task, error) {
	db := db.GetDB()
	task := entries.Task{}
	nullAbleTask := entries.NullAbleTask{}
	id := c.Param("id")

	if err := c.BindJSON(&task); err != nil {
		return task, err
	}

	query := "UPDATE tasks SET title=?, complete=?, explanation=? WHERE id = ?"

	if _, err := db.Exec(query, task.Title, task.Complete, task.Explanation, id); err != nil {
		return task, err
	}

	if err := db.Get(&nullAbleTask, "SELECT * FROM tasks where id = ?", id); err != nil {
		log.Println(err.Error())
		return task, err
	}

	task = entries.Task{
		ID:          nullAbleTask.ID,
		Title:       nullAbleTask.Title,
		Complete:    nullAbleTask.Complete,
		Explanation: nullAbleTask.Explanation.String,
		Due_date:    nullAbleTask.Due_date.Time,
		CreatedAt:   nullAbleTask.CreatedAt,
		UpdatedAt:   nullAbleTask.UpdatedAt,
	}
	return task, nil
}

func (model UserModel) DeleteModel(c *gin.Context) error {
	db := db.GetDB()
	id := c.Param("id")

	query := "DELETE FROM tasks WHERE id = ?"
	if _, err := db.Exec(query, id); err != nil {
		return err
	}

	return nil
}
