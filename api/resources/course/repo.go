package course

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
)

var ErrCourseNotFound = errors.New("course not found")

// Repository is responsible for communicating with the API's database.
type Repository interface {
	FetchCourses() ([]Course, error)
	FetchCourseByID(id int) (Course, error)
	InsertCourse(name string) error
	UpdateCourseByID(id int, name string) error
}

type repository struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) FetchCourses() ([]Course, error) {
	rows, err := r.db.Query(`SELECT * FROM course`)

	// PGX returns an error if a query returns no results. We'll return
	// an empty error + no error
	if err != nil && errors.Is(err, pgx.ErrNoRows) {
		return []Course{}, nil
	}
	if err != nil {
		return []Course{}, fmt.Errorf("unable to complete query, error: %w", err)
	}

	// Map results to Course structs
	return mapRowsToStruct(rows)
}

func (r *repository) FetchCourseByID(id int) (Course, error) {
	rows, err := r.db.Query(`SELECT * FROM course WHERE id = $1`, id)

	if err != nil && errors.Is(err, pgx.ErrNoRows) {
		return Course{}, ErrCourseNotFound
	}
	if err != nil {
		return Course{}, fmt.Errorf("unable to complete query, error: %w", err)
	}

	courses, err := mapRowsToStruct(rows)
	if err != nil {
		return Course{}, err
	}

	// We expect only 1 record returned whe querying by ID
	return courses[0], err
}

func (r *repository) InsertCourse(name string) error {
	statement, err := r.db.Prepare(`INSERT INTO course VALUES($1)`)
	if err != nil {
		return fmt.Errorf("unable to prepare course insert, error: %w", err)
	}
	defer statement.Close()

	_, err = statement.Exec(name)
	if err != nil {
		return fmt.Errorf("course insert failed, error: %w", err)
	}
	return nil
}

func (r *repository) UpdateCourseByID(id int, name string) error {
	statement, err := r.db.Prepare(`UPDATE course SET name = $1 WHERE id = $2`)
	if err != nil {
		return fmt.Errorf("unable to prepare course update, error: %w", err)
	}
	defer statement.Close()

	result, err := statement.Exec(name, id)
	if err != nil {
		return fmt.Errorf("course update failed, error: %w", err)
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("unable to fetch how many courses were updated, error: %w", err)
	}
	if affected == 0 {
		return ErrCourseNotFound
	}
	return nil
}

func mapRowsToStruct(rows *sql.Rows) ([]Course, error) {
	courses := []Course{}
	for rows.Next() {
		var id int
		var name string

		if err := rows.Scan(&id, &name); err != nil {
			return []Course{}, fmt.Errorf("error mapping Course row to struct, error: %s", err.Error())
		}

		c := Course{
			ID:   id,
			Name: name,
		}
		courses = append(courses, c)
	}
	return courses, nil
}
