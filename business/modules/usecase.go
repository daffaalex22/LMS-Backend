package modules

import (
	"backend/helper/err"
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

func (usecase *ModulesUseCase) ModulesAdd(ctx context.Context, domain Domain) (Domain, error) {
	if domain.CourseId == 0 {
		return Domain{}, err.ErrCourseIdEmpty
	}
	if domain.Title == "" {
		return Domain{}, err.ErrTitleEmpty
	}
	if domain.Order == 0 {
		return Domain{}, err.ErrOrderEmpty
	}

	dataCourse, err2 := usecase.repo.CheckCourse(ctx, domain.CourseId)
	if err2 != nil {
		return Domain{}, err.ErrIdCourse
	}
	domain.Course = dataCourse

	enroll, result := usecase.repo.ModulesAdd(ctx, domain)
	if result != nil {
		return Domain{}, result
	}
	return enroll, nil
}
