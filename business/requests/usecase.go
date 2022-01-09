package requests

import (
	"backend/helper/err"
	"context"
	"fmt"
	"time"
)

type RequestsUseCase struct {
	//repo
	repo RequestsRepoInterface
	ctx  time.Duration
}

func NewUseCase(elmRepo RequestsRepoInterface, contextTimeout time.Duration) RequestsUseCaseInterface {
	return &RequestsUseCase{
		repo: elmRepo,
		ctx:  contextTimeout,
	}
}

func (usecase *RequestsUseCase) RequestsGetAll(ctx context.Context) ([]Domain, error) {
	dataStudent, result := usecase.repo.RequestsGetAll(ctx)
	if result != nil {
		return []Domain{}, result
	}
	return dataStudent, nil
}

func (usecase *RequestsUseCase) RequestsAdd(ctx context.Context, domain Domain) (Domain, error) {
	if domain.StudentId == 0 {
		return Domain{}, err.ErrStudentIdEmpty
	}
	if domain.CourseId == 0 {
		return Domain{}, err.ErrCourseIdEmpty
	}
	if domain.Status == "" {
		fmt.Println("Status Dianggap Kosong")
		return Domain{}, err.ErrStatusEmpty
	}
	if domain.Message == "" {
		return Domain{}, err.ErrMessageEmpty
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

	request, result := usecase.repo.RequestsAdd(ctx, domain)
	if result != nil {
		return Domain{}, result
	}
	return request, nil
}

func (usecase *RequestsUseCase) RequestsUpdate(ctx context.Context, domain Domain, id uint) (Domain, error) {

	if domain.Status == "" {
		return Domain{}, err.ErrStatusEmpty
	}

	request, result := usecase.repo.RequestsUpdate(ctx, domain, id)
	if result != nil {
		return Domain{}, result
	}
	return request, nil
}

func (usecase *RequestsUseCase) RequestsGetByCourseId(ctx context.Context, courseId uint) ([]Domain, error) {
	modules, err := usecase.repo.RequestsGetByCourseId(ctx, courseId)
	if err != nil {
		return []Domain{}, err
	}
	return modules, nil
}
