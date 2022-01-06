package difficulties

import (
	"context"
	"time"
)

type difficultiesUsecase struct {
	repo           Repository
	contextTimeout time.Duration
}

func NewDifficultyUsecase(timeout time.Duration, cr Repository) Usecase {
	return &difficultiesUsecase{
		contextTimeout: timeout,
		repo:           cr,
	}
}

func (cu *difficultiesUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	resp, err := cu.repo.GetAll(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return resp, nil
}
