package repositories

import (
	"aivisual-core/internal/domain"
	"aivisual-core/internal/infra/db"
	"database/sql"
	"log"
)

// EventRepository 事件数据仓库接口
type EventRepository struct {
	db *db.DB
}

// NewEventRepository 创建新的事件数据仓库实例
func NewEventRepository(database *db.DB) *EventRepository {
	return &EventRepository{db: database}
}

// Create 创建事件记录
func (r *EventRepository) Create(event *domain.Event) error {
	model := db.EventFromDomain(event)

	query := `
		INSERT INTO events (id, device_id, event_type, time_stamp, description, created_at)
		VALUES (?, ?, ?, ?, ?, ?)
	`

	_, err := r.db.Exec(query,
		model.ID, model.DeviceID, model.EventType, model.TimeStamp, model.Description, model.CreatedAt)

	if err != nil {
		log.Printf("创建事件记录失败: %v", err)
		return err
	}

	return nil
}

// GetAll 获取事件记录列表
func (r *EventRepository) GetAll() ([]*domain.Event, error) {
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
		var model db.EventModel
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

// GetByID 根据ID获取事件记录
func (r *EventRepository) GetByID(id string) (*domain.Event, error) {
	query := `
		SELECT id, device_id, event_type, time_stamp, description, created_at
		FROM events
		WHERE id = ?
	`

	var model db.EventModel
	err := r.db.QueryRow(query, id).Scan(
		&model.ID, &model.DeviceID, &model.EventType, &model.TimeStamp, &model.Description, &model.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return model.ToDomain(), nil
}