package request

import "backend/business/enrollments"

type EnrollAdd struct {
	Student_Id uint `json:"studentid"`
	Course_Id  uint `json:"courseid"`
}

func (elm *EnrollAdd) ToDomain() enrollments.Domain {
	return enrollments.Domain{
		Student_Id: elm.Student_Id,
		Course_Id:  elm.Course_Id,
	}
}
