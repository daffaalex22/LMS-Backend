package modules

import (
	"backend/business/course"
	"context"
	"time"
)

type Domain struct {
	Id       uint
	CourseId uint
	Title    string
	Order    int
	CreateAt time.Time
	UpdateAt time.Time
	Course   course.Domain
}

type ModulesUseCaseInterface interface {
	ModulesGetAll(ctx context.Context) ([]Domain, error)
	ModulesAdd(ctx context.Context, domain Domain) (Domain, error)
	ModulesUpdate(ctx context.Context, domain Domain, id uint) (Domain, error)
	ModulesDelete(ctx context.Context, id uint) error
	ModulesGetById(ctx context.Context, id uint) (Domain, error)
	ModulesGetByCourseId(ctx context.Context, courseId uint) ([]Domain, error)
}

type ModulesRepoInterface interface {
	ModulesGetAll(ctx context.Context) ([]Domain, error)
	ModulesAdd(ctx context.Context, domain Domain) (Domain, error)
	CheckCourse(ctx context.Context, id uint) (course.Domain, error)
	ModulesUpdate(ctx context.Context, domain Domain, id uint) (Domain, error)
	ModulesDelete(ctx context.Context, id uint) error
	ModulesGetById(ctx context.Context, id uint) (Domain, error)
	ModulesGetByCourseId(ctx context.Context, courseId uint) ([]Domain, error)
}
