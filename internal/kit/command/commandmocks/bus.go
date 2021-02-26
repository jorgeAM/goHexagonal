package commandmocks

import (
	"context"

	"github.com/jorgeAM/goHexagonal/internal/kit/command"
	"github.com/stretchr/testify/mock"
)

type Bus struct {
	mock.Mock
}

func (b *Bus) Dispatch(ctx context.Context, cmd command.Command) error {
	args := b.Called(ctx, cmd)

	return args.Error(0)
}

func (b *Bus) Register(cmdType command.Type, handler command.Handler) {
	b.Called(cmdType, handler)
}
