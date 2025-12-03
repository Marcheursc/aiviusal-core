package service

import (
	"aivisual-core/internal/domain"
	"encoding/json"
	"time"
)

// Converter 负责将不同格式的报警消息转换为GA/T 1400标准格式
type Converter struct{}

// NewConverter 创建一个新的转换器实例
func NewConverter() *Converter {
	return &Converter{}
}

// ConvertRabbitMQMessage 将RabbitMQ消息转换为GA/T 1400标准的报警对象
func (c *Converter) ConvertRabbitMQMessage(data []byte) (*domain.Alarm, error) {
	// 解析原始消息格式（这里假设是JSON格式）
	var rawMessage map[string]interface{}
	if err := json.Unmarshal(data, &rawMessage); err != nil {
		return nil, err
	}

	// 构造GA/T 1400标准的报警对象
	alarm := &domain.Alarm{
		ID:        getString(rawMessage, "id", generateID()),
		TimeStamp: getTime(rawMessage, "timestamp", time.Now()),
		DeviceID:  getString(rawMessage, "device_id", ""),
		AlarmType: getString(rawMessage, "alarm_type", "UNKNOWN"),
		AlarmLevel: getInt(rawMessage, "alarm_level", 1),
		Description: getString(rawMessage, "description", ""),
		ImageURL: getString(rawMessage, "image_url", ""),
		VideoURL: getString(rawMessage, "video_url", ""),
		Location: domain.Location{
			Latitude:  getFloat(rawMessage, "latitude", 0.0),
			Longitude: getFloat(rawMessage, "longitude", 0.0),
			Address:   getString(rawMessage, "address", ""),
		},
		Status: getString(rawMessage, "status", "ACTIVE"),
	}

	return alarm, nil
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

// 辅助函数：从map中获取整数值
func getInt(m map[string]interface{}, key string, defaultValue int) int {
	if val, ok := m[key]; ok {
		if num, ok := val.(float64); ok { // JSON中的数字默认是float64
			return int(num)
		}
	}
	return defaultValue
}

// 辅助函数：从map中获取浮点数值
func getFloat(m map[string]interface{}, key string, defaultValue float64) float64 {
	if val, ok := m[key]; ok {
		if num, ok := val.(float64); ok {
			return num
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