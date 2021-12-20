package response

import (
	"backend/business/course"
	_categoriesResponse "backend/controllers/categories/response"
	_teacherResponse "backend/controllers/teacher/response"
	"time"
)

type AddResponse struct {
	Id           uint                                 `json:"id"`
	Title        string                               `json:"title"`
	Thumbnail    string                               `json:"thumbnail"`
	Description  string                               `json:"description"`
	CategoryId   uint                                 `json:"categoryId"`
	TeacherId    uint                                 `json:"teacherId"`
	DifficultyId uint                                 `json:"difficultyId"`
	Category     _categoriesResponse.CategoryResponse `json:"categories"`
	Teacher      _teacherResponse.TeacherProfile      `json:"teacher"`
	CreatedAt    time.Time                            `json:"created_at"`
	UpdatedAt    time.Time                            `json:"updated_at"`
}

func FromDomain(domain course.Domain) AddResponse {
	return AddResponse{
		Id:           domain.Id,
		Title:        domain.Title,
		Thumbnail:    domain.Thumbnail,
		Description:  domain.Description,
		CategoryId:   domain.CategoryId,
		Category:     _categoriesResponse.FromDomain(domain.Category),
		TeacherId:    domain.TeacherId,
		Teacher:      _teacherResponse.FromDomainProfile(domain.Teacher),
		DifficultyId: domain.DifficultyId,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
	}
}
