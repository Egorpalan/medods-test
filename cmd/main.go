package main

import (
	"net/http"
	"test-medods/internal/configs"
	"test-medods/internal/handlers"
	"test-medods/internal/repository"
	"test-medods/internal/service"
	"test-medods/pkg/logger"
)

func main() {
	// Загрузка конфигурации
	cfg, err := config.LoadConfig("configs/local.env")
	if err != nil {
		logger.ErrorLogger.Fatalf("Failed to load config: %v", err)
	}

	// Инициализация репозитория
	repo, err := repository.NewPostgresRepository(cfg.DatabaseURL)
	if err != nil {
		logger.ErrorLogger.Fatalf("Failed to initialize repository: %v", err)
	}

	// Инициализация сервисов
	authService := service.NewAuthService(repo)

	// Инициализация обработчиков
	authHandler := handlers.NewAuthHandler(authService)

	// Настройка маршрутов
	http.HandleFunc("/auth/token", authHandler.GenerateToken)
	http.HandleFunc("/auth/refresh", authHandler.RefreshToken)

	// Запуск HTTP сервера
	logger.InfoLogger.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		logger.ErrorLogger.Fatalf("Failed to start server: %v", err)
	}
}
