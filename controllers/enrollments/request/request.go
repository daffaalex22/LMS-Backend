package request

import "backend/business/enrollments"

type EnrollAdd struct {
	StudentId uint `json:"studentId"`
	CourseId  uint `json:"courseId"`
}

func (elm *EnrollAdd) ToDomain() enrollments.Domain {
	return enrollments.Domain{
		StudentId: elm.StudentId,
		CourseId:  elm.CourseId,
	}
}
