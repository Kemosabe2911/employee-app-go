package dto

type CreateDepartment struct {
	Name           string `json:"name"`
	DepartmentRoom string `json:"department_room"`
	DepartmentCode string `json:"department_code"`
	Website        string `json:"website"`
}

type UpdateDepartment struct {
	Name           string `json:"name"`
	DepartmentRoom string `json:"department_room"`
	DepartmentCode string `json:"department_code"`
	Website        string `json:"website"`
}
