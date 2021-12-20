package response

import (
	"backend/business/categories"
	"backend/business/course"
	"time"
)

type AddResponse struct {
	Id           uint              `json:"id"`
	Title        string            `json:"title"`
	Thumbnail    string            `json:"thumbnail"`
	Description  string            `json:"description"`
	CategoryId   uint              `json:"categoryId"`
	TeacherId    uint              `json:"teacherId"`
	DifficultyId uint              `json:"difficultyId"`
	Category     categories.Domain `json:"categories"`
	CreatedAt    time.Time         `json:"created_at"`
	UpdatedAt    time.Time         `json:"updated_at"`
}

func FromDomain(domain course.Domain) AddResponse {
	return AddResponse{
		Id:          domain.Id,
		Title:       domain.Title,
		Thumbnail:   domain.Thumbnail,
		Description: domain.Description,
		CategoryId:  domain.CategoryId,
		Category: categories.Domain{
			Id:        domain.Category.Id,
			Title:     domain.Category.Title,
			CreatedAt: domain.Category.CreatedAt,
			UpdateAt:  domain.Category.UpdateAt,
		},
		TeacherId:    domain.TeacherId,
		DifficultyId: domain.DifficultyId,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
	}
}
