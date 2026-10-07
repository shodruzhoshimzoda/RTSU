package handler

import (
	"log/slog"
	"net/http"
	"rtsu-students/internal/services/faculty"

	"github.com/gin-gonic/gin"
)





type FacultyHandler struct {
	log 	*slog.Logger
	service faculty.FacultyService
}


func NewFacultyHandler(log *slog.Logger, service faculty.FacultyService) FacultyHandler {
	return FacultyHandler{
		log: log,
		service: service,
	}
}


// GetFaculties - получение списка факультетов
func (f *FacultyHandler) GetFaculties(c *gin.Context) {

	faculties, err := f.service.GetFaculties(c.Request.Context())
	if err != nil {
		f.log.Error("GetFaculties: ", slog.Attr{Key: "error", Value: slog.StringValue(err.Error())})

		c.JSON(http.StatusInternalServerError, gin.H{
			"error":"failed to fetch faculties",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":map[string]any{
			"faculties": faculties,
		},
	})

	
}