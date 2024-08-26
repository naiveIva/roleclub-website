package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"roleclub-website/game-scheduler/internal/repository"
	"roleclub-website/game-scheduler/models"
	"roleclub-website/game-scheduler/pkg/customerrors"

	"github.com/jackc/pgerrcode"
	"github.com/lib/pq"
)

type GameManagerRepository struct {
	db *sql.DB
}

func NewGameManagerRepository(db *sql.DB) *GameManagerRepository {
	return &GameManagerRepository{
		db: db,
	}
}

func (gm *GameManagerRepository) AddGame(ctx context.Context, game *models.Game) error {
	const fn = "repository.postgres.game-manager.AddGame"

	_, err := gm.db.ExecContext(ctx,
		fmt.Sprintf(
			`INSERT INTO %s
			(game_name, authors,
			description, complexity,
			players_per_session, masters_per_session,
			roles, game_state)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
			repository.GamesTable,
		),
		game.Name, game.Authors,
		game.Description, game.Complexity,
		game.PlayersPerSession, game.MastersPerSession,
		(*pq.StringArray)(&game.Roles), game.State.String(),
	)

	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok && pgErr.Code == pgerrcode.UniqueViolation {
			err = customerrors.ErrGameAlreadyExists
		}
		return fmt.Errorf("%s: %w", fn, err)
	}

	return nil
}

func (gm *GameManagerRepository) GetGame(ctx context.Context, gamename string) (*models.Game, error) {
	const fn = "repository.postgres.game-manager.GetGame"

	game := &models.Game{}
	var gameState string

	row := gm.db.QueryRowContext(ctx,
		fmt.Sprintf(
			`SELECT * FROM %s WHERE game_name=$1`,
			repository.GamesTable,
		),
		gamename,
	)

	err := row.Scan(
		&game.Name,
		&game.Authors,
		&game.Description,
		&game.Complexity,
		&game.PlayersPerSession,
		&game.MastersPerSession,
		(*pq.StringArray)(&game.Roles),
		gameState,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = customerrors.ErrGameNotFound
		}
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	game.State, err = models.StringToGameState(gameState)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	return game, nil
}

func (gm *GameManagerRepository) DeleteGame(ctx context.Context, gamename string) error {
	const fn = "repository.postgres.game-manager.DeleteGame"

	_, err := gm.db.ExecContext(ctx,
		fmt.Sprintf(
			`DELETE FROM %s
			WHERE game_name=$1`,
			repository.GamesTable,
		),
		gamename,
	)

	if err != nil {
		return fmt.Errorf("%s: %w", fn, err)
	}

	return nil
}
