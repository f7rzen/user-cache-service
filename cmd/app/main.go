package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/f7rzen/user-cache-service/internal/cache"
	"github.com/f7rzen/user-cache-service/internal/client"
	"github.com/f7rzen/user-cache-service/internal/config"
	"github.com/f7rzen/user-cache-service/internal/handler"
	"github.com/f7rzen/user-cache-service/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	httpClient := &http.Client{
		Timeout: cfg.HTTPClientTimeout,
	}

	userClient := client.NewUserClient(cfg.ExternalAPIURL, httpClient)
	userCache := cache.NewCache()
	userService := service.NewUserService(userClient, userCache, cfg.CacheTTL)
	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	api := router.Group("/api/v1")
	{
		users := api.Group("/users")
		{
			users.GET("", userHandler.GetUsers)
			users.GET("/:id", userHandler.GetUserByID)
		}
	}

	addr := ":" + cfg.AppPort

	server := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	go func() {
		logger.Info("server starting", "addr", addr)

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("failed to start server", "error", err)
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	logger.Info("server shutting down")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Error("server shutdown error", "error", err)
		os.Exit(1)
	}

	logger.Info("server stopped")
}
