package repositories

import (
	"github.com/arthurc0102/gin-auth/app/models"
	"github.com/arthurc0102/gin-auth/db"
)

// GetUserByUsername return user
func GetUserByUsername(username string) (user models.User, exists bool) {
	db.Connection.Where("username = ?", username).First(&user)
	exists = !db.Connection.NewRecord(user)
	return
}
