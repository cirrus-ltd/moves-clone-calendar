package database

import (
	"fmt"
	"log"

	"github.com/Cirrus-Ltd/moves-clone-calendar/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// gormの初期化のみを行う
// gormではマイグレーションはやらないgolang-migrateを使用してスキーマのバージョン管理とマイグレーションを行う
// GORMを使用して日常的なデータベース操作（CRUD操作）を行う
func NewDB(cfg config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.Database.Host, cfg.Database.User, cfg.Database.Password, cfg.Database.DBName, cfg.Database.Port, cfg.Database.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	}, Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		return nil, err
	}
	log.Println("Database connection established")
	return db, nil
}
