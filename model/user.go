package model

type User struct {
	Id        uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Email     string `json:"email" gorm:"unique"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
}
