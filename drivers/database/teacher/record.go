package teacher

import (
	"backend/business/teacher"
	"time"

	"gorm.io/gorm"
)

type Teacher struct {
	Id       uint `gorm:"primaryKey"`
	Name     string
	Password string
	Email    string `gorm:"unique"`
	Avatar   string
	Phone    int
	Address  string
	CreateAt time.Time      `gorm:"autoCreateTime"`
	UpdateAt time.Time      `gorm:"autoUpdateTime"`
	DeleteAt gorm.DeletedAt `gorm:"index"`
}

func (tch Teacher) ToDomain() teacher.Domain {
	return teacher.Domain{
		Id:       tch.Id,
		Name:     tch.Name,
		Password: tch.Password,
		Email:    tch.Email,
		Avatar:   tch.Avatar,
		Phone:    tch.Phone,
		Address:  tch.Address,
		CreateAt: tch.CreateAt,
		UpdateAt: tch.UpdateAt,
		DeleteAt: tch.DeleteAt,
	}
}

func FromDomain(domain teacher.Domain) Teacher {
	return Teacher{
		Id:       domain.Id,
		Name:     domain.Name,
		Password: domain.Password,
		Email:    domain.Email,
		Avatar:   domain.Avatar,
		Phone:    domain.Phone,
		Address:  domain.Address,
		CreateAt: domain.CreateAt,
		UpdateAt: domain.UpdateAt,
		DeleteAt: domain.DeleteAt,
	}
}
