package dto

type CreateEmployeeRequest struct {
	Name         string `json:"name"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Age          int    `json:"age"`
	DepartmentID int    `json:"department_id"`
	RoleID       int    `json:"role_id"`
	AddressID    int    `json:"address_id"`
}
