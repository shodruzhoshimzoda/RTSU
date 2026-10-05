package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

// Config - единая структура хряняшая все осноные конфиги приложения
type Config struct {
	Env        string `yaml:"env"`
	HTTPServer `yaml:"http-server"`
}

// HTTPServer - представляет из себя структуру которая хранит основные конфиги для сервера
type HTTPServer struct {
	CORSOrigin string `yaml:"cors-origin"`
	Host       string `yaml:"host" env-default:"localhost"`
	Port       int    `yaml:"port" env-default:"8080"`
}

// LoadConfig загружает конфигурацию из файла, путь к которому указан в переменной окружения CONFIG_FILE_PATH.
func LoadConfig() *Config {

	err := godotenv.Load() //

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	configFilePath := os.Getenv("CONFIG_FILE_PATH")

	if configFilePath == "" {
		log.Fatal("CONFIG_FILE_PATH environment variable is not set")
	}

	// Проверка наличия файла конфигурации по указанному пути
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		log.Fatalf("Configuration file does not exist at path: %s", configFilePath)
	}

	var config Config

	// Чтение и парсинг файла конфигурации

	if err := cleanenv.ReadConfig(configFilePath, &config); err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	return &config

}
