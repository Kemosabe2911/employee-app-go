package model

type Role struct {
	Id   uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Role string `json:"role"`
}
