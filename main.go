package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Структура для JSON-ответа
type HealthResponse struct {
	Status string `json:"status"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	// Проверяем, что метод запроса GET
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Создаем объект ответа
	response := HealthResponse{
		Status: "OK",
	}

	// Устанавливаем заголовок ответа как JSON
	w.Header().Set("Content-Type", "application/json")

	// Кодируем структуру в JSON и отправляем ответ
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func main() {
	// Регистрируем обработчик для пути /health/
	http.HandleFunc("/health/", healthHandler)

	// Запускаем сервер на порту 8000
	fmt.Println("Server is running on port 8000...")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
