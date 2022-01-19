package student

import (
	"backend/business/student"
	"time"

	"gorm.io/gorm"
)

type Student struct {
	Id       uint `gorm:"primaryKey"`
	Name     string
	Password string
	Email    string `gorm:"unique"`
	Avatar   string
	Phone    string
	Address  string
	CreateAt time.Time      `gorm:"autoCreateTime"`
	UpdateAt time.Time      `gorm:"autoUpdateTime"`
	DeleteAt gorm.DeletedAt `gorm:"index"`
}

func (std Student) ToDomain() student.Domain {
	return student.Domain{
		Id:       std.Id,
		Name:     std.Name,
		Password: std.Password,
		Email:    std.Email,
		Avatar:   std.Avatar,
		Phone:    std.Phone,
		Address:  std.Address,
		CreateAt: std.CreateAt,
		UpdateAt: std.UpdateAt,
		DeleteAt: std.DeleteAt,
	}
}

func FromDomain(domain student.Domain) Student {
	return Student{
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
