package handlers

import (
	"go-aora-api/internal/handlers/requests"
	"go-aora-api/internal/repository"
	"go-aora-api/internal/services"
	"go-aora-api/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AoraHandler struct {
	aoraService *services.AoraService
}

func NewAoraHandler(aoraService *services.AoraService) *AoraHandler {
	return &AoraHandler{ aoraService: aoraService }
}

func (h *AoraHandler) CreateAoraHandler(c *gin.Context) {

	var inputs requests.AoraCreateRequest

	if err := c.ShouldBindBodyWithJSON(&inputs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
	}

	if err := utils.ValidateStruct(inputs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	aora := repository.CreateAoraData{
		Name:        inputs.Name,
		VideoURL: 	 inputs.VideoURL,
		VideoThumb:  inputs.VideoThumb,
		Author: 	 inputs.Author,
		AuthorPhoto: inputs.AuthorPhoto,
		Description: inputs.Description,
	}

	aoraCreated, err := h.aoraService.CreateAoraService(aora)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Creation failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{ "aroa": aoraCreated })
	return
}