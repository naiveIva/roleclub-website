package service

import (
	"auth/internal/config"
	"auth/internal/repository"
	"auth/models"
	"context"
	"log/slog"
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
