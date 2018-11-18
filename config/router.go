package config

import (
	"github.com/arthurc0102/gin-auth/app/actions"
	"github.com/arthurc0102/gin-auth/app/middlewares"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes register route of this project to server
func RegisterRoutes(server *gin.Engine) {
	server.GET("", actions.Root)

	auth := server.Group("auth")
	{
		auth.POST("login", actions.Login)
		auth.POST("register", actions.Register)
		auth.POST("refresh", actions.Refresh)
	}

	authorized := server.Group("", middlewares.ParseJWT())
	{
		authorized.GET("hello", actions.Hello)
	}
}
