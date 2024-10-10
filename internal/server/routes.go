package server

import (
	"go-aora-api/pkg/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()


	r.GET("/", s.HelloWorldHandler)

	// apiGroup := r.Group("/api")
	// apiGroup.Use(middleware.AuthMiddleware()) 
	// handlers.UserRoutes(apiGroup)
	// handlers.TestRoutes(apiGroup)

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
