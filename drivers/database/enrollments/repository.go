package enrollments

import (
	"backend/business/enrollments"
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

	checkStudent := repo.db.Table("students").Where("id = ?", newEnroll.StudentId).Find(&newEnroll.Student)
	if checkStudent.RowsAffected == 0 {
		return enrollments.Domain{}, err.ErrStudentNotFound
	}

	checkCourse := repo.db.Table("courses").Where("id = ?", newEnroll.CourseId).Find(&newEnroll.Course)
	if checkCourse.RowsAffected == 0 {
		return enrollments.Domain{}, err.ErrCourseNotFound
	}

	//fire to databases
	resultAdd := repo.db.Create(&newEnroll)
	if resultAdd.Error != nil {
		return enrollments.Domain{}, resultAdd.Error
	}
	return newEnroll.ToDomain(), nil
}
