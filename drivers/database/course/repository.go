package course

import (
	"backend/business/categories"
	"backend/business/course"
	"backend/business/difficulties"
	"backend/business/teacher"
	"backend/helper/err"
	"context"
	"time"

	"gorm.io/gorm"
)

type MysqlCoursesRepository struct {
	DB *gorm.DB
}

func NewMysqlCategoryRepository(db *gorm.DB) course.Repository {
	return &MysqlCoursesRepository{
		DB: db,
	}
}

func (rep *MysqlCoursesRepository) Create(ctx context.Context, domain course.Domain) (course.Domain, error) {
	newCourse := FromDomain(domain)
	newCourse.CreatedAt = time.Now()

	checkCategories := rep.DB.Table("categories").Where("id = ?", newCourse.CategoryId).Find(&newCourse.Category)
	if checkCategories.RowsAffected == 0 {
		return course.Domain{}, err.ErrCategoryNotFound
	}

	checkTeacher := rep.DB.Table("teachers").Where("id = ?", newCourse.TeacherId).Find(&newCourse.Teacher)
	if checkTeacher.RowsAffected == 0 {
		return course.Domain{}, err.ErrTeacherNotFound
	}

	checkdifficulties := rep.DB.Table("difficulties").Where("id = ?", newCourse.DifficultyId).Find(&newCourse.Difficulty)
	if checkdifficulties.RowsAffected == 0 {
		return course.Domain{}, err.ErrDifficultyNotFound
	}

	//fire to databases
	resultAdd := rep.DB.Create(&newCourse)
	if resultAdd.Error != nil {
		return course.Domain{}, resultAdd.Error
	}
	return newCourse.ToDomain(), nil
}

func (rep *MysqlCoursesRepository) GetAll(ctx context.Context) ([]course.BatchesDomain, error) {
	//Get all data from databases
	listCourses := []CourseInBatches{}

	rawQuery := `SELECT C.id, C.title, C.thumbnail, C.description, 
					C.category_id, C.difficulty_id, C.teacher_id, 
					AVG(E.rating) AS rating, A.title AS category,
					D.title AS difficulty
				FROM courses C 
				LEFT JOIN enrollments E ON C.id = E.course_id
				LEFT JOIN categories A ON A.id = C.category_id
				LEFT JOIN difficulties D ON D.id = C.difficulty_id
				GROUP BY C.id`

	result := rep.DB.Raw(rawQuery).Scan(&listCourses)
	if result.Error != nil {
		return []course.BatchesDomain{}, result.Error
	}

	if result.RowsAffected == 0 {
		return []course.BatchesDomain{}, err.ErrCourseNotFound
	}

	//convert from Repo to Domain List
	listDomain := BatchesToDomain(listCourses)
	return listDomain, nil
}

func (rep *MysqlCoursesRepository) SearchCourses(ctx context.Context, title string, category string, difficulty string) ([]course.BatchesDomain, error) {
	//Get all data from databases
	listCourses := []CourseInBatches{}

	rawQuery := `SELECT C.id, C.title, C.thumbnail, C.description, 
					C.category_id, C.difficulty_id, C.teacher_id, 
					AVG(E.rating) AS rating, A.title AS category,
					D.title AS difficulty
				FROM courses C 
				LEFT JOIN enrollments E ON C.id = E.course_id
				LEFT JOIN categories A ON A.id = C.category_id
				LEFT JOIN difficulties D ON D.id = C.difficulty_id
				WHERE A.title LIKE ? AND D.title LIKE ? AND C.title LIKE ?
				GROUP BY C.id`

	result := rep.DB.Raw(rawQuery, "%"+category+"%", "%"+difficulty+"%", "%"+title+"%").Scan(&listCourses)
	if result.Error != nil {
		return []course.BatchesDomain{}, result.Error
	}

	if result.RowsAffected == 0 {
		return []course.BatchesDomain{}, err.ErrCourseNotFound
	}

	//convert from Repo to Domain List
	listDomain := BatchesToDomain(listCourses)
	return listDomain, nil
}

func (rep *MysqlCoursesRepository) GetCourseById(ctx context.Context, id uint) (course.Domain, error) {
	var targetTable Course

	rawQuery := `SELECT C.id, C.title, C.thumbnail, C.description, C.category_id, 
					C.difficulty_id, C.teacher_id, AVG(E.rating) AS rating 
				FROM courses C 
				LEFT JOIN enrollments E ON C.id = E.course_id 
				WHERE C.id = ? 
				GROUP BY C.id`

	checkCourse := rep.DB.Raw(rawQuery, id).Scan(&targetTable)
	if checkCourse.RowsAffected == 0 {
		return course.Domain{}, err.ErrCourseNotFound
	}
	return targetTable.ToDomain(), nil
}

func (rep *MysqlCoursesRepository) Update(ctx context.Context, domain course.Domain) (course.Domain, error) {
	var targetTable Course
	newCourse := FromDomain(domain)

	//fire to databases
	resultUpdate := rep.DB.Model(&targetTable).Where("id = ?", domain.Id).Updates(newCourse)
	if resultUpdate.Error != nil {
		return course.Domain{}, resultUpdate.Error
	}
	return newCourse.ToDomain(), nil
}

func (rep *MysqlCoursesRepository) CheckTeacher(ctx context.Context, id uint) (teacher.Domain, error) {
	var targetTable Course

	checkTeacher := rep.DB.Table("teachers").Where("id = ?", id).Find(&targetTable.Teacher)
	if checkTeacher.RowsAffected == 0 {
		return teacher.Domain{}, err.ErrTeacherNotFound
	}
	return targetTable.Teacher.ToDomain(), nil
}

func (rep *MysqlCoursesRepository) CheckCategories(ctx context.Context, id uint) (categories.Domain, error) {
	var targetTable Course

	checkCategories := rep.DB.Table("categories").Where("id = ?", id).Find(&targetTable.Category)
	if checkCategories.RowsAffected == 0 {
		return categories.Domain{}, err.ErrCategoryNotFound
	}
	return targetTable.Category.ToDomain(), nil
}

func (rep *MysqlCoursesRepository) CheckDifficulties(ctx context.Context, id uint) (difficulties.Domain, error) {
	var targetTable Course

	checkCategories := rep.DB.Table("difficulties").Where("id = ?", id).Find(&targetTable.Difficulty)
	if checkCategories.RowsAffected == 0 {
		return difficulties.Domain{}, err.ErrCategoryNotFound
	}
	return targetTable.Difficulty.ToDomain(), nil
}

func (rep *MysqlCoursesRepository) Delete(ctx context.Context, id uint) error {
	var targetDelete Course

	//fire soft delete
	delete := rep.DB.Where("id = ?", id).Delete(&targetDelete)
	if delete.Error != nil {
		return delete.Error
	}
	return delete.Error
}

func (rep *MysqlCoursesRepository) GetCoursesByCourseIds(ctx context.Context, courseIds []uint) ([]course.BatchesDomain, error) {
	var targetTable []CourseInBatches
	rawQuery := `SELECT C.id, C.title, C.thumbnail, C.description, 
					C.category_id, C.difficulty_id, C.teacher_id, 
					AVG(E.rating) AS rating, A.title AS category,
					D.title AS difficulty
				FROM courses C 
				LEFT JOIN enrollments E ON C.id = E.course_id
				LEFT JOIN categories A ON A.id = C.category_id
				LEFT JOIN difficulties D ON D.id = C.difficulty_id
				WHERE C.id IN ?
				GROUP BY C.id`

	checkCourse := rep.DB.Raw(rawQuery, courseIds).Scan(&targetTable)
	if checkCourse.RowsAffected == 0 {
		return []course.BatchesDomain{}, err.ErrCourseNotFound
	}
	return BatchesToDomain(targetTable), nil
}

func (rep *MysqlCoursesRepository) GetEnrollmentsByStudentId(ctx context.Context, studentId uint) ([]course.CourseEnrollmentDomain, error) {
	var enrollments []CourseEnrollment

	getEnrollments := rep.DB.Table("enrollments").Where("student_id = ?", studentId).Find(&enrollments)
	if getEnrollments.RowsAffected == 0 {
		return []course.CourseEnrollmentDomain{}, err.ErrEnrollmentsNotFound
	}
	return EnrollmentsToDomain(enrollments), nil
}

func (rep *MysqlCoursesRepository) GetCourseByTeacherId(ctx context.Context, teacherId uint) ([]course.BatchesDomain, error) {
	var targetTable []CourseInBatches
	rawQuery := `SELECT C.id, C.title, C.thumbnail, C.description, 
					C.category_id, C.difficulty_id, C.teacher_id, 
					AVG(E.rating) AS rating, A.title AS category,
					D.title AS difficulty
				FROM courses C 
				LEFT JOIN enrollments E ON C.id = E.course_id
				LEFT JOIN categories A ON A.id = C.category_id
				LEFT JOIN difficulties D ON D.id = C.difficulty_id
				WHERE C.teacher_id = ?
				GROUP BY C.id`

	checkCourse := rep.DB.Raw(rawQuery, teacherId).Scan(&targetTable)
	if checkCourse.RowsAffected == 0 {
		return []course.BatchesDomain{}, err.ErrCourseNotFound
	}
	return BatchesToDomain(targetTable), nil
}
