package mock

import "Go-API-Tech-Challenge/api/resources/course"

type Repository struct {
	FetchCoursesFunc     func() ([]course.Course, error)
	FetchCourseByIDFunc  func(id int64) (course.Course, error)
	InsertCourseFunc     func(name string) (course.Course, error)
	UpdateCourseByIDFunc func(id int64, name string) error
	DeleteCourseByIDFunc func(id int64) error
}

func (r *Repository) FetchCourses() ([]course.Course, error) {
	return r.FetchCoursesFunc()
}

func (r *Repository) FetchCourseByID(id int64) (course.Course, error) {
	return r.FetchCourseByIDFunc(id)
}

func (r *Repository) InsertCourse(name string) (course.Course, error) {
	return r.InsertCourseFunc(name)
}

func (r *Repository) UpdateCourseByID(id int64, name string) error {
	return r.UpdateCourseByIDFunc(id, name)
}

func (r *Repository) DeleteCourseByID(id int64) error {
	return r.DeleteCourseByIDFunc(id)
}
