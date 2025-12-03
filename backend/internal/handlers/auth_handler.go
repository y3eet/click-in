package handlers

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
	"github.com/y3eet/click-in/internal/auth"
	"github.com/y3eet/click-in/internal/config"
	"github.com/y3eet/click-in/internal/models"
	"github.com/y3eet/click-in/internal/services"
)

type AuthHandler struct {
	userService         *services.UserService
	refreshTokenService *services.RefreshTokenService
	cfg                 *config.Config
	jwt                 *auth.JWTManager
}

type ExchangeRequestBody struct {
	ExchangeToken string `json:"exchange_token"`
}

func NewAuthHandler(userService *services.UserService, refreshToksetService *services.RefreshTokenService, cfg *config.Config) *AuthHandler {
	return &AuthHandler{userService: userService, refreshTokenService: refreshToksetService, cfg: cfg, jwt: auth.NewJWT(cfg)}
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
	var userDB = models.User{
		Email:      user.Email,
		Username:   user.Name,
		AvatarURL:  user.AvatarURL,
		Provider:   user.Provider,
		ProviderID: user.UserID,
	}
	err = a.userService.UpsertUser(&userDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to upsert user"})
		return
	}
	exchangeToken, err := a.jwt.EncodeExchangeToken(userDB.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	c.Redirect(http.StatusTemporaryRedirect, a.cfg.FrontendURL+"/auth/callback?exchange_token="+exchangeToken)
}

func (a *AuthHandler) Exchange(c *gin.Context) {
	var exchangeReqBody ExchangeRequestBody
	if err := c.ShouldBindJSON(&exchangeReqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON: " + err.Error()})
		return
	}

	claims, err := a.jwt.DecodeExchangeToken(exchangeReqBody.ExchangeToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid access: " + err.Error()})
		return
	}

	user, err := a.userService.GetUserByID(claims.UserID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "error getting user: " + err.Error()})
		return
	}

	accessToken, err := a.jwt.EncodeAccessToken(*user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to sign access token: " + err.Error()})
		return
	}
	refreshToken, err := a.jwt.EncodeRefreshToken(*user)

	c.SetCookie("access_token", accessToken, int(time.Hour.Seconds()), "/", "", a.cfg.IsProd, true)
	c.SetCookie("refresh_token", refreshToken, int(time.Hour.Seconds()*7*24), "/", "", a.cfg.IsProd, true)

	c.JSON(http.StatusOK, gin.H{
		"message":      "exchange successful",
		"access_token": accessToken,
		"user":         user,
	})
}

func (a *AuthHandler) Logout(c *gin.Context) {
	gothic.Logout(c.Writer, c.Request)
	c.JSON(http.StatusOK, gin.H{"message": "logged out"})
}

func setProviderContext(c *gin.Context, provider string) {
	ctx := context.WithValue(c.Request.Context(), gothic.ProviderParamKey, provider)
	query := c.Request.URL.Query()
	query.Set("provider", provider)
	c.Request.URL.RawQuery = query.Encode()
	c.Request = c.Request.WithContext(ctx)
}
