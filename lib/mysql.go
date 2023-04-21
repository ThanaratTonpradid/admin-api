package lib

import (
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type MySQLOptions struct {
	DSN      string
	LogLevel logger.LogLevel
}

type MySQL struct {
	DB *gorm.DB
}

func NewMySQL(opts MySQLOptions) *MySQL {
	db, err := gorm.Open(mysql.Open(opts.DSN), &gorm.Config{
		Logger: NewMySQLLogger(opts.LogLevel),
	})
	if err != nil {
		log.Fatalf("Failed to connect mysql: %s", err)
	}
	return &MySQL{DB: db}
}

func NewMySQLLogger(level logger.LogLevel) logger.Interface {
	config := logger.Config{
		SlowThreshold:             time.Second,
		LogLevel:                  level,
		IgnoreRecordNotFoundError: true,
		Colorful:                  false,
	}
	return logger.New(NewLibLogger(), config)
}

func NewMySQLLogLevel(level string) logger.LogLevel {
	if strings.ToLower(level) == "debug" {
		return logger.Info
	}
	return logger.Error
}

func (m *MySQL) Close() {
	db, err := m.DB.DB()
	if err != nil {
		log.Errorf("MySQL error: %s", err)
		return
	}
	if err := db.Close(); err != nil {
		log.Errorf("MySQL error in closing the connection: %s", err)
	}
}

func (MySQL) Paginate(page int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
