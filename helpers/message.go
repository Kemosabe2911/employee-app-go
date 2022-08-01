package helpers

import (
	"errors"
)

//errors
var (
	ErrInvalidRequestError      = errors.New("unable to process request")
	ErrProductNotFoundError     = errors.New("product not found")
	ErrUnableToSaveOrder        = errors.New("unable to save order")
	ErrUnableToSaveOrderDetails = errors.New("unable to save orderdetails")
	ErrOrderNotFoundError       = errors.New("orders not found")
	ErrRoleNotFoundError        = errors.New("roles not found")
	ErrDepartmentNotFoundError  = errors.New("departments not found")
)
