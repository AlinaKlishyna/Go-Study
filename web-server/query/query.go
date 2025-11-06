package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// query-параметры - данные, которые приходят в URL после знака ?
// параметры пути (:name)   -  /books/10 — конкретная книга с id = 10
// query-параметры (?name=value)  -  /books?author=Tolstoy — фильтр: “все книги Толстого”

func main() {
	router := gin.Default()

	router.GET("/tea", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest") // - c.DefaultQuery("param", "значение по умолчанию"), если параметр отсутствует — вернёт значение по умолчанию
		lastname := c.Query("lastname")                   // - берёт параметр из URL, если его нет, возвращает пустую строку ("")

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	router.Run(":8080")
}
