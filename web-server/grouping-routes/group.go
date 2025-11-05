package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// POST http://localhost:8080/first/login

// Представь, что у тебя сайт с разделами:
// /user/login
// /user/profile
// /admin/login
// /admin/dashboard
// Можно всё писать по отдельности, а можно красиво сгруппировать:
func main() {
	router := gin.Default()

	// simple group 1
	{
		first := router.Group("/first")
		first.GET("/login", loginEndPoint)
		first.POST("/login", loginEndPoint)
		first.POST("/submit", submitEndPoint)
		first.POST("/read", readEndPoint)
	}

	// simple froup 2
	{
		second := router.Group("/second")
		second.POST("/login", loginEndPoint)
		second.POST("/submit", submitEndPoint)
		second.POST("/read", readEndPoint)
	}

	router.Run(":8080")
}

func loginEndPoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "login - ok",
	})
}

func submitEndPoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "submit - ok",
	})
}

func readEndPoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "read - ok",
	})
}
