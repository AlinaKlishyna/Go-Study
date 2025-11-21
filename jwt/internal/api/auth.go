package api

import (
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	Auth interface {
		GenerateToken(username, password string) (string, error)
	}
}

type signInInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *AuthHandler) SignIn(c *gin.Context) {
	var input signInInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "invalid json"})
		return
	}

	token, err := h.Auth.GenerateToken(input.Username, input.Password)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"token": token})
}
