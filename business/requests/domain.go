package requests

import (
	"backend/business/course"
	"backend/business/student"
	"context"
	"time"
)

type Domain struct {
	StudentId uint
	CourseId  uint
	Rating    int
	Review    string
	CreateAt  time.Time
	UpdateAt  time.Time
	Student   student.Domain
	Course    course.Domain
}

type EnrollmentsUseCaseInterface interface {
	EnrollmentGetAll(ctx context.Context) ([]Domain, error)
	EnrollmentAdd(ctx context.Context, domain Domain) (Domain, error)
	EnrollUpdate(ctx context.Context, domain Domain) (Domain, error)
	EnrollGetByCourseId(ctx context.Context, courseId uint) ([]Domain, error)
}

type EnrollmentsRepoInterface interface {
	EnrollmentGetAll(ctx context.Context) ([]Domain, error)
	EnrollmentAdd(ctx context.Context, domain Domain) (Domain, error)
	EnrollUpdate(ctx context.Context, domain Domain, studentId uint, courseId uint) (Domain, error)
	EnrollGetByCourseId(ctx context.Context, courseId uint) ([]Domain, error)
	CheckStudent(ctx context.Context, id uint) (student.Domain, error)
	CheckCourse(ctx context.Context, id uint) (course.Domain, error)
}
