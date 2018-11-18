package services

import (
	"errors"

	"github.com/arthurc0102/gin-auth/app/models"
	"github.com/arthurc0102/gin-auth/app/repositories"
	"github.com/arthurc0102/gin-auth/app/serializers"
	"github.com/arthurc0102/gin-auth/app/utils"
)

// Login services
func Login(serializer serializers.Login) (models.User, error) {
	user, exists := repositories.GetUserByUsername(serializer.Username)

	if !exists || !utils.CheckPassword(serializer.Password, user.Password) {
		return user, errors.New("Login fail")
	}

	return user, nil
}
