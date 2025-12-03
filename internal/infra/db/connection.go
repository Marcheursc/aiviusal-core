package db

import (
	"aivisual-core/internal/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// DB 数据库连接实例
type DB struct {
	*sql.DB
}

// NewDB 创建新的数据库连接
func NewDB(cfg *config.DatabaseConfig) (*DB, error) {
	// 构建 MySQL 连接字符串
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)

	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return nil, err
	}

	// 测试连接
	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("数据库连接成功")
	return &DB{db}, nil
}

// RunMigrations 运行数据库迁移
func (db *DB) RunMigrations() error {
	statements := CreateTableStatements()
	for _, stmt := range statements {
		_, err := db.Exec(stmt)
		if err != nil {
			return fmt.Errorf("执行数据库迁移失败: %v", err)
		}
	}

	log.Println("数据库迁移完成")
	return nil
}