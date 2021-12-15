package categories

import (
	"backend/business/categories"
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
	err := rep.Conn.Find(&allCategories).Error
	if err != nil {
		return []categories.Domain{}, err
	}
	convert := ToDomainList(allCategories)
	return convert, nil
}
