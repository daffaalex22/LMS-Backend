package course

import (
	"backend/business/course"
	"context"
	"time"

	"gorm.io/gorm"
)

type MysqlCoursesRepository struct {
	Conn *gorm.DB
}

func NewMysqlCategoryRepository(conn *gorm.DB) course.Repository {
	return &MysqlCoursesRepository{
		Conn: conn,
	}
}

func (rep *MysqlCoursesRepository) Create(ctx context.Context, domain course.Domain) (course.Domain, error) {
	newCourse := FromDomain(domain)
	newCourse.CreatedAt = time.Now()

	//fire to databases
	resultAdd := rep.Conn.Create(&newCourse)
	if resultAdd.Error != nil {
		return course.Domain{}, resultAdd.Error
	}
	return newCourse.ToDomain(), nil
}
