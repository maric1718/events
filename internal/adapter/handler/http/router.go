package http

import (
	"events_app/internal/adapter/config"
	"log/slog"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
)

type Router struct {
	*gin.Engine
}

// Create new HTTP router
func NewRouter(
	config *config.HTTP,
	marketHandler MarketHandler,
	eventHandler EventHandler,
) (*Router, error) {

	ginConfig := cors.DefaultConfig()

	ginConfig.AllowOrigins = config.AllowedOrigins
	ginConfig.AllowMethods = config.AllowedHeaders

	router := gin.New()
	router.Use(sloggin.New(slog.Default()), gin.Recovery(), cors.New(ginConfig))

	v1 := router.Group("/v1")
	{
		market := v1.Group("/markets")
		{
			market.GET("", marketHandler.GetMarkets)
		}

		event := v1.Group("/events")
		{
			event.GET("", eventHandler.GetEvents)
		}
	}

	return &Router{
		router,
	}, nil
}

// Start the HTTP server
func (r *Router) Serve(listenAddr string) error {
	return r.Run(listenAddr)
}
