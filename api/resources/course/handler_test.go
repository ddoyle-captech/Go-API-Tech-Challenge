package course_test

import (
	"Go-API-Tech-Challenge/api/resources/course"
	"Go-API-Tech-Challenge/test/mock"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListCourses_Happy(t *testing.T) {
	tests := map[string]struct {
		courses  []course.Course
		expected int
	}{
		"if repo returns courses, return courses and 200 OK": {
			courses: []course.Course{
				{1, "Introduction to Quantum Physics"},
				{2, "Organic Chemistry"},
				{3, "Modern European History"},
			},
		},
		"if repo returns nothing, return empty list and 200 OK": {
			courses: []course.Course{},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			r := &mock.Repository{
				FetchCoursesFunc: func() ([]course.Course, error) {
					return test.courses, nil
				},
			}
			h := course.NewHandler(r)

			// Set-up mock HTTP communication
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/api/course", nil)

			// Execute the HTTP request
			h.ListCourses(w, req)

			if w.Result().StatusCode != http.StatusOK {
				t.Errorf("expected response status code: %d, received: %d", http.StatusOK, w.Result().StatusCode)
			}
		})
	}
}

func TestListCourses_Sad(t *testing.T) {
	r := &mock.Repository{
		FetchCoursesFunc: func() ([]course.Course, error) {
			return []course.Course{}, errors.New("unable to connect to database")
		},
	}

	h := course.NewHandler(r)

	// Set-up mock HTTP communication
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/course", nil)

	// Execute the HTTP request
	h.ListCourses(w, req)

	if w.Result().StatusCode != http.StatusInternalServerError {
		t.Errorf("expected response status code: %d, received: %d", http.StatusInternalServerError, w.Result().StatusCode)
	}
}
