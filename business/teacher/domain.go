package teacher

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

type TeacherUseCaseInterface interface {
	TeacherLogin(domain Domain, ctx context.Context) (Domain, error)
	TeacherRegister(domain *Domain, ctx context.Context) (Domain, error)
	TeacherGetProfile(ctx context.Context, id uint) (Domain, error)
	TeacherUpdate(ctx context.Context, domain Domain, id uint) (Domain, error)
}

type TeacherRepoInterface interface {
	TeacherLogin(domain Domain, ctx context.Context) (Domain, error)
	TeacherRegister(domain *Domain, ctx context.Context) (Domain, error)
	TeacherGetProfile(ctx context.Context, id uint) (Domain, error)
	TeacherUpdate(ctx context.Context, domain Domain, id uint) (Domain, error)
}