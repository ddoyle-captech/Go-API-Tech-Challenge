package course

import "database/sql"

type Repository interface {
	FetchCourses() []Course
}

type repository struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) FetchCourses() []Course {
	panic("not implemented")
}
