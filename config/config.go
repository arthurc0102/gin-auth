package config

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// Load config from .env file
func Load() {
	_ = godotenv.Load()

	if os.Getenv("GIN_MODE") != "debug" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Check env has been set
	checkEnv("SECRET_KEY")
}
