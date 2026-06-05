package routes

import (
	"gin-users-api/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine,
	userHandler *handlers.UserHandler) {

	api := r.Group("/api")

	{
		api.GET("/users", userHandler.GetUsers)
		api.GET("/users/:id", userHandler.GetUser)

		api.POST("/users", userHandler.CreateUser)

		api.PUT("/users/:id", userHandler.UpdateUser)

		api.DELETE("/users/:id", userHandler.DeleteUser)
	}
}
