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
		utils.ErrorResponse(c, "Invalid input", http.StatusBadRequest)
        return
	}

	if err := utils.ValidateStruct(inputs); err != nil {
		utils.ErrorResponse(c, err.Error(), http.StatusBadRequest)
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
		utils.ErrorResponse(c, "creation failed", http.StatusBadRequest)
		return
	}

	utils.SuccessResponse(c, aoraCreated,"Success", http.StatusCreated)
	return
}

func (h *AoraHandler) FindById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		utils.ErrorResponse(c, "Invalid input", http.StatusBadRequest)
		return
	}

	aora, err := h.aoraService.FindByIdAoraService(id)
	if err != nil {
		utils.ErrorResponse(c, "Not found", http.StatusBadRequest)
		return
	}

	utils.SuccessResponse(c, aora,"Success", http.StatusCreated)
}

func (h *AoraHandler) FindAllAora(c *gin.Context) {
	aoras := h.aoraService.FindAllAoraService()
	utils.SuccessResponse(c, aoras,"Success", http.StatusCreated)
	return
}


func (h *AoraHandler) UpdateAora(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		utils.ErrorResponse(c, "Invalid id", http.StatusBadRequest)
		return
	}

	inputs := models.Aora{}
	if err := c.ShouldBindJSON(&inputs); err != nil {
        utils.ErrorResponse(c, "Invalid inputs", http.StatusBadRequest)
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
		utils.ErrorResponse(c, "Something wrong", http.StatusBadRequest)
		return
	}

	utils.SuccessResponse(c, aora,"Success", http.StatusCreated)

}