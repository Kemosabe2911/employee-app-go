package config

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type GoogleConfig struct {
	GoogleLoginConfig oauth2.Config
}

var AppConfig GoogleConfig

const OauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

func LoadConfig() {
	AppConfig.GoogleLoginConfig = oauth2.Config{
		ClientID:     GetConfig().GoogleClientID,
		ClientSecret: GetConfig().GoogleClientSecret,
		RedirectURL:  "http://localhost:8080/google/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	}
}
