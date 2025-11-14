package main

import (
	"fmt"
	"log"
	"net/http"

	"sportteam/internal/db"
	"sportteam/internal/handlers"
)

func main() {
	// Инициализация подключения к БД
	db.Init()
	defer db.DB.Close()

	// Маршрут для JSON-данных
	http.HandleFunc("/players", handlers.HandlePlayers)

	// Статические файлы (index.html, css/js)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	fmt.Println("Сервер запущен на http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Ошибка HTTP-сервера: %v", err)
	}
}
