package actions

import (
	"net/http"

	"github.com/arthurc0102/gin-auth/app/models"
	"github.com/arthurc0102/gin-auth/app/serializers"
	"github.com/arthurc0102/gin-auth/app/services"
	"github.com/arthurc0102/gin-auth/app/utils"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

// Login actions
func Login(c *gin.Context) {
	var serializer serializers.Login

	if err := c.ShouldBind(&serializer); err != nil {
		c.JSON(http.StatusBadRequest, utils.HandleError(serializer, err))
		return
	}

	user, err := services.Login(serializer)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	token, err := services.CreateJWT(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// Register actions
func Register(c *gin.Context) {
	var serializer serializers.Register

	if err := c.ShouldBind(&serializer); err != nil {
		c.JSON(http.StatusBadRequest, utils.HandleError(serializer, err))
		return
	}

	user := models.User{}
	copier.Copy(&user, &serializer)

	if err := user.Save(); err != nil {
		c.JSON(http.StatusBadRequest, utils.HandleError(user, err))
		return
	}

	token, err := services.CreateJWT(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// Refresh action
func Refresh(c *gin.Context) {
	var serializer serializers.RefreshJWT

	if err := c.ShouldBind(&serializer); err != nil {
		c.JSON(http.StatusBadRequest, utils.HandleError(serializer, err))
		return
	}

	token, err := services.RefreshJWT(serializer.Token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
