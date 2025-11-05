package customconfig

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default() // Это готовый HTTP-сервер, у которого уже есть встроенные обработчики ошибок, логирование и возможность добавлять маршруты
	// ":8080" → слушай все подключения на порту 8080
	router.Run(":8080")
	http.ListenAndServe(":8080", router) // Она запускает обычный сервер HTTP на порту 8080
}
