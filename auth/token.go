package auth

import (
	"time"

	"github.com/Kemosabe2911/employee-app-go/config"
	"github.com/Kemosabe2911/employee-app-go/logger"
	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT(email string, isAdmin bool) (string, error) {
	var signingKey = []byte(config.GetConfig().JwtSecretKey)

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["email"] = email
	claims["isAdmin"] = isAdmin
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	logger.Info(token)

	tokenString, err := token.SignedString(signingKey)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}
