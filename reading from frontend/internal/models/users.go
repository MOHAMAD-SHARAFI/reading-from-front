package models

type User struct {
	Name  string `json:"name"`
	Email string `json:"email" gorm:"uniqueindex"`
	Age   int    `json:"age"`
}
