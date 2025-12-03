package domain

import "time"

// Alarm 报警信息
type Alarm struct {
	ID            string    `json:"id"`
	TimeStamp     time.Time `json:"time_stamp"`
	DeviceID      string    `json:"device_id"`
	AlarmType     string    `json:"alarm_type"` // LOITERING, GATHERING, LEAVE, BANNER
	AlarmLevel    int       `json:"alarm_level"`
	Description   string    `json:"description"`
	ImageURL      string    `json:"image_url,omitempty"`
	VideoURL      string    `json:"video_url,omitempty"`
	Location      Location  `json:"location"`
	Status        string    `json:"status"` // ACTIVE, PROCESSED, CLOSED
}

// Location 位置信息
type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Address   string  `json:"address,omitempty"`
}