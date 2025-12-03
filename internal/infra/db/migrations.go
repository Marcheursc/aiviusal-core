package db

// CreateTableStatements 返回创建表的SQL语句
func CreateTableStatements() []string {
	return []string{
		`
		CREATE TABLE IF NOT EXISTS alarms (
			id VARCHAR(255) PRIMARY KEY,
			time_stamp DATETIME NOT NULL,
			device_id VARCHAR(255) NOT NULL,
			alarm_type VARCHAR(100) NOT NULL,
			alarm_level INTEGER NOT NULL,
			description TEXT,
			image_url TEXT,
			video_url TEXT,
			latitude DOUBLE,
			longitude DOUBLE,
			address TEXT,
			status VARCHAR(50) NOT NULL,
			created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
		);
		`,
		`
		CREATE TABLE IF NOT EXISTS events (
			id VARCHAR(255) PRIMARY KEY,
			device_id VARCHAR(255) NOT NULL,
			event_type VARCHAR(100) NOT NULL,
			time_stamp DATETIME NOT NULL,
			description TEXT,
			created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
		);
		`,
		`
		CREATE INDEX idx_alarms_device_id ON alarms(device_id);
		`,
		`
		CREATE INDEX idx_alarms_time_stamp ON alarms(time_stamp);
		`,
		`
		CREATE INDEX idx_events_device_id ON events(device_id);
		`,
		`
		CREATE INDEX idx_events_time_stamp ON events(time_stamp);
		`,
	}
}