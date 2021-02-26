package inmemory

import (
	"context"

	"github.com/jorgeAM/goHexagonal/internal/kit/command"
)

type commandBus struct {
	handlers map[command.Type]command.Handler
}

func NewCommadBus() command.Bus {
	return commandBus{
		handlers: make(map[command.Type]command.Handler),
	}
}

func (c commandBus) Dispatch(ctx context.Context, cmd command.Command) error {
	cmdType := cmd.Type()

	handler, ok := c.handlers[cmdType]

	if !ok {
		return nil
	}

	return handler.Handle(ctx, cmd)
}

func (c commandBus) Register(cmdType command.Type, handler command.Handler) {
	c.handlers[cmdType] = handler
}
