package enrollments

import (
	"backend/business/enrollments"
	"backend/drivers/database/course"
	"backend/drivers/database/student"
	"time"

	"gorm.io/gorm"
)

type Enrollments struct {
	Student_Id uint            `gorm:"primaryKey"`
	Student    student.Student `gorm:"foreignKey:Student_Id"`
	Course_Id  uint            `gorm:"primaryKey"`
	Course     course.Course   `gorm:"foreignKey:Course_Id"`
	Rating     int
	Review     string
	CreateAt   time.Time
	UpdateAt   time.Time
	DeleteAt   gorm.DeletedAt `gorm:"index"`
}

func (elm Enrollments) ToDomain() enrollments.Domain {
	return enrollments.Domain{
		Student_Id: elm.Student_Id,
		Course_Id:  elm.Course_Id,
		Rating:     elm.Rating,
		Review:     elm.Review,
		CreateAt:   elm.CreateAt,
		UpdateAt:   elm.UpdateAt,
	}
}

func FromDomain(domain enrollments.Domain) Enrollments {
	return Enrollments{
		Student_Id: domain.Student_Id,
		Course_Id:  domain.Course_Id,
		Rating:     domain.Rating,
		Review:     domain.Review,
		CreateAt:   domain.CreateAt,
		UpdateAt:   domain.UpdateAt,
	}
}

func ToDomainList(dataelm []Enrollments) []enrollments.Domain {
	All := []enrollments.Domain{}
	for _, v := range dataelm {
		All = append(All, v.ToDomain())
	}
	return All
}
