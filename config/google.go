package config

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func SetGoogleConfig() *oauth2.Config {
	googleConfig := &oauth2.Config{
		ClientID:     GetConfig().GoogleClientID,
		ClientSecret: GetConfig().GoogleClientSecret,
		RedirectURL:  "http://localhost:8080/v1/google/callback",
		Scopres: []string{
			"https://googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	}
}
