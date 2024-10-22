package services

import (
	"go-aora-api/internal/models"
	"go-aora-api/internal/repository"
)

type AoraService struct {
	repo *repository.AoraRepository
}

func NewAoraService(repo *repository.AoraRepository) *AoraService {
	return &AoraService{ repo: repo }
}

func (s *AoraService) CreateAoraService(data repository.CreateAoraData) (models.Aora, error) {
	return s.repo.CreateAora(data);
}

func (s *AoraService) UpdateAoraService(id int, data repository.UpdateDataAora) (models.Aora, error) {
	return s.repo.UpdateAora(id, data);
}

func (s *AoraService) FindByIdAoraService(id int) (models.Aora, error) {
	return s.repo.FindByIdAora(id)
}

func (s *AoraService) FindAllAoraService() []models.Aora {
	return s.repo.FindAllAora()
}