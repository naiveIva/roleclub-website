package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"roleclub-website/game-scheduler/internal/repository"
	"roleclub-website/game-scheduler/models"
	"time"
)

type ScheduleAccesser struct {
	db *sql.DB
}

func NewScheduleAccesser(db *sql.DB) *ScheduleAccesser {
	return &ScheduleAccesser{
		db: db,
	}
}

func (sa *ScheduleAccesser) GetSchedule(ctx context.Context, from time.Time, to time.Time) ([]*models.Event, error) {
	const fn = "repository.postgres.schedule-accesser.GetSchedule"

	schedule := make([]*models.Event, 0)

	rows, err := sa.db.QueryContext(ctx,
		fmt.Sprintf(
			`SELECT * FROM %s
			WHERE event_date BETWEEN $1 AND $2
			ORDER BY event_date ASC`,
			repository.EventsTable,
		),
		from, to,
	)

	if err != nil {
		return nil, fmt.Errorf("%s: %v", fn, err)
	}
	defer rows.Close()

	for rows.Next() {
		event := &models.Event{}
		err = rows.Scan(
			&event.ID,
			&event.Gamename,
			&event.Date,
			&event.NumOfSessions,
			&event.IsSubscriptionOpen,
		)

		if err != nil {
			return nil, fmt.Errorf("%s: %v", fn, err)
		}
		schedule = append(schedule, event)
		
	}

	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("%s: %v", fn, err)
	}

	return schedule, nil
}
