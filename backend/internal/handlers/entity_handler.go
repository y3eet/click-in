package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/y3eet/click-in/internal/auth"
	"github.com/y3eet/click-in/internal/models"
	"github.com/y3eet/click-in/internal/services"
)

type EntityHandler struct {
	service *services.EntityService
}

type CreateEntityRequest struct {
	Name     string `json:"name" binding:"required"`
	ImageKey string `json:"image_key" binding:"required"`
	Mp3Key   string `json:"mp3_key"`
}

func NewEntityHandler(service *services.EntityService) *EntityHandler {
	return &EntityHandler{service: service}
}

func (h EntityHandler) CreateEntity(c *gin.Context) {
	claims, ok := auth.GetClaims(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	var req CreateEntityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	entity := &models.Entity{
		Name:     req.Name,
		ImageKey: req.ImageKey,
		Mp3Key:   req.Mp3Key,
		UserID:   claims.User.ID,
	}
	if err := h.service.CreateNewEntity(entity); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create entity"})
		return
	}
	c.JSON(http.StatusOK, entity)
}

func (h EntityHandler) GetAllEntity(c *gin.Context) {
	entities, err := h.service.GetAllEntity()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve entities"})
		return
	}
	c.JSON(http.StatusOK, entities)
}
