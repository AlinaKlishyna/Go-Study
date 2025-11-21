package api

import (
	"github.com/AlinaKlishyna/Go-Study.git/jwt/core"
	"github.com/AlinaKlishyna/Go-Study.git/jwt/internal/auth"
	"github.com/gin-gonic/gin"
)

func NewRouter(domain *core.Domain) *gin.Engine {
	router := gin.Default()

	authService := auth.NewAuthorization(domain)

	userHandler := &UserHandler{
		Domain: domain,
		Auth:   authService,
	}

	authHandler := &AuthHandler{
		Auth: authService,
	}

	user := router.Group("/users")
	userHandler.Register(user)

	router.POST("/login", authHandler.SignIn)

	return router
}
