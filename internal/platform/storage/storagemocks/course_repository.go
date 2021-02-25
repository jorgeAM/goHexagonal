package storagemocks

import (
	"context"

	"github.com/stretchr/testify/mock"

	mooc "github.com/jorgeAM/goHexagonal/internal"
)

type CourseRepository struct {
	mock.Mock
}

func (r *CourseRepository) Save(ctx context.Context, course mooc.Course) error {
	args := r.Called(ctx, course)

	var res error

	if rf, ok := args.Get(0).(func(context.Context, mooc.Course) error); ok {
		res = rf(ctx, course)
	} else {
		res = args.Error(0)
	}

	return res
}
