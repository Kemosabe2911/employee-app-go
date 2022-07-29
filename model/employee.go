package model

type Employee struct {
	Id           uint       `json:"id" gorm:"primaryKey;autoIncrement"`
	Name         string     `json:"name"`
	Username     string     `json:"username" gorm:"unique"`
	Email        string     `json:"email" gorm:"unique"`
	Age          int        `json:"age"`
	IsActive     bool       `json:"is_active"`
	IdProof      string     `json:"id_proof"`
	DepartmentID int        `json:"department_id"`
	RoleID       int        `json:"role_id"`
	AddressID    int        `json:"address_id"`
	Department   Department `gorm:"foreignKey:DepartmentID"`
	Role         Role       `gorm:"foreignKey:RoleID"`
	Address      Address    `gorm:"foreignKey:AddressID"`
}
