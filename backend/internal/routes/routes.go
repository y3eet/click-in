package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/y3eet/click-in/internal/config"
	"github.com/y3eet/click-in/internal/handlers"
	"github.com/y3eet/click-in/internal/repositories"
	"github.com/y3eet/click-in/internal/services"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB, cfg *config.Config) {
	// Initialize layers
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	authHandler := handlers.NewAuthHandler(userService, cfg)

	api := r.Group("/api")
	{
		users := api.Group("/users")
		{
			users.GET("/:id", userHandler.GetUser)
			users.GET("", userHandler.GetAllUsers)
		}

	}
	auth := r.Group("/auth")
	{
		auth.GET("/:provider", authHandler.Login)
		auth.GET("/:provider/callback", authHandler.Callback)
		auth.POST("/:provider/callback", authHandler.Callback)
		auth.POST("/logout", authHandler.Logout)
	}
}
