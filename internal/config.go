package internal

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	DBHost     string `env:"DB_HOST"`
	DBName     string `env:"DB_NAME"`
	DBUsername string `env:"DB_USERNAME"`
	DBPassword string `env:"DB_PASSWORD"`
	DBPort     int    `env:"DB_PORT" envDefault:"5432"`
}

func (c Config) Dsn() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", c.DBUsername, c.DBPassword, c.DBHost, c.DBPort, c.DBName)
}

func LoadConfig() (Config, error) {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return Config{}, fmt.Errorf("Failed to load config: %w", err)
	}
	return cfg, nil
}
