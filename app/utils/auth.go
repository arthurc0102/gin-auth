package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword hash password
func HashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		log.Fatalln(err)
	}

	return string(hash)
}

// CheckPassword check password match or not
func CheckPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
