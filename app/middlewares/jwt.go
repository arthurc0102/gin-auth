package middlewares

import (
	"net/http"
	"strings"

	"github.com/arthurc0102/gin-auth/app/services"

	"github.com/gin-gonic/gin"
)

// ParseJWT check token
func ParseJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := strings.Replace(c.GetHeader("Authorization"), "Bearer ", "", -1)
		token, err := services.ParseJWT(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.Set("payload", token.Claims)
	}
}
