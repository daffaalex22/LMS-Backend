package categories

import (
	"backend/business/categories"
	"backend/helpers/err"
	"context"

	"gorm.io/gorm"
)

type MysqlCategoryRepository struct {
	Conn *gorm.DB
}

func NewMysqlCategoryRepository(conn *gorm.DB) categories.Repository {
	return &MysqlCategoryRepository{
		Conn: conn,
	}
}

func (rep *MysqlCategoryRepository) GetAll(ctx context.Context) ([]categories.Domain, error) {
	var allCategories []Category
	result := rep.Conn.Find(&allCategories)

	if result.RowsAffected == 0 {
		return []categories.Domain{}, err.ErrCategoryNotFound
	}

	if result.Error != nil {
		return []categories.Domain{}, result.Error
	}
	convert := ToDomainList(allCategories)
	return convert, nil
}
