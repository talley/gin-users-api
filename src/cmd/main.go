package main

import (
	"gin-users-api/config"
	"gin-users-api/handlers"
	"gin-users-api/repositories"
	"gin-users-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	config.ConnectDB()

	repo := repositories.NewUserRepository(config.DB)

	handler := handlers.NewUserHandler(repo)

	router := gin.Default()

	routes.RegisterRoutes(router, handler)

	router.Run(":8085")
}
