package models

import jwt "github.com/golang-jwt/jwt/v4"

//Token struct declaration
type Token struct {
	UserID uint
	Email  string
	Name   string
	Role   string
	Verify string
	*jwt.StandardClaims
}
