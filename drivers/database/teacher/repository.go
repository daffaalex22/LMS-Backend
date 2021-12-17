package teacher

import (
	_middleware "backend/app/middleware"
	"backend/business/teacher"
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type TeacherRepository struct {
	db   *gorm.DB
	jwts *_middleware.ConfigsJWT
}

func NewTeacherRepository(gormDb *gorm.DB, configsJWT *_middleware.ConfigsJWT) teacher.TeacherRepoInterface {
	return &TeacherRepository{
		db:   gormDb,
		jwts: configsJWT,
	}
}
func (repo *TeacherRepository) TeacherRegister(domain *teacher.Domain, ctx context.Context) (teacher.Domain, error) {
	tchDb := FromDomain(*domain)
	err := repo.db.Create(&tchDb)
	if err.Error != nil {
		return teacher.Domain{}, err.Error
	}
	return tchDb.ToDomain(), nil
}

func (repo *TeacherRepository) TeacherLogin(domain teacher.Domain, ctx context.Context) (teacher.Domain, error) {
	tchDb := FromDomain(domain)

	err := repo.db.Where("email = ?", tchDb.Email).First(&tchDb).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return teacher.Domain{}, errors.New("email not found")
		}
		return teacher.Domain{}, errors.New("error in database")
	}
	JwtCustomClaimsTch := _middleware.JwtCustomClaimsTch{
		uint(tchDb.Id),
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}
	tchResponse := tchDb.ToDomain()
	tchResponse.Token = repo.jwts.GenerateJWTtch(JwtCustomClaimsTch)
	return tchResponse, nil
}

func (repo *TeacherRepository) TeacherGetProfile(ctx context.Context, id uint) (teacher.Domain, error) {
	var tchDb Teacher
	err := repo.db.Find(&tchDb, "id = ?", id)
	if err.Error != nil {
		return teacher.Domain{}, err.Error
	}
	return tchDb.ToDomain(), nil
}

func (repo *TeacherRepository) TeacherUpdate(ctx context.Context, domain teacher.Domain, id uint) (teacher.Domain, error) {
	var tchDb Teacher
	err := repo.db.Find(&tchDb, id)
	if err.Error != nil {
		return teacher.Domain{}, err.Error
	}
	if err.RowsAffected == 0 {
		return tchDb.ToDomain(), errors.New("record not found")
	}
	err1 := repo.db.Where("email = ?", domain.Email).First(&tchDb)
	if err1.RowsAffected != 0 {
		return teacher.Domain{}, errors.New("email has applied")
	}
	err2 := repo.db.Model(&tchDb).Updates(FromDomain(domain))
	if err2.Error != nil {
		return teacher.Domain{}, err2.Error
	}
	return tchDb.ToDomain(), nil
}
