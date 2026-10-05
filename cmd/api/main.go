package main

import (
	"fmt"
	"os"
	"rtsu-students/internal/config"
	"rtsu-students/internal/delivery"
	"rtsu-students/pkg/logger"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {

	// Получение всех конфигураций приложения
	cfg := config.LoadConfig()

	// Инициализация логгера
	log := logger.SetupLogger(cfg.Env)

	router := gin.New()

	router.Use(gin.Recovery()) // для перехвата паник
	router.Use(log)

	h := delivery.NewHandler(cfg)
	router.GET("/health", h.HealthCheck)

	if err := router.Run(cfg.HTTPServer.Host + ":" + strconv.Itoa(cfg.HTTPServer.Port)); err != nil {
		fmt.Printf("Error while starting server: %v", err)
		os.Exit(1)
	}

}
