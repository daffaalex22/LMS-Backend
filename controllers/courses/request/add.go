package request

import "backend/business/course"

type AddRequest struct {
	Title        string `json:"title"`
	Thumbnail    string `json:"thumbnail"`
	Description  string `json:"description"`
	CategoryId   uint   `json:"categoryId"`
	TeacherId    uint   `json:"teacherId"`
	DifficultyId uint   `json:"difficultyId"`
}

func (req *AddRequest) ToDomain() course.Domain {
	return course.Domain{
		Title:        req.Title,
		Thumbnail:    req.Thumbnail,
		Description:  req.Description,
		CategoryId:   req.CategoryId,
		TeacherId:    req.TeacherId,
		DifficultyId: req.DifficultyId,
	}
}
