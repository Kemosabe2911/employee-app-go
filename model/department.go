package model

type Department struct {
	Id                  string            `gorm:"primaryKey"`
	Name                string            `json:"name"`
	DepartmentDetailsID string            `json:"department_details_id"`
	Department          DepartmentDetails `gorm:"foreignKey:DepartmentDetailsID"`
}

type DepartmentDetails struct {
	Id             string `gorm:"primarykey"`
	DepartmentRoom string `json:"department_room"`
	DepartmentCode string `json:"department_code"`
	Website        string `json:"website"`
}
