package service

import (
	"context"
	"log/slog"
	"roleclub-website/game-scheduler/internal/config"
	"roleclub-website/game-scheduler/internal/repository"
	"roleclub-website/game-scheduler/models"
	"roleclub-website/game-scheduler/pkg/logger"
	"time"
)

type ScheduleAccesserService struct {
	log        *slog.Logger
	cfg        *config.Config
	repository *repository.Repository
}

func NewScheduleAccesserService(log *slog.Logger, cfg *config.Config, rep *repository.Repository) *ScheduleAccesserService {
	return &ScheduleAccesserService{
		log:        log,
		cfg:        cfg,
		repository: rep,
	}
}

func (sa *ScheduleAccesserService) GetSchedule(ctx context.Context, from time.Time, to time.Time) ([]*models.Event, error) {
	const fn = "service.schedule-manager.GetSchedule"

	log := sa.log.With(
		slog.String("fn", fn),
		slog.String("start date", from.String()),
		slog.String("end date", to.String()),
	)

	schedule, err := sa.repository.GetSchedule(ctx, from, to)
	if err != nil {
		log.Error("failed to get the schedule", logger.Error(err))
		return nil, err
	}

	log.Info("schedule was found successfully")

	return schedule, nil
}
