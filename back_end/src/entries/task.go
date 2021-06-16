package entries

type Task struct {
	ID string `gorm:"primaryKey"`
  Name string
  Category string
}