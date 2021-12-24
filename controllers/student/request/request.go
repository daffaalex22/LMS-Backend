package request

import "backend/business/student"

type StudentLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type StudentRegister struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type StudentUpdate struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
	Avatar          string `json:"avatar"`
	Phone           int    `json:"phone"`
	Address         string `json:"address"`
}

func (std *StudentLogin) ToDomainLogin() *student.Domain {
	return &student.Domain{
		Email:    std.Email,
		Password: std.Password,
	}
}

func (std *StudentRegister) ToDomainRegist() *student.Domain {
	return &student.Domain{
		Name:     std.Name,
		Email:    std.Email,
		Password: std.Password,
	}
}

func (std *StudentUpdate) ToDomainUpdate() *student.Domain {
	return &student.Domain{
		Name:            std.Name,
		Email:           std.Email,
		Password:        std.Password,
		ConfirmPassword: std.ConfirmPassword,
		Avatar:          std.Avatar,
		Phone:           std.Phone,
		Address:         std.Address,
	}
}
