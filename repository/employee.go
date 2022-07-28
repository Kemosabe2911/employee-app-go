package repository

import (
	"github.com/Kemosabe2911/employee-app-go/logger"
	"github.com/Kemosabe2911/employee-app-go/model"
	"gorm.io/gorm"
)

type EmployeeRepository interface {
	CreateEmployee(model.Employee) (model.Employee, error)
	GetAllEmployees() ([]model.Employee, error)
	GetEmployeeById(string) (model.Employee, error)
	CreateAddress(model.Address) (model.Address, error)
	GetAddressById(int) (model.Address, error)
	DeleteEmployee(string) error
	UpdateEmployee(string, model.Employee) (model.Employee, error)
	UpdateAddress(string, model.Address) (model.Address, error)
	GetEmployeeByEmail(string) (model.Employee, error)
	// UploadIdProof(id string, newFileName string) (model.Employee, error)
}

type employeeRepository struct {
	DB *gorm.DB
}

func CreateEmployeeRepository(db *gorm.DB) *employeeRepository {
	return &employeeRepository{
		DB: db,
	}
}

func (er *employeeRepository) CreateEmployee(employee model.Employee) (model.Employee, error) {
	logger.Info("Start CreateEmployee in Repo")
	err := er.DB.Create(&employee).Preload("Department").Preload("Role").Preload("Address").Preload("Department.Department").First(&employee).Error
	logger.Info("End CreateEmployee in Repo")
	return employee, err
}

func (er *employeeRepository) GetAllEmployees() ([]model.Employee, error) {
	logger.Info("Start GetAllEmployees in Repo")
	var employee []model.Employee
	err := er.DB.Preload("Address").Preload("Role").Preload("Department").Preload("Department.Department").Find(&employee).Error
	logger.Info("End GetAllEmployees in Repo")
	return employee, err
}

func (er *employeeRepository) GetEmployeeById(id string) (model.Employee, error) {
	logger.Info("Started GetEmployeeById in Repo")
	var employee model.Employee
	err := er.DB.Preload("Address").Preload("Role").Preload("Department").Preload("Department.Department").First(&employee, "id = ?", id).Error
	logger.Info("Ended GetEmployeeById in Repo")
	return employee, err
}

func (er *employeeRepository) CreateAddress(address model.Address) (model.Address, error) {
	logger.Info("Start CreateAddress in Repo")
	err := er.DB.Create(&address).Error
	logger.Info("End CreateAddress in Repo")
	return address, err
}

func (er *employeeRepository) GetAddressById(id int) (model.Address, error) {
	logger.Info("Started GetAddressById in Repo")
	var address model.Address
	err := er.DB.First(&address, "id = ?", id).Error
	logger.Info("Ended GetAddressById in Repo")
	return address, err
}

func (er *employeeRepository) DeleteEmployee(id string) error {
	logger.Info("Start DeleteEmployee in Repo")
	var employee model.Employee
	err := er.DB.Delete(&employee, "id = ?", id).Error
	logger.Info("End DeleteEmployee in Repo")
	return err
}

func (er *employeeRepository) UpdateEmployee(id string, employee model.Employee) (model.Employee, error) {
	logger.Info("Started UpdateEmployee in Repo")
	var employeeData model.Employee
	err := er.DB.Preload("Address").Preload("Role").Preload("Department").Preload("Department.Department").First(&employeeData, "id = ?", id).Error
	if err != nil {
		logger.Error("Employee not found")
		return employeeData, err
	}
	logger.Info(employeeData, employee)
	// err = er.DB.Session(&gorm.Session{FullSaveAssociations: true}).Where("id = ?", id).Updates(&employee).Preload("Address").Preload("Role").Preload("Department").Preload("Department.Department").First(&employee, "id = ?", id).Error
	err = er.DB.Model(&employeeData).Updates(map[string]interface{}{
		"name":       employee.Name,
		"Username":   employee.Username,
		"Email":      employee.Email,
		"age":        employee.Age,
		"is_active":  employee.IsActive,
		"Department": employee.Department,
		"Role":       employee.Role,
		"Address":    employee.Address,
	}).Preload("Address").Preload("Role").Preload("Department").Preload("Department.Department").First(&employee, "id = ?", id).Error
	logger.Info("Ended UpdateEmployee in Repo")
	return employee, err
}

func (er *employeeRepository) UpdateAddress(id string, address model.Address) (model.Address, error) {
	logger.Info("Started UpdateAddress in Repo")
	err := er.DB.Where("id = ?", id).Updates(&address).Error
	logger.Info("Ended UpdateAddress in Repo")
	return address, err
}

func (er *employeeRepository) GetEmployeeByEmail(email string) (model.Employee, error) {
	logger.Info("Started GetEmployeeById in Repo")
	var employee model.Employee
	err := er.DB.Preload("Address").Preload("Role").Preload("Department").Preload("Department.Department").First(&employee, "email = ?", email).Error
	logger.Info("Ended GetEmployeeById in Repo")
	return employee, err
}

// func (er *employeeRepository) UploadIdProof(id string, newFileName string) (model.Employee, error) {
// 	logger.Info("Start UploadIdProof in Repo")
// 	var employeeData model.Employee
// 	if err := er.DB.Where("id = ?", id).First(&employeeData).Error; err != nil {
// 		logger.Error("Error while fetching from employee repo", err.Error())
// 		return model.Employee{}, err
// 	}
// 	err := er.DB.Model(&employeeData).Update("IdProof", newFileName).Preload("Address").Preload("Role").Preload("Department").Preload("Department.Department").First(&employeeData).Error
// 	logger.Info("End UploadIdProof in Repo")
// 	return employeeData, err
// }
