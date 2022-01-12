package response

import (
	"backend/business/types"
	"time"
)

type TypeResponse struct {
	Id        uint      `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"createdAt"`
	UpdateAt  time.Time `json:"updateAt"`
}

func FromDomain(domain types.Domain) TypeResponse {
	return TypeResponse{
		Id:        domain.Id,
		Title:     domain.Title,
		CreatedAt: domain.CreatedAt,
		UpdateAt:  domain.UpdateAt,
	}
}

func FromDomainList(domain []types.Domain) []TypeResponse {
	list := []TypeResponse{}
	for _, value := range domain {
		list = append(list, FromDomain(value))
	}
	return list
}
