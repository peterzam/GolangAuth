package controllers

import (
	"Golang-User-Auth/models"
	"Golang-User-Auth/utils"
	"os"
	"strconv"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func Authenticator(token, email, password string) map[string]interface{} {
	if token == "" && (email != "" || password != "") {
		timeout, _ := strconv.Atoi(os.Getenv("tokenlifetime"))
		tokenSecret := os.Getenv("tokenSecret")
		user := &models.User{}

		if err := db.Where("Email = ?", email).First(user).Error; err != nil {
			var resp = map[string]interface{}{"status": "false", "message": "Email address not found"}
			//Loggable:Email not found
			return resp
		}

		expiresAt := time.Now().Add(time.Duration(timeout) * time.Second).Unix()
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
			var resp = map[string]interface{}{"status": "false", "message": "Invalid login credentials. Please try again"}
			//Loggable:Wrong Password
			return resp
		}

		tk := &models.Token{
			UserID: user.ID,
			Name:   user.Name,
			Email:  user.Email,
			Verify: user.Verify,
			Role:   user.Role,
			StandardClaims: &jwt.StandardClaims{
				ExpiresAt: expiresAt,
			},
		}

		token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)

		tokenString, err := token.SignedString([]byte(tokenSecret))
		if err != nil {
			Now := strconv.FormatInt(time.Now().Unix(), 16)
			var resp = map[string]interface{}{"status": "false", "message": "Internal error,Please contact Admin with this key (" + Now + ")"}
			//Loggable:Cannot encrypt token
			return resp
		}

		var resp = map[string]interface{}{"status": user.Role, "message": "Logged in", "token": tokenString}
		//Loggable:Logged in with any role.
		return resp

	} else if token == "" && (email == "" || password == "") {
		var resp = map[string]interface{}{"status": "blank", "message": ""}
		//Loggable:Logged in with any role.
		return resp
	} else {
		if utils.JWTCertify(token) {
			var resp = map[string]interface{}{"status": utils.GetRole(token), "message": "Logged in", "token": token}
			//Loggable:Logged in with any role.
			return resp
		} else {
			var resp = map[string]interface{}{"status": "false", "message": "Session expired or Invalid Cookies"}
			return resp
		}
	}

}
