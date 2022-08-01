package model

type Address struct {
	Id     uint   `json:"id" gorm:"primaryKey;autoautoIncrement"`
	Street string `json:"street"`
	City   string `json:"city"`
	State  string `json:"state"`
}
