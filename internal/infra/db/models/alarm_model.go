package models

import (
	"aivisual-core/internal/domain"
	"database/sql"
	"time"
)

// AlarmModel 报警数据模型
type AlarmModel struct {
	ID          string         `db:"id"`
	TimeStamp   time.Time      `db:"time_stamp"`
	DeviceID    string         `db:"device_id"`
	AlarmType   string         `db:"alarm_type"`
	AlarmLevel  int            `db:"alarm_level"`
	Description string         `db:"description"`
	ImageURL    sql.NullString `db:"image_url"`
	VideoURL    sql.NullString `db:"video_url"`
	Latitude    float64        `db:"latitude"`
	Longitude   float64        `db:"longitude"`
	Address     sql.NullString `db:"address"`
	Status      string         `db:"status"`
	CreatedAt   time.Time      `db:"created_at"`
	UpdatedAt   time.Time      `db:"updated_at"`
}

// ToDomain 转换为领域模型
func (a *AlarmModel) ToDomain() *domain.Alarm {
	alarm := &domain.Alarm{
		ID:          a.ID,
		TimeStamp:   a.TimeStamp,
		DeviceID:    a.DeviceID,
		AlarmType:   a.AlarmType,
		AlarmLevel:  a.AlarmLevel,
		Description: a.Description,
		Location: domain.Location{
			Latitude:  a.Latitude,
			Longitude: a.Longitude,
		},
		Status: a.Status,
	}

	if a.ImageURL.Valid {
		alarm.ImageURL = a.ImageURL.String
	}

	if a.VideoURL.Valid {
		alarm.VideoURL = a.VideoURL.String
	}

	if a.Address.Valid {
		alarm.Location.Address = a.Address.String
	}

	return alarm
}

// FromDomain 从领域模型转换
func FromDomain(alarm *domain.Alarm) *AlarmModel {
	model := &AlarmModel{
		ID:          alarm.ID,
		TimeStamp:   alarm.TimeStamp,
		DeviceID:    alarm.DeviceID,
		AlarmType:   alarm.AlarmType,
		AlarmLevel:  alarm.AlarmLevel,
		Description: alarm.Description,
		Latitude:    alarm.Location.Latitude,
		Longitude:   alarm.Location.Longitude,
		Status:      alarm.Status,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if alarm.ImageURL != "" {
		model.ImageURL = sql.NullString{String: alarm.ImageURL, Valid: true}
	}

	if alarm.VideoURL != "" {
		model.VideoURL = sql.NullString{String: alarm.VideoURL, Valid: true}
	}

	if alarm.Location.Address != "" {
		model.Address = sql.NullString{String: alarm.Location.Address, Valid: true}
	}

	return model
}