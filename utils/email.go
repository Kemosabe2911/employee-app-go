package utils

import "net/mail"

func ValidMailAddress(email string) (string, bool) {
	validEmail, err := mail.ParseAddress(email)
	if err != nil {
		return "", false
	}
	return validEmail.Address, true
}
