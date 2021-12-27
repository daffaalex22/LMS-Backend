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
	dataStudent, result := usecase.repo.EnrollmentGetAll(ctx)
	if result != nil {
		return []Domain{}, result
	}
	return dataStudent, nil
}

func (usecase *EnrollmentUseCase) EnrollmentAdd(ctx context.Context, domain Domain) (Domain, error) {
	if domain.StudentId == 0 {
		return Domain{}, err.ErrStudentIdEmpty
	}
	if domain.CourseId == 0 {
		return Domain{}, err.ErrCourseIdEmpty
	}

	dataStudent, err1 := usecase.repo.CheckStudent(ctx, domain.StudentId)
	if err1 != nil {
		return Domain{}, err.ErrIdStudent
	}
	domain.Student = dataStudent

	dataCourse, err2 := usecase.repo.CheckCourse(ctx, domain.CourseId)
	if err2 != nil {
		return Domain{}, err.ErrIdCourse
	}
	domain.Course = dataCourse

	enroll, result := usecase.repo.EnrollmentAdd(ctx, domain)
	if result != nil {
		return Domain{}, result
	}
	return enroll, nil
}

func (usecase *EnrollmentUseCase) EnrollUpdate(ctx context.Context, domain Domain) (Domain, error) {
	// domain.StudentId = studentId
	// domain.CourseId = courseId
	if domain.StudentId == 0 {
		return Domain{}, err.ErrCourseIdEmpty
	}
	if domain.CourseId == 0 {
		return Domain{}, err.ErrCourseIdEmpty
	}
	if domain.Rating == 0 {
		return Domain{}, err.ErrRatingEmpty
	}
	if domain.Review == "" {
		return Domain{}, err.ErrReviewEmpty
	}
	dataStudent, err1 := usecase.repo.CheckStudent(ctx, domain.StudentId)
	if err1 != nil {
		return Domain{}, err.ErrIdStudent
	}
	domain.Student = dataStudent

	dataCourse, err2 := usecase.repo.CheckCourse(ctx, domain.CourseId)
	if err2 != nil {
		return Domain{}, err.ErrIdCourse
	}
	domain.Course = dataCourse

	enroll, result := usecase.repo.EnrollUpdate(ctx, domain, domain.StudentId, domain.CourseId)
	if result != nil {
		return Domain{}, result
	}
	return enroll, nil
}
