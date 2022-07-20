package model

type Role struct {
	Id   uint   `gorm:"primaryKey;autoIncrement"`
	Role string `json:"role"`
}
