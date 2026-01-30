package client

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/sword-fisher-fly/ai-alert/internal/global"
	"github.com/sword-fisher-fly/ai-alert/internal/models"

	"github.com/glebarez/sqlite"
	"github.com/zeromicro/go-zero/core/logc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBConfig struct {
	Type    string // database type: mysql or sqlite
	Host    string
	Port    string
	User    string
	Pass    string
	DBName  string
	Timeout string
	Path    string // SQLite database file path
}

func NewDBClient(config DBConfig) *gorm.DB {
	var db *gorm.DB
	var err error

	if config.Type == "" {
		config.Type = "mysql"
	}

	switch config.Type {
	case "sqlite":
		db, err = initSQLiteDB(config)
	case "mysql":
		db, err = initMySQLDB(config)
	default:
		logc.Errorf(context.Background(), "unsupported database type: %s", config.Type)
		return nil
	}

	if err != nil {
		logc.Errorf(context.Background(), "failed to connect database: %s", err.Error())
		return nil
	}

	// Create tables automatically
	err = db.AutoMigrate(
		&models.AiContentRecord{},
	)
	if err != nil {
		logc.Error(context.Background(), err.Error())
		return nil
	}

	if global.Config.Server.Mode == "debug" {
		db.Debug()
	} else {
		db.Logger = logger.Default.LogMode(logger.Silent)
	}

	return db
}

// initMySQLDB is used to initialize MySQL database connection.
func initMySQLDB(config DBConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4,utf8&parseTime=True&loc=Local&timeout=%s",
		config.User,
		config.Pass,
		config.Host,
		config.Port,
		config.DBName,
		config.Timeout)

	logc.Infof(context.Background(), "connecting to MySQL database: %s:%s/%s", config.Host, config.Port, config.DBName)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

// initSQLiteDB is used to initialize SQLite database connection.
func initSQLiteDB(config DBConfig) (*gorm.DB, error) {
	// Set default db file path for SQLite
	if config.Path == "" {
		config.Path = "data/ai.db"
	}

	// Make sure that the directory exists
	dir := filepath.Dir(config.Path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create directory %s: %w", dir, err)
	}

	logc.Infof(context.Background(), "connecting to SQLite database: %s", config.Path)
	return gorm.Open(sqlite.Open(config.Path), &gorm.Config{})
}
