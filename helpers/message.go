package helpers

import (
	"errors"
)

//errors
var (
	ErrInvalidRequest           = errors.New("unable to process request")
	ErrProductNotFound          = errors.New("product not found")
	ErrUnableToSaveOrder        = errors.New("unable to save order")
	ErrUnableToSaveOrderDetails = errors.New("unable to save orderdetails")
	ErrOrderNotFound            = errors.New("orders not found")
	ErrRoleNotFound             = errors.New("roles not found")
	ErrDepartmentNotFound       = errors.New("departments not found")
)
