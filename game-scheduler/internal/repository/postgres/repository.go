package postgres

import (
	"database/sql"
	"roleclub-website/game-scheduler/internal/repository"
)

func NewRepository(db *sql.DB) *repository.Repository {
	return &repository.Repository{
		GameManager:      NewGameManagerRepository(db),
		EventManager:     NewEventManagerRepository(db),
		ScheduleAccesser: NewScheduleAccesser(db),
	}
}
