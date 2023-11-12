package main

import (
	"github.com/PrrP17/crud-go.git/src/configuration/logger"
	"github.com/PrrP17/crud-go.git/src/contreoller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {

	logger.Info("stater app")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
