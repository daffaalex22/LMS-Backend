package request

import "backend/business/course"

type UpdateRequest struct {
	Id           uint   `json:"id"`
	Title        string `json:"title"`
	Thumbnail    string `json:"thumbnail"`
	Description  string `json:"description"`
	CategoryId   uint   `json:"categoryId"`
	TeacherId    uint   `json:"teacherId"`
	DifficultyId uint   `json:"difficultyId"`
}

func (req *UpdateRequest) ToDomain() course.Domain {
	return course.Domain{
		Id:           req.Id,
		Title:        req.Title,
		Thumbnail:    req.Thumbnail,
		Description:  req.Description,
		CategoryId:   req.CategoryId,
		TeacherId:    req.TeacherId,
		DifficultyId: req.DifficultyId,
	}
}
