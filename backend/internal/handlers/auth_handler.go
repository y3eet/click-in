package handlers

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

type AuthHandler struct{}

func NewAuthHnadler() *AuthHandler {
	return &AuthHandler{}
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
	c.JSON(http.StatusOK, gin.H{
		"status": "authenticated",
		"user":   user,
	})
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
