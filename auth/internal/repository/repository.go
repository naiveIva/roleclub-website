package repository

import (
	"auth/models"
)

type Repository struct {
	Authorization
	// UserAlteration
}

type Authorization interface {
	CreateUser(user *models.User) error
	GetUser(tel_number string) (*models.User, error)
}

type UserAlteration interface {
	IncrementPlayedGames(playerID int) error
	SetPlayedGames(numOfGames int, playerID int) error
	SetStatus(status string, playerID int) error
}
