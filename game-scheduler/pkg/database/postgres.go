package database

import (
	"database/sql"
	"fmt"
	"roleclub-website/game-scheduler/internal/config"

	_ "github.com/lib/pq"
)

const (
	UsersTable = "players"

	PlayerStatus = "player"
	MasterStatus = "master"
	AdminStatus  = "admin"
)

func NewPostgresDB(cfg *config.Config) (*sql.DB, error) {
	const fn = "pkg.database.NewPostgresDB"

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
