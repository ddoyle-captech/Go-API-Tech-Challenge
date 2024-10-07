package mock

import "Go-API-Tech-Challenge/api/resources/course"

type Repository struct {
	FetchCoursesFunc    func() ([]course.Course, error)
	FetchCourseByIDFunc func(id int) (course.Course, error)
	UpdateCourseFunc    func(id int, name string) error
}

func (r *Repository) FetchCourses() ([]course.Course, error) {
	return r.FetchCoursesFunc()
}

func (r *Repository) FetchCourseByID(id int) (course.Course, error) {
	return r.FetchCourseByIDFunc(id)
}

func (r *Repository) UpdateCourse(id int, name string) error {
	return r.UpdateCourseFunc(id, name)
}
