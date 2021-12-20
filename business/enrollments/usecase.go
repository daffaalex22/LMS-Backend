package enrollments

import (
	"context"
	"time"
)

type EnrollmentUseCase struct {
	//repo
	repo EnrollmentsRepoInterface
	ctx  time.Duration
}

func NewUseCase(elmRepo EnrollmentsRepoInterface, contextTimeout time.Duration) EnrollmentsUseCaseInterface {
	return &EnrollmentUseCase{
		repo: elmRepo,
		ctx:  contextTimeout,
	}
}

func (usecase *EnrollmentUseCase) EnrollmentGetAll(ctx context.Context) ([]Domain, error) {
	return usecase.repo.EnrollmentGetAll(ctx)
}
