package repository

import (
	"errors"
	"fmt"
	"go-aora-api/internal/database"
	"go-aora-api/internal/models"

	"gorm.io/gorm"
)

type AoraRepository struct {}

type CreateAoraData struct {
	Name		string	
	VideoURL	string	
	VideoThumb	string	
	Author		string	
	AuthorPhoto string	
	Description	string	
}

func (repo *AoraRepository) CreateAora(data CreateAoraData) (models.Aora, error) {
	aora := models.Aora{
		Name:		data.Name,
		VideoURL:	data.VideoURL,
		VideoThumb: data.VideoThumb,
		Author:		data.Author,
		AuthorPhoto: data.AuthorPhoto,
		Description: data.Description,
	}

	database.DB.Create(&aora)
	return aora, nil;
}

func (repo *AoraRepository) FindByIdAora(id int) (models.Aora, error) {
	aora := models.Aora{ID: uint(id)}

	result := database.DB.First(&aora)
	if result.Error != nil {
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            return models.Aora{}, fmt.Errorf("test with ID %d not found", id)
        }
        return models.Aora{}, result.Error
    }

	return aora, nil;
}

func (repo *AoraRepository) FindAllAora() []models.Aora {
	aoras := []models.Aora{}

	database.DB.Find(&aoras)
	return aoras;
}

type UpdateDataAora struct {
	Name		*string	
	VideoURL	*string	
	VideoThumb	*string	
	Author		*string	
	AuthorPhoto *string	
	Description	*string
}

func (repo *AoraRepository) UpdateAora(id int, data UpdateDataAora) (models.Aora, error) {
	aora, err := repo.FindByIdAora(id);

	if err != nil {
		return models.Aora{}, err
	}

	if data.Name != nil {
		aora.Name = *data.Name
	}

	if data.VideoThumb != nil {
		aora.VideoThumb = *data.VideoThumb
	}
	if data.VideoURL != nil {
		aora.VideoURL = *data.VideoURL
	}
	if data.Author != nil {
		aora.Author = *data.Author
	}
	if data.AuthorPhoto != nil {
		aora.AuthorPhoto = *data.AuthorPhoto
	}

	if data.Description != nil {
		aora.Description = *data.Description
	}

	database.DB.Save(&aora)
	return models.Aora{}, nil
}