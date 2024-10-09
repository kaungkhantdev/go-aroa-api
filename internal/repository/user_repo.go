package repository

import "go-aora-api/internal/models"

type UserRepository struct {}


type CreateData struct {
	Name	string
	Email 	string
	Avatar	string
}

func (repo *UserRepository) Create(data CreateData) models.User {
	user := models.User{
		Name: data.Name,
		Email: data.Email,
		Avatar: data.Avatar,
	}

	return user;
}