package request

import "backend/business/requests"

type RequestsAdd struct {
	StudentId uint   `json:"studentId"`
	CourseId  uint   `json:"courseId"`
	TypeId    uint   `json:"typeId"`
	Status    string `json:"status"`
	Message   string `json:"message"`
}

type RequestsUpdate struct {
	Id     uint   `json:"id"`
	Status string `json:"status"`
}

func (req *RequestsAdd) ToDomain() requests.Domain {
	return requests.Domain{
		StudentId: req.StudentId,
		CourseId:  req.CourseId,
		TypeId:    req.TypeId,
		Status:    req.Status,
		Message:   req.Message,
	}
}

func (req *RequestsUpdate) ToDomain() requests.Domain {
	return requests.Domain{
		Id:     req.Id,
		Status: req.Status,
	}
}
