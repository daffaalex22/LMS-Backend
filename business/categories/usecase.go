package categories

import (
	"context"
	"time"
)

type categoriesUsecase struct {
	repo           Repository
	contextTimeout time.Duration
}

func NewCategoryUsecase(timeout time.Duration, cr Repository) Usecase {
	return &categoriesUsecase{
		contextTimeout: timeout,
		repo:           cr,
	}
}

func (cu *categoriesUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	resp, err := cu.repo.GetAll(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return resp, nil
}
