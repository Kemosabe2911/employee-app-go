package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/Kemosabe2911/employee-app-go/config"
	"github.com/Kemosabe2911/employee-app-go/logger"
	"github.com/dgrijalva/jwt-go"
)

var signingKey = []byte(config.GetConfig().JwtSecretKey)

type JWTClaim struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

type TokenStruct struct {
	Access  string
	Refresh string
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

func ValidateToken(signedToken string) (string, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return signingKey, nil
		})
	if err != nil || !token.Valid {
		logger.Infof("Error parsing access token or access token invalid: %s", err.Error())
		return "", errors.New("invalid token")
	}

	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return "", err
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return "", err
	}
	return claims.Email, nil
}

func GenerateAccessAndRefreshToken(email string) (interface{}, error) {
	access_token, err := GenerateAccessToken(email)
	if err != nil {
		logger.Error("Error while creating Access Token")
		return nil, err
	}
	logger.Info(access_token)

	refresh_token, err := GenerateRefreshToken(email)
	if err != nil {
		logger.Error("Error while creating Refresh Token")
		return nil, err
	}
	logger.Info(refresh_token)

	tokens := TokenStruct{
		Access:  access_token,
		Refresh: refresh_token,
	}
	return tokens, nil
}
