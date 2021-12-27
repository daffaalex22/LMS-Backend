package request

import "backend/business/enrollments"

type EnrollAdd struct {
	StudentId uint `json:"studentId"`
	CourseId  uint `json:"courseId"`
}

type EnrollUpdate struct {
	StudentId uint   `json:"studentId"`
	CourseId  uint   `json:"courseId"`
	Rating    int    `json:"rating"`
	Review    string `json:"review"`
}

func (elm *EnrollAdd) ToDomain() enrollments.Domain {
	return enrollments.Domain{
		StudentId: elm.StudentId,
		CourseId:  elm.CourseId,
	}
}

func (elm *EnrollUpdate) ToDomain() enrollments.Domain {
	return enrollments.Domain{
		StudentId: elm.StudentId,
		CourseId:  elm.CourseId,
		Rating:    elm.Rating,
		Review:    elm.Review,
	}
}
