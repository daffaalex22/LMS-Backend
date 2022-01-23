package requests

import (
	"backend/business/requests"
	"backend/drivers/database/course"
	"backend/drivers/database/student"
	"backend/drivers/database/types"
	"time"

	"gorm.io/gorm"
)

type Requests struct {
	Id        uint `gorm:"primaryKey"`
	StudentId uint
	Student   student.Student `gorm:"foreignKey:StudentId"`
	CourseId  uint
	Course    course.Course `gorm:"foreignKey:CourseId"`
	TypeId    uint
	Type      types.Types `gorm:"foreignKey:TypeId"`
	Status    string
	Message   string
	CreateAt  time.Time      `gorm:"autoCreateTime"`
	UpdateAt  time.Time      `gorm:"autoUpdateTime"`
	DeleteAt  gorm.DeletedAt `gorm:"index"`
}

func (req Requests) ToDomain() requests.Domain {
	return requests.Domain{
		Id:        req.Id,
		TypeId:    req.TypeId,
		StudentId: req.StudentId,
		CourseId:  req.CourseId,
		Student:   req.Student.ToDomain(),
		Course:    req.Course.ToDomain(),
		Type:      req.Type.ToDomain(),
		Status:    req.Status,
		Message:   req.Message,
		CreateAt:  req.CreateAt,
		UpdateAt:  req.UpdateAt,
	}
}

func FromDomain(domain requests.Domain) Requests {
	return Requests{
		TypeId:    domain.TypeId,
		Type:      types.FromDomain(domain.Type),
		StudentId: domain.StudentId,
		Student:   student.FromDomain(domain.Student),
		CourseId:  domain.CourseId,
		Course:    course.FromDomain(domain.Course),
		Status:    domain.Status,
		Message:   domain.Message,
		CreateAt:  domain.CreateAt,
		UpdateAt:  domain.UpdateAt,
	}
}

func ToDomainList(dataRequests []Requests) []requests.Domain {
	All := []requests.Domain{}
	for _, v := range dataRequests {
		All = append(All, v.ToDomain())
	}
	return All
}
