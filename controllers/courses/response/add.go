package response

import (
	"backend/business/course"
	_categoriesResponse "backend/controllers/categories/response"
	_difficultyResponse "backend/controllers/difficulties/response"
	_teacherResponse "backend/controllers/teacher/response"
	"time"
)

type Response struct {
	Id           uint                                   `json:"id"`
	Title        string                                 `json:"title"`
	Thumbnail    string                                 `json:"thumbnail"`
	Description  string                                 `json:"description"`
	Rating       float32                                `json:"rating"`
	CategoryId   uint                                   `json:"categoryId"`
	TeacherId    uint                                   `json:"teacherId"`
	DifficultyId uint                                   `json:"difficultyId"`
	Category     _categoriesResponse.CategoryResponse   `json:"categories"`
	Teacher      _teacherResponse.TeacherProfile        `json:"teacher"`
	Difficulty   _difficultyResponse.DifficultyResponse `json:"difficulty"`
	CreatedAt    time.Time                              `json:"created_at"`
	UpdatedAt    time.Time                              `json:"updated_at"`
}

type BatchesResponse struct {
	Id           uint      `json:"id"`
	Title        string    `json:"title"`
	Thumbnail    string    `json:"thumbnail"`
	Description  string    `json:"description"`
	Rating       float32   `json:"rating"`
	CategoryId   uint      `json:"categoryId"`
	TeacherId    uint      `json:"teacherId"`
	DifficultyId uint      `json:"difficultyId"`
	Category     string    `json:"category"`
	Teacher      string    `json:"teacher"`
	Difficulty   string    `json:"difficulty"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func FromDomain(domain course.Domain) Response {
	return Response{
		Id:           domain.Id,
		Title:        domain.Title,
		Thumbnail:    domain.Thumbnail,
		Description:  domain.Description,
		Rating:       domain.Rating,
		CategoryId:   domain.CategoryId,
		Category:     _categoriesResponse.FromDomain(domain.Category),
		TeacherId:    domain.TeacherId,
		Teacher:      _teacherResponse.FromDomainProfile(domain.Teacher),
		DifficultyId: domain.DifficultyId,
		Difficulty:   _difficultyResponse.FromDomain(domain.Difficulty),
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
	}
}

func BatchesFromDomain(domain course.BatchesDomain) BatchesResponse {
	return BatchesResponse{
		Id:           domain.Id,
		Title:        domain.Title,
		Thumbnail:    domain.Thumbnail,
		Description:  domain.Description,
		Rating:       domain.Rating,
		CategoryId:   domain.CategoryId,
		Category:     domain.Category,
		TeacherId:    domain.TeacherId,
		Teacher:      domain.Teacher,
		DifficultyId: domain.DifficultyId,
		Difficulty:   domain.Difficulty,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
	}
}

func FromDomainList(domain []course.Domain) []Response {
	list := []Response{}
	for _, v := range domain {
		list = append(list, FromDomain(v))
	}
	return list
}

func BatchesFromDomainList(domain []course.BatchesDomain) []BatchesResponse {
	list := []BatchesResponse{}
	for _, v := range domain {
		list = append(list, BatchesFromDomain(v))
	}
	return list
}
