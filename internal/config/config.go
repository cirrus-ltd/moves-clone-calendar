package config

import (
	"log"

	"github.com/caarlos0/env/v10"
)

type Database struct {
	User     string `env:"DB_USER"`
	DBName   string `env:"DB_NAME"`
	SSLMode  string `env:"DB_SSLMODE"`
	Password string `env:"DB_PASSWORD"`
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
	var cfg *Config
	if err := env.Parse(&cfg); err != nil {
		log.Printf("Error loading .env file: %v", err)
		return nil, err
	}
	return cfg, nil
}
