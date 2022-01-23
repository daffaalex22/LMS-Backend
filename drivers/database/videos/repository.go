package videos

import (
	"backend/business/modules"
	"backend/business/videos"
	"backend/helper/err"
	"context"

	"gorm.io/gorm"
)

type VideosRepository struct {
	db *gorm.DB
}

func NewVideosRepository(gormDb *gorm.DB) videos.VideosRepoInterface {
	return &VideosRepository{
		db: gormDb,
	}
}

func (repo *VideosRepository) VideosAdd(ctx context.Context, domain videos.Domain) (videos.Domain, error) {
	newVideos := FromDomain(domain)
	resultAdd := repo.db.Create(&newVideos)
	if resultAdd.Error != nil {
		return videos.Domain{}, resultAdd.Error
	}
	return newVideos.ToDomain(), nil
}

func (repo *VideosRepository) VideosUpdate(ctx context.Context, domain videos.Domain, id uint) (videos.Domain, error) {
	var targetTable Videos
	newVideos := FromDomain(domain)
	resultUpdate := repo.db.Model(&targetTable).Where("id = ?", id).Updates(newVideos)
	if resultUpdate.Error != nil {
		return videos.Domain{}, resultUpdate.Error
	}
	return newVideos.ToDomain(), nil
}

func (repo *VideosRepository) VideosGetByModuleId(ctx context.Context, moduleId uint) ([]videos.Domain, error) {
	var targetTable []Videos
	resultGet := repo.db.Preload("Module").Where("module_id = ?", moduleId).Find(&targetTable)
	if resultGet.Error != nil {
		return []videos.Domain{}, resultGet.Error
	}
	if resultGet.RowsAffected == 0 {
		return ToDomainList(targetTable), err.ErrNotFound
	}
	return ToDomainList(targetTable), nil
}

func (repo *VideosRepository) VideosDelete(ctx context.Context, id uint) error {
	var targetTable Videos
	result := repo.db.Delete(&targetTable, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return err.ErrNotFound
	}
	return nil
}

func (repo *VideosRepository) VideosGetById(ctx context.Context, id uint) (videos.Domain, error) {
	var targetTable Videos
	result := repo.db.Preload("Module").First(&targetTable, id)
	if result.Error != nil {
		return videos.Domain{}, result.Error
	}
	if result.RowsAffected == 0 {
		return videos.Domain{}, err.ErrNotFound
	}
	return targetTable.ToDomain(), nil
}

func (repo *VideosRepository) CheckModule(ctx context.Context, id uint) (modules.Domain, error) {
	var targetTable Videos
	checkModule := repo.db.Table("modules").Where("id = ?", id).Find(&targetTable.Module)
	if checkModule.RowsAffected == 0 {
		return modules.Domain{}, err.ErrCourseNotFound
	}
	return targetTable.Module.ToDomain(), nil
}
