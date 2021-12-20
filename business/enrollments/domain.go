package enrollments

import (
	"context"
	"time"
)

type Domain struct {
	Student_Id uint
	Course_Id  uint
	Rating     int
	Review     string
	CreateAt   time.Time
	UpdateAt   time.Time
}

type EnrollmentsUseCaseInterface interface {
	EnrollmentGetAll(ctx context.Context) ([]Domain, error)
}

type EnrollmentsRepoInterface interface {
	EnrollmentGetAll(ctx context.Context) ([]Domain, error)
}
