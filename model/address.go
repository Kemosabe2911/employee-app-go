package model

type Address struct {
	Id     uint   `gorm:"primaryKey;autoautoIncrement"`
	Street string `json:"street"`
	City   string `json:"city"`
	State  string `json:"state"`
}
