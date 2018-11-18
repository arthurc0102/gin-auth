package main

import (
	"fmt"
	"os"

	"github.com/arthurc0102/gin-auth/config"
	"github.com/arthurc0102/gin-auth/db"
	"github.com/arthurc0102/gin-auth/db/migrate"
	"github.com/gin-gonic/gin"
)

func init() {
	commandList := []string{
		"run",
		"migrate",
	}

	if len(os.Args) < 2 {
		fmt.Println("Choice a command:")

		for _, cmd := range commandList {
			fmt.Printf("    %s\n", cmd)
		}

		os.Exit(0)
	}
}

func main() {
	config.Load()

	db.Connect()
	defer db.Close()

	switch os.Args[1] {
	case "migrate":
		migrate.Do()
		fmt.Println("Migrate done")
	case "run":
		server := gin.Default()

		config.RegisterRoutes(server)
		config.RegisterValidators()

		server.Run()
	}
}
