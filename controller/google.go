package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Kemosabe2911/employee-app-go/config"
	"github.com/Kemosabe2911/employee-app-go/logger"
)

type GoogleResponse struct {
	Id            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Picture       string `json:"picture"`
	Hd            string `json:"hd"`
}

func GoogleLogin(w http.ResponseWriter, r *http.Request) {
	// if c.Request.Method != "GET" {
	// 	c.JSON(500, "Method not allowed")
	// 	return
	// }

	// url := config.AppConfig.GoogleLoginConfig.AuthCodeURL("random")
	// logger.Info(url)
	// c.Redirect(http.StatusTemporaryRedirect, url)

	if r.Method != "GET" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	url := config.AppConfig.GoogleLoginConfig.AuthCodeURL("random")
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func GoogleCallback(w http.ResponseWriter, r *http.Request) {
	// if c.Request.Method != "GET" {
	// 	c.JSON(500, "Method not allowed")
	// 	return
	// }

	// state := c.PostForm("state")
	// logger.Info(state)

	if r.Method != "GET" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	state := r.FormValue("state")
	logger.Info(state)
	code := r.FormValue("code")
	w.Header().Add("content-type", "application/json")

	if state != "random" {
		http.Redirect(w, r, "/v1/google/login", http.StatusTemporaryRedirect)
		fmt.Fprintf(w, "invalid oauth google state")
		return
	}

	// Exchange Auth Code for Tokens
	token, err := config.AppConfig.GoogleLoginConfig.Exchange(
		context.Background(), code)

	// ERROR : Auth Code Exchange Failed
	if err != nil {
		fmt.Fprintf(w, "falied code exchange: %s", err.Error())
		return
	}

	// Fetch User Data from google server
	response, err := http.Get(config.OauthGoogleUrlAPI + token.AccessToken)

	// ERROR : Unable to get user data from google
	if err != nil {
		fmt.Fprintf(w, "failed getting user info: %s", err.Error())
		return
	}

	// Parse user data JSON Object
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Fprintf(w, "failed read response: %s", err.Error())
		return
	}
	// data := string(contents)
	// logger.Info(data)

	var resp GoogleResponse
	err = json.Unmarshal(contents, &resp)

	if err != nil {
		panic(err)
		// logger.Errorf("Can't unmarshal %+v ", err)
		// return
	}

	logger.Info(resp)
	logger.Info(resp.Email)
	// logger.Info(resp.Email)
	// send back response to browser
	fmt.Fprintln(w, string(contents))
}
