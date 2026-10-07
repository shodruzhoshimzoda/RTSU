package main

import (
	"context"
	"os"
	"rtsu-students/internal/config"
	"rtsu-students/internal/repository/postgres"
	"rtsu-students/pkg/logger"
	"strconv"

	"github.com/gin-gonic/gin"
)

func healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  "OK",
		"message": "Server is running",
	})
}

func main() {

	// Получение всех конфигураций приложения
	cfg := config.LoadConfig()

	log := logger.SetupLogger(cfg.Env)
	ginLogger := logger.SetupGinLogger(cfg.Env)

	ctx := context.Background()

	// Подключение к базе данных
	db, err := postgres.ConnectionToDB(ctx, cfg.GetDatabaseDSN())
	if err != nil {
		log.Error("Failed to connect to database", "error", err)
		os.Exit(1)
	}

	defer db.Close()

	log.Info("connection to database was successfully ")

	router := gin.New()

	router.Use(gin.Recovery()) // для перехвата паник
	router.Use(ginLogger)

	router.GET("/health", healthCheck)

	if err := router.Run(cfg.HTTPServer.Host + ":" + strconv.Itoa(cfg.HTTPServer.Port)); err != nil {
		log.Error("Failed to start server", "error", err)
		os.Exit(1)
	}

}
