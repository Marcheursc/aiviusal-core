package routes

import (
	"aivisual-core/internal/domain"
	"aivisual-core/internal/infra/db/repositories"
	"aivisual-core/internal/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// WVPHandler 处理WVP-Pro设备事件的webhook
type WVPHandler struct {
	converter *service.Converter
	alarmRepo *repositories.AlarmRepository
	eventRepo *repositories.EventRepository
}

// NewWVPHandler 创建新的WVP处理器
func NewWVPHandler(converter *service.Converter, alarmRepo *repositories.AlarmRepository, eventRepo *repositories.EventRepository) *WVPHandler {
	return &WVPHandler{
		converter: converter,
		alarmRepo: alarmRepo,
		eventRepo: eventRepo,
	}
}

// HandleWVPEvent 处理WVP-Pro设备事件
func (h *WVPHandler) HandleWVPEvent(c *gin.Context) {
	var eventData map[string]interface{}
	if err := c.ShouldBindJSON(&eventData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	// 解析事件类型
	eventType := ""
	if val, ok := eventData["eventType"]; ok {
		eventType = val.(string)
	}

	// 根据事件类型处理
	switch eventType {
	case "ONLINE":
		h.handleOnlineEvent(c, eventData)
	case "OFFLINE":
		h.handleOfflineEvent(c, eventData)
	case "HEARTBEAT":
		h.handleHeartbeatEvent(c, eventData)
	default:
		// 尝试作为报警事件处理
		h.handleAlarmEvent(c, eventData)
	}
}

// handleOnlineEvent 处理设备上线事件
func (h *WVPHandler) handleOnlineEvent(c *gin.Context, eventData map[string]interface{}) {
	event := &domain.Event{
		ID:          getString(eventData, "eventId", generateID()),
		DeviceID:    getString(eventData, "deviceId", ""),
		EventType:   "ONLINE",
		TimeStamp:   getTime(eventData, "timestamp", time.Now()),
		Description: "Device came online",
	}

	// 存储事件到数据库
	if err := h.eventRepo.Create(event); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to store event",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Online event processed",
		"data":    event,
	})
}

// handleOfflineEvent 处理设备下线事件
func (h *WVPHandler) handleOfflineEvent(c *gin.Context, eventData map[string]interface{}) {
	event := &domain.Event{
		ID:          getString(eventData, "eventId", generateID()),
		DeviceID:    getString(eventData, "deviceId", ""),
		EventType:   "OFFLINE",
		TimeStamp:   getTime(eventData, "timestamp", time.Now()),
		Description: "Device went offline",
	}

	// 存储事件到数据库
	if err := h.eventRepo.Create(event); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to store event",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Offline event processed",
		"data":    event,
	})
}

// handleHeartbeatEvent 处理心跳事件
func (h *WVPHandler) handleHeartbeatEvent(c *gin.Context, eventData map[string]interface{}) {
	event := &domain.Event{
		ID:          getString(eventData, "eventId", generateID()),
		DeviceID:    getString(eventData, "deviceId", ""),
		EventType:   "HEARTBEAT",
		TimeStamp:   getTime(eventData, "timestamp", time.Now()),
		Description: "Device heartbeat received",
	}

	// 存储事件到数据库
	if err := h.eventRepo.Create(event); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to store event",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Heartbeat event processed",
		"data":    event,
	})
}

// handleAlarmEvent 处理报警事件
func (h *WVPHandler) handleAlarmEvent(c *gin.Context, eventData map[string]interface{}) {
	// 这里只是简化处理，实际应该解析完整的报警数据
	event := &domain.Event{
		ID:          getString(eventData, "eventId", generateID()),
		DeviceID:    getString(eventData, "deviceId", ""),
		EventType:   "ALARM",
		TimeStamp:   getTime(eventData, "timestamp", time.Now()),
		Description: "Device alarm received",
	}

	// 存储事件到数据库
	if err := h.eventRepo.Create(event); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to store event",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Alarm event processed",
		"data":    event,
	})
}

// 辅助函数：从map中获取字符串值
func getString(m map[string]interface{}, key string, defaultValue string) string {
	if val, ok := m[key]; ok {
		if str, ok := val.(string); ok {
			return str
		}
	}
	return defaultValue
}

// 辅助函数：从map中获取时间值
func getTime(m map[string]interface{}, key string, defaultValue time.Time) time.Time {
	if val, ok := m[key]; ok {
		if str, ok := val.(string); ok {
			if t, err := time.Parse(time.RFC3339, str); err == nil {
				return t
			}
		}
	}
	return defaultValue
}

// 辅助函数：生成唯一ID
func generateID() string {
	// 简化实现，实际应该使用UUID或其他唯一ID生成方法
	return time.Now().Format("20060102150405")
}