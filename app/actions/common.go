package actions

import (
	"fmt"
	"net/http"

	"github.com/arthurc0102/gin-auth/app/utils"

	"github.com/gin-gonic/gin"
)

// Root of this project
func Root(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "It works!",
	})
}

// Hello actions
func Hello(c *gin.Context) {
	payload := c.MustGet("payload").(*utils.JWTClaims)

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Hello %s, welcome back", payload.Username),
	})
}
