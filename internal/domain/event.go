package domain

import "time"

// Event 设备事件
type Event struct {
	ID          string    `json:"id"`
	DeviceID    string    `json:"device_id"`
	EventType   string    `json:"event_type"` // ONLINE, OFFLINE, HEARTBEAT
	TimeStamp   time.Time `json:"time_stamp"`
	Description string    `json:"description"`
}