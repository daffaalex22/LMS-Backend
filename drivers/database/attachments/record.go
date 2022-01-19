package attachments

import (
	"backend/business/attachments"
	"time"

	"gorm.io/gorm"
)

type Attachments struct {
	Id          uint `gorm:"primaryKey"`
	ContentType string
	ContentId   uint
	VideoId     uint
	ReadingId   uint
	// QuizId     uint
	Title    string
	Url      string
	CreateAt time.Time      `gorm:"autoCreateTime"`
	UpdateAt time.Time      `gorm:"autoUpdateTime"`
	DeleteAt gorm.DeletedAt `gorm:"index"`
}

func (attach Attachments) ToDomain() attachments.Domain {
	return attachments.Domain{
		Id:          attach.Id,
		ContentType: attach.ContentType,
		ContentId:   attach.ContentId,
		VideoId:     attach.VideoId,
		// ReadingId: attach.ReadingId,
		Title:    attach.Title,
		Url:      attach.Url,
		CreateAt: attach.CreateAt,
		UpdateAt: attach.UpdateAt,
	}
}

func FromDomain(domain attachments.Domain) Attachments {
	return Attachments{
		Id:          domain.Id,
		ContentType: domain.ContentType,
		ContentId:   domain.ContentId,
		VideoId:     domain.VideoId,
		// ReadingId: domain.ReadingId,
		Title:    domain.Title,
		Url:      domain.Url,
		CreateAt: domain.CreateAt,
		UpdateAt: domain.UpdateAt,
	}
}

func ToDomainList(datamds []Attachments) []attachments.Domain {
	All := []attachments.Domain{}
	for _, v := range datamds {
		All = append(All, v.ToDomain())
	}
	return All
}
