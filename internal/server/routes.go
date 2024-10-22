package server

import (
	"fmt"
	"go-aora-api/internal/handlers"
	"go-aora-api/pkg/jwt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()


	r.GET("/", s.HelloWorldHandler)
	r.POST("/verify", s.VerifyToken)

	apiGroup := r.Group("/api")
	// apiGroup.Use(middleware.AuthMiddleware()) 
	handlers.AuthRoutes(apiGroup)
	// handlers.TestRoutes(apiGroup)

	/** AORA ROUTES */
	handlers.AoraRoutes(apiGroup)

	return r
}

func (s *Server) HelloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)

	userId := 1;
	token, err := jwt.GenerateJWT(userId);

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	resp["token"] = token;
	resp["message"] = "Hello World"

	c.JSON(http.StatusOK, resp)
}

func (s *Server) VerifyToken(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
		c.Abort()
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ");
	if tokenString == authHeader {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Bearer token required"})
		c.Abort()
		return
	}

	claims, err := jwt.VerifyJWT(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}

	resp := map[string]string{}
	resp["auth_header"] = authHeader;
	resp["token"] = tokenString;
	
	userId := fmt.Sprintf("User Id: %d", claims.UserId)
	resp["user_id"] = userId
	
	c.JSON(http.StatusOK, resp)
}
