package request

import "backend/business/teacher"

type TeacherLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TeacherRegister struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type TeacherUpdate struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	NewPassword string `json:"newPassword"`
	Avatar      string `json:"avatar"`
	Phone       string `json:"phone"`
	Address     string `json:"address"`
	BackGround  string `json:"background"`
}

func (tch *TeacherLogin) ToDomainLogin() *teacher.Domain {
	return &teacher.Domain{
		Email:    tch.Email,
		Password: tch.Password,
	}
}

func (tch *TeacherRegister) ToDomainRegist() *teacher.Domain {
	return &teacher.Domain{
		Name:     tch.Name,
		Email:    tch.Email,
		Password: tch.Password,
	}
}

func (tch *TeacherUpdate) ToDomainUpdate() *teacher.Domain {
	return &teacher.Domain{
		Name:        tch.Name,
		Email:       tch.Email,
		Password:    tch.Password,
		NewPassword: tch.NewPassword,
		Avatar:      tch.Avatar,
		Phone:       tch.Phone,
		Address:     tch.Address,
		BackGround:  tch.BackGround,
	}
}
