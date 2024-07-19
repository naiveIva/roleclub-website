package service

import (
	"auth/internal/config"
	"auth/internal/repository"
	"auth/models"
	"context"
	"errors"
	"log/slog"
)

var (
	ErrUserNotFound        = errors.New("user not found")
	ErrUserAlreadyExists   = errors.New("user already exists")
	ErrorWrongPassword     = errors.New("wrong password")
)

type Authorization interface {
	RegisterUser(ctx context.Context, user *models.User) error
	Login(ctx context.Context, telNumber, password string) (string, error)
}

type Service struct {
	Authorization
}

func NewService(log *slog.Logger, cfg *config.Config, rep *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthorizationService(log, cfg, rep),
	}
}
