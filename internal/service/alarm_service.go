package service

import (
	"aivisual-core/internal/domain"
)

// AlarmService 报警服务
type AlarmService struct {
	converter *Converter
}

// NewAlarmService 创建新的报警服务实例
func NewAlarmService(converter *Converter) *AlarmService {
	return &AlarmService{
		converter: converter,
	}
}

// ProcessAlarmMessage 处理报警消息
func (s *AlarmService) ProcessAlarmMessage(data []byte) (*domain.Alarm, error) {
	return s.converter.ConvertRabbitMQMessage(data)
}