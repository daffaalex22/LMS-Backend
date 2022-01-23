package response

import (
	"backend/business/requests"
	_courseReponse "backend/controllers/courses/response"
	_studentReponse "backend/controllers/student/response"
	_typeReponse "backend/controllers/types/response"
	"time"
)

type RequestsResponse struct {
	Id        uint                           `json:"id"`
	StudentId uint                           `json:"studentId"`
	CourseId  uint                           `json:"courseId"`
	TypeId    uint                           `json:"typeId"`
	Status    string                         `json:"status"`
	Message   string                         `json:"message"`
	Student   _studentReponse.StudentProfile `json:"student"`
	Course    _courseReponse.Response        `json:"course"`
	Type      _typeReponse.TypeResponse      `json:"type"`
	CreateAt  time.Time                      `json:"createdAt"`
	UpdateAt  time.Time                      `json:"updateAt"`
}

func FromDomain(domain requests.Domain) RequestsResponse {
	return RequestsResponse{
		Id:        domain.Id,
		TypeId:    domain.TypeId,
		Type:      _typeReponse.FromDomain(domain.Type),
		StudentId: domain.StudentId,
		Student:   _studentReponse.FromDomainProfile(domain.Student),
		CourseId:  domain.CourseId,
		Course:    _courseReponse.FromDomain(domain.Course),
		Status:    domain.Status,
		Message:   domain.Message,
		CreateAt:  domain.CreateAt,
		UpdateAt:  domain.UpdateAt,
	}
}

func FromDomainList(domain []requests.Domain) []RequestsResponse {
	list := []RequestsResponse{}
	for _, value := range domain {
		list = append(list, FromDomain(value))
	}
	return list
}
