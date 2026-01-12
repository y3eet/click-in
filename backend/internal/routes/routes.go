package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/y3eet/click-in/internal/config"
	"github.com/y3eet/click-in/internal/handlers"
	"github.com/y3eet/click-in/internal/middleware"
	"github.com/y3eet/click-in/internal/repositories"
	"github.com/y3eet/click-in/internal/services"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB, cfg *config.Config) {

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	clickableRepo := repositories.NewClickableRepository(db)
	clickableService := services.NewClickableService(clickableRepo)
	clickableHandler := handlers.NewClickableHandler(clickableService)

	clickRepo := repositories.NewClickRepository(db)
	clickService := services.NewClickService(clickRepo)
	clickHandler := handlers.NewClickHandler(clickService)

	refreshTokenRepo := repositories.NewRefreshTokenRepository(db)
	refreshTokenService := services.NewRefreshTokenService(refreshTokenRepo)

	authHandler := handlers.NewAuthHandler(userService, refreshTokenService, cfg)

	r.GET("/health", func(ctx *gin.Context) {
		DBSQL, err := db.DB()
		if err != nil || DBSQL.Ping() != nil {
			ctx.JSON(
				500,
				gin.H{"status": "Database connection error"},
			)
			return
		}
		ctx.JSON(
			200,
			gin.H{"status": "OK"},
		)
	})

	api := r.Group("/api")
	{
		users := api.Group("/users")
		users.Use(middleware.AuthMiddleware)
		{
			users.GET("/:id", userHandler.GetUser)
			users.GET("", userHandler.GetAllUsers)
		}

		clickable := api.Group("/clickable")
		clickable.Use(middleware.AuthMiddleware)
		{
			clickable.POST("", clickableHandler.CreateClickable)
			clickable.GET("", clickableHandler.GetAllClickable)
			clickable.GET("/:id", clickableHandler.GetClickableById)
		}

		click := api.Group("/click")
		click.Use(middleware.AuthMiddleware)
		{
			click.POST("", clickHandler.CreateClick)
			clickEvent := click.Group("/event")
			{
				clickEvent.GET("/count/:clickable_id", clickHandler.StreamClickCountByClickableID)
			}

		}

		file := api.Group("/file")
		{
			file.POST("/upload", handlers.FileUpload)
			file.GET("/:key", handlers.ViewFile)
		}

	}
	auth := r.Group("/auth")
	{
		auth.GET("/:provider", authHandler.Login)
		auth.GET("/:provider/callback", authHandler.Callback)
		auth.POST("/:provider/callback", authHandler.Callback)
		auth.POST("/exchange", authHandler.Exchange)
		auth.POST("/refresh", authHandler.RefreshToken)
		auth.GET("/current-user", authHandler.CurrentUser)
		auth.POST("/logout", authHandler.Logout)
	}
	ws := r.Group("/ws")
	ws.Use(middleware.AuthMiddleware)
	{
		ws.GET("", handlers.WebSocketHandler)
	}

}
