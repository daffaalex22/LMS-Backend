package readings

import (
	"backend/business/readings"
	"backend/drivers/database/modules"
	"time"

	"gorm.io/gorm"
)

type Readings struct {
	Id         uint `gorm:"primaryKey"`
	ModuleId   uint
	Module     modules.Modules `gorm:"foreignKey:ModuleId"`
	Title      string
	Content    string
	Order      int
	Quiz       string
	Attachment string
	CreateAt   time.Time      `gorm:"autoCreateTime"`
	UpdateAt   time.Time      `gorm:"autoUpdateTime"`
	DeleteAt   gorm.DeletedAt `gorm:"index"`
}

func (reading Readings) ToDomain() readings.Domain {
	return readings.Domain{
		Id:         reading.Id,
		ModuleId:   reading.ModuleId,
		Module:     reading.Module.ToDomain(),
		Title:      reading.Title,
		Order:      reading.Order,
		Content:    reading.Content,
		Quiz:       reading.Quiz,
		Attachment: reading.Attachment,
		CreateAt:   reading.CreateAt,
		UpdateAt:   reading.UpdateAt,
	}
}

func FromDomain(domain readings.Domain) Readings {
	return Readings{
		Id:         domain.Id,
		ModuleId:   domain.ModuleId,
		Module:     modules.FromDomain(domain.Module),
		Title:      domain.Title,
		Content:    domain.Content,
		Order:      domain.Order,
		Quiz:       domain.Quiz,
		Attachment: domain.Attachment,
		CreateAt:   domain.CreateAt,
		UpdateAt:   domain.UpdateAt,
	}
}

func ToDomainList(datamds []Readings) []readings.Domain {
	All := []readings.Domain{}
	for _, v := range datamds {
		All = append(All, v.ToDomain())
	}
	return All
}
