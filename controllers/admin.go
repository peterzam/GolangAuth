package controllers

import (
	"Golang-User-Auth/models"
	"Golang-User-Auth/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func AdminDashBoard(r *gin.Context) {
	token, _ := r.Cookie("token")
	resp := Authenticator(token, "", "")

	temp, _ := json.Marshal(resp["status"])
	status, _ := strconv.Unquote(string(temp))

	if status == "user" {
		r.Redirect(http.StatusTemporaryRedirect, "/dashboard")
	} else if status == "admin" {
		r.HTML(http.StatusOK, "admin.html", gin.H{
			"title": "Admin Dashboard",
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
		r.Redirect(http.StatusTemporaryRedirect, "/")
	}
}

func ListUser(r *gin.Context) {
	token, _ := r.Cookie("token") //Loggable:Cookie is broken
	if utils.AdminVerify(token) {
		var users []models.User
		db.Preload(clause.Associations).Find(&users)
		r.JSON(http.StatusOK, users)
	} else {
		r.Redirect(http.StatusTemporaryRedirect, "/")

	}

}

func UpdateUser(r *gin.Context) {
	user := &models.User{}
	id := r.Param("id")
	db.First(&user, id)
	json.NewDecoder(r.Request.Body).Decode(user)
	db.Save(&user)
}

func DeleteUser(r *gin.Context) {
	token, _ := r.Cookie("token") //Loggable:Cookie is broken
	if utils.AdminVerify(token) {
		id := r.Param("id")
		var user models.User
		db.First(&user, id)
		//db.Delete(&user)
		db.Unscoped().Delete(&user) //permenent delete
	} else {
		r.Redirect(http.StatusTemporaryRedirect, "/")

	}

}
