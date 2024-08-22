package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"roleclub-website/game-scheduler/internal/repository"
	"roleclub-website/game-scheduler/models"

	"github.com/google/uuid"
)

type EventManagerRepository struct {
	db *sql.DB
}

func NewEventManagerRepository(db *sql.DB) *EventManagerRepository {
	return &EventManagerRepository{
		db: db,
	}
}

func (em *EventManagerRepository) AddEvent(ctx context.Context, event *models.Event) error {
	const fn = "repository.postgres.event-manager.AddEvent"

	_, err := em.db.ExecContext(ctx,
		fmt.Sprintf(
			`INSERT INTO %s
			(id, game_name, event_date,
			num_of_sessions, is_subscription_open)
			VALUES ($1, $2, $3, $4, $5)`,
			repository.EventsTable,
		),
		uuid.New(), event.Gamename, event.Date,
		event.NumOfSessions, event.IsSubscriptionOpen,
	)

	if err != nil {
		return fmt.Errorf("%s: %v", fn, err)
	}

	return nil
}

func (em *EventManagerRepository) GetEvent(ctx context.Context, id uuid.UUID) (*models.Event, error) {
	const fn = "repository.postgres.event-manager.GetEvent"

	event := &models.Event{}

	row := em.db.QueryRowContext(ctx,
		fmt.Sprintf(
			`SELECT * FROM %s WHERE id=$1`,
			repository.EventsTable,
		),
		id,
	)

	err := row.Scan(
		&event.ID,
		&event.Gamename,
		&event.Date,
		&event.NumOfSessions,
		&event.IsSubscriptionOpen,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("%s: %v", fn, repository.ErrEventNotFound)
	}

	if err != nil {
		return nil, fmt.Errorf("%s: %v", fn, err)
	}

	return event, nil
}

func (em *EventManagerRepository) DeleteEvent(ctx context.Context, id uuid.UUID) error {
	const fn = "repository.postgres.event-manager.DeleteEvent"

	_, err := em.db.ExecContext(ctx,
		fmt.Sprintf(
			`DELETE * FROM %s
			WHERE id=$1`,
			repository.EventsTable,
		),
		id,
	)

	if err != nil {
		return fmt.Errorf("%s: %v", fn, err)
	}

	return nil
}
