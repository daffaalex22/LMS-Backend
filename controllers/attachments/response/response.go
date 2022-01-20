package response

import (
	"backend/business/attachments"
	"time"
)

type AttachmentsResponse struct {
	Id          uint      `json:"id"`
	ContentType string    `json:"contentType"`
	ContentId   uint      `json:"contentId"`
	VideoId     uint      `json:"videoId"`
	ReadingId   uint      `json:"readingId"`
	Title       string    `json:"title"`
	Url         string    `json:"url"`
	CreateAt    time.Time `json:"createdAt"`
	UpdateAt    time.Time `json:"updateAt"`
}

func FromDomain(domain attachments.Domain) AttachmentsResponse {
	return AttachmentsResponse{
		Id:          domain.Id,
		ContentType: domain.ContentType,
		ContentId:   domain.ContentId,
		VideoId:     domain.VideoId,
		ReadingId:   domain.ReadingId,
		Title:       domain.Title,
		Url:         domain.Url,
		CreateAt:    domain.CreateAt,
		UpdateAt:    domain.UpdateAt,
	}
}

func FromDomainList(domain []attachments.Domain) []AttachmentsResponse {
	list := []AttachmentsResponse{}
	for _, value := range domain {
		list = append(list, FromDomain(value))
	}
	return list
}
