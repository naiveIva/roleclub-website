package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Env      string   `yaml:"env" env-default:"local"`
	Database Database `yaml:"database"`
	Server   Server   `yaml:"server"`
}

type Database struct {
	Host     string `yaml:"host" env-default:"localhost"`
	Port     string `yaml:"port" env-default:"5432"`
	SSLMode  string `yaml:"sslmode" env-default:"disabled"`
	DBName   string `yaml:"dbname" env:"POSTGRES_DB"`
	Username string `yaml:"username" env:"POSTGRES_USER"`
	Password string `yaml:"password" env:"POSTGRES_PASSWORD"`
}

type Server struct {
	Host string `yaml:"host" env-default:"localhost"`
	Port string `yaml:"port" env-default:"8081"`
}

func MustInit(cfg_name string) *Config {
	var cfg Config
	if err := cleanenv.ReadConfig(cfg_name, &cfg); err != nil {
		log.Fatal("cannot read config")
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
	cfg.Database.DBName = os.Getenv("POSTGRES_DB")
	cfg.Database.Username = os.Getenv("POSTGRES_USER")
	cfg.Database.Password = os.Getenv("POSTGRES_PASSWORD")
	return &cfg
}
