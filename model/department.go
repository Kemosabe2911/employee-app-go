package model

type Department struct {
	Id                  uint              `json:"id" gorm:"primaryKey;autoIncrement"`
	Name                string            `json:"name"`
	DepartmentDetailsID int               `json:"department_details_id"`
	Department          DepartmentDetails `gorm:"foreignKey:DepartmentDetailsID"`
}

type DepartmentDetails struct {
	Id             uint   `json:"id" gorm:"primarykey;autoIncrement"`
	DepartmentRoom string `json:"department_room"`
	DepartmentCode string `json:"department_code" gorm:"unique"`
	Website        string `json:"website"`
}
