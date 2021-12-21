package enrollments

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
}

type EnrollmentsRepoInterface interface {
	EnrollmentGetAll(ctx context.Context) ([]Domain, error)
	EnrollmentAdd(ctx context.Context, domain Domain) (Domain, error)
}
