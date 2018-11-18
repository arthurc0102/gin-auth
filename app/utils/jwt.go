package utils

import (
	jwt "github.com/dgrijalva/jwt-go"
)

// JWTClaims jwt claims
type JWTClaims struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Gender    string `json:"gender"`
	Age       uint   `json:"age"`

	jwt.StandardClaims
}
