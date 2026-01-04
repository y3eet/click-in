package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/y3eet/click-in/internal/auth"
	"github.com/y3eet/click-in/internal/models"
	"github.com/y3eet/click-in/internal/services"
)

type ClickHandler struct {
	service *services.ClickService
}

func NewClickHandler(service *services.ClickService) *ClickHandler {
	return &ClickHandler{service: service}
}

func (h *ClickHandler) CreateClick(c *gin.Context) {
	claims, _ := auth.GetClaims(c)

	var clickInput struct {
		ClickableID uint `json:"clickable_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&clickInput); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	click := &models.Click{
		ClickableID: clickInput.ClickableID,
		UserID:      claims.User.ID,
	}

	h.service.CreateNewClick(click)
	c.JSON(201, gin.H{"message": "click recorded"})
}

func (h *ClickHandler) CountClicksByClickableID(clickableID uint) (int64, error) {
	return h.service.CountClicksByClickableID(clickableID)
}

func (h *ClickHandler) CountClicksByUserID(userID uint) (int64, error) {
	return h.service.CountClicksByUserID(userID)
}

func (h *ClickHandler) CountClicksByClickableAndUser(clickableID uint, userID uint) (int64, error) {
	return h.service.CountClicksByClickableAndUser(clickableID, userID)
}
