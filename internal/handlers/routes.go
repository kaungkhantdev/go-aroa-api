package handlers

import (
	"go-aora-api/internal/middleware"
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

func AoraRoutes(router *gin.RouterGroup) {
	
	aoraRepo := &repository.AoraRepository{}
	aoraService := services.NewAoraService(aoraRepo)
	
	aoraHandler := NewAoraHandler(aoraService)

	aoraGroup := router.Group("aoras")
	aoraGroup.Use(middleware.AuthMiddleware())

	aoraGroup.POST("", aoraHandler.CreateAoraHandler)
	aoraGroup.GET("/:id", aoraHandler.FindById)
	aoraGroup.GET("", aoraHandler.FindAllAora)
	aoraGroup.PATCH("/:id", aoraHandler.UpdateAora)
}