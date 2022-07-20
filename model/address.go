package model

type Address struct {
	Id     string `gorm:"primaryKey"`
	Street string `json:"street"`
	City   string `json:"city"`
	State  string `json:"state"`
}
