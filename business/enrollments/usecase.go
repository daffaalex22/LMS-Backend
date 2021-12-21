package enrollments

import (
	"backend/helper/err"
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

func (usecase *EnrollmentUseCase) EnrollmentAdd(ctx context.Context, domain Domain) (Domain, error) {
	if domain.Student_Id == 0 {
		return Domain{}, err.ErrStudentIdEmpty
	}
	if domain.Course_Id == 0 {
		return Domain{}, err.ErrCourseIdEmpty
	}
	enroll, err1 := usecase.repo.EnrollmentAdd(ctx, domain)
	if err1 != nil {
		return Domain{}, err1
	}
	return enroll, nil
}
