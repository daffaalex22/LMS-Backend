package categories

import (
	"backend/business/categories"
	"backend/helper/err"
	"context"

	"gorm.io/gorm"
)

type MysqlCategoryRepository struct {
	db *gorm.DB
}

func NewMysqlCategoryRepository(dbs *gorm.DB) categories.Repository {
	return &MysqlCategoryRepository{
		db: dbs,
	}
}

func (rep *MysqlCategoryRepository) GetAll(ctx context.Context) ([]categories.Domain, error) {
	var allCategories []Category
	result := rep.db.Find(&allCategories)

	if result.RowsAffected == 0 {
		return []categories.Domain{}, err.ErrCategoryNotFound
	}

	if result.Error != nil {
		return []categories.Domain{}, result.Error
	}
	convert := ToDomainList(allCategories)
	return convert, nil
}
