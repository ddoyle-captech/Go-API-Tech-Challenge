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
		sqlmock.NewRows(courseColumns).
			AddRow(1, "CS 408 Human Augmentics").
			AddRow(2, "BIO 101 Intro to Biology").
			AddRow(3, "MATH 401 Differential Equations"),
	)

	r := course.NewRepo(db)
	courses, err := r.FetchCourses()

	if err != nil {
		t.Errorf("expected nil error, received: %s", err.Error())
	}
	if len(courses) != 3 {
		t.Errorf("expected a list of %d courses, but received %d", 3, len(courses))
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("query did not fulfill expectations: %s", err.Error())
	}
}

func TestFetchCourseByID_Happy(t *testing.T) {
	expected := course.Course{
		ID:   1,
		Name: "CS 408 Human Augmentics",
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("unexpected error creating a mock DB connection, error: %s", err.Error())
	}
	defer db.Close()

	mock.ExpectQuery("SELECT \\* FROM course WHERE id = \\$1").
		WithArgs(expected.ID).
		WillReturnRows(
			sqlmock.NewRows(courseColumns).AddRow(expected.ID, expected.Name),
		)

	r := course.NewRepo(db)
	result, err := r.FetchCourseByID(1)

	if err != nil {
		t.Errorf("expected nil error, received: %s", err.Error())
	}
	if result.ID != expected.ID {
		t.Errorf("expected course ID: %d, received: %d", expected.ID, result.ID)
	}
	if result.Name != expected.Name {
		t.Errorf("expected course name: %s, received: %s", expected.Name, result.Name)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("query did not fulfill expectations: %s", err.Error())
	}
}

func TestUpdateCourse_Happy(t *testing.T) {
	courseID := 1
	courseName := "UI Design"

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("unexpected error creating a mock DB connection, error: %s", err.Error())
	}
	defer db.Close()

	mock.ExpectPrepare("UPDATE course SET name = \\$1 WHERE id = \\$2").
		ExpectExec().
		WithArgs(courseName, courseID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	r := course.NewRepo(db)

	err = r.UpdateCourse(courseID, courseName)
	if err != nil {
		t.Errorf("expected a nil error, received: %+v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("query did not fulfill expectations: %s", err.Error())
	}
}
