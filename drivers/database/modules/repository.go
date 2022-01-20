package modules

import (
	"backend/business/course"
	"backend/business/modules"
	"backend/helper/err"
	"context"

	"gorm.io/gorm"
)

type ModulesRepository struct {
	db *gorm.DB
}

func NewModulesRepository(gormDb *gorm.DB) modules.ModulesRepoInterface {
	return &ModulesRepository{
		db: gormDb,
	}
}
func (repo *ModulesRepository) ModulesGetAll(ctx context.Context) ([]modules.Domain, error) {
	var mdsDb []Modules
	result := repo.db.Preload("Course").Find(&mdsDb)
	if result.RowsAffected == 0 {
		return []modules.Domain{}, err.ErrModulesNotFound
	}

	if result.Error != nil {
		return []modules.Domain{}, result.Error
	}
	return ToDomainList(mdsDb), nil
}

func (repo *ModulesRepository) ModulesAdd(ctx context.Context, domain modules.Domain) (modules.Domain, error) {
	newModules := FromDomain(domain)
	resultAdd := repo.db.Create(&newModules)
	if resultAdd.Error != nil {
		return modules.Domain{}, resultAdd.Error
	}
	return newModules.ToDomain(), nil
}

func (repo *ModulesRepository) ModulesUpdate(ctx context.Context, domain modules.Domain, id uint) (modules.Domain, error) {
	var targetTable Modules
	newModules := FromDomain(domain)
	resultUpdate := repo.db.Model(&targetTable).Where("id = ?", id).Updates(newModules)
	if resultUpdate.Error != nil {
		return modules.Domain{}, resultUpdate.Error
	}
	return newModules.ToDomain(), nil
}

func (repo *ModulesRepository) ModulesGetByCourseId(ctx context.Context, courseId uint) ([]modules.Domain, error) {
	var targetTable []Modules
	resultGet := repo.db.Preload("Course").Where("course_id = ?", courseId).Find(&targetTable)
	if resultGet.Error != nil {
		return []modules.Domain{}, resultGet.Error
	}
	if resultGet.RowsAffected == 0 {
		return ToDomainList(targetTable), err.ErrNotFound
	}
	return ToDomainList(targetTable), nil
}

func (repo *ModulesRepository) ModulesDelete(ctx context.Context, id uint) error {
	var targetTable Modules
	result := repo.db.Delete(&targetTable, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return err.ErrNotFound
	}
	return nil
}

func (repo *ModulesRepository) ModulesGetById(ctx context.Context, id uint) (modules.Domain, error) {
	var targetTable Modules
	result := repo.db.Preload("Course").First(&targetTable, id)
	if result.Error != nil {
		return modules.Domain{}, result.Error
	}
	if result.RowsAffected == 0 {
		return modules.Domain{}, err.ErrNotFound
	}
	return targetTable.ToDomain(), nil
}

func (repo *ModulesRepository) CheckCourse(ctx context.Context, id uint) (course.Domain, error) {
	var targetTable Modules
	checkCourse := repo.db.Table("courses").Where("id = ?", id).Find(&targetTable.Course)
	if checkCourse.RowsAffected == 0 {
		return course.Domain{}, err.ErrCourseNotFound
	}
	return targetTable.Course.ToDomain(), nil
}
