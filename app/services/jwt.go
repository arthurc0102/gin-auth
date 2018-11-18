package services

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/arthurc0102/gin-auth/app/models"
	"github.com/arthurc0102/gin-auth/app/utils"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/copier"
)

var secretKey = os.Getenv("SECRET_KEY")

// CreateJWT token
func CreateJWT(user models.User) (string, error) {
	claims := utils.JWTClaims{}
	claims.StandardClaims = jwt.StandardClaims{
		ExpiresAt: time.Now().Add(30 * time.Minute).Unix(),
		IssuedAt:  time.Now().Unix(),
	}

	copier.Copy(&claims, &user)
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := jwtToken.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return token, nil
}

// ParseJWT parse token
func ParseJWT(tokenString string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &utils.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("Token is not valid")
	}

	return token, nil
}

// RefreshJWT refresh token
func RefreshJWT(tokenString string) (string, error) {
	token, err := ParseJWT(tokenString)
	if err != nil {
		return "", err
	}

	user := models.User{}
	claims := token.Claims.(*utils.JWTClaims)
	copier.Copy(&user, claims)

	newToken, err := CreateJWT(user)
	if err != nil {
		return "", err
	}

	return newToken, nil
}
