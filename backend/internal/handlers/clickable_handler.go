package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/y3eet/click-in/internal/auth"
	"github.com/y3eet/click-in/internal/models"
	"github.com/y3eet/click-in/internal/services"
)

type ClickableHandler struct {
	service *services.ClickableService
}

type CreateClickableRequest struct {
	Name     string `json:"name" binding:"required"`
	ImageKey string `json:"image_key" binding:"required"`
	Mp3Key   string `json:"mp3_key"`
}

func NewClickableHandler(service *services.ClickableService) *ClickableHandler {
	return &ClickableHandler{service: service}
}

func (h ClickableHandler) CreateClickable(c *gin.Context) {
	claims, ok := auth.GetClaims(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	var req CreateClickableRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	clickable := &models.Clickable{
		Name:     req.Name,
		ImageKey: req.ImageKey,
		Mp3Key:   req.Mp3Key,
		UserID:   claims.User.ID,
	}
	if err := h.service.CreateNewClickable(clickable); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create clickable"})
		return
	}
	c.JSON(http.StatusOK, clickable)
}

func (h ClickableHandler) GetAllClickable(c *gin.Context) {
	entities, err := h.service.GetAllClickable()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve entities"})
		return
	}
	c.JSON(http.StatusOK, entities)
}
