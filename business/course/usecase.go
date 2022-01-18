package course

import (
	"backend/helper/err"
	"backend/helper/konversi"
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
		return Domain{}, err.ErrTitleEmpty
	}
	if domain.CategoryId == 0 {
		return Domain{}, err.ErrCategoryIdEmpty
	}
	if domain.TeacherId == 0 {
		return Domain{}, err.ErrTeacherIdEmpty
	}

	course, err := uc.Repo.Create(ctx, domain)
	if err != nil {
		return Domain{}, err
	}
	return course, nil
}

func (uc *courseUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	course, err := uc.Repo.GetAll(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return course, nil
}

func (uc *courseUsecase) GetCourseById(ctx context.Context, id string) (Domain, error) {
	if id == "" {
		return Domain{}, err.ErrIdEmpty
	}

	uintId, err := konversi.StringToUint(id)
	if err != nil {
		return Domain{}, err
	}

	course, err := uc.Repo.GetCourseById(ctx, uintId)
	if err != nil {
		return Domain{}, err
	}

	course.Category, err = uc.Repo.CheckCategories(ctx, course.CategoryId)
	if err != nil {
		return Domain{}, err
	}

	course.Teacher, err = uc.Repo.CheckTeacher(ctx, course.TeacherId)
	if err != nil {
		return Domain{}, err
	}

	course.Difficulty, err = uc.Repo.CheckDifficulties(ctx, course.DifficultyId)
	if err != nil {
		return Domain{}, err
	}

	return course, nil
}

func (uc *courseUsecase) Update(ctx context.Context, id string, domain Domain) (Domain, error) {
	if id == "" {
		return Domain{}, err.ErrIdEmpty
	}
	if domain.Title == "" {
		return Domain{}, err.ErrTitleEmpty
	}
	if domain.CategoryId == 0 {
		return Domain{}, err.ErrCategoryIdEmpty
	}
	if domain.TeacherId == 0 {
		return Domain{}, err.ErrTeacherIdEmpty
	}

	uintId, err := konversi.StringToUint(id)
	if err != nil {
		return Domain{}, err
	}
	domain.Id = uintId

	//check course
	_, err1 := uc.Repo.GetCourseById(ctx, domain.Id)
	if err1 != nil {
		return Domain{}, err1
	}

	//check categories
	dataCategories, err := uc.Repo.CheckCategories(ctx, domain.CategoryId)
	if err != nil {
		return Domain{}, err
	}
	//masukin data categories di course
	domain.Category = dataCategories

	//check teacher
	dataTeacher, err := uc.Repo.CheckTeacher(ctx, domain.TeacherId)
	if err != nil {
		return Domain{}, err
	}
	//masukin data ke teacher di course
	domain.Teacher = dataTeacher

	course, err := uc.Repo.Update(ctx, domain)
	if err != nil {
		return Domain{}, err
	}
	return course, nil
}

func (uc *courseUsecase) Delete(ctx context.Context, id string) (Domain, error) {
	if id == "" {
		return Domain{}, err.ErrIdEmpty
	}

	uintId, err := konversi.StringToUint(id)
	if err != nil {
		return Domain{}, err
	}

	dataCourse, err := uc.Repo.GetCourseById(ctx, uintId)
	if err != nil {
		return Domain{}, err
	}

	errDeleting := uc.Repo.Delete(ctx, uintId)
	if errDeleting != nil {
		return Domain{}, err
	}
	return dataCourse, nil
}

func (uc *courseUsecase) GetCourseByStudentId(ctx context.Context, studentId uint) ([]Domain, error) {
	var courseIds []uint

	enrollments, err := uc.Repo.GetEnrollmentsByStudentId(ctx, studentId)
	if err != nil {
		return []Domain{}, err
	}

	for i := 0; i < len(enrollments); i++ {
		courseIds = append(courseIds, enrollments[i].CourseId)
	}

	course, err := uc.Repo.GetCoursesByCourseIds(ctx, courseIds)
	if err != nil {
		return []Domain{}, err
	}
	return course, nil
}

func (uc *courseUsecase) GetCourseByTeacherId(ctx context.Context, teacherId uint) ([]Domain, error) {

	course, err := uc.Repo.GetCourseByTeacherId(ctx, teacherId)
	if err != nil {
		return []Domain{}, err
	}
	return course, nil
}
