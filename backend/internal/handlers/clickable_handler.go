package handlers

import (
	"net/http"
	"strconv"

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
	clickable, err := h.service.GetClickableByName(req.Name)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error finding clickable: " + err.Error()})
		return
	}
	if clickable.ID != 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "Clickable already exists"})
		return
	}
	clickable = &models.Clickable{
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
	clickable, err := h.service.GetAllClickable()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve clickable"})
		return
	}
	c.JSON(http.StatusOK, clickable)
}

func (h ClickableHandler) GetClickableById(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Clickable ID: " + err.Error()})
		return
	}
	clickable, err := h.service.GetClickableByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Clickable not found: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, clickable)
}
