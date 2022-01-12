package types

import (
	"backend/business/types"
	"time"

	"gorm.io/gorm"
)

type Types struct {
	ID        uint `gorm:"primaryKey"`
	Title     string
	CreatedAt time.Time
	UpdateAt  time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (typ *Types) ToDomain() types.Domain {
	return types.Domain{
		Id:        typ.ID,
		Title:     typ.Title,
		CreatedAt: typ.CreatedAt,
		UpdateAt:  typ.UpdateAt,
	}
}
func FromDomain(domain types.Domain) Types {
	return Types{
		ID:        domain.Id,
		Title:     domain.Title,
		CreatedAt: domain.CreatedAt,
		UpdateAt:  domain.UpdateAt,
	}
}

func ToDomainList(record []Types) []types.Domain {
	var returnValue []types.Domain
	for _, value := range record {
		returnValue = append(returnValue, value.ToDomain())
	}
	return returnValue
}
