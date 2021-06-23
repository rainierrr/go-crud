package entries

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Name     string `json:"name"`
	Category string `json:"category"`
}
