package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// DB — глобальная переменная с пулом соединений
var DB *sql.DB

// Init открывает соединение с PostgreSQL
func Init() {
	connStr := "host=localhost port=5432 user=sport_user password=sport_pass123 dbname=sportteam_db sslmode=disable"

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Ошибка при открытии БД: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("Не удалось подключиться к БД: %v", err)
	}

	fmt.Println("Подключение к БД успешно")
}
