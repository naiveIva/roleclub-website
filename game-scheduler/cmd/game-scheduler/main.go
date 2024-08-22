package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"roleclub-website/game-scheduler/internal/config"
	"roleclub-website/game-scheduler/internal/repository"
	"roleclub-website/game-scheduler/internal/repository/postgres"
	"roleclub-website/game-scheduler/models"
	"roleclub-website/game-scheduler/pkg/database"
	"roleclub-website/game-scheduler/pkg/logger"
	"time"
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

	db, err := database.NewPostgresDB(cfg)
	if err != nil {
		log.Error("failed to open sql connection: %v", logger.Error(err))
	}

	repo := postgres.NewRepository(db)

	// CreateGames(repo)

	GetGames(repo)

	// CreateEvents(repo)

	// sch, err := repo.GetSchedule(context.TODO(),
	// 	time.Date(2024, time.September, 1, 0, 0, 0, 0, time.UTC),
	// 	time.Date(2025, time.June, 1, 0, 0, 0, 0, time.UTC),
	// )

	// if err != nil {
	// 	log.Error("failed to get schedule:", err)
		
	// } else {
	// 	fmt.Println("-----")
	// 	for _, event := range sch {
	// 		fmt.Println(event)
	// 	}
	// }

	// sch, err = repo.GetSchedule(context.TODO(),
	// 	time.Date(2023, time.September, 1, 0, 0, 0, 0, time.Local),
	// 	time.Date(2024, time.June, 1, 0, 0, 0, 0, time.Local),
	// )

	// if err != nil {
	// 	log.Error("failed to get schedule:", err)
	// } else {
	// 	fmt.Println("-----")
	// 	for _, event := range sch {
	// 		fmt.Println(event)
	// 	}
	// }
}

func CreateGames(repo *repository.Repository) {
	games := []models.Game{
		{
			Name:              "уточка",
			Authors:           "Витя Зарипов",
			Description:       "игра про уточку",
			Complexity:        3,
			PlayersPerSession: 12,
			MastersPerSession: 1,
			Roles:             []string{"ведьма", "крысолов"},
		},
		{
			Name:              "логос",
			Authors:           "Ксюша Борщова",
			Description:       "игра про закрытую школу",
			Complexity:        4,
			PlayersPerSession: 12,
			MastersPerSession: 2,
			Roles:             []string{"биолог", "ученик", "ученица"},
		},
	}

	for _, game := range games {
		err := repo.AddGame(context.TODO(), &game)
		if err != nil {
			log.Println("failed to add game:", err)
		}
	}
}

func GetGames(repo *repository.Repository) {
	games := []string{"уточка", "логос", "бабубэ"}
	for _, game := range games {
		g, err := repo.GetGame(context.TODO(), game)
		if err != nil {
			log.Println("failed to find game:", err)
		} else {
			fmt.Println("FOUND")
			fmt.Println(g)
		}
	}
}

func CreateEvents(repo *repository.Repository) {
	events := []models.Event{
		{
			Gamename:           "уточка",
			Date:               time.Date(2024, time.September, 12, 16, 0, 0, 0, time.Local),
			NumOfSessions:      3,
			IsSubscriptionOpen: false,
		},
		{
			Gamename:           "логос",
			Date:               time.Date(2024, time.September, 30, 16, 0, 0, 0, time.Local),
			NumOfSessions:      3,
			IsSubscriptionOpen: false,
		},
	}

	for _, event := range events {
		err := repo.AddEvent(context.TODO(), &event)
		if err != nil {
			log.Println("failed to add event:", err)
		} else {
			log.Println("added event successfully:", event.Gamename)
		}
	}
}
