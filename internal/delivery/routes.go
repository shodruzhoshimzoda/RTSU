package delivery

import (
	"net/http"
	"rtsu-students/internal/config"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	cfg config.Config
}

func NewHandler(cfg *config.Config) Handler {
	return Handler{
		cfg: *cfg,
	}
}

func (h *Handler) HealthCheck(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{
			"status": "ok",
			"env":    h.cfg.Env,
		},
	)
}
