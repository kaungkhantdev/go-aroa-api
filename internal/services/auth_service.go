package services

import (
	"go-aora-api/internal/models"
	"go-aora-api/internal/repository"
)

type AuthService struct {
	repo *repository.UserRepository
}

func NewAuthService(repo *repository.UserRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Register(data repository.CreateData) models.User {
	return s.repo.Create(data);
}