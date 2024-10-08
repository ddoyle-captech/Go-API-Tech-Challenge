package course_test

import (
	"Go-API-Tech-Challenge/api/resources/course"
	"Go-API-Tech-Challenge/test/mock"
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
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
		FetchCourseByIDFunc: func(id int64) (course.Course, error) {
			return course.Course{
				ID:   int64(1),
				Name: "My favorite class",
			}, nil
		},
	}
	h := course.NewHandler(r)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/course/1", nil)
	req = addURLParamToRequest(req, "id", "1")

	h.GetCourse(w, req)

	if w.Result().StatusCode != http.StatusOK {
		t.Errorf("expected response status code: %d, received: %d", http.StatusOK, w.Result().StatusCode)
	}
}

func TestGetCourse_Sad(t *testing.T) {
	tests := map[string]struct {
		courseID string
		err      error
		expected int
	}{
		"if course ID is invalid, return a 400 Bad Request": {
			courseID: "vdugavdyuwt9878",
			err:      nil,
			expected: http.StatusBadRequest,
		},
		"if no courses are found, return a 404 Not Found": {
			courseID: "1",
			err:      course.ErrCourseNotFound,
			expected: http.StatusNotFound,
		},
		"if repo returns an error, return a 500 Internal Server Error": {
			courseID: "1",
			err:      errors.New("database is missing"),
			expected: http.StatusInternalServerError,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			r := &mock.Repository{
				FetchCourseByIDFunc: func(id int64) (course.Course, error) {
					return course.Course{}, test.err
				},
			}
			h := course.NewHandler(r)

			w := httptest.NewRecorder()
			url := fmt.Sprintf("/api/course/%s", test.courseID)

			req := httptest.NewRequest(http.MethodGet, url, nil)
			req = addURLParamToRequest(req, "id", test.courseID)

			h.GetCourse(w, req)

			if w.Result().StatusCode != test.expected {
				t.Errorf("expected response status code: %d, received: %d", test.expected, w.Result().StatusCode)
			}
		})
	}
}

func TestCreateCourse_Happy(t *testing.T) {
	r := &mock.Repository{
		InsertCourseFunc: func(name string) (course.Course, error) {
			return course.Course{
				ID:   int64(1),
				Name: name,
			}, nil
		},
	}
	body := strings.NewReader("{\"name\": \"UI Programming\"}")
	h := course.NewHandler(r)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/course", body)

	h.CreateCourse(w, req)

	if w.Result().StatusCode != http.StatusOK {
		t.Errorf("expected response status code: %d, received: %d", http.StatusOK, w.Result().StatusCode)
	}
}

func TestCreateCourse_Sad(t *testing.T) {
	tests := map[string]struct {
		request  string
		err      error
		expected int
	}{
		"if course update request is invalid, return a 400 Bad Request": {
			request:  "dawdawdwadwdwadwad dwdwd12",
			err:      nil,
			expected: http.StatusBadRequest,
		},
		"if repo returns an error, return a 500 Internal Server Error": {
			request:  "{\"name\": \"UI Programming\"}",
			err:      errors.New("database is missing"),
			expected: http.StatusInternalServerError,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			r := &mock.Repository{
				InsertCourseFunc: func(name string) (course.Course, error) {
					return course.Course{
						ID:   int64(1),
						Name: name,
					}, test.err
				},
			}
			h := course.NewHandler(r)

			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/api/course", strings.NewReader(test.request))

			h.CreateCourse(w, req)
			if w.Result().StatusCode != test.expected {
				t.Errorf("expected response status code: %d, received: %d", test.expected, w.Result().StatusCode)
			}
		})
	}
}

func TestUpdateCourse_Happy(t *testing.T) {
	r := &mock.Repository{
		UpdateCourseByIDFunc: func(id int64, name string) error {
			return nil
		},
	}
	body := strings.NewReader("{\"name\": \"UI Programming\"}")
	h := course.NewHandler(r)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPut, "/api/course/1", body)
	req = addURLParamToRequest(req, "id", "1")

	h.UpdateCourse(w, req)

	if w.Result().StatusCode != http.StatusOK {
		t.Errorf("expected response status code: %d, received: %d", http.StatusOK, w.Result().StatusCode)
	}
}

func TestUpdateCourse_Sad(t *testing.T) {
	tests := map[string]struct {
		id       string
		request  string
		err      error
		expected int
	}{
		"if course ID is invalid, return a 400 Bad Request": {
			id:       "dnajdbahwbdhawdbhk",
			expected: http.StatusBadRequest,
		},
		"if course update request is invalid, return a 400 Bad Request": {
			id:       "1",
			request:  "dawdawdwadwdwadwad dwdwd12",
			expected: http.StatusBadRequest,
		},
		"if no courses are found, return a 404 Not Found": {
			id:       "1",
			request:  "{\"name\": \"UI Programming\"}",
			err:      course.ErrCourseNotFound,
			expected: http.StatusNotFound,
		},
		"if repo returns an error, return a 500 Internal Server Error": {
			id:       "1",
			request:  "{\"name\": \"UI Programming\"}",
			err:      errors.New("database is missing"),
			expected: http.StatusInternalServerError,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			r := &mock.Repository{
				UpdateCourseByIDFunc: func(id int64, name string) error {
					return test.err
				},
			}
			h := course.NewHandler(r)

			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPut, "/api/course/"+test.id, strings.NewReader(test.request))
			req = addURLParamToRequest(req, "id", test.id)

			h.UpdateCourse(w, req)
			if w.Result().StatusCode != test.expected {
				t.Errorf("expected response status code: %d, received: %d", test.expected, w.Result().StatusCode)
			}
		})
	}
}

// Because Chi is used to set/parse the URL params, we need to create a Chi context and manually add the URL param value
// when testing the handler directly. Typically, the router handles this for us.
func addURLParamToRequest(r *http.Request, key, value string) *http.Request {
	chiCtx := chi.NewRouteContext()
	r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, chiCtx))
	chiCtx.URLParams.Add(key, value)
	return r
}
