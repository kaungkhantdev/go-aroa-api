package handlers

import (
	"go-aora-api/internal/handlers/requests"
	"go-aora-api/internal/repository"
	"go-aora-api/internal/services"
	"go-aora-api/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)


type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Register(c *gin.Context) {

	var inputs requests.AuthRegisterRequest

	if err := c.ShouldBindBodyWithJSON(&inputs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
	}


	if err := utils.ValidateStruct(inputs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := repository.CreateData{
		Name: 		inputs.Name,
		Password: 	inputs.Password,
		Email: 		inputs.Email,
		Avatar: 	inputs.Avatar,
	}

	token, err := h.authService.Register(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Registration failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Registration successful", "token": token})

}

func (h *AuthHandler) Login(c *gin.Context) {
	var inputs requests.AuthLoginRequest

	if err := c.ShouldBindBodyWithJSON(&inputs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := utils.ValidateStruct(inputs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := services.LoginData{
		Email:    inputs.Email,
		Password: inputs.Password,
	}
	token, err := h.authService.Login(user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "user": token})
}