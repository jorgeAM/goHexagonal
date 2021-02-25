package creating

import (
	"context"

	mooc "github.com/jorgeAM/goHexagonal/internal"
)

type CourseService struct {
	courseRepository mooc.CourseRepository
}

func NewCourseService(repository mooc.CourseRepository) CourseService {
	return CourseService{
		courseRepository: repository,
	}
}

func (s CourseService) CreateCourse(ctx context.Context, id, name, duration string) error {
	course, err := mooc.NewCourse(id, name, duration)

	if err != nil {
		return err
	}

	return s.courseRepository.Save(ctx, *course)
}
