package student

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	Id              uint
	Name            string
	Password        string
	ConfirmPassword string
	Email           string
	Avatar          string
	Phone           int
	Address         string
	Token           string
	CreateAt        time.Time
	UpdateAt        time.Time
	DeleteAt        gorm.DeletedAt `gorm:"index"`
}

type StudentUseCaseInterface interface {
	Login(domain Domain, ctx context.Context) (Domain, error)
	Register(domain *Domain, ctx context.Context) (Domain, error)
	GetProfile(ctx context.Context, id uint) (Domain, error)
	StudentUpdate(ctx context.Context, domain Domain, id uint) (Domain, error)
}

type StudentRepoInterface interface {
	Login(domain Domain, ctx context.Context) (Domain, error)
	Register(domain *Domain, ctx context.Context) (Domain, error)
	GetProfile(ctx context.Context, id uint) (Domain, error)
	StudentUpdate(ctx context.Context, domain Domain, id uint) (Domain, error)
}
