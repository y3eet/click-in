package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

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

func (h *ClickHandler) StreamClickCountByClickableID(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	c.Writer.Header().Set("Transfer-Encoding", "chunked")

	clickableID := c.Param("clickable_id")
	fmt.Printf("Starting click count stream for clickable ID: %s\n", clickableID)

	clickableIDUint64, err := strconv.ParseUint(clickableID, 10, 0)
	if err != nil {
		c.String(http.StatusBadRequest, "invalid clickable_id")
		return
	}
	clickableIDUint := uint(clickableIDUint64)

	// Make sure Gin flushes instead of buffering
	flusher, ok := c.Writer.(http.Flusher)
	if !ok {
		c.String(http.StatusInternalServerError, "Streaming unsupported")
		return
	}
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	ctx := c.Request.Context()

	for {
		select {
		case t := <-ticker.C:
			count, err := h.service.CountClicksByClickableID(clickableIDUint)
			if err != nil {
				return
			}
			payload := gin.H{
				"timestamp": t.Format(time.RFC3339),
				"count":     count,
			}

			b, err := json.Marshal(payload)
			if err != nil {
				return
			}

			fmt.Fprintf(c.Writer, "data: %s\n\n", b)
			flusher.Flush()

		case <-ctx.Done():
			return
		}
	}
}

// func (h *ClickHandler) CountClicksByClickableID(clickableID uint) (int64, error) {
// 	return h.service.CountClicksByClickableID(clickableID)
// }

// func (h *ClickHandler) CountClicksByUserID(userID uint) (int64, error) {
// 	return h.service.CountClicksByUserID(userID)
// }

// func (h *ClickHandler) CountClicksByClickableAndUser(clickableID uint, userID uint) (int64, error) {
// 	return h.service.CountClicksByClickableAndUser(clickableID, userID)
// }
