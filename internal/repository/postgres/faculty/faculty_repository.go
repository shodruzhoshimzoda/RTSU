package faculty

import (
	"context"
	"errors"
	"fmt"
	"rtsu-students/internal/domain"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type FacultyRepository struct {
	dbPool *pgxpool.Pool
}

func NewFacultyRepository(db *pgxpool.Pool) FacultyRepository {
	return FacultyRepository{
		dbPool: db,
	}

}

func (f *FacultyRepository) GetFaculties(ctx context.Context) ([]*domain.Faculty, error) {

	query := "SELECT uuid, name, description, created_at, updated_at FROM faculties"

	var faculties = make([]*domain.Faculty, 0)

	rows, err := f.dbPool.Query(ctx, query)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return faculties, nil

		}

		return nil, fmt.Errorf("get faculties: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var faculty domain.Faculty
		err := rows.Scan(
			&faculty.UUID,
			&faculty.Name,
			&faculty.Description,
			&faculty.CreatedAt,
			&faculty.UpdatedAt,

		)
		if err != nil {
			return nil, fmt.Errorf("get faculties: %w", err)

		}

		faculties = append(faculties, &faculty)

	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("get faculties: %w", err)
	}

	return faculties, nil

}
