package entries

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type NullAbleTask struct {
	ID          uuid.UUID      `json:"id"`
	Title       string         `json:"title"`
	Explanation sql.NullString `json:"explanation"`
	Complete    bool           `json:"complete"`
	Due_date    sql.NullTime   `json:"due_date"`
	CreatedAt   time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at" db:"updated_at"`
}

type Task struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title" query:"title"`
	Explanation string    `json:"explanation"`
	Complete    bool      `json:"complete"`
	Due_date    time.Time `json:"due_date"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}
