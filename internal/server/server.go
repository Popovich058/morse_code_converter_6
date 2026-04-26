package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"six-sprint/internal/handlers"
)

// Представляем HTTP-сервер с логгером
type Server struct {
	logger *log.Logger
	server *http.Server
}

// Создаём и настраиваем экземпляр сервера
// Принимаем логгер и возвращаем готовый к запуску сервер
func New(logger *log.Logger) *Server {
	// Создаём роутер
	mux := http.NewServeMux()

	// Регистрируем хендлеры
	mux.HandleFunc("/", handlers.RootHandler)
	mux.HandleFunc("/upload", handlers.UploadHandler)

	// Настраиваем HTTP-сервер
	httpServer := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	// Возвращаем экземпляр сервера
	return &Server{
		logger: logger,
		server: httpServer,
	}
}

// Запускаем сервер
// Возвращаем ошибку, если сервер не удалось запустить
func (s *Server) Start() error {
	s.logger.Printf("Server starting on %s", s.server.Addr)
	return s.server.ListenAndServe()
}

// Останавливаем сервер
// Принимаем контекст для управления таймаутом
func (s *Server) Shutdown(ctx context.Context) error {
	s.logger.Printf("Server shutting down")
	return s.server.Shutdown(ctx)
}
