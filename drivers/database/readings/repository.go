package readings

import (
	"backend/business/modules"
	"backend/business/readings"
	"backend/helper/err"
	"context"

	"gorm.io/gorm"
)

type ReadingsRepository struct {
	db *gorm.DB
}

func NewReadingsRepository(gormDb *gorm.DB) readings.ReadingsRepoInterface {
	return &ReadingsRepository{
		db: gormDb,
	}
}

func (repo *ReadingsRepository) ReadingsGetById(ctx context.Context, id uint) (readings.Domain, error) {
	var targetTable Readings
	result := repo.db.Preload("Module").First(&targetTable, id)
	if result.Error != nil {
		return readings.Domain{}, result.Error
	}
	if result.RowsAffected == 0 {
		return readings.Domain{}, err.ErrNotFound
	}
	return targetTable.ToDomain(), nil
}

func (repo *ReadingsRepository) ReadingsAdd(ctx context.Context, domain readings.Domain) (readings.Domain, error) {
	newReadings := FromDomain(domain)
	resultAdd := repo.db.Create(&newReadings)
	if resultAdd.Error != nil {
		return readings.Domain{}, resultAdd.Error
	}
	return newReadings.ToDomain(), nil
}

func (repo *ReadingsRepository) ReadingsUpdate(ctx context.Context, domain readings.Domain, id uint) (readings.Domain, error) {
	var targetTable Readings
	newReadings := FromDomain(domain)
	resultUpdate := repo.db.Model(&targetTable).Where("id = ?", id).Updates(newReadings)
	if resultUpdate.Error != nil {
		return readings.Domain{}, resultUpdate.Error
	}
	return newReadings.ToDomain(), nil
}

func (repo *ReadingsRepository) ReadingsGetByModuleId(ctx context.Context, moduleId uint) ([]readings.Domain, error) {
	var targetTable []Readings
	resultGet := repo.db.Preload("Module").Where("module_id = ?", moduleId).Find(&targetTable)
	if resultGet.Error != nil {
		return []readings.Domain{}, resultGet.Error
	}
	if resultGet.RowsAffected == 0 {
		return ToDomainList(targetTable), err.ErrNotFound
	}
	return ToDomainList(targetTable), nil
}

func (repo *ReadingsRepository) ReadingsDelete(ctx context.Context, id uint) error {
	var targetTable Readings
	result := repo.db.Delete(&targetTable, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return err.ErrNotFound
	}
	return nil
}

func (repo *ReadingsRepository) CheckModule(ctx context.Context, id uint) (modules.Domain, error) {
	var targetTable Readings
	checkModule := repo.db.Table("modules").Where("id = ?", id).Find(&targetTable.Module)
	if checkModule.RowsAffected == 0 {
		return modules.Domain{}, err.ErrCourseNotFound
	}
	return targetTable.Module.ToDomain(), nil
}
