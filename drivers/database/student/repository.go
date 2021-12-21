package student

import (
	_middleware "backend/app/middleware"
	"backend/business/student"
	"backend/helper/err"
	"backend/helper/password"
	"context"
	"time"

	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type StudentRepository struct {
	db  *gorm.DB
	jwt *_middleware.ConfigJWT
}

func NewStudentRepository(gormDb *gorm.DB, configsJWT *_middleware.ConfigJWT) student.StudentRepoInterface {
	return &StudentRepository{
		db:  gormDb,
		jwt: configsJWT,
	}
}
func (repo *StudentRepository) Register(domain *student.Domain, ctx context.Context) (student.Domain, error) {
	stdDb := FromDomain(*domain)
	err1 := repo.db.Where("email = ?", stdDb.Email).First(&stdDb)
	if err1.RowsAffected != 0 {
		return student.Domain{}, err.ErrEmailHasApplied
	}
	result := repo.db.Create(&stdDb)
	if result.Error != nil {
		return student.Domain{}, result.Error
	}
	return stdDb.ToDomain(), nil

}
func (repo *StudentRepository) Login(domain student.Domain, ctx context.Context) (student.Domain, error) {
	var students Student
	stdDb := FromDomain(domain)
	result := repo.db.Where("email = ?", stdDb.Email).First(&students).Error
	if result != nil {
		if result == gorm.ErrRecordNotFound {
			return student.Domain{}, err.ErrEmailNotExist
		}
		return student.Domain{}, err.ErrStudentNotFound
	}
	if password.CheckSamePassword(stdDb.Password, students.Password) {
		JwtCustomClaims := _middleware.JwtCustomClaims{
			uint(students.Id),
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			},
		}
		stdResponse := students.ToDomain()
		stdResponse.Token = repo.jwt.GenerateJWT(JwtCustomClaims)
		return stdResponse, nil

	}
	return student.Domain{}, err.ErrWrongPassword
}
func (repo *StudentRepository) GetProfile(ctx context.Context, id uint) (student.Domain, error) {
	var stdDb Student
	result := repo.db.Find(&stdDb, "id = ?", id)
	if result.Error != nil {
		return student.Domain{}, result.Error
	}
	if result.RowsAffected == 0 {
		return stdDb.ToDomain(), err.ErrNotFound
	}
	return stdDb.ToDomain(), nil
}

func (repo *StudentRepository) StudentUpdate(ctx context.Context, domain student.Domain, id uint) (student.Domain, error) {
	var stdDb Student
	err1 := repo.db.Find(&stdDb, id)
	if err1.Error != nil {
		return student.Domain{}, err1.Error
	}
	if err1.RowsAffected == 0 {
		return stdDb.ToDomain(), err.ErrNotFound
	}
	err2 := repo.db.Where("email = ?", domain.Email).First(&stdDb)
	if err2.RowsAffected != 0 {
		return student.Domain{}, err.ErrEmailHasApplied
	}
	result := repo.db.Model(&stdDb).Updates(FromDomain(domain))
	if result.Error != nil {
		return student.Domain{}, result.Error
	}
	return stdDb.ToDomain(), nil
}
