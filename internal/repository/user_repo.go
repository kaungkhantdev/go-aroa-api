package repository

import (
	"errors"
	"fmt"
	"go-aora-api/internal/database"
	"go-aora-api/internal/models"
	"go-aora-api/pkg/hash"

	"gorm.io/gorm"
)

type UserRepository struct {}


type CreateData struct {
	Name		string
	Email 		string
	Avatar		string
	Password	string
}

func (repo *UserRepository) Create(data CreateData) (models.User, error) {
	hashedPassword, err := hash.HashPassword(data.Password)
	if err != nil {
		return models.User{}, err
	}

	user := models.User{
		Name: data.Name,
		Email: data.Email,
		Avatar: data.Avatar,
		Password: hashedPassword,
	}

	database.DB.Create(&user)
	return user, nil;
}

type UpdateData struct {
	Name		*string
	Email 		*string
	Avatar		*string
}

func (repo *UserRepository) Update(id int, data UpdateData) (models.User, error) {
	user, err := repo.FindById(id);

	if err != nil {
		return models.User{}, err
	}

	if data.Name != nil {
		user.Name = *data.Name
	}

	if data.Email != nil {
		user.Email = *data.Email
	}

	if data.Avatar != nil {
		user.Avatar = *data.Avatar
	}

	database.DB.Save(&user)
	return models.User{}, nil
}

func (repo *UserRepository) FindById(id int) (models.User, error) {
	user := models.User{ID: uint(id)}

	result := database.DB.First(&user)
	if result.Error != nil {
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            return models.User{}, fmt.Errorf("test with ID %d not found", id)
        }
        return models.User{}, result.Error
    }

	return user, nil;
}


func (repo *UserRepository) FindAll() []models.User {
	users := []models.User{}

	database.DB.Find(&users)
	return users;
}


func (repo *UserRepository) FindByEmail(email string) (models.User, error) {
	user := models.User{}

	result := database.DB.Where("email=?", email).First(&user);
	if result.Error != nil {
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            return models.User{}, fmt.Errorf("user with email %v not found", email)
        }
        return models.User{}, result.Error
    }

	return user, nil;
}
