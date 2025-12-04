package handlers

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
	"github.com/mssola/user_agent"

	"github.com/y3eet/click-in/internal/auth"
	"github.com/y3eet/click-in/internal/config"
	"github.com/y3eet/click-in/internal/constants"
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
	userDB := models.User{
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
	userAgents := c.Request.UserAgent()
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
	ua := user_agent.New(userAgents)

	browserName, browserVersion := ua.Browser()
	platform := ua.Platform()
	os := ua.OS()
	is_mobile := ua.Mobile()
	ip := c.ClientIP()
	refreshToken, err := a.jwt.EncodeRefreshToken(*user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to sign refresh token: " + err.Error()})
		return
	}

	// Create refresh token in DB (will auto-delete existing token for same user agent)
	err = a.refreshTokenService.CreateRefreshToken(&models.RefreshToken{
		Token:          refreshToken,
		UserID:         user.ID,
		Browser:        browserName,
		BrowserVersion: browserVersion,
		Platform:       platform,
		OS:             os,
		IsMobile:       is_mobile,
		UserAgent:      userAgents,
		IPAddress:      ip,
		ExpiresAt:      time.Now().Add(7 * 24 * time.Hour),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create refresh token: " + err.Error()})
		return
	}

	c.SetCookie("access_token", accessToken, int(constants.AccessTokenTTL.Seconds()), "/", "", a.cfg.IsProd, true)
	c.SetCookie("refresh_token", refreshToken, int(constants.RefreshTokenTTL.Seconds()), "/", "", a.cfg.IsProd, true)

	c.JSON(http.StatusOK, gin.H{
		"message":      "exchange successful",
		"access_token": accessToken,
		"user":         user,
	})
}

func (a *AuthHandler) CurrentUser(c *gin.Context) {
	accessToken, err := c.Cookie("access_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}
	payload, err := a.jwt.DecodeAccessToken(accessToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Error decoding access token: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, payload)
}

func (a *AuthHandler) Logout(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}
	err = a.refreshTokenService.DeleteRefreshTokenByToken(refreshToken)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}
	c.SetCookie("access_token", "", -1, "/", "", a.cfg.IsProd, true)
	c.SetCookie("refresh_token", "", -1, "/", "", a.cfg.IsProd, true)
	c.JSON(http.StatusOK, gin.H{"message": "logged out"})
}

func (a *AuthHandler) RefreshToken(c *gin.Context) {

	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthenticated, no refresh token found"})
		return
	}

	oldToken, err := a.refreshTokenService.GetRefreshTokenByToken(refreshToken)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Unauthenticated, no refresh token found in db"})
		return
	}

	if oldToken.ExpiresAt.Before(time.Now()) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Refresh token expired"})
		_ = a.refreshTokenService.DeleteRefreshTokenByToken(oldToken.Token)
		return
	}

	decodedToken, err := a.jwt.DecodeRefreshToken(oldToken.Token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Refresh token expired"})
		return
	}

	user, err := a.userService.GetUserByID(decodedToken.User.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	newAccessToken, err := a.jwt.EncodeAccessToken(*user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Error encoding access token"})
		return
	}
	newRefreshToken, err := a.jwt.EncodeRefreshToken(*user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Error encoding refresh token"})
		return
	}

	err = a.refreshTokenService.RefreshToken(oldToken, newRefreshToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update refresh token: " + err.Error()})
		return
	}

	c.SetCookie("access_token", newAccessToken, int(constants.AccessTokenTTL.Seconds()), "/", "", a.cfg.IsProd, true)
	c.SetCookie("refresh_token", newRefreshToken, int(constants.RefreshTokenTTL.Seconds()), "/", "", a.cfg.IsProd, true)
	c.JSON(http.StatusOK, gin.H{"message": "Successfully refreshed token", "user": user})
}

func setProviderContext(c *gin.Context, provider string) {
	ctx := context.WithValue(c.Request.Context(), gothic.ProviderParamKey, provider)
	query := c.Request.URL.Query()
	query.Set("provider", provider)
	c.Request.URL.RawQuery = query.Encode()
	c.Request = c.Request.WithContext(ctx)
}
