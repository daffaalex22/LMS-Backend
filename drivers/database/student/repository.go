package student

import (
	"backend/business/student"
	"context"
	"errors"

	"gorm.io/gorm"
)

type StudentRepository struct {
	db *gorm.DB
}

func NewStudentRepository(gormDb *gorm.DB) student.StudentRepoInterface {
	return &StudentRepository{
		db: gormDb,
	}
}
func (repo *StudentRepository) Register(domain *student.Domain, ctx context.Context) (student.Domain, error) {
	stdDb := FromDomain(*domain)
	err := repo.db.Create(&stdDb)
	if err.Error != nil {
		return student.Domain{}, err.Error
	}
	return stdDb.ToDomain(), nil

}
func (repo *StudentRepository) Login(domain student.Domain, ctx context.Context) (student.Domain, error) {
	stdDb := FromDomain(domain)

	err := repo.db.Where("email = ?", stdDb.Email).First(&stdDb).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return student.Domain{}, errors.New("email not found")
		}
		return student.Domain{}, errors.New("error in database")
	}
	return stdDb.ToDomain(), nil
}
func (repo *StudentRepository) GetProfile(ctx context.Context, id uint) (student.Domain, error) {
	var stdDb Student
	err := repo.db.Find(&stdDb, "id = ?", id)
	if err.Error != nil {
		return student.Domain{}, err.Error
	}
	return stdDb.ToDomain(), nil
}
