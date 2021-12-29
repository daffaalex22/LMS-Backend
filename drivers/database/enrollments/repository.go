package enrollments

import (
	"backend/business/course"
	"backend/business/enrollments"
	"backend/business/student"
	"backend/helper/err"
	"context"

	"gorm.io/gorm"
)

type EnrollmentsRepository struct {
	db *gorm.DB
}

func NewEnrollmentsRepository(gormDb *gorm.DB) enrollments.EnrollmentsRepoInterface {
	return &EnrollmentsRepository{
		db: gormDb,
	}
}
func (repo *EnrollmentsRepository) EnrollmentGetAll(ctx context.Context) ([]enrollments.Domain, error) {
	var elmDb []Enrollments
	err1 := repo.db.Preload("Student").Preload("Course").Find(&elmDb)
	if err1.RowsAffected == 0 {
		return []enrollments.Domain{}, err.ErrEnrollNotFound
	}

	if err1.Error != nil {
		return []enrollments.Domain{}, err1.Error
	}
	return ToDomainList(elmDb), nil
}

func (repo *EnrollmentsRepository) EnrollmentAdd(ctx context.Context, domain enrollments.Domain) (enrollments.Domain, error) {
	newEnroll := FromDomain(domain)

	//fire to databases
	resultAdd := repo.db.Create(&newEnroll)
	if resultAdd.Error != nil {
		return enrollments.Domain{}, resultAdd.Error
	}
	return newEnroll.ToDomain(), nil
}

func (repo *EnrollmentsRepository) EnrollUpdate(ctx context.Context, domain enrollments.Domain, studentId uint, courseId uint) (enrollments.Domain, error) {
	var targetTable Enrollments
	newEnroll := FromDomain(domain)
	resultUpdate := repo.db.Model(&targetTable).Where("student_id = ? AND course_id = ?", studentId, courseId).Updates(newEnroll)
	if resultUpdate.Error != nil {
		return enrollments.Domain{}, resultUpdate.Error
	}
	return newEnroll.ToDomain(), nil
}

func (repo *EnrollmentsRepository) EnrollGetByCourseId(ctx context.Context, courseId uint) ([]enrollments.Domain, error) {
	var targetTable []Enrollments
	resultGet := repo.db.Preload("Student").Preload("Course").Where("course_id = ?", courseId).Find(&targetTable)
	if resultGet.Error != nil {
		return []enrollments.Domain{}, resultGet.Error
	}
	if resultGet.RowsAffected == 0 {
		return ToDomainList(targetTable), err.ErrNotFound
	}
	return ToDomainList(targetTable), nil
}

func (repo *EnrollmentsRepository) CheckStudent(ctx context.Context, id uint) (student.Domain, error) {
	var targetTable Enrollments
	checkStudent := repo.db.Table("students").Where("id = ?", id).Find(&targetTable.Student)
	if checkStudent.RowsAffected == 0 {
		return student.Domain{}, err.ErrStudentNotFound
	}
	return targetTable.Student.ToDomain(), nil
}

func (repo *EnrollmentsRepository) CheckCourse(ctx context.Context, id uint) (course.Domain, error) {
	var targetTable Enrollments
	checkCourse := repo.db.Table("courses").Where("id = ?", id).Find(&targetTable.Course)
	if checkCourse.RowsAffected == 0 {
		return course.Domain{}, err.ErrCourseNotFound
	}
	return targetTable.Course.ToDomain(), nil
}
