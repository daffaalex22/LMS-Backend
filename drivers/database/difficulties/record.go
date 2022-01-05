package difficulties

import (
	"backend/business/difficulties"
	"time"

	"gorm.io/gorm"
)

type Difficulty struct {
	ID        uint `gorm:"primaryKey"`
	Title     string
	CreatedAt time.Time
	UpdateAt  time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (difficulty *Difficulty) ToDomain() difficulties.Domain {
	return difficulties.Domain{
		Id:        difficulty.ID,
		Title:     difficulty.Title,
		CreatedAt: difficulty.CreatedAt,
		UpdateAt:  difficulty.UpdateAt,
	}
}
func FromDomain(domain difficulties.Domain) Difficulty {
	return Difficulty{
		ID:        domain.Id,
		Title:     domain.Title,
		CreatedAt: domain.CreatedAt,
		UpdateAt:  domain.UpdateAt,
	}
}

func ToDomainList(record []Difficulty) []difficulties.Domain {
	var returnValue []difficulties.Domain
	for _, value := range record {
		returnValue = append(returnValue, value.ToDomain())
	}
	return returnValue
}
