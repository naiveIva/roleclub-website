package service

import (
	"auth/internal/repository"
	"auth/models"

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

func (auth *AuthorizationService) CreateUser(user models.User) error {
	hash, err := HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hash
	return auth.repository.CreateUser(&user)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}