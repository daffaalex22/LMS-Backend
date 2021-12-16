package response

import (
	"backend/business/categories"
	"time"
)

type CategoryResponse struct {
	Id        uint      `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"createdAt"`
	UpdateAt  time.Time `json:"updateAt"`
}

func FromDomain(domain categories.Domain) CategoryResponse {
	return CategoryResponse{
		Id:        domain.Id,
		Title:     domain.Title,
		CreatedAt: domain.CreatedAt,
		UpdateAt:  domain.UpdateAt,
	}
}

func FromDomainList(domain []categories.Domain) []CategoryResponse {
	list := []CategoryResponse{}
	for _, value := range domain {
		list = append(list, FromDomain(value))
	}
	return list
}
