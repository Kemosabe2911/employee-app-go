package helpers

import (
	"errors"
)

//errors
var (
	InvalidRequestError         = errors.New("unable to process request")
	EmployeeNotFoundError       = errors.New("employee not found")
	UnableToSaveEmployee        = errors.New("unable to save employee")
	UnableToSaveEmployeeDetails = errors.New("unable to save employeedetails")
)
