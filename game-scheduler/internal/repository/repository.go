package repository

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"

	"roleclub-website/game-scheduler/models"
)

var (
	ErrGameNotFound      = errors.New("game not found")
	ErrGameAlreadyExists = errors.New("game already exists")
	// ErrInvalidGameState  = errors.New("invalid game state")

	ErrEventNotFound = errors.New("event not found")
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
