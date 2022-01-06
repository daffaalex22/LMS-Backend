package difficulties

import (
	"backend/business/difficulties"
	"backend/helper/err"
	"context"

	"gorm.io/gorm"
)

type MysqlDifficultyRepository struct {
	db *gorm.DB
}

func NewMysqlDifficultyRepository(dbs *gorm.DB) difficulties.Repository {
	return &MysqlDifficultyRepository{
		db: dbs,
	}
}

func (rep *MysqlDifficultyRepository) GetAll(ctx context.Context) ([]difficulties.Domain, error) {
	var allDifficulties []Difficulty
	result := rep.db.Find(&allDifficulties)

	if result.RowsAffected == 0 {
		return []difficulties.Domain{}, err.ErrDifficultyNotFound
	}

	if result.Error != nil {
		return []difficulties.Domain{}, result.Error
	}
	convert := ToDomainList(allDifficulties)
	return convert, nil
}
