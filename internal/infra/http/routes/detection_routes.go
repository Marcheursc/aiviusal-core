package routes

import (
	"aivisual-core/internal/domain"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// DetectionHandler 处理检测相关的请求
type DetectionHandler struct {
	// 可以添加检测服务依赖
}

// NewDetectionHandler 创建新的检测处理器
func NewDetectionHandler() *DetectionHandler {
	return &DetectionHandler{}
}

// GetDetections 获取目标检测结果
func (h *DetectionHandler) GetDetections(c *gin.Context) {
	// 模拟数据，实际应该从数据库查询
	detections := []domain.ObjectDetection{
		{
			ID:         "det001",
			TimeStamp:  time.Now().Add(-10 * time.Minute),
			DeviceID:   "camera_001",
			ObjectType: "PERSON",
			Confidence: 0.95,
			BoundingBox: domain.Box{
				X:      100,
				Y:      150,
				Width:  80,
				Height: 120,
			},
			TrackingID:  "track_001",
			SnapshotURL: "/snapshots/det001.jpg",
		},
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    detections,
	})
}