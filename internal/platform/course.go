package mooc

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

var ErrInvalidCourseID = errors.New("invalid course ID")

type CourseID struct {
	value string
}

func NewCourseID(value string) (*CourseID, error) {
	v, err := uuid.Parse(value)

	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrInvalidCourseID, value)
	}

	return &CourseID{value: v.String()}, nil
}

func (id *CourseID) String() string {
	return id.value
}

var ErrEmptyCourseName = errors.New("the field Course Name can not be empty")

type CourseName struct {
	value string
}

func NewCourseName(value string) (*CourseName, error) {
	if value == "" {
		return nil, ErrEmptyCourseName
	}

	return &CourseName{value: value}, nil
}

func (name *CourseName) String() string {
	return name.value
}

var ErrEmptyDuration = errors.New("the field Duration can not be empty")

type CourseDuration struct {
	value string
}

func NewCourseDuration(value string) (*CourseDuration, error) {
	if value == "" {
		return nil, ErrEmptyDuration
	}

	return &CourseDuration{value: value}, nil
}

func (duration *CourseDuration) String() string {
	return duration.value
}

type CourseRepository interface {
	Save(ctx context.Context, course Course) error
}

type Course struct {
	id       CourseID
	name     CourseName
	duration CourseDuration
}

func NewCourse(id, name, duration string) (*Course, error) {
	idVO, err := NewCourseID(id)

	if err != nil {
		return nil, err
	}

	nameVO, err := NewCourseName(name)

	if err != nil {
		return nil, err
	}

	durationVO, err := NewCourseDuration(duration)

	return &Course{
		id:       *idVO,
		name:     *nameVO,
		duration: *durationVO,
	}, nil
}

func (c Course) ID() *CourseID {
	return &c.id
}

func (c Course) Name() *CourseName {
	return &c.name
}

func (c Course) Duration() *CourseDuration {
	return &c.duration
}
