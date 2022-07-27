package auth

import (
	"time"

	"github.com/Kemosabe2911/employee-app-go/config"
	"github.com/dgrijalva/jwt-go"
)

var signingKey = []byte(config.GetConfig().JwtSecretKey)

type JWTClaim struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func generateJwt(email string, expTime time.Time) (string, error) {
	claims := &JWTClaim{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GenerateAccessToken(email string) (string, error) {

	expirationTime := time.Now().Add(15 * time.Minute)
	tokenString, err := generateJwt(email, expirationTime)
	return tokenString, err

}

func GenerateRefreshToken(email string) (string, error) {

	expirationTime := time.Now().Add(24 * 90 * time.Hour)
	tokenString, err := generateJwt(email, expirationTime)
	return tokenString, err

}
