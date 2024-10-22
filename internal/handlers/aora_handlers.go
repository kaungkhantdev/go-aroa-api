package handlers

import "go-aora-api/internal/services"

type AoraHandler struct {
	aoraService *services.AoraService
}

func NewAoraHandler(aoraService *services.AoraService) *AoraHandler {
	return &AoraHandler{ aoraService: aoraService }
}