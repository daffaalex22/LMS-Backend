package requests

import (
	"backend/business/course"
	"backend/business/student"
	"context"
	"time"
)

type Domain struct {
	Id        uint
	StudentId uint
	CourseId  uint
	TypeId    uint
	Status    string
	Message   string
	CreateAt  time.Time
	UpdateAt  time.Time
	Student   student.Domain
	Course    course.Domain
	// Type
}

type RequestsUseCaseInterface interface {
	RequestsGetAll(ctx context.Context) ([]Domain, error)
	RequestsAdd(ctx context.Context, domain Domain) (Domain, error)
	RequestsUpdate(ctx context.Context, domain Domain, id uint) (Domain, error)
	RequestsGetByCourseId(ctx context.Context, courseId uint) ([]Domain, error)
}

type RequestsRepoInterface interface {
	RequestsGetAll(ctx context.Context) ([]Domain, error)
	RequestsAdd(ctx context.Context, domain Domain) (Domain, error)
	RequestsUpdate(ctx context.Context, domain Domain, id uint) (Domain, error)
	RequestsGetByCourseId(ctx context.Context, courseId uint) ([]Domain, error)
	CheckStudent(ctx context.Context, id uint) (student.Domain, error)
	CheckCourse(ctx context.Context, id uint) (course.Domain, error)
}
