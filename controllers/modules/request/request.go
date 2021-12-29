package request

import "backend/business/modules"

type ModulesAdd struct {
	CourseId uint   `json:"courseid"`
	Title    string `json:"title"`
	Order    int    `json:"order"`
}
type ModulesUpdate struct {
	CourseId uint   `json:"courseid"`
	Title    string `json:"title"`
	Order    int    `json:"order"`
}

func (mds *ModulesAdd) ToDomain() modules.Domain {
	return modules.Domain{
		CourseId: mds.CourseId,
		Title:    mds.Title,
		Order:    mds.Order,
	}
}
func (mds *ModulesUpdate) ToDomain() modules.Domain {
	return modules.Domain{
		CourseId: mds.CourseId,
		Title:    mds.Title,
		Order:    mds.Order,
	}
}
