package model

type Employee struct {
	Id           uint       `gorm:"primaryKey;autoIncrement"`
	Name         string     `json:"name"`
	Username     string     `gorm:"unique"`
	Email        string     `gorm:"unique"`
	Age          int        `json:"age"`
	IsActive     bool       `json:"isActive"`
	DepartmentID int        `json:"department_id"`
	RoleID       int        `json:"role_id"`
	AddressID    int        `json:"address_id"`
	Department   Department `gorm:"foreignKey:DepartmentID"`
	Role         Role       `gorm:"foreignKey:RoleID"`
	Address      Address    `gorm:"foreignKey:AddressID"`
}
