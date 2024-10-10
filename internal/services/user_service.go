package services

import (
	"go-aora-api/internal/models"
	"go-aora-api/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

// constructor 
func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUserService(data repository.CreateData) models.User {
	return s.repo.Create(data);
}