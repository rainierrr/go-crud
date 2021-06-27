package entries

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type NullableTask struct {
	ID          uuid.UUID      `json:"id"`
	Name        string         `json:"name"`
	Category    sql.NullString `json:"category"`
	Explanation sql.NullString `json:"explanation"`
	CreatedAt   time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at" db:"updated_at"`
}

type Task struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Category    string    `json:"category"`
	Explanation string    `json:"explanation"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}
