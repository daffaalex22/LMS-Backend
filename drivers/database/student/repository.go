package student

import (
	_middleware "backend/app/middleware"
	"backend/business/student"
	"context"
	"errors"
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
	err := repo.db.Where("email = ?", stdDb.Email).First(&stdDb)
	if err.RowsAffected != 0 {
		return student.Domain{}, errors.New("email has applied")
	}
	err1 := repo.db.Create(&stdDb)
	if err1.Error != nil {
		return student.Domain{}, errors.New("email has applied")
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
	JwtCustomClaims := _middleware.JwtCustomClaims{
		uint(stdDb.Id),
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}
	stdResponse := stdDb.ToDomain()
	stdResponse.Token = repo.jwt.GenerateJWT(JwtCustomClaims)
	return stdResponse, nil
}
func (repo *StudentRepository) GetProfile(ctx context.Context, id uint) (student.Domain, error) {
	var stdDb Student
	err := repo.db.Find(&stdDb, "id = ?", id)
	if err.Error != nil {
		return student.Domain{}, err.Error
	}
	return stdDb.ToDomain(), nil
}

func (repo *StudentRepository) StudentUpdate(ctx context.Context, domain student.Domain, id uint) (student.Domain, error) {
	var stdDb Student
	err := repo.db.Find(&stdDb, id)
	if err.Error != nil {
		return student.Domain{}, err.Error
	}
	if err.RowsAffected == 0 {
		return stdDb.ToDomain(), errors.New("record not found")
	}
	err1 := repo.db.Where("email = ?", domain.Email).First(&stdDb)
	if err1.RowsAffected != 0 {
		return student.Domain{}, errors.New("email has applied")
	}
	err2 := repo.db.Model(&stdDb).Updates(FromDomain(domain))
	if err2.Error != nil {
		return student.Domain{}, err2.Error
	}
	return stdDb.ToDomain(), nil
}
