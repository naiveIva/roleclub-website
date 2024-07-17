package postgres

import (
	"database/sql"
	"auth/internal/repository"
)

func NewRepository(db *sql.DB) *repository.Repository {
	return &repository.Repository{
		Authorization: NewAuthRepository(db),
	}
}
