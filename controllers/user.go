package controllers

import (
	"Golang-User-Auth/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UserDashBoard(r *gin.Context) {
	token, _ := r.Cookie("token")
	resp := Authenticator(token, "", "")

	temp, _ := json.Marshal(resp["status"])
	status, _ := strconv.Unquote(string(temp))

	if status == "user" {
		verification := utils.GetVerify(token)
		if verification == "true" {
			r.HTML(http.StatusOK, "user.html", gin.H{
				"title": "User Dashboard",
				"ShowEmail": func() string {
					return utils.GetEmail(token)
				},
				"ShowName": func() string {
					return utils.GetName(token)
				},
				"LogOut": func() {
					Logout(r)
				},
			})
		} else {
			r.HTML(http.StatusOK, "user_verify.html", gin.H{
				"title": "User Dashboard",
				"ShowEmail": func() string {
					return utils.GetEmail(token)
				},
				"ShowName": func() string {
					return utils.GetName(token)
				},
				"LogOut": func() {
					Logout(r)
				},
			})
		}

	} else if status == "admin" {
		r.Redirect(http.StatusTemporaryRedirect, "/admin")
	} else {
		r.Redirect(http.StatusTemporaryRedirect, "/")
	}
}
