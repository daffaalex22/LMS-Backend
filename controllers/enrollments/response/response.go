package response

import (
	"backend/business/enrollments"
	_courseReponse "backend/controllers/courses/response"
	_studentReponse "backend/controllers/student/response"
	"time"
)

type EnrollmentsResponse struct {
	StudentId uint                           `json:"studentid"`
	CourseId  uint                           `json:"courseid"`
	Rating    int                            `json:"rating"`
	Review    string                         `json:"review"`
	Student   _studentReponse.StudentProfile `json:"student"`
	Course    _courseReponse.AddResponse     `json:"course"`
	CreateAt  time.Time                      `json:"createdAt"`
	UpdateAt  time.Time                      `json:"updateAt"`
}

func FromDomain(domain enrollments.Domain) EnrollmentsResponse {
	return EnrollmentsResponse{
		StudentId: domain.StudentId,
		Student:   _studentReponse.FromDomainProfile(domain.Student),
		CourseId:  domain.CourseId,
		Course:    _courseReponse.FromDomain(domain.Course),
		Rating:    domain.Rating,
		Review:    domain.Review,
		CreateAt:  domain.CreateAt,
		UpdateAt:  domain.UpdateAt,
	}
}

func FromDomainList(domain []enrollments.Domain) []EnrollmentsResponse {
	list := []EnrollmentsResponse{}
	for _, value := range domain {
		list = append(list, FromDomain(value))
	}
	return list
}
