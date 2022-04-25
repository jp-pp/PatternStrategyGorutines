package repository

import (
	"context"
	"fmt"
	"time"
)

type Repository interface {
	QueryHandlerInterface
	ErrorQueryHandlerInterface
}

type QueryHandlerInterface interface {
	HandleQuery(ctx context.Context) (bool, error)
}

type QueryHandler struct {

}

func (h *QueryHandler) HandleQuery(ctx context.Context) (bool, error) {

	var seconds time.Duration = 50

	time.Sleep(seconds)

	fmt.Printf(
		"FirstHandle executed FirstQuery on %v seconds\n",
		seconds,
	)

	return true, nil
}

type ErrorQueryHandlerInterface interface {
	HandleErrorQuery(ctx context.Context) (bool, error)
}

type ErrorQueryHandler struct {

}

func (h *ErrorQueryHandler) HandleErrorQuery(ctx context.Context) (bool, error) {

	var (
		seconds time.Duration = 25
		err error
	)


	time.Sleep(25)

	err = fmt.Errorf(
		"FirstHandle finished executing SecondQuery with an error after %v seconds\n",
		seconds,
	)

	return false, err
}
