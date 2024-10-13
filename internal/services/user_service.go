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

func (s *UserService) CreateUserService(data repository.CreateData) (models.User, error) {
	return s.repo.Create(data);
}

func (s *UserService) CheckEmail(email string) bool {
	_, err := s.repo.FindByEmail(email);
	return err == nil
}

func (s *UserService) FindByEmail(email string) (models.User, error) {
	return s.repo.FindByEmail(email);
}