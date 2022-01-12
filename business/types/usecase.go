package types

import (
	"context"
	"time"
)

type typesUsecase struct {
	repo           Repository
	contextTimeout time.Duration
}

func NewTypeUsecase(timeout time.Duration, cr Repository) Usecase {
	return &typesUsecase{
		contextTimeout: timeout,
		repo:           cr,
	}
}

func (cu *typesUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	resp, err := cu.repo.GetAll(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return resp, nil
}
