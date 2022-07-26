package utils

import "golang.org/x/crypto/bcrypt"

func VerfityPassword(password string, confirm string) bool {
	if password == confirm {
		return true
	} else {
		return false
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
