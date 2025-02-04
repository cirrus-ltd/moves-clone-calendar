package migrations

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/Cirrus-Ltd/moves-clone-calendar/internal/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations(cfg *config.Config) {
	// 環境変数のデバッグ出力
	log.Printf("DB_USER: %s, DB_PASSWORD: %s, DB_HOST: %s, DB_PORT: %s, DB_NAME: %s, DB_SSLMODE: %s",
		cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.DBName, cfg.Database.SSLMode)

	// DSNの作成
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.DBName, cfg.Database.SSLMode)

	log.Printf("DSN: %s", dsn)
	// リトライロジックの追加
	var m *migrate.Migrate
	var err error
	for i := 0; i < 10; i++ {
		m, err = migrate.New(
			"file://internal/infrastructure/migrations",
			dsn)
		if err == nil {
			break
		}
		log.Printf("Failed to create migrate instance: %v. Retrying in 5 seconds...", err)
		time.Sleep(5 * time.Second)
	}
	if err != nil {
		log.Fatalf("Failed to create migrate instance after retries: %v", err)
	}

	// マイグレーションの実行
	if err := m.Up(); errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	log.Println("Migrations ran successfully")
}
