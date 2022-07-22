package dto

type CreateEmployeeRequest struct {
	Name         string `json:"name"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	Age          int    `json:"age"`
	DepartmentID int    `json:"department_id"`
	RoleID       int    `json:"role_id"`
	Street       string `json:"street"`
	City         string `json:"city"`
	State        string `json:"state"`
}

type UpdateEmployeeRequest struct {
	Name         string `json:"name"`
	Username     string `json:"username"`
	Age          int    `json:"age"`
	IsActive     bool   `json:"is_active"`
	DepartmentID int    `json:"department_id"`
	RoleID       int    `json:"role_id"`
	Street       string `json:"street"`
	City         string `json:"city"`
	State        string `json:"state"`
}
