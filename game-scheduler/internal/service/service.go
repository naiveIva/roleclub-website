package service

import (
	"context"
	"log/slog"
	"roleclub-website/game-scheduler/internal/config"
	"roleclub-website/game-scheduler/internal/repository"
	"roleclub-website/game-scheduler/models"
	"time"

	"github.com/google/uuid"
)

type Service struct {
	GameManager
	EventManager
	ScheduleAccesser
}

type GameManager interface {
	AddGame(ctx context.Context, game *models.Game) error
	GetGame(ctx context.Context, gamename string) (*models.Game, error)
	DeleteGame(ctx context.Context, gamename string) error
}

type EventManager interface {
	AddEvent(ctx context.Context, event *models.Event) error
	GetEvent(ctx context.Context, id uuid.UUID) (*models.Event, error)
	DeleteEvent(ctx context.Context, id uuid.UUID) error
}

type ScheduleAccesser interface {
	GetSchedule(ctx context.Context, from time.Time, to time.Time) ([]*models.Event, error)
}

func NewService(log *slog.Logger, cfg *config.Config, rep *repository.Repository) *Service {
	return &Service{
		GameManager:      NewGameManagerService(log, cfg, rep),
		EventManager:     NewEventManagerService(log, cfg, rep),
		ScheduleAccesser: NewScheduleAccesserService(log, cfg, rep),
	}
}
