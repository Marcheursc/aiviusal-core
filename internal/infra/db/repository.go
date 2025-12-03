package db

// Repository 数据仓库组合
type Repository struct {
	Alarms *AlarmRepository
	Events *EventRepository
}

// NewRepository 创建新的数据仓库实例
func NewRepository(db *DB) *Repository {
	return &Repository{
		Alarms: NewAlarmRepository(db),
		Events: NewEventRepository(db),
	}
}

// CreateAlarm 创建报警记录
func (r *Repository) CreateAlarm(alarm *domain.Alarm) error {
	model := FromDomain(alarm)

	query := `
		INSERT INTO alarms (
			id, time_stamp, device_id, alarm_type, alarm_level, description, 
			image_url, video_url, latitude, longitude, address, status, created_at, updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
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

// GetAlarms 获取报警记录列表
func (r *Repository) GetAlarms() ([]*domain.Alarm, error) {
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
		var model AlarmModel
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

// GetAlarmByID 根据ID获取报警记录
func (r *Repository) GetAlarmByID(id string) (*domain.Alarm, error) {
	query := `
		SELECT id, time_stamp, device_id, alarm_type, alarm_level, description,
		       image_url, video_url, latitude, longitude, address, status, created_at, updated_at
		FROM alarms
		WHERE id = $1
	`

	var model AlarmModel
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

// CreateEvent 创建事件记录
func (r *Repository) CreateEvent(event *domain.Event) error {
	model := EventFromDomain(event)

	query := `
		INSERT INTO events (id, device_id, event_type, time_stamp, description, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := r.db.Exec(query,
		model.ID, model.DeviceID, model.EventType, model.TimeStamp, model.Description, model.CreatedAt)

	if err != nil {
		log.Printf("创建事件记录失败: %v", err)
		return err
	}

	return nil
}

// GetEvents 获取事件记录列表
func (r *Repository) GetEvents() ([]*domain.Event, error) {
	query := `
		SELECT id, device_id, event_type, time_stamp, description, created_at
		FROM events
		ORDER BY time_stamp DESC
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []*domain.Event
	for rows.Next() {
		var model EventModel
		err := rows.Scan(
			&model.ID, &model.DeviceID, &model.EventType, &model.TimeStamp, &model.Description, &model.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		events = append(events, model.ToDomain())
	}

	return events, nil
}