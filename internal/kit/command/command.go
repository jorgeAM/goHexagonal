package command

import "context"

type Type string

type Command interface {
	Type() Type
}

type Handler interface {
	Handle(ctx context.Context, cmd Command) error
}

type Bus interface {
	Dispatch(ctx context.Context, cmd Command) error
	Register(Type, Handler)
}
