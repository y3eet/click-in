package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/y3eet/click-in/internal/auth"
	"github.com/y3eet/click-in/internal/config"
	"github.com/y3eet/click-in/internal/database"
	"github.com/y3eet/click-in/internal/routes"
)

func main() {
	cfg := config.Load()
	auth.NewAuth(cfg)

	db, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	r := gin.Default()
	routes.SetupRoutes(r, db, cfg)

	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
