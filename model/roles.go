package model

type Role struct {
	Id   string `gorm:"primaryKey"`
	Role string `json:"role"`
}
