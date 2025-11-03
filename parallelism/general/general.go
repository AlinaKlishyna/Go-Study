package main

import (
	"fmt"
	"math/rand" // для генерации случайных чисел (имитация разных временных задержек)
	"time"      // работа со временем, измерение, паузы (Sleep)
)

func main() {
	t := time.Now() // Засекаем время и сохраняем текущее время
	// Настраиваем генератор случайных чисел
	// Seed() - задаёт «начальное зерно» для генератора случайных чисел, чтобы результаты были разные при каждом запуске
	rand.Seed(t.UnixNano()) // UnixNano() возвращает текущее время в наносекундах, а это — идеальный источник случайности

	// Сначала парсится YouTube, потом Example.
	// Пока что не параллельно, а строго друг за другом
	go parseURL("http://youtube.com/")
	parseURL("http://example.com/")

	fmt.Printf("Parsing completed. Time Elapased: %.2f seconds\n", time.Since(t).Seconds())
}

// simulation
func parseURL(url string) {
	for i := 0; i < 5; i++ {
		// Генерируется случайная задержка от 500 до 999 миллисекунд
		latency := rand.Intn(500) + 500 // rand.Intn(500) даёт число от 0 до 499, добавляем 500 → диапазон 500–999

		time.Sleep(time.Duration(latency) * time.Millisecond) // делает паузу — программа «засыпает» на это время

		fmt.Printf("Parsing <%s> - Step %d - Latency %d\n", url, i+1, latency)
	}
}
