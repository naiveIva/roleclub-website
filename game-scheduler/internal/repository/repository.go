package repository

import (
	"context"
	"time"

	"github.com/google/uuid"

	"roleclub-website/game-scheduler/models"
)

const (
	GamesTable  = "games"
	EventsTable = "scheduled_events"
)

type Repository struct {
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
