package enrollments

import (
	"backend/business/enrollments"
	"backend/helper/err"
	"context"

	"gorm.io/gorm"
)

type EnrollmentsRepository struct {
	db *gorm.DB
}

func NewEnrollmentsRepository(gormDb *gorm.DB) enrollments.EnrollmentsRepoInterface {
	return &EnrollmentsRepository{
		db: gormDb,
	}
}
func (repo *EnrollmentsRepository) EnrollmentGetAll(ctx context.Context) ([]enrollments.Domain, error) {
	var elmDb []Enrollments
	err1 := repo.db.Find(&elmDb)
	if err1.RowsAffected == 0 {
		return []enrollments.Domain{}, err.ErrNotFound
	}

	if err1.Error != nil {
		return []enrollments.Domain{}, err1.Error
	}
	return AllEnrollments(elmDb), nil
}
