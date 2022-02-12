package utils

import (
	"Golang-User-Auth/models"
	"os"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
)

func GetID(token string) string {
	tokenSecret := os.Getenv("tokenSecret")
	tk := &models.Token{}
	jwt.ParseWithClaims(token, tk, func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenSecret), nil
	})
	ID := strconv.FormatUint(uint64(tk.UserID), 10)
	return ID
}

func GetEmail(token string) string {
	tokenSecret := os.Getenv("tokenSecret")
	tk := &models.Token{}
	jwt.ParseWithClaims(token, tk, func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenSecret), nil
	})
	return tk.Email
}

func GetName(token string) string {
	tokenSecret := os.Getenv("tokenSecret")
	tk := &models.Token{}
	jwt.ParseWithClaims(token, tk, func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenSecret), nil
	})
	return tk.Name
}

func GetRole(token string) string {
	tokenSecret := os.Getenv("tokenSecret")
	tk := &models.Token{}
	jwt.ParseWithClaims(token, tk, func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenSecret), nil
	})
	return tk.Role
}

func GetVerify(token string) string {
	tokenSecret := os.Getenv("tokenSecret")
	tk := &models.Token{}
	jwt.ParseWithClaims(token, tk, func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenSecret), nil
	})
	return tk.Verify
}
