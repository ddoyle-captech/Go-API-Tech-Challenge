package course

type Repository interface {
	FetchCourses() []Course
}

type repository struct {
}

func NewRepo() Repository {
	return &repository{}
}

func (r *repository) FetchCourses() []Course {
	panic("not implemented")
}
