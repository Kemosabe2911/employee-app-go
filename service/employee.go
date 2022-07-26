package service

import (
	"strconv"

	"github.com/Kemosabe2911/employee-app-go/dto"
	"github.com/Kemosabe2911/employee-app-go/logger"
	"github.com/Kemosabe2911/employee-app-go/model"
	"github.com/Kemosabe2911/employee-app-go/repository"
	"gorm.io/gorm"
)

type EmployeeService interface {
	CreateEmployee(employeeRequest dto.CreateEmployeeRequest) *model.APIResponse
	GetAllEmployees() *model.APIResponse
	GetEmployeeById(string) *model.APIResponse
	DeleteEmployee(string) *model.APIResponse
	UpdateEmployee(string, dto.UpdateEmployeeRequest) *model.APIResponse
	UploadIdProof(id string, newFileName string) *model.APIResponse
}

type employeeService struct {
	employeeRepository repository.EmployeeRepository
	roleRepository     repository.RoleRepository
	// departmentRepository repository.DepartmentRepository
	DB *gorm.DB
}

func CreateEmployeeService(db *gorm.DB) *employeeService {
	return &employeeService{
		employeeRepository: repository.CreateEmployeeRepository(db),
		roleRepository:     repository.CreateRoleRepository(db),
		DB:                 db,
	}
}

func (es *employeeService) CreateEmployee(employeeRequest dto.CreateEmployeeRequest) *model.APIResponse {
	logger.Info("Start CreateEmployee in Service")
	address := model.Address{
		Street: employeeRequest.Street,
		City:   employeeRequest.City,
		State:  employeeRequest.State,
	}

	// address, err1 := es.employeeRepository.CreateAddress(address)
	// if err1 != nil {
	// 	logger.Error("Error while inserting address")
	// 	return &model.APIResponse{
	// 		StatusCode: 404,
	// 		Data: &model.ErrorStatus{
	// 			Message: "Cannot save address",
	// 		},
	// 	}
	// }

	// role, err2 := es.roleRepository.GetRoleById(strconv.Itoa(employeeRequest.RoleID))
	// if err2 != nil {
	// 	logger.Error("Error while getting role")
	// 	return &model.APIResponse{
	// 		StatusCode: 404,
	// 		Data: &model.ErrorStatus{
	// 			Message: "Cannot retrive role",
	// 		},
	// 	}
	// }

	employee := model.Employee{
		Name:         employeeRequest.Name,
		Username:     employeeRequest.Username,
		Email:        employeeRequest.Email,
		Age:          employeeRequest.Age,
		IsActive:     true,
		DepartmentID: employeeRequest.DepartmentID,
		RoleID:       employeeRequest.RoleID,
		Address:      address,
	}
	logger.Info(employee)

	employee, err := es.employeeRepository.CreateEmployee(employee)
	if err != nil {
		logger.Error("Error while saving product")
		return &model.APIResponse{
			StatusCode: 404,
			Data: &model.ErrorStatus{
				Message: "Cannot save employee",
			},
		}
	}
	logger.Info("Saved employee")
	return &model.APIResponse{
		StatusCode: 201,
		Data:       employee,
	}
}

func (es *employeeService) GetAllEmployees() *model.APIResponse {
	logger.Info("Start GetAllEmployees in Service")
	employee, err := es.employeeRepository.GetAllEmployees()
	if err != nil {
		logger.Error("Error in service")
		return &model.APIResponse{
			StatusCode: 404,
			Data: &model.ErrorStatus{
				Message: "Employees not found",
			},
		}
	}

	// var employeeAll []model.Employee

	// for i, data := range employee {
	// 	logger.Info(i, data, data.AddressID, data.RoleID)
	// 	address, err1 := es.employeeRepository.GetAddressById(data.AddressID)
	// 	if err1 != nil {
	// 		logger.Error("Error in service")
	// 		return &model.APIResponse{
	// 			StatusCode: 404,
	// 			Data: &model.ErrorStatus{
	// 				Message: "Address not found",
	// 			},
	// 		}
	// 	}
	// 	data.Address = address

	// 	role, err2 := es.roleRepository.GetRoleById(strconv.Itoa(data.RoleID))
	// 	if err2 != nil {
	// 		logger.Error("Error in service")
	// 		return &model.APIResponse{
	// 			StatusCode: 404,
	// 			Data: &model.ErrorStatus{
	// 				Message: "Role not found",
	// 			},
	// 		}
	// 	}
	// 	data.Role = role
	// 	employeeAll = append(employeeAll, data)
	// }

	// logger.Info(employeeAll)
	logger.Info("End GetAllEmployees in Service")
	return &model.APIResponse{
		StatusCode: 200,
		Data:       employee,
	}
}

func (es *employeeService) GetEmployeeById(id string) *model.APIResponse {
	logger.Info("Started GetEmployeeById in Service")
	employee, err := es.employeeRepository.GetEmployeeById(id)
	if err != nil {
		logger.Error("Error in service")
		return &model.APIResponse{
			StatusCode: 404,
			Data: &model.ErrorStatus{
				Message: "Employee not found",
			},
		}
	}

	// address, err1 := es.employeeRepository.GetAddressById(employee.AddressID)
	// if err1 != nil {
	// 	logger.Error("Error in service")
	// 	return &model.APIResponse{
	// 		StatusCode: 404,
	// 		Data: &model.ErrorStatus{
	// 			Message: "Address not found",
	// 		},
	// 	}
	// }
	// employee.Address = address

	// role, err2 := es.roleRepository.GetRoleById(strconv.Itoa(employee.RoleID))
	// if err2 != nil {
	// 	logger.Error("Error in service")
	// 	return &model.APIResponse{
	// 		StatusCode: 404,
	// 		Data: &model.ErrorStatus{
	// 			Message: "Role not found",
	// 		},
	// 	}
	// }
	// employee.Role = role

	logger.Info("End GetEmployeeById in Service")
	return &model.APIResponse{
		StatusCode: 200,
		Data:       employee,
	}
}

func (es *employeeService) DeleteEmployee(id string) *model.APIResponse {
	logger.Info("Start DeleteEmployee in Service")
	employee, err := es.employeeRepository.GetEmployeeById(id)
	logger.Info(employee)
	if err != nil {
		logger.Error("Error in service")
		return &model.APIResponse{
			StatusCode: 404,
			Data: &model.ErrorStatus{
				Message: "Employee not found",
			},
		}
	}

	err = es.employeeRepository.DeleteEmployee(id)
	if err != nil {
		logger.Error("Error while deleting employee")
		return &model.APIResponse{
			StatusCode: 400,
			Data:       "Failed to delete",
		}
	}

	logger.Info("Deleted Employee")
	return &model.APIResponse{
		StatusCode: 200,
		Data:       "Successfully Deleted",
	}
}

func (es *employeeService) UpdateEmployee(id string, employeeRequest dto.UpdateEmployeeRequest) *model.APIResponse {
	logger.Info("Start UpdateEmployee - Service")
	employee, err := es.employeeRepository.GetEmployeeById(id)
	logger.Info(employee)
	email := employee.Email
	if err != nil {
		logger.Error("Error in service")
		return &model.APIResponse{
			StatusCode: 404,
			Data: &model.ErrorStatus{
				Message: "Employee not found",
			},
		}
	}
	address := model.Address{
		Street: employeeRequest.Street,
		City:   employeeRequest.City,
		State:  employeeRequest.State,
	}
	// logger.Info(address)
	address, err = es.employeeRepository.UpdateAddress(strconv.Itoa(employee.AddressID), address)
	if err != nil {
		logger.Error("Error while updating address")
		return &model.APIResponse{
			StatusCode: 404,
			Data: &model.ErrorStatus{
				Message: "Cannot update address",
			},
		}
	}

	employee = model.Employee{
		Name:         employeeRequest.Name,
		Username:     employeeRequest.Username,
		Email:        email,
		Age:          employeeRequest.Age,
		IsActive:     employeeRequest.IsActive,
		DepartmentID: employeeRequest.DepartmentID,
		RoleID:       employeeRequest.RoleID,
		Address:      address,
	}
	logger.Info(employee)

	employee, ok := es.employeeRepository.UpdateEmployee(id, employee)
	if ok != nil {
		logger.Error("Error while address employee")
		return &model.APIResponse{
			StatusCode: 404,
			Data: &model.ErrorStatus{
				Message: "Cannot update employee",
			},
		}
	}
	logger.Info("Updated employee")
	return &model.APIResponse{
		StatusCode: 201,
		Data:       employee,
	}
}

func (es *employeeService) UploadIdProof(id string, newFileName string) *model.APIResponse {
	logger.Info("Start UploadIdProof in Service")
	employeeData, err := es.employeeRepository.UploadIdProof(id, newFileName)
	if err != nil {
		logger.Error("Error while updating file name field")
		return &model.APIResponse{
			StatusCode: 404,
			Data: &model.ErrorStatus{
				Message: "Cannot update file name",
			},
		}
	}
	logger.Info("End UploadIdProof in Service")
	return &model.APIResponse{
		StatusCode: 201,
		Data:       employeeData,
	}
}
