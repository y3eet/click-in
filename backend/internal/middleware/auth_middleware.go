package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/y3eet/click-in/internal/auth"
	"github.com/y3eet/click-in/internal/config"
)

func AuthMiddleware(c *gin.Context) {
	accessToken, err := c.Cookie("access_token")
	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
		return
	}
	cfg := config.Cfg
	jwt := auth.NewJWT(cfg)
	claims, err := jwt.DecodeAccessToken(accessToken)
	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
		return
	}
	c.Set("claims", claims)
	c.Next()
}
