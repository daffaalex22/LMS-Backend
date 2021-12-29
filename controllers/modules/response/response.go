package response

import (
	"backend/business/modules"
	_courseReponse "backend/controllers/courses/response"
	"time"
)

type ModulesResponse struct {
	Id       uint                    `json:"id"`
	CourseId uint                    `json:"courseId"`
	Title    string                  `json:"title"`
	Order    int                     `json:"order"`
	Course   _courseReponse.Response `json:"course"`
	CreateAt time.Time               `json:"createdAt"`
	UpdateAt time.Time               `json:"updateAt"`
}

func FromDomain(domain modules.Domain) ModulesResponse {
	return ModulesResponse{
		Id:       domain.Id,
		CourseId: domain.CourseId,
		Title:    domain.Title,
		Order:    domain.Order,
		Course:   _courseReponse.FromDomain(domain.Course),
		CreateAt: domain.CreateAt,
		UpdateAt: domain.UpdateAt,
	}
}

func FromDomainList(domain []modules.Domain) []ModulesResponse {
	list := []ModulesResponse{}
	for _, value := range domain {
		list = append(list, FromDomain(value))
	}
	return list
}
