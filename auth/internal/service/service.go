package service

import (
	"auth/internal/repository"
	"auth/models"
	"errors"
)

var (
	ErrUserNotFound        = errors.New("user not found")
	ErrUserAlreadyExists   = errors.New("user already exists")
	ErrInvalidPlayerStatus = errors.New("invalid player status")
	ErrorWrongPassword     = errors.New("wrong password")
)

type Authorization interface {
	CreateUser(user *models.User) error
	Login(telNumber, password string) (*models.User, error)
}

type Service struct {
	Authorization
}

func NewService(rep *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthorizationService(rep),
	}
}
