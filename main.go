package main

import (
	"PatternStrategyGorutines/repository"
	"PatternStrategyGorutines/service"
	"context"
	"fmt"
)

var (
	repo 		repository.Repository
	builder		service.BuilderInterface
)

func init() {
	repo = repository.Repository(&struct{
		repository.QueryHandlerInterface
		repository.ErrorQueryHandlerInterface
	}{
		&repository.QueryHandler{},
		&repository.ErrorQueryHandler{},
	})

	builder = &service.Builder{Repository: repo}
}

func main()  {
	var (
		response bool
		err error
	)
	ctx := context.Background()

	response, err = builder.Build(ctx)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response)
	}
}