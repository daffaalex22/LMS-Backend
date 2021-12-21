package enrollments

import (
	"backend/business/enrollments"
	"backend/drivers/database/course"
	"backend/drivers/database/student"
	"time"

	"gorm.io/gorm"
)

type Enrollments struct {
	StudentId uint            `gorm:"primaryKey"`
	Student   student.Student `gorm:"foreignKey:StudentId"`
	CourseId  uint            `gorm:"primaryKey"`
	Course    course.Course   `gorm:"foreignKey:CourseId"`
	Rating    int
	Review    string
	CreateAt  time.Time      `gorm:"autoCreateTime"`
	UpdateAt  time.Time      `gorm:"autoUpdateTime"`
	DeleteAt  gorm.DeletedAt `gorm:"index"`
}

func (elm Enrollments) ToDomain() enrollments.Domain {
	return enrollments.Domain{
		StudentId: elm.StudentId,
		CourseId:  elm.CourseId,
		Student:   elm.Student.ToDomain(),
		Course:    elm.Course.ToDomain(),
		Rating:    elm.Rating,
		Review:    elm.Review,
		CreateAt:  elm.CreateAt,
		UpdateAt:  elm.UpdateAt,
	}
}

func FromDomain(domain enrollments.Domain) Enrollments {
	return Enrollments{
		StudentId: domain.StudentId,
		CourseId:  domain.CourseId,
		Rating:    domain.Rating,
		Review:    domain.Review,
		CreateAt:  domain.CreateAt,
		UpdateAt:  domain.UpdateAt,
	}
}

func ToDomainList(dataelm []Enrollments) []enrollments.Domain {
	All := []enrollments.Domain{}
	for _, v := range dataelm {
		All = append(All, v.ToDomain())
	}
	return All
}
