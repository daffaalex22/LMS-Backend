package response

import (
	"backend/business/student"
	"time"

	"gorm.io/gorm"
)

type StudentResponse struct {
	Id       uint           `json:"id"`
	Name     string         `json:"name"`
	Email    string         `json:"email"`
	Avatar   string         `json:"avatar"`
	Phone    int            `json:"phone"`
	Address  string         `json:"address"`
	Token    string         `json:"token"`
	CreateAt time.Time      `json:"createAt"`
	UpdateAt time.Time      `json:"updateAt"`
	DeleteAt gorm.DeletedAt `json:"deleteAt"`
}

type ResponseRegister struct {
	Id       uint           `json:"id"`
	Name     string         `json:"name"`
	Email    string         `json:"email"`
	Avatar   string         `json:"avatar"`
	Phone    int            `json:"phone"`
	Address  string         `json:"address"`
	CreateAt time.Time      `json:"createAt"`
	UpdateAt time.Time      `json:"updateAt"`
	DeleteAt gorm.DeletedAt `json:"deleteAt"`
}
type ResponseUpdate struct {
	Id       uint           `json:"id"`
	Name     string         `json:"name"`
	Email    string         `json:"email"`
	Avatar   string         `json:"avatar"`
	Phone    int            `json:"phone"`
	Address  string         `json:"address"`
	CreateAt time.Time      `json:"createAt"`
	UpdateAt time.Time      `json:"updateAt"`
	DeleteAt gorm.DeletedAt `json:"deleteAt"`
}

type StudentProfile struct {
	Id       uint           `json:"id"`
	Name     string         `json:"name"`
	Email    string         `json:"email"`
	Avatar   string         `json:"avatar"`
	Phone    int            `json:"phone"`
	Address  string         `json:"address"`
	CreateAt time.Time      `json:"createAt"`
	UpdateAt time.Time      `json:"updateAt"`
	DeleteAt gorm.DeletedAt `json:"deleteAt"`
}

func FromDomainLogin(domain student.Domain) StudentResponse {
	return StudentResponse{
		Id:       domain.Id,
		Name:     domain.Name,
		Email:    domain.Email,
		Avatar:   domain.Avatar,
		Phone:    domain.Phone,
		Address:  domain.Address,
		Token:    domain.Token,
		CreateAt: domain.CreateAt,
		UpdateAt: domain.UpdateAt,
		DeleteAt: domain.DeleteAt,
	}
}

func FromDomainToRegist(domain student.Domain) ResponseRegister {
	return ResponseRegister{
		Id:       domain.Id,
		Name:     domain.Name,
		Email:    domain.Email,
		Avatar:   domain.Avatar,
		Phone:    domain.Phone,
		Address:  domain.Address,
		CreateAt: domain.CreateAt,
		UpdateAt: domain.UpdateAt,
		DeleteAt: domain.DeleteAt,
	}
}

func FromDomainToUpdate(domain student.Domain) ResponseUpdate {
	return ResponseUpdate{
		Id:       domain.Id,
		Name:     domain.Name,
		Email:    domain.Email,
		Avatar:   domain.Avatar,
		Phone:    domain.Phone,
		Address:  domain.Address,
		CreateAt: domain.CreateAt,
		UpdateAt: domain.UpdateAt,
		DeleteAt: domain.DeleteAt,
	}
}

func FromDomainProfile(domain student.Domain) StudentProfile {
	return StudentProfile{
		Id:       domain.Id,
		Name:     domain.Name,
		Email:    domain.Email,
		Avatar:   domain.Avatar,
		Phone:    domain.Phone,
		Address:  domain.Address,
		CreateAt: domain.CreateAt,
		UpdateAt: domain.UpdateAt,
		DeleteAt: domain.DeleteAt,
	}
}
