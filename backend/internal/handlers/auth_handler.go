package handlers

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
	"github.com/y3eet/click-in/internal/config"
	"github.com/y3eet/click-in/internal/models"
	"github.com/y3eet/click-in/internal/services"
)

type AuthHandler struct {
	userService *services.UserService
	cfg         *config.Config
}

func NewAuthHandler(userService *services.UserService, cfg *config.Config) *AuthHandler {
	return &AuthHandler{userService: userService, cfg: cfg}
}
func (a *AuthHandler) Login(c *gin.Context) {
	provider := strings.TrimSpace(c.Param("provider"))
	if provider == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "provider is required"})
		return
	}

	setProviderContext(c, provider)

	if gothUser, err := gothic.CompleteUserAuth(c.Writer, c.Request); err == nil {
		c.JSON(200, gin.H{
			"status": "already_authenticated",
			"user":   gothUser,
		})
		return
	}
	// Otherwise start OAuth login
	gothic.BeginAuthHandler(c.Writer, c.Request)
}

func (a *AuthHandler) Callback(c *gin.Context) {
	provider := strings.TrimSpace(c.Param("provider"))
	if provider == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "provider is required"})
		return
	}

	setProviderContext(c, provider)

	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//Upsert user to DB here

	err = a.userService.UpsertUser(&models.User{
		Email:      user.Email,
		Username:   user.Name,
		AvatarURL:  user.AvatarURL,
		Provider:   user.Provider,
		ProviderID: user.UserID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to upsert user"})
		return
	}
}

func (a *AuthHandler) Logout(c *gin.Context) {
	gothic.Logout(c.Writer, c.Request)
	c.JSON(200, gin.H{"message": "logged out"})
}

func setProviderContext(c *gin.Context, provider string) {
	ctx := context.WithValue(c.Request.Context(), gothic.ProviderParamKey, provider)
	query := c.Request.URL.Query()
	query.Set("provider", provider)
	c.Request.URL.RawQuery = query.Encode()
	c.Request = c.Request.WithContext(ctx)
}
