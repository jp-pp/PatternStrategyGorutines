package entity

import (
	"context"
	"fmt"
)

type Interface interface {
	IsDone() bool
	Error() error
	Build(ctx context.Context, cancelFunc context.CancelFunc)
}

type Entity struct {
	isDone bool
	error error
	HandleFunc func(ctx context.Context) (bool, error)
}

func (e *Entity) IsDone() bool {
	return e.isDone
}

func (e *Entity) Error() error {
	return e.error
}

func (e *Entity) Build(ctx context.Context, cancelFunc context.CancelFunc) {

	select {
	case <-ctx.Done():
		fmt.Println("OneEntity return error")
	default:
		e.isDone, e.error = e.HandleFunc(ctx)

		if e.error != nil {
			cancelFunc()
		}
	}
}