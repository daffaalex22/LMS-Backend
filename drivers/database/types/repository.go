package types

import (
	"backend/business/types"
	"backend/helper/err"
	"context"

	"gorm.io/gorm"
)

type MysqlTypeRepository struct {
	db *gorm.DB
}

func NewMysqlTypeRepository(dbs *gorm.DB) types.Repository {
	return &MysqlTypeRepository{
		db: dbs,
	}
}

func (rep *MysqlTypeRepository) GetAll(ctx context.Context) ([]types.Domain, error) {
	var allTypes []Types
	result := rep.db.Find(&allTypes)

	if result.RowsAffected == 0 {
		return []types.Domain{}, err.ErrTypeNotFound
	}

	if result.Error != nil {
		return []types.Domain{}, result.Error
	}
	convert := ToDomainList(allTypes)
	return convert, nil
}
