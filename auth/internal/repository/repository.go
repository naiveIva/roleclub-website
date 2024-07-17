package repository

import (
	"errors"
	"auth/models"
)

var (
	ErrUserNotFound        = errors.New("user not found")
	ErrUserAlreadyExists   = errors.New("user already exists")
	ErrInvalidPlayerStatus = errors.New("invalid player status")
)

type Repository struct {
	Authorization
}

type Authorization interface {
	CreateUser(user *models.User) error
	GetUser(tel_number, password string) (*models.User, error)
	// IncrementPlayedGames(playerID int) error
	// SetPlayedGames(numOfGames int, playerID int) error
	// SetStatus(status string, playerID int) error
}
