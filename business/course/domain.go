package course

import (
	"backend/business/categories"
	"backend/business/difficulties"
	"backend/business/student"
	"backend/business/teacher"
	"context"
	"time"
)

type Domain struct {
	Id           uint
	Title        string
	Thumbnail    string
	Description  string
	CategoryId   uint
	TeacherId    uint
	DifficultyId uint
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Category     categories.Domain
	Teacher      teacher.Domain
	Difficulty   difficulties.Domain
}

type CourseEnrollmentDomain struct {
	StudentId uint
	CourseId  uint
	Rating    int
	Review    string
	CreatedAt time.Time
	UpdatedAt time.Time
	Student   student.Domain
	Course    Domain
}

type Usecase interface {
	Create(ctx context.Context, domain Domain) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	GetCourseById(ctx context.Context, id string) (Domain, error)
	GetCourseByStudentId(ctx context.Context, studentId uint) ([]Domain, error)
	GetCourseByTeacherId(ctx context.Context, teacherId uint) ([]Domain, error)
	Update(ctx context.Context, id string, domain Domain) (Domain, error)
	Delete(ctx context.Context, id string) (Domain, error)
}

type Repository interface {
	Create(ctx context.Context, domain Domain) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	Update(ctx context.Context, domain Domain) (Domain, error)
	GetCourseById(ctx context.Context, id uint) (Domain, error)
	GetCourseByStudentId(ctx context.Context, courseIds []uint) ([]Domain, error)
	GetCourseByTeacherId(ctx context.Context, teacherId uint) ([]Domain, error)
	Delete(ctx context.Context, id uint) error

	CheckTeacher(ctx context.Context, id uint) (teacher.Domain, error)
	CheckCategories(ctx context.Context, id uint) (categories.Domain, error)
	GetEnrollmentsByStudentId(ctx context.Context, studentId uint) ([]CourseEnrollmentDomain, error)
}
