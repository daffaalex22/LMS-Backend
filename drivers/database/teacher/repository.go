package teacher

import (
	_middleware "backend/app/middleware"
	"backend/business/teacher"
	"backend/helper/err"
	"backend/helper/password"
	"context"
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
	err1 := repo.db.Where("email = ?", tchDb.Email).First(&tchDb)
	if err1.RowsAffected != 0 {
		return teacher.Domain{}, err.ErrEmailHasApplied
	}
	result := repo.db.Create(&tchDb)
	if result.Error != nil {
		return teacher.Domain{}, result.Error
	}
	return tchDb.ToDomain(), nil
}

func (repo *TeacherRepository) TeacherLogin(domain teacher.Domain, ctx context.Context) (teacher.Domain, error) {
	var teachers Teacher
	tchDb := FromDomain(domain)
	err1 := repo.db.Where("email = ?", tchDb.Email).First(&teachers).Error
	if err1 != nil {
		if err1 == gorm.ErrRecordNotFound {
			return teacher.Domain{}, err.ErrEmailNotExist
		}
		return teacher.Domain{}, err.ErrTeacherNotFound
	}
	if password.CheckSamePassword(tchDb.Password, teachers.Password) {
		JwtCustomClaimsTch := _middleware.JwtCustomClaimsTch{
			uint(teachers.Id),
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			},
		}
		tchResponse := teachers.ToDomain()
		tchResponse.Token = repo.jwts.GenerateJWTtch(JwtCustomClaimsTch)
		return tchResponse, nil
	}
	return teacher.Domain{}, err.ErrWrongPassword
}

func (repo *TeacherRepository) TeacherGetProfile(ctx context.Context, id uint) (teacher.Domain, error) {
	var tchDb Teacher
	result := repo.db.Find(&tchDb, "id = ?", id)
	if result.Error != nil {
		return teacher.Domain{}, result.Error
	}
	if result.RowsAffected == 0 {
		return tchDb.ToDomain(), err.ErrNotFound
	}
	return tchDb.ToDomain(), nil
}

func (repo *TeacherRepository) TeacherUpdate(ctx context.Context, domain teacher.Domain, id uint) (teacher.Domain, error) {
	var tchDb Teacher
	err1 := repo.db.Find(&tchDb, id)
	if err1.Error != nil {
		return teacher.Domain{}, err1.Error
	}
	if err1.RowsAffected == 0 {
		return tchDb.ToDomain(), err.ErrNotFound
	}
	err2 := repo.db.Where("email = ?", domain.Email).First(&tchDb)
	if err2.RowsAffected != 0 {
		return teacher.Domain{}, err.ErrEmailHasApplied
	}
	result := repo.db.Model(&tchDb).Updates(FromDomain(domain))
	if result.Error != nil {
		return teacher.Domain{}, result.Error
	}
	return tchDb.ToDomain(), nil
}
