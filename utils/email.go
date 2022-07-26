package utils

import (
	"net/mail"

	"github.com/Kemosabe2911/employee-app-go/logger"
)

func ValidMailAddress(email string) bool {
	validEmail, err := mail.ParseAddress(email)
	if err != nil {
		return false
	}
	logger.Info(validEmail.Address)
	return true
}
