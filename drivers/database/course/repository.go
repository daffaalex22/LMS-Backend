package course

import (
	"backend/business/course"
	"backend/helper/err"
	"backend/helper/konversi"
	"context"
	"time"

	"gorm.io/gorm"
)

type MysqlCoursesRepository struct {
	DB *gorm.DB
}

func NewMysqlCategoryRepository(db *gorm.DB) course.Repository {
	return &MysqlCoursesRepository{
		DB: db,
	}
}

func (rep *MysqlCoursesRepository) Create(ctx context.Context, domain course.Domain) (course.Domain, error) {
	newCourse := FromDomain(domain)
	newCourse.CreatedAt = time.Now()

	checkCategories := rep.DB.Table("categories").Where("id = ?", newCourse.CategoryId).Find(&newCourse.Category)
	if checkCategories.RowsAffected == 0 {
		return course.Domain{}, err.ErrCategoryNotFound
	}

	checkTeacher := rep.DB.Table("teachers").Where("id = ?", newCourse.TeacherId).Find(&newCourse.Teacher)
	if checkTeacher.RowsAffected == 0 {
		return course.Domain{}, err.ErrTeacherNotFound
	}

	//fire to databases
	resultAdd := rep.DB.Create(&newCourse)
	if resultAdd.Error != nil {
		return course.Domain{}, resultAdd.Error
	}
	return newCourse.ToDomain(), nil
}

func (rep *MysqlCoursesRepository) GetAll(ctx context.Context) ([]course.Domain, error) {
	//Get all data from databases
	listCourses := []Course{}
	resultAdd := rep.DB.Preload("Category").Preload("Teacher").Find(&listCourses)
	if resultAdd.Error != nil {
		return []course.Domain{}, resultAdd.Error
	}

	if resultAdd.RowsAffected == 0 {
		return []course.Domain{}, err.ErrCourseNotFound
	}

	//convert from Repo to Domain List
	listDomain := ToDomainList(listCourses)
	return listDomain, nil
}

func (rep *MysqlCoursesRepository) Update(ctx context.Context, id string, domain course.Domain) (course.Domain, error) {
	var targetTable Course
	newCourse := FromDomain(domain)
	newCourse.UpdatedAt = time.Now()

	idUint, errConvert := konversi.StringToUint(id)
	if errConvert != nil {
		return course.Domain{}, errConvert
	}
	newCourse.Id = idUint

	checkCourse := rep.DB.Where("id = ?", newCourse.Id).Find(&targetTable)
	if checkCourse.RowsAffected == 0 {
		return course.Domain{}, err.ErrCourseNotFound
	}

	checkCategories := rep.DB.Table("categories").Where("id = ?", newCourse.CategoryId).Find(&newCourse.Category)
	if checkCategories.RowsAffected == 0 {
		return course.Domain{}, err.ErrCategoryNotFound
	}

	checkTeacher := rep.DB.Table("teachers").Where("id = ?", newCourse.TeacherId).Find(&newCourse.Teacher)
	if checkTeacher.RowsAffected == 0 {
		return course.Domain{}, err.ErrTeacherNotFound
	}

	//fire to databases
	resultUpdate := rep.DB.Model(&targetTable).Updates(newCourse)
	if resultUpdate.Error != nil {
		return course.Domain{}, resultUpdate.Error
	}
	return newCourse.ToDomain(), nil
}
