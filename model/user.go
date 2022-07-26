package model

type User struct {
	Id        uint   `gorm:"primaryKey;autoIncrement"`
	Email     string `gorm:"unique"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
}
