package controllers

import (
	"Golang-User-Auth/models"
	"Golang-User-Auth/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RegisterHomePage(r *gin.Context) {
	r.HTML(http.StatusOK, "register.html", gin.H{
		"title": "Register",
	})

}

//CreateUser function -- create a new user
func RegisterForm(r *gin.Context) {

	user := &models.User{}
	//json.NewDecoder(r.Request.Body).Decode(user)
	user.Email = r.PostForm("email")
	user.Password = r.PostForm("password")
	user.Name = r.PostForm("name")
	user.Role = "user"
	user.Verify = "false"

	if utils.CheckName(user.Name) == false {
		r.HTML(http.StatusOK, "error_user_create.html", gin.H{
			"title": "Error User Creation",
			"ShowMessage": func() string {
				return "The name you entered is not valid."
			},
		})
	} else if utils.CheckEmail(user.Email) == false {
		r.HTML(http.StatusOK, "error_user_create.html", gin.H{
			"title": "Error User Creation",
			"ShowMessage": func() string {
				return "The email you entered is not a vaild email."
			},
		})
	} else if utils.CheckPassword(user.Password) == false {
		r.HTML(http.StatusOK, "error_user_create.html", gin.H{
			"title": "Error User Creation",
			"ShowMessage": func() string {
				return "Password must contains a Number, an UpperCase, a LowerCase, a Special Character and more than 8 characters."
			},
		})
	} else {
		pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			//fmt.Println(err)
			err := ErrorResponse{
				Err: "Password Encryption  failed",
			}
			r.JSON(http.StatusInternalServerError, err)
		}

		user.Password = string(pass)

		createdUser := db.Create(user)
		var errMessage = createdUser.Error

		if createdUser.Error != nil {
			fmt.Println(errMessage)

			r.HTML(http.StatusOK, "error_user_create.html", gin.H{
				"title": "Error User Creation",
				"ShowMessage": func() string {
					return "The email " + user.Email + " has already registered"
				},
			})
		} else {
			r.HTML(http.StatusOK, "user_created.html", gin.H{
				"title":   "User Created",
				"message": "User Created Successfully",
				"ShowName": func() string {
					return user.Name
				},
			})
		}
	}

}
