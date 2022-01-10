package requests

import (
	"backend/business/course"
	"backend/business/requests"
	"backend/business/student"
	_coursedb "backend/drivers/database/course"
	"backend/helper/err"
	"context"

	"gorm.io/gorm"
)

type RequestsRepository struct {
	db *gorm.DB
}

func NewRequestsRepository(gormDb *gorm.DB) requests.RequestsRepoInterface {
	return &RequestsRepository{
		db: gormDb,
	}
}
func (repo *RequestsRepository) RequestsGetAll(ctx context.Context) ([]requests.Domain, error) {
	var elmDb []Requests
	err1 := repo.db.Preload("Student").Preload("Course").Preload("Type").Find(&elmDb)
	if err1.RowsAffected == 0 {
		return []requests.Domain{}, err.ErrRequestsNotFound
	}

	if err1.Error != nil {
		return []requests.Domain{}, err1.Error
	}
	return ToDomainList(elmDb), nil
}

func (repo *RequestsRepository) RequestGetById(ctx context.Context, id uint) (requests.Domain, error) {
	var response Requests

	err := repo.db.Preload("Student").Preload("Course").Preload("Type").Where("id = ?", id).First(&response)
	if err.Error != nil {
		return requests.Domain{}, err.Error
	}

	return response.ToDomain(), nil
}

func (repo *RequestsRepository) RequestsGetByStudentId(ctx context.Context, studentId uint) ([]requests.Domain, error) {
	var response []Requests

	err := repo.db.Preload("Student").Preload("Course").Preload("Type").Where("student_id = ?", studentId).Find(&response)
	if err.Error != nil {
		return []requests.Domain{}, err.Error
	}

	return ToDomainList(response), nil
}

func (repo *RequestsRepository) RequestsAdd(ctx context.Context, domain requests.Domain) (requests.Domain, error) {
	newRequests := FromDomain(domain)

	//fire to databases
	resultAdd := repo.db.Create(&newRequests)
	if resultAdd.Error != nil {
		return requests.Domain{}, resultAdd.Error
	}
	return newRequests.ToDomain(), nil
}

func (repo *RequestsRepository) RequestsUpdate(ctx context.Context, domain requests.Domain, id uint) (requests.Domain, error) {
	var response Requests
	newRequests := FromDomain(domain)
	resultUpdate := repo.db.Model(&response).Where("id = ?", id).Updates(newRequests)
	if resultUpdate.Error != nil {
		return requests.Domain{}, resultUpdate.Error
	}
	err := repo.db.Where("id = ?", id).First(&response)
	if err.Error != nil {
		return requests.Domain{}, err.Error
	}
	return response.ToDomain(), nil
}

func (repo *RequestsRepository) RequestsGetByCourseId(ctx context.Context, courseId uint) ([]requests.Domain, error) {
	var targetTable []Requests
	resultGet := repo.db.Preload("Student").Preload("Course").Preload("Type").Where("course_id = ?", courseId).Find(&targetTable)
	if resultGet.Error != nil {
		return []requests.Domain{}, resultGet.Error
	}
	if resultGet.RowsAffected == 0 {
		return ToDomainList(targetTable), err.ErrNotFound
	}
	return ToDomainList(targetTable), nil
}

func (repo *RequestsRepository) CheckStudent(ctx context.Context, id uint) (student.Domain, error) {
	var targetTable Requests
	checkStudent := repo.db.Table("students").Where("id = ?", id).Find(&targetTable.Student)
	if checkStudent.RowsAffected == 0 {
		return student.Domain{}, err.ErrStudentNotFound
	}
	return targetTable.Student.ToDomain(), nil
}

func (repo *RequestsRepository) CheckCourse(ctx context.Context, id uint) (course.Domain, error) {
	var targetTable Requests
	checkCourse := repo.db.Table("courses").Where("id = ?", id).Find(&targetTable.Course)
	if checkCourse.RowsAffected == 0 {
		return course.Domain{}, err.ErrCourseNotFound
	}
	return targetTable.Course.ToDomain(), nil
}

func (repo *RequestsRepository) GetCoursesByTeacherId(ctx context.Context, teacherId uint) ([]course.Domain, error) {
	var targetTable []_coursedb.Course
	checkCourse := repo.db.Table("courses").Where("teacher_id = ?", teacherId).Find(&targetTable)
	if checkCourse.RowsAffected == 0 {
		return []course.Domain{}, err.ErrCourseNotFound
	}
	return _coursedb.ToDomainList(targetTable), nil
}

func (repo *RequestsRepository) RequestsGetByCourseIds(ctx context.Context, courseIds []uint) ([]requests.Domain, error) {
	var targetTable []Requests
	checkRequests := repo.db.Preload("Student").Preload("Course").Preload("Type").Where("course_id IN ?", courseIds).Find(&targetTable)
	if checkRequests.RowsAffected == 0 {
		return []requests.Domain{}, err.ErrRequestsNotFound
	}
	return ToDomainList(targetTable), nil
}
