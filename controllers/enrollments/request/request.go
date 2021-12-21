package request

import "backend/business/enrollments"

type EnrollAdd struct {
	StudentId uint `json:"studentid"`
	CourseId  uint `json:"courseid"`
}

func (elm *EnrollAdd) ToDomain() enrollments.Domain {
	return enrollments.Domain{
		StudentId: elm.StudentId,
		CourseId:  elm.CourseId,
	}
}
