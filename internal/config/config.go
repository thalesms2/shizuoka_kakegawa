package config

import (
	"log"
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	Port string `env:"PORT" envDefault:"8080"`
	DBHost string `env:"DB_HOST" envDefault:"localhost"`
	DBPort string `env:"DB_PORT" envDefault:"5432"`
	DBUser string `env:"DB_USER" envRequired:"true"`
	DBPassword string `env:"DB_USER" envRequired:"true"`
	DBName string `env:"DB_USER" envRequired:"true"`
}

func LoadConfig() *Config {
	_ = godotenv.Load()

	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		log.Fatalf("Error on load environment variables: %v", err)
	}

	return cfg
}
