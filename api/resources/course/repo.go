package course

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Repository interface {
	FetchCourses() ([]Course, error)
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

	rows, err := r.db.Query(`SELECT * FROM courses`)

	// PGX returns an error if a query returns no results. We'll return
	// an empty error + no error
	if err != nil && errors.Is(err, pgx.ErrNoRows) {
		return []Course{}, nil
	}
	if err != nil {
		return []Course{}, fmt.Errorf("unable to complete query, error: %s", err.Error())
	}

	// Map results to Course structs
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
