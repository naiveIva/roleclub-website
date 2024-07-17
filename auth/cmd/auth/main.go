package main

import (
	"flag"
	"auth/internal/config"
	"auth/internal/logger"
	"auth/internal/repository/postgres"
	"auth/models"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config", "config/config.yaml", "path to config file")
}

func main() {
	flag.Parse()

	cfg := config.MustInit(configPath)

	log := logger.Init()

	db, err := postgres.NewPostgresDB(cfg)
	if err != nil {
		log.Error("failed to open sql connection: %v", err)
	}

	repo := postgres.NewRepository(db)

	users := []models.User{
		{
			FirstName:    "Иоанна",
			LastName:     "Мысниченко",
			FatherName:   "Николаевна",
			TelNumber:    "12345",
			Password:     "qwerty",
			IsHSEStudent: false,
		},
		{
			FirstName:    "Ангелина",
			LastName:     "Никитич",
			FatherName:   "Батьковна",
			TelNumber:    "67890",
			Password:     "ффффффф",
			IsHSEStudent: true,
		},
		{
			FirstName:    "Ксения",
			LastName:     "Борщева",
			FatherName:   "Батьковна",
			TelNumber:    "000",
			Password:     "fffjjjjfjfjf",
			IsHSEStudent: false,
		},
	}

	for _, user := range users {
		err = repo.CreateUser(&user)
		if err != nil {
			log.Error(err.Error())
		}
	}

	// TODO: init storage

	// TODO: init router

	// TODO: run server
}
