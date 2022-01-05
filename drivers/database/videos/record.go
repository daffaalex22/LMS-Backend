package videos

import (
	"backend/business/videos"
	"backend/drivers/database/modules"
	"time"

	"gorm.io/gorm"
)

type Videos struct {
	Id       uint `gorm:"primaryKey"`
	ModuleId uint
	Module   modules.Modules `gorm:"foreignKey:ModuleId"`
	Title    string
	Caption  string
	Url      string
	Order    int
	CreateAt time.Time      `gorm:"autoCreateTime"`
	UpdateAt time.Time      `gorm:"autoUpdateTime"`
	DeleteAt gorm.DeletedAt `gorm:"index"`
}

func (video Videos) ToDomain() videos.Domain {
	return videos.Domain{
		Id:       video.Id,
		ModuleId: video.ModuleId,
		Module:   video.Module.ToDomain(),
		Title:    video.Title,
		Order:    video.Order,
		Caption:  video.Caption,
		Url:      video.Url,
		CreateAt: video.CreateAt,
		UpdateAt: video.UpdateAt,
	}
}

func FromDomain(domain videos.Domain) Videos {
	return Videos{
		Id:       domain.Id,
		ModuleId: domain.ModuleId,
		Module:   modules.FromDomain(domain.Module),
		Title:    domain.Title,
		Caption:  domain.Caption,
		Order:    domain.Order,
		Url:      domain.Url,
		CreateAt: domain.CreateAt,
		UpdateAt: domain.UpdateAt,
	}
}

func ToDomainList(datamds []Videos) []videos.Domain {
	All := []videos.Domain{}
	for _, v := range datamds {
		All = append(All, v.ToDomain())
	}
	return All
}
