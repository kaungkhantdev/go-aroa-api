package handlers

import (
	"go-aora-api/internal/repository"
	"go-aora-api/internal/services"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.RouterGroup) {

	userRepo := &repository.UserRepository{}
	userService := services.NewUserService(userRepo)
	authService := services.NewAuthService(userService)

	authHandler := NewAuthHandler(authService)

	router.POST("/login", authHandler.Login)
	router.POST("/register", authHandler.Register)
}