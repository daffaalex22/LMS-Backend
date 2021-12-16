package categories

import (
	"backend/business/categories"
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID        uint `gorm:"primaryKey"`
	Title     string
	CreatedAt time.Time
	UpdateAt  time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (category *Category) ToDomain() categories.Domain {
	return categories.Domain{
		Id:        category.ID,
		Title:     category.Title,
		CreatedAt: category.CreatedAt,
		UpdateAt:  category.UpdateAt,
	}
}

func ToDomainList(record []Category) []categories.Domain {
	var returnValue []categories.Domain
	for _, value := range record {
		returnValue = append(returnValue, value.ToDomain())
	}
	return returnValue
}
