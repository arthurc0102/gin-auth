package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// NotFound response http 404
func NotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"message": "Can't found resources",
	})
}
