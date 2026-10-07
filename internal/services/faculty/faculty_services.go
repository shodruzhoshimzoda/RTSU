package faculty

import (
	"context"
	"rtsu-students/internal/domain"
)


// FacultyRepo - интерфейс репозитория для работы с факультетами. 
type FacultyRepo interface {
	 GetFaculties(ctx context.Context) ([]*domain.Faculty, error)
}

type FacultyService struct {
	facultyRepo FacultyRepo
}

func NewFacultyService(facultyRepo FacultyRepo) *FacultyService {
	return &FacultyService{
		facultyRepo: facultyRepo,
	}
}

func (s *FacultyService) GetFaculties(ctx context.Context) ([]*domain.Faculty, error) {
	return s.facultyRepo.GetFaculties(ctx)
}
