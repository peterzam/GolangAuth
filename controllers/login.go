package controllers

import (
	"Golang-User-Auth/utils"
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type ErrorResponse struct {
	Err string
}

type error interface {
	Error() string
}

var db = utils.ConnectDB()

var e = godotenv.Load()

func HomePage(r *gin.Context) {
	token, _ := r.Cookie("token")
	resp := Authenticator(token, "", "")

	temp, _ := json.Marshal(resp["status"])
	status, _ := strconv.Unquote(string(temp))

	if status == "user" {
		r.Redirect(http.StatusTemporaryRedirect, "/dashboard")
	} else if status == "admin" {
		r.Redirect(http.StatusTemporaryRedirect, "/admin")
	} else {
		LoginForm(r)
	}
}

func LoginForm(r *gin.Context) {

	email := r.PostForm("email")
	password := r.PostForm("password")

	resp := Authenticator("", email, password)
	temp, _ := json.Marshal(resp["status"])
	status, _ := strconv.Unquote(string(temp))
	SetSession(resp, r)
	if status == "user" {
		r.Redirect(http.StatusMovedPermanently, "/dashboard")
	} else if status == "admin" {
		r.Redirect(http.StatusMovedPermanently, "/admin")
	} else {
		{
			r.HTML(http.StatusOK, "index.html", gin.H{
				"title": "GoLang Authentication",
				"ShowMessage": func() string {
					temp, _ = json.Marshal(resp["message"])
					message, _ := strconv.Unquote(string(temp))
					return message
				},
			})
		}
	}

}

func SetSession(resp map[string]interface{}, r *gin.Context) {

	temp, _ := json.Marshal(resp["token"]) //Loggable:Response is broken
	token, _ := strconv.Unquote(string(temp))

	timeout, _ := strconv.Atoi(os.Getenv("tokenlifetime"))
	sitename := os.Getenv("sitename")
	tokenSecure, _ := strconv.ParseBool(os.Getenv("tokenSecure"))
	tokenHTTPOnly, _ := strconv.ParseBool(os.Getenv("tokenHTTPOnly"))

	r.SetCookie("token", token, timeout, "/", sitename, tokenSecure, tokenHTTPOnly)
}

func Logout(r *gin.Context) {

	CleanToken(r)
	r.Redirect(http.StatusTemporaryRedirect, "/")
}

func CleanToken(r *gin.Context) {
	sitename := os.Getenv("sitename")
	tokenHTTPOnly, _ := strconv.ParseBool(os.Getenv("tokenHTTPOnly"))
	tokenSecure, _ := strconv.ParseBool(os.Getenv("tokenSecure"))
	r.SetCookie("token", "", -1, "/", sitename, tokenSecure, tokenHTTPOnly)
	r.Redirect(http.StatusTemporaryRedirect, "/")
}
