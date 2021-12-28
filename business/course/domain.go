package course

import (
	"backend/business/categories"
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
	// Difficulties
}

type Usecase interface {
	Create(ctx context.Context, domain Domain) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	GetCourseById(ctx context.Context, id string) (Domain, error)
	Update(ctx context.Context, id string, domain Domain) (Domain, error)
	Delete(ctx context.Context, id string) (Domain, error)
}

type Repository interface {
	Create(ctx context.Context, domain Domain) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	Update(ctx context.Context, domain Domain) (Domain, error)
	GetCourseById(ctx context.Context, id uint) (Domain, error)
	Delete(ctx context.Context, id uint) error

	CheckTeacher(ctx context.Context, id uint) (teacher.Domain, error)
	CheckCategories(ctx context.Context, id uint) (categories.Domain, error)
}
