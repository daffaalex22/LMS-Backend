package types

import (
	"context"
	"time"
)

type Domain struct {
	Id        uint
	Title     string
	CreatedAt time.Time
	UpdateAt  time.Time
}

type Usecase interface {
	GetAll(ctx context.Context) ([]Domain, error)
}

type Repository interface {
	GetAll(ctx context.Context) ([]Domain, error)
}
