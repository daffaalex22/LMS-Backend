package response

import (
	"backend/business/readings"
	_modules "backend/controllers/modules/response"
	"time"
)

type ReadingsResponse struct {
	Id         uint                     `json:"id"`
	ModuleId   uint                     `json:"moduleId"`
	Module     _modules.ModulesResponse `json:"module"`
	Title      string                   `json:"title"`
	Content    string                   `json:"content"`
	Order      int                      `json:"order"`
	Quiz       string                   `json:"quiz"`
	Attachment string                   `json:"attachment"`
	CreateAt   time.Time                `json:"createdAt"`
	UpdateAt   time.Time                `json:"updateAt"`
}

func FromDomain(domain readings.Domain) ReadingsResponse {
	return ReadingsResponse{
		Id:         domain.Id,
		ModuleId:   domain.ModuleId,
		Module:     _modules.FromDomain(domain.Module),
		Title:      domain.Title,
		Content:    domain.Content,
		Order:      domain.Order,
		Quiz:       domain.Quiz,
		Attachment: domain.Attachment,
		CreateAt:   domain.CreateAt,
		UpdateAt:   domain.UpdateAt,
	}
}

func FromDomainList(domain []readings.Domain) []ReadingsResponse {
	list := []ReadingsResponse{}
	for _, value := range domain {
		list = append(list, FromDomain(value))
	}
	return list
}
