package main

import (
	"context"
	"os"
	"rtsu-students/internal/config"
	"rtsu-students/internal/delivery/handler"
	"rtsu-students/internal/repository/postgres"
	facultyrepository "rtsu-students/internal/repository/postgres/faculty"
	facultyservice "rtsu-students/internal/services/faculty"
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


	// Иницализация логгера
	router := gin.New()
	router.Use(gin.Recovery()) // для перехвата паник
	router.Use(ginLogger)	   // в качестве логгера булем использовать свой 

	router.GET("/health", healthCheck)	


	// Инициализация репозитория, сервиса и хендлера для работы с факультетами
	facultyRepo := facultyrepository.NewFacultyRepository(db)
	facultyService := facultyservice.NewFacultyService(&facultyRepo)
	facultyHandler := handler.NewFacultyHandler(log, *facultyService)


	router.GET("/faculties", facultyHandler.GetFaculties)
	


	if err := router.Run(cfg.HTTPServer.Host + ":" + strconv.Itoa(cfg.HTTPServer.Port)); err != nil {
		log.Error("Failed to start server", "error", err)
		os.Exit(1)
	}

}

