package routes

import (
	"daimyo/backend-api/controllers"
	"daimyo/backend-api/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.POST("/api/register", controllers.Register)
	r.POST("/api/login", controllers.Login)
	r.GET("/api/users", middlewares.AuthMiddleware(), controllers.FindUsers)
	r.POST("/api/users", middlewares.AuthMiddleware(), controllers.CreateUser)
	r.GET("/api/users/:id", middlewares.AuthMiddleware(), controllers.FindUserById)
	r.PUT("/api/users/:id", middlewares.AuthMiddleware(), controllers.UpdateUser)
	r.DELETE("/api/users/:id", middlewares.AuthMiddleware(), controllers.DeleteUser)

	return r
}
