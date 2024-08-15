package migrations

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations() {
	// 環境変数の読み込み
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("POSTGRES_DB")
	dbSSLMode := os.Getenv("DB_SSLMODE")

	// 環境変数のデバッグ出力
	log.Printf("DB_USER: %s, DB_PASSWORD: %s, DB_HOST: %s, DB_PORT: %s, DB_NAME: %s, DB_SSLMODE: %s",
		dbUser, dbPassword, dbHost, dbPort, dbName, dbSSLMode)

	// DSNの作成
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		dbUser, dbPassword, dbHost, dbPort, dbName, dbSSLMode)

	// リトライロジックの追加
	var m *migrate.Migrate
	var err error
	for i := 0; i < 10; i++ {
		m, err = migrate.New(
			"file://migrations",
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
