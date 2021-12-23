package course

import (
	"backend/business/course"
	"backend/drivers/database/categories"
	"backend/drivers/database/teacher"
	"time"

	"gorm.io/gorm"
)

type Course struct {
	Id           uint `gorm:"primaryKey"`
	Title        string
	Thumbnail    string
	Description  string
	CategoryId   uint
	Category     categories.Category `gorm:"foreignKey:CategoryId"`
	TeacherId    uint
	Teacher      teacher.Teacher `gorm:"foreignKey:TeacherId"`
	DifficultyId uint
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
	DeleteAt     gorm.DeletedAt
}

func (courses *Course) ToDomain() course.Domain {
	return course.Domain{
		Id:           courses.Id,
		Title:        courses.Title,
		Thumbnail:    courses.Thumbnail,
		Description:  courses.Description,
		CategoryId:   courses.CategoryId,
		Category:     courses.Category.ToDomain(),
		TeacherId:    courses.TeacherId,
		Teacher:      courses.Teacher.ToDomain(),
		DifficultyId: courses.DifficultyId,
		CreatedAt:    courses.CreatedAt,
		UpdatedAt:    courses.UpdatedAt,
	}
}

func FromDomain(domain course.Domain) Course {
	return Course{
		Id:           domain.Id,
		Title:        domain.Title,
		Thumbnail:    domain.Thumbnail,
		Description:  domain.Description,
		CategoryId:   domain.CategoryId,
		Category:     categories.FromDomain(domain.Category),
		TeacherId:    domain.TeacherId,
		Teacher:      teacher.FromDomain(domain.Teacher),
		DifficultyId: domain.DifficultyId,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
	}
}

func ToDomainList(courses []Course) []course.Domain {
	list := []course.Domain{}
	for _, v := range courses {
		list = append(list, v.ToDomain())
	}
	return list
}
