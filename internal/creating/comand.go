package creating

import (
	"context"
	"errors"

	"github.com/jorgeAM/goHexagonal/internal/kit/command"
)

const CourseCommandType command.Type = "command.creating.cours"

type courseCommand struct {
	id       string
	name     string
	duration string
}

func NewCourseCommand(id, name, duration string) command.Command {
	return courseCommand{
		id:       id,
		name:     name,
		duration: duration,
	}
}

func (c courseCommand) Type() command.Type {
	return CourseCommandType
}

type CourseCommandHandler struct {
	service CourseService
}

func NewCourseCommandHandler(service CourseService) command.Handler {
	return CourseCommandHandler{
		service: service,
	}
}

func (c CourseCommandHandler) Handle(ctx context.Context, cmd command.Command) error {
	createCourseCmd, ok := cmd.(courseCommand)

	if !ok {
		return errors.New("unexpected command")
	}

	return c.service.CreateCourse(ctx, createCourseCmd.id, createCourseCmd.name, createCourseCmd.duration)
}
