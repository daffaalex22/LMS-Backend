package attachments

import (
	"backend/business/attachments"
	"backend/business/modules"
	"backend/drivers/database/readings"
	"backend/drivers/database/videos"
	"backend/helper/err"
	"context"

	"gorm.io/gorm"
)

type AttachmentsRepository struct {
	db *gorm.DB
}

func NewAttachmentsRepository(gormDb *gorm.DB) attachments.AttachmentsRepoInterface {
	return &AttachmentsRepository{
		db: gormDb,
	}
}

func (repo *AttachmentsRepository) AttachmentsAdd(ctx context.Context, domain attachments.Domain) (attachments.Domain, error) {
	newAttachments := FromDomain(domain)
	resultAdd := repo.db.Create(&newAttachments)
	if resultAdd.Error != nil {
		return attachments.Domain{}, resultAdd.Error
	}
	return newAttachments.ToDomain(), nil
}

func (repo *AttachmentsRepository) AttachmentsUpdate(ctx context.Context, domain attachments.Domain, id uint) (attachments.Domain, error) {
	var targetTable Attachments
	newAttachments := FromDomain(domain)
	resultUpdate := repo.db.Model(&targetTable).Where("id = ?", id).Updates(newAttachments)
	if resultUpdate.Error != nil {
		return attachments.Domain{}, resultUpdate.Error
	}
	return newAttachments.ToDomain(), nil
}

func (repo *AttachmentsRepository) AttachmentsGetById(ctx context.Context, id uint) (attachments.Domain, error) {
	var targetTable Attachments
	resultGet := repo.db.Where("id = ?", id).Find(&targetTable)
	if resultGet.Error != nil {
		return attachments.Domain{}, resultGet.Error
	}
	if resultGet.RowsAffected == 0 {
		return attachments.Domain{}, err.ErrNotFound
	}
	return targetTable.ToDomain(), nil
}

func (repo *AttachmentsRepository) AttachmentsDelete(ctx context.Context, id uint) error {
	var targetTable Attachments
	result := repo.db.Delete(&targetTable, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return err.ErrNotFound
	}
	return nil
}

func (repo *AttachmentsRepository) CheckContent(ctx context.Context, contentType string, id uint) (interface{}, error) {
	var targetTable Attachments
	checkModule := repo.db.Table("modules").Where("id = ?", id).Find(&targetTable)
	if checkModule.RowsAffected == 0 {
		return modules.Domain{}, err.ErrCourseNotFound
	}
	return targetTable.ToDomain(), nil
}

func (repo *AttachmentsRepository) CheckVideos(ctx context.Context, videoId uint) error {
	var targetTable videos.Videos
	checkContent := repo.db.Table("videos").Where("id = ?", videoId).Find(&targetTable)
	if checkContent.Error != nil {
		return err.ErrVideosNotFound
	}
	return nil
}

func (repo *AttachmentsRepository) CheckReadings(ctx context.Context, readingId uint) error {
	var targetTable readings.Readings
	checkContent := repo.db.Table("readings").Where("id = ?", readingId).Find(&targetTable)
	if checkContent.Error != nil {
		return err.ErrReadingsNotFound
	}
	return nil
}
