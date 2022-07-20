package model

type Project struct {
	Id          string `gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsActive    bool   `json:"isActive"`
}
