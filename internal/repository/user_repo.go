package repository

import (
	"go-aora-api/internal/database"
	"go-aora-api/internal/models"
)

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

	database.DB.Create(&user)
	return user;
}

func (repo *UserRepository) FindById(id int) models.User {
	user := models.User{ID: uint(id)}

	database.DB.First(&user)
	return user;
}


func (repo *UserRepository) FindAll() []models.User {
	users := []models.User{}

	database.DB.Find(&users)
	return users;
}