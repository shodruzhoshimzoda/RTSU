package logger

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lmittmann/tint"
)

const (
	EnvProduction  = "prod"
	EnvDevelopment = "dev"
	EnvLocal       = "local"
)

// SetupLogger возвразает логгер  для приложения с настройками в зависимости от окружения (env).
func SetupLogger(env string) *slog.Logger {
	options := &slog.HandlerOptions{Level: slog.LevelInfo}
	var handler slog.Handler

	if env == EnvProduction || env == EnvDevelopment {
		handler = slog.NewJSONHandler(os.Stdout, options)
	} else {
		if env == EnvLocal {
			options.Level = slog.LevelDebug
		}
		options.AddSource = env == EnvDevelopment
		handler = tint.NewTextHandler(os.Stdout, &tint.Options{
			Level:     options.Level,
			AddSource: options.AddSource,
		})
	}

	return slog.New(handler)
}

// SetupGinLogger возаращает мидлвея для логгера для  Http обработчиков
func SetupGinLogger(env string) gin.HandlerFunc {

	if env == EnvProduction || env == EnvDevelopment {

		// В ПРОДАКШЕНЕ: Выводим логи Gin в формате JSON
		return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
			logMap := map[string]any{
				"time":       param.TimeStamp.Format(time.RFC3339),
				"status":     param.StatusCode,
				"latency_ns": param.Latency.Nanoseconds(),
				"client_ip":  param.ClientIP,
				"method":     param.Method,
				"path":       param.Path,
				"error":      param.ErrorMessage,
			}

			data, _ := json.Marshal(logMap)
			return string(data) + "\n"
		})
	} else {

		// ЛОКАЛЬНО: Используем дефолтный читаемый и цветной формат Gin
		return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
			var statusColor, methodColor, resetColor string
			if param.IsOutputColor() {
				statusColor = param.StatusCodeColor()
				methodColor = param.MethodColor()
				resetColor = param.ResetColor()
			}

			return fmt.Sprintf("[GIN] %s |%s %3d %s| %12v | %15s |%s %-7s %s %#v\n%s",
				param.TimeStamp.Format("2006/01/02 - 15:04:05"),
				statusColor, param.StatusCode, resetColor,
				param.Latency,
				param.ClientIP,
				methodColor, param.Method, resetColor,
				param.Path,
				param.ErrorMessage,
			)
		})
	}
}
