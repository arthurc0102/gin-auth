package migrate

import (
	"github.com/arthurc0102/gin-auth/app/models"
	"github.com/arthurc0102/gin-auth/db"
)

// Do migrate
func Do() {
	db.Connection.AutoMigrate(
		&models.User{},
	)
}
