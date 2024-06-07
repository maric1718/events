package main

import (
	"events_app/internal/adapter/config"
	"events_app/internal/adapter/handler/http"
	"events_app/internal/adapter/logger"
	"events_app/internal/adapter/storage/file_system/repository"
	"events_app/internal/adapter/storage/kafka"
	"events_app/internal/core/service"
	"fmt"

	"log/slog"
)

func main() {
	// Load environment variables
	config, err := config.New()
	if err != nil {
		slog.Error("Error loading environment variables", "error", err)
		return
	}

	// Set logger
	logger.Set(config.App)

	slog.Info("Starting the application", "app", config.App.Name, "env", config.App.Env)

	repository.StartCronjobs()

	marketRepo := repository.NewMarketRepository()
	marketService := service.NewMarketService(marketRepo)
	marketHandler := http.NewMarketHandler(marketService)

	eventRepo := repository.NewEventRepository()
	eventService := service.NewEventService(eventRepo)
	eventHandler := http.NewEventHandler(eventService)

	// Init market data
	if err = marketRepo.Init(); err != nil {
		slog.Error("Error initializing data", "error", err)
		return
	}

	// Init event data
	if err = eventRepo.Init(); err != nil {
		slog.Error("Error initializing data", "error", err)
		return
	}

	// Starting kafka
	kafka.StartKafka(config.Kafka.Brokers, config.Kafka.Topics)

	// Init router
	router, err := http.NewRouter(
		config.HTTP,
		*marketHandler,
		*eventHandler,
	)
	if err != nil {
		slog.Error("Error initializing router", "error", err)
		return
	}

	// Start server
	listenAddr := fmt.Sprintf("%s:%s", config.HTTP.URL, config.HTTP.Port)

	slog.Info("Starting the HTTP server on port", "listen_address", config.HTTP.Port)

	if err = router.Serve(listenAddr); err != nil {
		slog.Error("Error starting the HTTP server", "error", err)
		return
	}
}
