package routes

import (
	"aivisual-core/internal/infra/db/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

// EventHandler 处理事件相关的请求
type EventHandler struct {
	eventRepo *repositories.EventRepository
}

// NewEventHandler 创建新的事件处理器
func NewEventHandler(eventRepo *repositories.EventRepository) *EventHandler {
	return &EventHandler{
		eventRepo: eventRepo,
	}
}

// GetEvents 获取事件列表
func (h *EventHandler) GetEvents(c *gin.Context) {
	events, err := h.eventRepo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "获取事件列表失败",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    events,
	})
}