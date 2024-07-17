package postgres

import (
	"fmt"
	"auth/internal/config"

	"database/sql"

	_ "github.com/lib/pq"
)

const (
	usersTable = "players"

	playerStatus = "player"
	masterStatus = "master"
	adminStatus  = "admin"
)

func NewPostgresDB(cfg *config.Config) (*sql.DB, error) {
	const fn = "storage.postgres.NewPostgresDB"

	db, err := sql.Open("postgres", fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.DBName,
		cfg.Database.Username, cfg.Database.Password, cfg.Database.SSLMode))

	if err != nil {
		return nil, fmt.Errorf("%s: %v", fn, err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("%s: %v", fn, err)
	}

	return db, nil
}
