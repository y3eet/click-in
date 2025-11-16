package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/y3eet/click-in/internal/handlers"
	"github.com/y3eet/click-in/internal/repositories"
	"github.com/y3eet/click-in/internal/services"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	// Initialize layers
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	v1 := r.Group("/api")
	{
		users := v1.Group("/users")
		{
			users.POST("", userHandler.CreateUser)
			users.GET("/:id", userHandler.GetUser)
			users.GET("", userHandler.GetAllUsers)
		}
	}
}
