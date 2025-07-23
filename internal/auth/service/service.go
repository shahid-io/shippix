package service

import (
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"github.com/shahid-io/shippix/internal/auth/repository"
	"github.com/shahid-io/shippix/internal/common/auth"
	"github.com/shahid-io/shippix/pkg/models"
)

type AuthService struct {
	userRepo *repository.UserRepository
}

func NewAuthService(repo *repository.UserRepository) *AuthService {
	return &AuthService{userRepo: repo}
}

func (s *AuthService) Login(emailOrPhone, password string) (string, error) {
	user, err := s.userRepo.GetByEmailOrPhone(emailOrPhone)
	if err != nil || user == nil {
		return "", errors.New("user not found")
	}

	if user.Password != "" {
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
			return "", errors.New("invalid password")
		}
	}

	return auth.GenerateToken(user.ID, user.Role)
}

func (s *AuthService) Signup(user *models.User) (string, error) {
	if user.Email == "" && user.Phone == "" {
		return "", errors.New("email or phone required")
	}

	if strings.TrimSpace(user.Password) != "" {
		hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
		if err != nil {
			return "", err
		}
		user.Password = string(hashed)
	}

	if err := s.userRepo.Create(user); err != nil {
		return "", err
	}

	return auth.GenerateToken(user.ID, user.Role)
}
