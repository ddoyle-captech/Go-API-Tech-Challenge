package mock

import "Go-API-Tech-Challenge/api/resources/course"

type Repository struct {
	FetchCoursesFunc     func() ([]course.Course, error)
	FetchCourseByIDFunc  func(id int) (course.Course, error)
	InsertCourseFunc     func(name string) error
	UpdateCourseByIDFunc func(id int, name string) error
}

func (r *Repository) FetchCourses() ([]course.Course, error) {
	return r.FetchCoursesFunc()
}

func (r *Repository) FetchCourseByID(id int) (course.Course, error) {
	return r.FetchCourseByIDFunc(id)
}

func (r *Repository) InsertCourse(name string) error {
	return r.InsertCourseFunc(name)
}

func (r *Repository) UpdateCourseByID(id int, name string) error {
	return r.UpdateCourseByIDFunc(id, name)
}
