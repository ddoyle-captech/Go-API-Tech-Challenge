package course_test

import (
	"Go-API-Tech-Challenge/api/resources/course"
	"Go-API-Tech-Challenge/test/mock"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
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

func TestGetCourse_Happy(t *testing.T) {
	r := &mock.Repository{
		FetchCourseByIDFunc: func(id int) (course.Course, error) {
			return course.Course{
				ID:   1,
				Name: "My favorite class",
			}, nil
		},
	}

	h := course.NewHandler(r)

	// Set up the chi router and register the handler
	router := chi.NewRouter()
	router.Get("/api/course/{id}", http.HandlerFunc(h.GetCourse))

	// Set-up mock HTTP communication
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "1")

	ctx := context.WithValue(context.Background(), chi.RouteCtxKey, rctx)

	w := httptest.NewRecorder()
	req := httptest.NewRequestWithContext(ctx, http.MethodGet, "/api/course/1", nil)

	// Execute the HTTP request
	router.ServeHTTP(w, req)

	if w.Result().StatusCode != http.StatusOK {
		t.Errorf("expected response status code: %d, received: %d", http.StatusOK, w.Result().StatusCode)
	}
}
