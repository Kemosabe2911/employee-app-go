package helpers

import (
	"errors"
)

//errors
var (
	InvalidRequestError      = errors.New("unable to process request")
	ProductNotFoundError     = errors.New("product not found")
	UnableToSaveOrder        = errors.New("unable to save order")
	UnableToSaveOrderDetails = errors.New("unable to save orderdetails")
	OrderNotFoundError       = errors.New("orders not found")
	RoleNotFoundError        = errors.New("roles not found")
	DepartmentNotFoundError  = errors.New("departments not found")
)
