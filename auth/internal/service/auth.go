package service

import (
	"auth/internal/repository"
	"auth/models"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type AuthorizationService struct {
	repository *repository.Repository
}

func NewAuthorizationService(rep *repository.Repository) *AuthorizationService {
	return &AuthorizationService{
		repository: rep,
	}
}

func (auth *AuthorizationService) CreateUser(user *models.User) error {
	hash, err := HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hash

	err = auth.repository.CreateUser(user)
	if errors.Is(err, repository.ErrUserAlreadyExists) {
		return ErrUserAlreadyExists
	}
	return err
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (auth *AuthorizationService) Login(telNumber, password string) (*models.User, error) {
	usr, err := auth.repository.GetUser(telNumber)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	if !CheckPasswordHash(password, usr.Password) {
		return nil, ErrorWrongPassword
	}
	// todo: create jwt token

	return usr, nil
}
