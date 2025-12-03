package routes

import (
	"aivisual-core/internal/infra/db/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AlarmHandler 处理报警相关的请求
type AlarmHandler struct {
	alarmRepo *repositories.AlarmRepository
}

// NewAlarmHandler 创建新的报警处理器
func NewAlarmHandler(alarmRepo *repositories.AlarmRepository) *AlarmHandler {
	return &AlarmHandler{
		alarmRepo: alarmRepo,
	}
}

// GetAlarms 获取报警列表
func (h *AlarmHandler) GetAlarms(c *gin.Context) {
	alarms, err := h.alarmRepo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "获取报警列表失败",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    alarms,
	})
}

// GetAlarmByID 根据ID获取特定报警
func (h *AlarmHandler) GetAlarmByID(c *gin.Context) {
	alarmID := c.Param("id")

	alarm, err := h.alarmRepo.GetByID(alarmID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "获取报警详情失败",
			"error":   err.Error(),
		})
		return
	}

	if alarm == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "未找到指定报警",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    alarm,
	})
}