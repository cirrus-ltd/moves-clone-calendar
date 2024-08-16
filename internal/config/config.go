package config

import (
	"log"

	"github.com/caarlos0/env/v10"
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
	var cfg *Config
	if err := env.Parse(&cfg); err != nil {
		log.Printf("Error loading .env file: %v", err)
		return nil, err
	}
	return cfg, nil
}
