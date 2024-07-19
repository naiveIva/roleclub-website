package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Env      string   `yaml:"env" env-default:"local"`
	Database Database `yaml:"database"`
	Server   Server   `yaml:"server"`
	Jwt      JWT      `yaml:"jwt"`
}

type Database struct {
	Host     string `yaml:"host" env-default:"localhost"`
	Port     string `yaml:"port" env-default:"5432"`
	DBName   string `yaml:"dbname"`
	Username string `yaml:"username"`
	Password string `yaml:"password" env:"DB_PASSWORD"`
	SSLMode  string `yaml:"sslmode" env-default:"disabled"`
}

type Server struct {
	Address string `yaml:"address" env-default:"localhost:8082"`
}

type JWT struct {
	JwtKey   string        `env:"JWT_KEY"`
	TokenTTL time.Duration `yaml:"token_ttl" env-default:"1h"`
}

func MustInit(cfg_name string) *Config {
	var cfg Config
	if err := cleanenv.ReadConfig(cfg_name, &cfg); err != nil {
		log.Fatal("cannot read config")
	}

	// getting password from .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
	cfg.Database.Password = os.Getenv("DB_PASSWORD")
	return &cfg
}
