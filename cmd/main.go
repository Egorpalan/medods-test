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
	cfg, err := config.LoadConfig("configs/local.env")
	if err != nil {
		logger.ErrorLogger.Fatalf("Failed to load config: %v", err)
	}

	repo, err := repository.NewPostgresRepository(cfg.DatabaseURL)
	if err != nil {
		logger.ErrorLogger.Fatalf("Failed to initialize repository: %v", err)
	}

	authService := service.NewAuthService(repo)

	authHandler := handlers.NewAuthHandler(authService)

	http.HandleFunc("/auth/token", authHandler.GenerateToken)
	http.HandleFunc("/auth/refresh", authHandler.RefreshToken)

	logger.InfoLogger.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		logger.ErrorLogger.Fatalf("Failed to start server: %v", err)
	}
}
