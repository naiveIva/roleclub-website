package service

import (
	"auth/internal/repository"
	"auth/models"
	"auth/pkg/logger"
	"context"
	"errors"
	"fmt"
	"log/slog"

	"golang.org/x/crypto/bcrypt"
)

type AuthorizationService struct {
	log        *slog.Logger
	repository *repository.Repository
}

func NewAuthorizationService(log *slog.Logger, rep *repository.Repository) *AuthorizationService {
	return &AuthorizationService{
		repository: rep,
		log: log,
	}
}

func (auth *AuthorizationService) RegisterUser(ctx context.Context, user *models.User) error {
	const fn = "service.auth.CreateUser"

	log := auth.log.With(
		slog.String("fn", fn),
		slog.String("user", fmt.Sprintf("%s, %s", user.FirstName, user.TelNumber)),
	)

	log.Info("registering user")

	hash, err := HashPassword(user.Password)
	if err != nil {
		log.Error("failed to generate password hash", logger.Error(err))
		return fmt.Errorf("%s: %w", fn, err)
	}
	user.Password = hash

	err = auth.repository.CreateUser(user)
	if err != nil {
		log.Error("failed to save user", logger.Error(err))
		if errors.Is(err, repository.ErrUserAlreadyExists) {
			err = ErrUserAlreadyExists
			return fmt.Errorf("%s: %w", fn, err)
		}
		return err
	}

	log.Info("user created successfully")

	return nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (auth *AuthorizationService) Login(ctx context.Context, telNumber, password string) (*models.User, error) {
	const fn = "service.auth.Login"

	log := auth.log.With(
		slog.String("fn", fn),
		slog.String("user tel number", telNumber),
	)

	log.Info("logging in the user")

	usr, err := auth.repository.GetUser(telNumber)
	if err != nil {
		log.Error("failed to find user", logger.Error(err))
		if errors.Is(err, repository.ErrUserNotFound) {
			return nil, fmt.Errorf("%s: %w", fn, err)
		}
		return nil, err
	}

	if !CheckPasswordHash(password, usr.Password) {
		err = ErrorWrongPassword
		log.Error("password verification failed", logger.Error(err))
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	log.Info("user logged in successfully")

	// todo: create jwt token

	return usr, nil
}
