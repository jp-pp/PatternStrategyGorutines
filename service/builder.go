package service

import (
	"PatternStrategyGorutines/entity"
	"PatternStrategyGorutines/repository"
	"context"
	"fmt"
)


type BuilderInterface interface {
	Build(ctx context.Context) (response bool, err error)
}

type Builder struct {
	Repository repository.Repository
}

func (b *Builder) Build(ctx context.Context) (response bool, err error) {

	var (
		errChan chan error
		entityCount = 8
		entityList []entity.Interface
	)

	entityList = make([]entity.Interface, entityCount)

	entityContext, cancel := context.WithCancel(ctx)

	defer cancel()

	errChan = make(chan error)

	for i := 0; i < entityCount; i++ {

		if i == 0 {

			go func() {

				LOOP:

					count := len(entityList)

					for i := 0; i < len(entityList); i++ {

						if entityList[i].Error() != nil {
							errChan <- entityList[i].Error()
							break
						} else if entityList[i].IsDone() {
							count--
							if count == 0 {
								response = true
								break
							}
							continue
						} else {
							goto LOOP
						}
					}
			}()
		}
		entityList[i] = &entity.Entity{HandleFunc: b.Repository.HandleQuery}
		go entityList[i].Build(entityContext, cancel)
	}

	err = <-errChan

	if err == nil {
		 fmt.Println(err)
	}

	return response, err
}

type SecondBuilder struct {

}