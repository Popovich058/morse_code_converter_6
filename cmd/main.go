package main

import (
	"log"
	"os"
	"net/http"

	"six-sprint/internal/server"
)

func main() {
	// Создаём логгер
	// Выводим в os.Stdout с префиксом "SERVER: " и стандартными флагами
	logger := log.New(os.Stdout, "SERVER: ", log.LstdFlags)

	// Создаём сервер, передавая логгер
	srv := server.New(logger)

	// Запускаем сервер
	logger.Println("Starting server...")
	if err := srv.Start(); err != nil {
		// Если ошибка не связана с корректным закрытием сервера
		if err != http.ErrServerClosed {
			logger.Fatalf("Failed to start server: %v", err)
		}
	}
}
