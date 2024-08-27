package service

import (
	"context"
	"fmt"
	"log/slog"
	"roleclub-website/game-scheduler/internal/config"
	"roleclub-website/game-scheduler/internal/repository"
	"roleclub-website/game-scheduler/models"
	"roleclub-website/game-scheduler/pkg/logger"

	"github.com/google/uuid"
)

type EventManagerService struct {
	log        *slog.Logger
	cfg        *config.Config
	repository *repository.Repository
}

func NewEventManagerService(log *slog.Logger, cfg *config.Config, rep *repository.Repository) *EventManagerService {
	return &EventManagerService{
		log:        log,
		cfg:        cfg,
		repository: rep,
	}
}

func (em *EventManagerService) AddEvent(ctx context.Context, event *models.Event) error {
	const fn = "service.event-manager.AddEvent"

	log := em.log.With(
		slog.String("fn", fn),
		slog.String("event", fmt.Sprintf("%s, %s", event.Gamename, event.Date)),
	)

	log.Info("adding event")

	err := em.repository.AddEvent(ctx, event)
	if err != nil {
		log.Error("failed to add event", logger.Error(err))
	}

	log.Info("event was added successfully")

	return nil
}

func (em *EventManagerService) GetEvent(ctx context.Context, id uuid.UUID) (*models.Event, error) {
	const fn = "service.event-manager.GetEvent"

	log := em.log.With(
		slog.String("fn", fn),
		slog.String("event id", id.String()),
	)

	log.Info("getting the event")

	event, err := em.repository.GetEvent(ctx, id)
	if err != nil {
		log.Error("failed to find the event", logger.Error(err))
		return nil, err
	}

	log.Info("event was found successfully")

	return event, nil
}

func (em *EventManagerService) DeleteEvent(ctx context.Context, id uuid.UUID) error {
	const fn = "service.event-manager.DeleteEvent"

	log := em.log.With(
		slog.String("fn", fn),
		slog.String("event id", id.String()),
	)

	log.Info("deleting the event")

	err := em.repository.DeleteEvent(ctx, id)
	if err != nil {
		log.Error("failed to delete the event", logger.Error(err))
		return err
	}

	log.Info("event was deleted successfully")

	return nil
}
