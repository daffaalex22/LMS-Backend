package response

import (
	"backend/business/teacher"
	"time"

	"gorm.io/gorm"
)

type TeacherResponse struct {
	Id         uint           `json:"id"`
	Name       string         `json:"name"`
	Email      string         `json:"email"`
	Avatar     string         `json:"avatar"`
	Phone      int            `json:"phone"`
	Address    string         `json:"address"`
	BackGround string         `json:"background"`
	Token      string         `json:"token"`
	CreateAt   time.Time      `json:"createAt"`
	UpdateAt   time.Time      `json:"updateAt"`
	DeleteAt   gorm.DeletedAt `json:"deleteAt"`
}

type TeacherProfile struct {
	Id         uint           `json:"id"`
	Name       string         `json:"name"`
	Email      string         `json:"email"`
	Avatar     string         `json:"avatar"`
	Phone      int            `json:"phone"`
	Address    string         `json:"address"`
	BackGround string         `json:"background"`
	CreateAt   time.Time      `json:"createAt"`
	UpdateAt   time.Time      `json:"updateAt"`
	DeleteAt   gorm.DeletedAt `json:"deleteAt"`
}

type TeacherResponseRegister struct {
	Id         uint           `json:"id"`
	Name       string         `json:"name"`
	Email      string         `json:"email"`
	Avatar     string         `json:"avatar"`
	Phone      int            `json:"phone"`
	Address    string         `json:"address"`
	BackGround string         `json:"background"`
	CreateAt   time.Time      `json:"createAt"`
	UpdateAt   time.Time      `json:"updateAt"`
	DeleteAt   gorm.DeletedAt `json:"deleteAt"`
}

type TeacherResponseUpdate struct {
	Id         uint           `json:"id"`
	Name       string         `json:"name"`
	Email      string         `json:"email"`
	Avatar     string         `json:"avatar"`
	Phone      int            `json:"phone"`
	Address    string         `json:"address"`
	BackGround string         `json:"background"`
	CreateAt   time.Time      `json:"createAt"`
	UpdateAt   time.Time      `json:"updateAt"`
	DeleteAt   gorm.DeletedAt `json:"deleteAt"`
}

func FromDomainLogin(domain teacher.Domain) TeacherResponse {
	return TeacherResponse{
		Id:         domain.Id,
		Name:       domain.Name,
		Email:      domain.Email,
		Avatar:     domain.Avatar,
		Phone:      domain.Phone,
		Address:    domain.Address,
		BackGround: domain.BackGround,
		Token:      domain.Token,
		CreateAt:   domain.CreateAt,
		UpdateAt:   domain.UpdateAt,
		DeleteAt:   domain.DeleteAt,
	}
}

func FromDomainProfile(domain teacher.Domain) TeacherProfile {
	return TeacherProfile{
		Id:         domain.Id,
		Name:       domain.Name,
		Email:      domain.Email,
		Avatar:     domain.Avatar,
		Phone:      domain.Phone,
		Address:    domain.Address,
		BackGround: domain.BackGround,
		CreateAt:   domain.CreateAt,
		UpdateAt:   domain.UpdateAt,
		DeleteAt:   domain.DeleteAt,
	}
}

func FromDomainToRegist(domain teacher.Domain) TeacherResponseRegister {
	return TeacherResponseRegister{
		Id:         domain.Id,
		Name:       domain.Name,
		Email:      domain.Email,
		Avatar:     domain.Avatar,
		Phone:      domain.Phone,
		Address:    domain.Address,
		BackGround: domain.BackGround,
		CreateAt:   domain.CreateAt,
		UpdateAt:   domain.UpdateAt,
		DeleteAt:   domain.DeleteAt,
	}
}

func FromDomainToUpdate(domain teacher.Domain) TeacherResponseUpdate {
	return TeacherResponseUpdate{
		Id:         domain.Id,
		Name:       domain.Name,
		Email:      domain.Email,
		Avatar:     domain.Avatar,
		Phone:      domain.Phone,
		Address:    domain.Address,
		BackGround: domain.BackGround,
		CreateAt:   domain.CreateAt,
		UpdateAt:   domain.UpdateAt,
		DeleteAt:   domain.DeleteAt,
	}
}
