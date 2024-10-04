package course_test

import (
	"Go-API-Tech-Challenge/api/resources/course"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

var courseColumns = []string{
	"id",
	"name",
}

func TestFetchCourses_Happy(t *testing.T) {
	// Set-up SQL database mocks
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("unexpected error creating a mock DB connection, error: %s", err.Error())
	}
	defer db.Close()

	mock.ExpectQuery("SELECT \\* FROM course").WillReturnRows(
		sqlmock.NewRows(courseColumns).AddRow(1, "CS 408 Human Augmentics"),
	)

	r := course.NewRepo(db)
	courses, err := r.FetchCourses()

	if err != nil {
		t.Errorf("expected nil error, received: %s", err.Error())
	}
	if len(courses) == 0 {
		t.Error("expected a list of courses, but received empty list")
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("query did not fulfill expectations: %s", err.Error())
	}
}

func TestFetchCourses_Sad(t *testing.T) {

}
