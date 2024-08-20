package config

import (
	"log"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type Database struct {
	User     string `env:"POSTGRES_USER"`
	DBName   string `env:"POSTGRES_DB"`
	SSLMode  string `env:"DB_SSLMODE"`
	Password string `env:"POSTGRES_PASSWORD"`
	Host     string `env:"DB_HOST"`
	Port     string `env:"DB_PORT"`
}
type Server struct {
	Port string `env:"SERVER_PORT"`
}

type Config struct {
	Database Database
	Server   Server
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Printf("Error loading .env file: %v", err)
		return nil, err
	}

	cfg := &Config{}
	if err := env.Parse(&cfg.Database); err != nil {
		log.Printf("Error loading database config: %v", err)
		return nil, err
	}
	if err := env.Parse(&cfg.Server); err != nil {
		log.Printf("Error loading server config: %v", err)
		return nil, err
	}

	// デバッグ出力
	log.Printf("Loaded config: %+v", cfg)

	return cfg, nil
}
