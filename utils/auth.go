package utils

import (
	"Golang-User-Auth/models"
	"os"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

//Exception struct
type Exception models.Exception

var e = godotenv.Load()

func JWTCertify(token string) bool {
	tokenSecret := os.Getenv("tokenSecret")
	tk := &models.Token{}
	_, err := jwt.ParseWithClaims(token, tk, func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenSecret), nil
	})
	if token == "" || err != nil {
		return false
	} else {
		return true
	}
}

func AdminVerify(token string) bool {
	role := GetRole(token)
	if role == "admin" && JWTCertify(token) == true {
		return true
	} else {
		return false
	}

}
