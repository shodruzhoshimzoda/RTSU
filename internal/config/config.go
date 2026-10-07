package config

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

// Config - единая структура хряняшая все осноные конфиги приложения
type Config struct {
	Env         string `yaml:"env"`
	HTTPServer  `yaml:"http-server"`
	DatabaseDSN `yaml:"db-conn"`
}

// DatabaseDSN - структура которая хранит конфиги для подключения к базе данных
type DatabaseDSN struct {
	Host     string `yaml:"host" env-default:"localhost"`
	Port     int    `yaml:"port" env-default:"5432"`
	User     string `yaml:"user" env-default:"postgres"`
	Database string `yaml:"database" env-default:"shop"`
	SSLMode  string `yaml:"sslmode" env-default:"disable"`
	Password string `yaml:"-" env-default:"DB_USER_PASSWORD"`
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

	dbPassword := os.Getenv("DB_USER_PASSWORD")

	if dbPassword == "" {
		log.Fatal("DB_USER_PASSWORD environment variable is not set")
	}

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

	config.Password = dbPassword // Set the database password from the environment variable

	return &config

}

func (db *DatabaseDSN) GetDatabaseDSN() string {
	dbDSN := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(db.User, db.Password), // coding spec symbols
		Host:   fmt.Sprintf("%s:%d", db.Host, db.Port),
		Path:   db.Database,
	}
	query := dbDSN.Query()
	query.Set("sslmode", db.SSLMode)
	dbDSN.RawQuery = query.Encode()
	return dbDSN.String()
}
