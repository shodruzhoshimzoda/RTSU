package logger

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	EnvProduction  = "production"
	EnvDevelopment = "development"
	EnvLocal       = "local"
)

// SetupLogger возвращает gin.HandlerFunc в зависимости от env
func SetupLogger(env string) gin.HandlerFunc {

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
