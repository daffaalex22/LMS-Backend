package modules

import (
	"backend/business/modules"
	"backend/drivers/database/course"
	"time"

	"gorm.io/gorm"
)

type Modules struct {
	Id       uint `gorm:"primaryKey"`
	CourseId uint
	Course   course.Course `gorm:"foreignKey:CourseId"`
	Title    string
	Order    int
	CreateAt time.Time      `gorm:"autoCreateTime"`
	UpdateAt time.Time      `gorm:"autoUpdateTime"`
	DeleteAt gorm.DeletedAt `gorm:"index"`
}

func (mds Modules) ToDomain() modules.Domain {
	return modules.Domain{
		Id:       mds.Id,
		CourseId: mds.CourseId,
		Course:   mds.Course.ToDomain(),
		Title:    mds.Title,
		Order:    mds.Order,
		CreateAt: mds.CreateAt,
		UpdateAt: mds.UpdateAt,
	}
}

func FromDomain(domain modules.Domain) Modules {
	return Modules{
		Id:       domain.Id,
		CourseId: domain.CourseId,
		Course:   course.FromDomain(domain.Course),
		Title:    domain.Title,
		Order:    domain.Order,
		CreateAt: domain.CreateAt,
		UpdateAt: domain.UpdateAt,
	}
}

func ToDomainList(datamds []Modules) []modules.Domain {
	All := []modules.Domain{}
	for _, v := range datamds {
		All = append(All, v.ToDomain())
	}
	return All
}
