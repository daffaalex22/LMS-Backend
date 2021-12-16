package course

import (
	"backend/business"
	"context"
	"time"
)

type courseUsecase struct {
	Repo           Repository
	contextTimeout time.Duration
}

func NewCourseUsecase(timeout time.Duration, repo Repository) Usecase {
	return &courseUsecase{
		contextTimeout: timeout,
		Repo:           repo,
	}
}

func (uc *courseUsecase) Create(ctx context.Context, domain Domain) (Domain, error) {
	if domain.Title == "" {
		return Domain{}, business.ErrTitleNotFound
	}

	course, err := uc.Repo.Create(ctx, domain)
	if err != nil {
		return Domain{}, err
	}
	return course, nil
}
