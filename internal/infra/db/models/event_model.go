package models

import (
	"aivisual-core/internal/domain"
	"time"
)

// EventModel 事件数据模型
type EventModel struct {
	ID          string    `db:"id"`
	DeviceID    string    `db:"device_id"`
	EventType   string    `db:"event_type"`
	TimeStamp   time.Time `db:"time_stamp"`
	Description string    `db:"description"`
	CreatedAt   time.Time `db:"created_at"`
}

// ToDomain 转换为领域模型
func (e *EventModel) ToDomain() *domain.Event {
	return &domain.Event{
		ID:          e.ID,
		DeviceID:    e.DeviceID,
		EventType:   e.EventType,
		TimeStamp:   e.TimeStamp,
		Description: e.Description,
	}
}

// FromDomain 从领域模型转换
func FromDomain(event *domain.Event) *EventModel {
	return &EventModel{
		ID:          event.ID,
		DeviceID:    event.DeviceID,
		EventType:   event.EventType,
		TimeStamp:   event.TimeStamp,
		Description: event.Description,
		CreatedAt:   time.Now(),
	}
}