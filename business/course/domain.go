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
}

type Repository interface {
	Create(ctx context.Context, domain Domain) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	GetCourseById(ctx context.Context, id uint) (Domain, error)
}
