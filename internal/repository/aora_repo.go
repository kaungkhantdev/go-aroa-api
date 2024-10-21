package repository

import (
	"go-aora-api/internal/database"
	"go-aora-api/internal/models"
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

func (repo *AoraRepository) Create(data CreateAoraData) (models.Aora, error) {
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