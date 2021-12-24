package modules

import (
	"context"
	"time"
)

type ModulesUseCase struct {
	//repo
	repo ModulesRepoInterface
	ctx  time.Duration
}

func NewUseCase(mdsRepo ModulesRepoInterface, contextTimeout time.Duration) ModulesUseCaseInterface {
	return &ModulesUseCase{
		repo: mdsRepo,
		ctx:  contextTimeout,
	}
}

func (usecase *ModulesUseCase) ModulesGetAll(ctx context.Context) ([]Domain, error) {
	return usecase.repo.ModulesGetAll(ctx)
}
