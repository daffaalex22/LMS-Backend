package modules

import (
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
