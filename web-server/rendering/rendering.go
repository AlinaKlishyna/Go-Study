package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Создаётся основной “маршрутизатор”, который будет слушать запросы и направлять их в нужные функции
	router := gin.Default() // создаем роутер по умолчанию - localhost:8080

	// /someJSON — адрес
	// gin.H — просто короткий способ создать словарь (map) map[string]interface{}
	// c.JSON(...) — отвечает клиенту в формате JSON
	router.GET("/someJSON", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{ // удобно, когда нужно просто быстро вернуть несколько ключей
			"message": "hey",
			"status":  http.StatusOK,
		})
	})

	// Создаётся структура (struct)
	// Gin, как всегда, создаёт c *gin.Context и передаёт тебе — внутри c ты можешь ответить пользователю
	router.GET("/moreJSON", func(c *gin.Context) { // Когда кто-то откроет /moreJSON, выполни вот эту функцию

		// Он говорит Gin (и вообще Go), что при превращении структуры в JSON
		var user struct {
			Name    string `json:"user"` // это поле должно называться user, а не Name в JSON
			Message string `json:"message"`
			Number  int    `json:"number"`
		}

		user.Name = "Alinka"
		user.Message = "Cio, come stai?"
		user.Number = 3791302696

		c.JSON(http.StatusAccepted, user) // лучше, когда ты хочешь описать модель данных (“Пользователь”, “Продукт”, “Книга”, “Сообщение”)
	})

	router.GET("/someXML", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{
			"message": "hey",
			"status":  http.StatusOK,
		})
	})

	router.GET("/someYAML", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{
			"message": "hey",
			"status":  http.StatusOK,
		})
	})

	router.GET("/someProtoBuf", func(c *gin.Context) {
		//reps := []int64{int64(1), int64(2)}
		//label := "test"

		//data := &protoexample.Test{
		//Label: &label,
		//Reps: reps,
		//}

		//c.ProtoBuf(http.StatusOK, data)
	})

	// Listen and serve on 0.0.0.0:8080
	router.Run(":8080")
}
