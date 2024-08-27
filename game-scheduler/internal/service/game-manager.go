package service

import (
	"context"
	"fmt"
	"log/slog"
	"roleclub-website/game-scheduler/internal/config"
	"roleclub-website/game-scheduler/internal/repository"
	"roleclub-website/game-scheduler/models"
	"roleclub-website/game-scheduler/pkg/logger"
)

type GameManagerService struct {
	log        *slog.Logger
	cfg        *config.Config
	repository *repository.Repository
}

func NewGameManagerService(log *slog.Logger, cfg *config.Config, rep *repository.Repository) *GameManagerService {
	return &GameManagerService{
		log:        log,
		cfg:        cfg,
		repository: rep,
	}
}

func (gm *GameManagerService) AddGame(ctx context.Context, game *models.Game) error {
	const fn = "service.game-manager.AddGame"

	log := gm.log.With(
		slog.String("fn", fn),
		slog.String("game", game.Name),
	)

	log.Info("adding game")

	err := gm.repository.AddGame(ctx, game)

	if err != nil {
		log.Error("failed to add game", logger.Error(err))
		return fmt.Errorf("%s: %w", fn, err)
	}

	log.Info("game added successfully")

	return nil
}

func (gm *GameManagerService) GetGame(ctx context.Context, gamename string) (*models.Game, error) {
	const fn = "service.game-manager.GetGame"

	log := gm.log.With(
		slog.String("fn", fn),
		slog.String("game", gamename),
	)

	log.Info("getting the game")

	game, err := gm.repository.GetGame(ctx, gamename)
	if err != nil {
		log.Error("failed to find the game", logger.Error(err))
		return nil, err
	}

	log.Info("game was found successfully")

	return game, nil
}

func (gm *GameManagerService) DeleteGame(ctx context.Context, gamename string) error {
	const fn = "service.game-manager.DeleteGame"

	log := gm.log.With(
		slog.String("fn", fn),
		slog.String("game", gamename),
	)

	log.Info("deleting the game")

	err := gm.repository.DeleteGame(ctx, gamename)
	if err != nil {
		log.Error("failed to delete the game", logger.Error(err))
		return err
	}

	log.Info("game was deleted successfully")

	return nil
}
