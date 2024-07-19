package service

import (
	"auth/internal/config"
	"auth/internal/repository"
	"auth/models"
	"auth/pkg/logger"
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthorizationService struct {
	log        *slog.Logger
	cfg        *config.Config
	repository *repository.Repository
}

func NewAuthorizationService(log *slog.Logger, cfg *config.Config, rep *repository.Repository) *AuthorizationService {
	return &AuthorizationService{
		log:        log,
		cfg:        cfg,
		repository: rep,
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

func (auth *AuthorizationService) Login(ctx context.Context, telNumber, password string) (string, error) {
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
			return "", fmt.Errorf("%s: %w", fn, err)
		}
		return "", err
	}

	if !CheckPasswordHash(password, usr.Password) {
		err = ErrorWrongPassword
		log.Error("password verification failed", logger.Error(err))
		return "", fmt.Errorf("%s: %w", fn, err)
	}

	log.Info("user logged in successfully")

	token, err := NewToken(usr, auth.cfg.Jwt.TokenTTL, auth.cfg.Jwt.JwtKey)
	if err != nil {
		log.Error("failed to generate jwt token", logger.Error(err))
		return "", fmt.Errorf("%s: %w", fn, err)
	}
	log.Info("jwt token created successfully")

	return token, nil
}

func NewToken(user *models.User, duration time.Duration, key string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub": user.ID,
			"exp": time.Now().Add(duration).Unix(),
		})

	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string, key string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return token, nil
}
