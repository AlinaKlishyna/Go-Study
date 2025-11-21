package api

import (
	"github.com/AlinaKlishyna/Go-Study.git/jwt/core"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Domain *core.Domain
	Auth   interface {
		Register(username, password string) (*core.User, error)
	}
}

type registerInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *UserHandler) Register(group *gin.RouterGroup) {
	group.POST("/register", func(c *gin.Context) {
		var body registerInput
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(400, gin.H{"error": "invalid json"})
			return
		}

		user, err := h.Auth.Register(body.Username, body.Password)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(201, user)
	})
}
