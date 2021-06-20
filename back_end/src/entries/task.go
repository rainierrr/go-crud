package entries

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	name     string
	category string
}
