package model

type Employee struct {
	Id           string     `gorm:"primaryKey"`
	Name         string     `json:"name"`
	Username     string     `json:"username"`
	Password     string     `json:"password"`
	Age          int        `json:"age"`
	IsActive     bool       `json:"isActive"`
	DepartmentID string     `json:"department_id"`
	RoleID       string     `json:"role_id"`
	AddressID    string     `json:"address_id"`
	Department   Department `gorm:"foreignKey:DepartmentID"`
	Role         Role       `gorm:"foreignKey:RoleID"`
	Address      Address    `gorm:"foreignKey:AddressID"`
}
