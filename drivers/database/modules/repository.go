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

	//fire to databases
	resultAdd := repo.db.Create(&newModules)
	if resultAdd.Error != nil {
		return modules.Domain{}, resultAdd.Error
	}
	return newModules.ToDomain(), nil
}

func (repo *ModulesRepository) CheckCourse(ctx context.Context, id uint) (course.Domain, error) {
	var targetTable Modules
	checkCourse := repo.db.Table("courses").Where("id = ?", id).Find(&targetTable.Course)
	if checkCourse.RowsAffected == 0 {
		return course.Domain{}, err.ErrCourseNotFound
	}
	return targetTable.Course.ToDomain(), nil
}
