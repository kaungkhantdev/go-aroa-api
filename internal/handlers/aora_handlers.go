package handlers

import (
	"go-aora-api/internal/handlers/requests"
	"go-aora-api/internal/models"
	"go-aora-api/internal/repository"
	"go-aora-api/internal/services"
	"go-aora-api/internal/utils"
	"net/http"
	"strconv"

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

func (h *AoraHandler) FindById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	aora, err := h.aoraService.FindByIdAoraService(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{ "aroa": aora })
	return
}

func (h *AoraHandler) FindAllAora(c *gin.Context) {
	aoras := h.aoraService.FindAllAoraService()
	c.JSON(http.StatusOK, gin.H{ "aroas": aoras })
	return
}


func (h *AoraHandler) UpdateAora(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	inputs := models.Aora{}
	if err := c.ShouldBindJSON(&inputs); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

	var data repository.UpdateDataAora
    if inputs.Name != "" {
        data.Name = &inputs.Name
    }
    if inputs.VideoURL != "" {
		data.VideoURL = &inputs.VideoURL
	}
	if inputs.VideoThumb != "" {
		data.VideoThumb = &inputs.VideoThumb
	}
	if inputs.Author != "" {
		data.Author = &inputs.Author
	}
	if inputs.AuthorPhoto != "" {
		data.AuthorPhoto = &inputs.AuthorPhoto
	}
	if inputs.Description != "" {
		data.Description = &inputs.Description
	}

	aora, err := h.aoraService.UpdateAoraService(id, data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Something wrong"})
		return
	}

	c.JSON(http.StatusOK, gin.H{ "data": aora })

}