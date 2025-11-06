package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// объявляешь два маршрута, которые реагируют на разные паттерны URL
func main() {
	router := gin.Default()

	// /user/:lastname
	//  : перед словом lastname означает — это динамический параметр.
	router.GET("/user/:lastname", func(c *gin.Context) { // http://localhost:8080/user/alina --> Hello alina
		lastname := c.Param("lastname")
		c.String(http.StatusOK, "Hello %s", lastname)
	})

	// *action — захватывает всё, что идёт после lastname, включая /
	router.GET("/user/:lastname/*action", func(c *gin.Context) { // http://localhost:8080/user/alina/eating/pizza --> alina is /eating/pizza
		lastname := c.Param("lastname")
		action := c.Param("action")
		message := lastname + " is " + action // в *action всегда включён начальный слэш /
		c.String(http.StatusOK, message)
	})

	router.GET("/user/full/:lastname/:name", func(c *gin.Context) {
		name := c.Param("name")
		lastname := c.Param("lastname")
		message := "Hello, " + name + " " + lastname + "!"
		c.String(http.StatusOK, message)
	})

	router.Run(":8080")
}
