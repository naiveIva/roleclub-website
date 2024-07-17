package service

import (
	"auth/internal/repository"
	"auth/models"
)

type Authorization interface {
	CreateUser(user models.User) error
}

type Service struct {
	Authorization
}

func NewService(rep *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthorizationService(rep),
	}
}