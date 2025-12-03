package repositories

import (
	"aivisual-core/internal/domain"
	"aivisual-core/internal/infra/db"
	"database/sql"
	"log"
)

// AlarmRepository 报警数据仓库接口
type AlarmRepository struct {
	db *db.DB
}

// NewAlarmRepository 创建新的报警数据仓库实例
func NewAlarmRepository(database *db.DB) *AlarmRepository {
	return &AlarmRepository{db: database}
}

// Create 创建报警记录
func (r *AlarmRepository) Create(alarm *domain.Alarm) error {
	model := db.FromDomain(alarm)

	query := `
		INSERT INTO alarms (
			id, time_stamp, device_id, alarm_type, alarm_level, description, 
			image_url, video_url, latitude, longitude, address, status, created_at, updated_at
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	_, err := r.db.Exec(query,
		model.ID, model.TimeStamp, model.DeviceID, model.AlarmType, model.AlarmLevel,
		model.Description, model.ImageURL, model.VideoURL, model.Latitude, model.Longitude,
		model.Address, model.Status, model.CreatedAt, model.UpdatedAt)

	if err != nil {
		log.Printf("创建报警记录失败: %v", err)
		return err
	}

	return nil
}

// GetAll 获取报警记录列表
func (r *AlarmRepository) GetAll() ([]*domain.Alarm, error) {
	query := `
		SELECT id, time_stamp, device_id, alarm_type, alarm_level, description,
		       image_url, video_url, latitude, longitude, address, status, created_at, updated_at
		FROM alarms
		ORDER BY time_stamp DESC
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var alarms []*domain.Alarm
	for rows.Next() {
		var model db.AlarmModel
		err := rows.Scan(
			&model.ID, &model.TimeStamp, &model.DeviceID, &model.AlarmType, &model.AlarmLevel,
			&model.Description, &model.ImageURL, &model.VideoURL, &model.Latitude, &model.Longitude,
			&model.Address, &model.Status, &model.CreatedAt, &model.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		alarms = append(alarms, model.ToDomain())
	}

	return alarms, nil
}

// GetByID 根据ID获取报警记录
func (r *AlarmRepository) GetByID(id string) (*domain.Alarm, error) {
	query := `
		SELECT id, time_stamp, device_id, alarm_type, alarm_level, description,
		       image_url, video_url, latitude, longitude, address, status, created_at, updated_at
		FROM alarms
		WHERE id = ?
	`

	var model db.AlarmModel
	err := r.db.QueryRow(query, id).Scan(
		&model.ID, &model.TimeStamp, &model.DeviceID, &model.AlarmType, &model.AlarmLevel,
		&model.Description, &model.ImageURL, &model.VideoURL, &model.Latitude, &model.Longitude,
		&model.Address, &model.Status, &model.CreatedAt, &model.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return model.ToDomain(), nil
}