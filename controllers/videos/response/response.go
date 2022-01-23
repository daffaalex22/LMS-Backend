package response

import (
	"backend/business/videos"
	_modules "backend/controllers/modules/response"
	"time"
)

type VideosResponse struct {
	Id         uint                     `json:"id"`
	ModuleId   uint                     `json:"moduleId"`
	Module     _modules.ModulesResponse `json:"module"`
	Title      string                   `json:"title"`
	Caption    string                   `json:"caption"`
	Url        string                   `json:"url"`
	Order      int                      `json:"order"`
	Quiz       string                   `json:"quiz"`
	Attachment string                   `json:"attachment"`
	CreateAt   time.Time                `json:"createdAt"`
	UpdateAt   time.Time                `json:"updateAt"`
}

func FromDomain(domain videos.Domain) VideosResponse {
	return VideosResponse{
		Id:         domain.Id,
		ModuleId:   domain.ModuleId,
		Module:     _modules.FromDomain(domain.Module),
		Title:      domain.Title,
		Caption:    domain.Caption,
		Url:        domain.Url,
		Order:      domain.Order,
		Quiz:       domain.Quiz,
		Attachment: domain.Attachment,
		CreateAt:   domain.CreateAt,
		UpdateAt:   domain.UpdateAt,
	}
}

func FromDomainList(domain []videos.Domain) []VideosResponse {
	list := []VideosResponse{}
	for _, value := range domain {
		list = append(list, FromDomain(value))
	}
	return list
}
