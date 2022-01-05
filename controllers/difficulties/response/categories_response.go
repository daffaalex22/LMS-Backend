package response

import (
	"backend/business/difficulties"
	"time"
)

type DifficultyResponse struct {
	Id        uint      `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"createdAt"`
	UpdateAt  time.Time `json:"updateAt"`
}

func FromDomain(domain difficulties.Domain) DifficultyResponse {
	return DifficultyResponse{
		Id:        domain.Id,
		Title:     domain.Title,
		CreatedAt: domain.CreatedAt,
		UpdateAt:  domain.UpdateAt,
	}
}

func FromDomainList(domain []difficulties.Domain) []DifficultyResponse {
	list := []DifficultyResponse{}
	for _, value := range domain {
		list = append(list, FromDomain(value))
	}
	return list
}
