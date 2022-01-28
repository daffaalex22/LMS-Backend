package course_test

import (
	"backend/business/categories"
	"backend/business/course"
	"backend/business/course/mocks"
	"backend/business/difficulties"
	"backend/business/student"
	"backend/business/teacher"
	"backend/helper/err"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var courseRepository mocks.Repository
var courseService course.Usecase
var courseDomain course.Domain
var coursesDomain []course.Domain
var batchedDomain course.BatchesDomain
var batchedsDomain []course.BatchesDomain
var categoryDomain categories.Domain
var teacherDomain teacher.Domain
var enrollmentDomain course.CourseEnrollmentDomain
var enrollmentsDomain []course.CourseEnrollmentDomain

func setup() {
	courseService = course.NewCourseUsecase(time.Hour*1, &courseRepository)
	courseDomain = course.Domain{
		Id:           1,
		Title:        "Unit Testing",
		Thumbnail:    "www.google.com",
		Description:  "Learning Unit Testing",
		Rating:       5,
		CategoryId:   1,
		TeacherId:    1,
		DifficultyId: 1,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	categoryDomain = categories.Domain{
		Id:    1,
		Title: "Engineering",
	}
	teacherDomain = teacher.Domain{
		Id:          1,
		Name:        "Bambang",
		Password:    "alterra123",
		NewPassword: "",
		Email:       "someone@alterra.id",
		Avatar:      "www.google.com",
		Phone:       "0878181818181",
		Address:     "dirumah",
		BackGround:  "orang biasa",
		Token:       "alterra123",
	}
	coursesDomain = append(coursesDomain, courseDomain)

	enrollmentDomain = course.CourseEnrollmentDomain{
		StudentId: 1,
		CourseId:  1,
		Rating:    4,
		Review:    "Mantap Jiwa Coursenya",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Course:    courseDomain,
		Student: student.Domain{
			Id:       1,
			Name:     "Daffa' Alexander",
			Password: "$2a$10$6UPRGYrF4cQCO4aHLcGmouO8jXLWW.KkAYiSlU1p3IGSUwZf/kjnq",
			Email:    "drivealex22@gmail.com",
			Avatar:   "https://firebasestorage.googleapis.com/v0/b/learning-management-syst-bf947.appspot.com/o/profile%2FS__66502658.jpg?alt=media&token=46ea7317-0cad-44d9-af22-6453fc26ee30",
			Phone:    "089602903213",
			Address:  "Bandung",
			Token:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiZXhwIjoxNjQyMDYzNzE3fQ.wOON-dfYvFHmUM5wVVNTeaOf1spuHYnkDiur45jnyJ4",
		},
	}
	enrollmentsDomain = append(enrollmentsDomain, enrollmentDomain)

	batchedDomain = course.BatchesDomain{
		Id:           1,
		Title:        "Unit Testing",
		Thumbnail:    "www.google.com",
		Description:  "Learning Unit Testing",
		Rating:       5,
		CategoryId:   1,
		TeacherId:    1,
		DifficultyId: 1,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	batchedsDomain = append(batchedsDomain, batchedDomain)
}

func TestCreate(t *testing.T) {
	setup()
	t.Run("Test case - 1 | Create - Succes", func(t *testing.T) {
		courseRepository.On("Create", mock.Anything, mock.Anything).Return(courseDomain, nil).Once()
		data, err := courseService.Create(context.Background(), courseDomain)

		assert.Nil(t, err)
		assert.Equal(t, data, courseDomain)
	})

	t.Run("Test case - 5 | Create - Failed | Error Create", func(t *testing.T) {
		courseRepository.On("Create", mock.Anything, mock.Anything).Return(courseDomain, errors.New("error test")).Once()
		_, errr := courseService.Create(context.Background(), courseDomain)

		assert.Error(t, errr)
	})

	t.Run("Test case - 2 | Create - Failed | Title Empty", func(t *testing.T) {
		courseRepository.On("Create", mock.Anything, mock.Anything).Return(nil, err.ErrTitleEmpty).Once()
		data, errr := courseService.Create(context.Background(), course.Domain{
			Id:           1,
			Thumbnail:    "www.google.com",
			Description:  "Learning Unit Testing",
			Rating:       5,
			CategoryId:   1,
			TeacherId:    1,
			DifficultyId: 1,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		})

		assert.NotNil(t, errr)
		assert.Equal(t, errr, err.ErrTitleEmpty)
		assert.Equal(t, data, course.Domain{})
	})
	t.Run("Test case - 3 | Create - Failed | CategoryId Empty", func(t *testing.T) {
		courseRepository.On("Create", mock.Anything, mock.Anything).Return(nil, err.ErrCategoryIdEmpty).Once()
		data, errr := courseService.Create(context.Background(), course.Domain{
			Id:           1,
			Title:        "Unit Testing",
			Thumbnail:    "www.google.com",
			Description:  "Learning Unit Testing",
			Rating:       5,
			CategoryId:   0,
			TeacherId:    1,
			DifficultyId: 1,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		})

		assert.NotNil(t, errr)
		assert.Equal(t, errr, err.ErrCategoryIdEmpty)
		assert.Equal(t, data, course.Domain{})
	})
	t.Run("Test case - 4 | Create - Failed | TeacherId Empty", func(t *testing.T) {
		courseRepository.On("Create", mock.Anything, mock.Anything).Return(nil, err.ErrTeacherIdEmpty).Once()
		data, errr := courseService.Create(context.Background(), course.Domain{
			Id:           1,
			Title:        "Unit Testing",
			Thumbnail:    "www.google.com",
			Description:  "Learning Unit Testing",
			Rating:       5,
			CategoryId:   1,
			TeacherId:    0,
			DifficultyId: 1,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		})

		assert.NotNil(t, errr)
		assert.Equal(t, errr, err.ErrTeacherIdEmpty)
		assert.Equal(t, data, course.Domain{})
	})

}

func TestGetAll(t *testing.T) {
	setup()
	t.Run("Test case - 1 | GetAll - Success", func(t *testing.T) {
		courseRepository.On("GetAll", mock.Anything, mock.Anything).Return([]course.BatchesDomain{}, nil).Once()
		data, errr := courseService.GetAll(context.Background())

		assert.Nil(t, errr)
		assert.Equal(t, data, []course.BatchesDomain{})
	})
	t.Run("Test case - 2 | GetAll - Failed || Repository Empty", func(t *testing.T) {
		courseRepository.On("GetAll", mock.Anything).Return([]course.BatchesDomain{}, err.ErrDataEmpty).Once()
		_, errr := courseService.GetAll(context.Background())

		assert.NotNil(t, errr)
	})
}

func TestSearchCourses(t *testing.T) {
	setup()
	t.Run("Test case - 1 | SearchCourses - Success", func(t *testing.T) {
		courseRepository.On("SearchCourses", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return([]course.BatchesDomain{}, nil).Once()
		data, errr := courseService.SearchCourses(context.Background(), "Unit Testing", "Software Engineering", "easy")

		assert.Nil(t, errr)
		assert.Equal(t, data, []course.BatchesDomain{})
	})
	t.Run("Test case - 2 | SearchCourses - Failed || Repository Empty", func(t *testing.T) {
		courseRepository.On("SearchCourses", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return([]course.BatchesDomain{}, err.ErrDataEmpty).Once()
		_, errr := courseService.SearchCourses(context.Background(), "Unit Testing", "Software Engineering", "easy")

		assert.NotNil(t, errr)
	})
}

func TestGetCourseById(t *testing.T) {
	setup()
	t.Run("Test case - 1 | GetCourseById - Success", func(t *testing.T) {
		courseRepository.On("GetCourseById", mock.Anything, mock.Anything).Return(courseDomain, nil).Once()
		courseRepository.On("CheckCategories", mock.Anything, mock.Anything).Return(categories.Domain{}, nil).Once()
		courseRepository.On("CheckTeacher", mock.Anything, mock.Anything).Return(teacher.Domain{}, nil).Once()
		courseRepository.On("CheckDifficulties", mock.Anything, mock.Anything).Return(difficulties.Domain{}, nil).Once()
		data, errr := courseService.GetCourseById(context.Background(), "1")

		assert.Nil(t, errr)
		assert.Equal(t, data, courseDomain)
	})
	t.Run("Test case - 4 | GetCourseById - Failed || course empty", func(t *testing.T) {
		courseRepository.On("GetCourseById", mock.Anything, mock.Anything).Return(course.Domain{}, err.ErrDataEmpty).Once()
		_, errr := courseService.GetCourseById(context.Background(), "1")

		assert.Error(t, errr)
	})
	t.Run("Test case - 5 | GetCourseById - Failed || categories empty", func(t *testing.T) {
		courseRepository.On("GetCourseById", mock.Anything, mock.Anything).Return(courseDomain, nil).Once()
		courseRepository.On("CheckCategories", mock.Anything, mock.Anything).Return(categories.Domain{}, err.ErrDataEmpty).Once()
		_, errr := courseService.GetCourseById(context.Background(), "1")

		assert.Error(t, errr)
	})
	t.Run("Test case - 6 | GetCourseById - Failed || teacher empty", func(t *testing.T) {
		courseRepository.On("GetCourseById", mock.Anything, mock.Anything).Return(courseDomain, nil).Once()
		courseRepository.On("CheckCategories", mock.Anything, mock.Anything).Return(categoryDomain, nil).Once()
		courseRepository.On("CheckTeacher", mock.Anything, mock.Anything).Return(teacher.Domain{}, err.ErrDataEmpty).Once()

		_, errr := courseService.GetCourseById(context.Background(), "1")

		assert.Error(t, errr)
	})
	t.Run("Test case - 7 | GetCourseById - Failed || difficulties empty", func(t *testing.T) {
		courseRepository.On("GetCourseById", mock.Anything, mock.Anything).Return(courseDomain, nil).Once()
		courseRepository.On("CheckCategories", mock.Anything, mock.Anything).Return(categoryDomain, nil).Once()
		courseRepository.On("CheckTeacher", mock.Anything, mock.Anything).Return(teacherDomain, nil).Once()
		courseRepository.On("CheckDifficulties", mock.Anything, mock.Anything).Return(difficulties.Domain{}, err.ErrDataEmpty).Once()

		_, errr := courseService.GetCourseById(context.Background(), "1")

		assert.Error(t, errr)
	})
	t.Run("Test case - 2 | GetCourseById - Failed || id empty", func(t *testing.T) {
		courseRepository.On("GetCourseById", mock.Anything, mock.Anything).Return(courseDomain, nil).Once()
		courseRepository.On("CheckCategories", mock.Anything, mock.Anything).Return(categories.Domain{}, nil).Once()
		courseRepository.On("CheckTeacher", mock.Anything, mock.Anything).Return(teacher.Domain{}, nil).Once()
		courseRepository.On("CheckDifficulties", mock.Anything, mock.Anything).Return(difficulties.Domain{}, nil).Once()
		data, errr := courseService.GetCourseById(context.Background(), "")

		assert.NotNil(t, errr)
		assert.Equal(t, data, course.Domain{})
	})
	t.Run("Test case - 3 | GetCourseById - Failed || convert Error", func(t *testing.T) {
		_, errr := courseService.GetCourseById(context.Background(), "a")

		assert.NotNil(t, errr)
	})

}

func TestUpdate(t *testing.T) {
	setup()

	t.Run("Test case 1 | Update - Success", func(t *testing.T) {
		courseRepository.On("Update",
			mock.Anything,
			mock.AnythingOfType("course.Domain")).Return(courseDomain, nil).Once()

		courseRepository.On("GetCourseById",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(courseDomain, nil).Once()

		courseRepository.On("CheckCategories",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(categoryDomain, nil).Once()

		courseRepository.On("CheckTeacher",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(teacherDomain, nil).Once()

		crs, err := courseService.Update(context.Background(),
			"1",
			courseDomain)

		assert.Nil(t, err)
		assert.Equal(t, courseDomain, crs)
	})

	t.Run("Test case 2 | Update - Failed || error id", func(t *testing.T) {
		crs, err := courseService.Update(context.Background(),
			"",
			courseDomain)

		assert.Error(t, err)
		assert.Equal(t, crs, course.Domain{})
	})

	t.Run("Test case 3 | Update - Failed || Title Empty", func(t *testing.T) {
		_, err := courseService.Update(context.Background(),
			"1",
			course.Domain{
				Title: "",
			})

		assert.NotNil(t, err)
	})

	t.Run("Test case 4 | Update - Failed || CategoryId Empty", func(t *testing.T) {
		_, err := courseService.Update(context.Background(),
			"1",
			course.Domain{
				Title: "Physics Fundamentals",
			})

		assert.NotNil(t, err)
	})

	t.Run("Test case 4 Update - Failed || TeacherId Empty", func(t *testing.T) {
		_, err := courseService.Update(context.Background(),
			"1",
			course.Domain{
				Title:      "Physics Fundamentals",
				CategoryId: 1,
			})

		assert.NotNil(t, err)
	})

	t.Run("Test case 6 | Update - Failed || Error CheckCategory", func(t *testing.T) {
		courseRepository.On("Update",
			mock.Anything,
			mock.AnythingOfType("course.Domain")).Return(course.Domain{}, err.ErrDataEmpty).Once()

		courseRepository.On("GetCourseById",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(courseDomain, nil).Once()

		courseRepository.On("CheckCategories",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(categoryDomain, err.ErrDataEmpty).Once()

		courseRepository.On("CheckTeacher",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(teacherDomain, nil).Once()

		crs, err := courseService.Update(context.Background(),
			"1",
			courseDomain)

		assert.Error(t, err)
		assert.Equal(t, crs, course.Domain{})
	})

	t.Run("Test case 7 | Update - Failed || Error CheckTeacher", func(t *testing.T) {
		courseRepository.On("Update",
			mock.Anything,
			mock.AnythingOfType("course.Domain")).Return(course.Domain{}, err.ErrDataEmpty).Once()

		courseRepository.On("GetCourseById",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(courseDomain, nil).Once()

		courseRepository.On("CheckCategories",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(categoryDomain, nil).Once()

		courseRepository.On("CheckTeacher",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(teacherDomain, err.ErrDataEmpty).Once()

		crs, err := courseService.Update(context.Background(),
			"1",
			courseDomain)

		assert.Error(t, err)
		assert.Equal(t, crs, course.Domain{})
	})

	t.Run("Test case 8 | Update - Failed || Error GetCourseById", func(t *testing.T) {
		courseRepository.On("Update",
			mock.Anything,
			mock.AnythingOfType("course.Domain")).Return(course.Domain{}, err.ErrDataEmpty).Once()

		courseRepository.On("GetCourseById",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(courseDomain, err.ErrDataEmpty).Once()

		courseRepository.On("CheckCategories",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(categoryDomain, nil).Once()

		courseRepository.On("CheckTeacher",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(teacherDomain, nil).Once()

		crs, err := courseService.Update(context.Background(),
			"1",
			courseDomain)

		assert.Error(t, err)
		assert.Equal(t, crs, course.Domain{})
	})

	t.Run("Test case 9 | Update - Failed || error id can't convert", func(t *testing.T) {
		crs, err := courseService.Update(context.Background(),
			"a",
			courseDomain)

		assert.Error(t, err)
		assert.Equal(t, crs, course.Domain{})
	})
	t.Run("Test case 10 | Update - failed || data course empty", func(t *testing.T) {
		courseRepository.On("Update",
			mock.Anything,
			mock.AnythingOfType("course.Domain")).Return(courseDomain, nil).Once()

		courseRepository.On("GetCourseById",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(course.Domain{}, err.ErrDataEmpty).Once()

		courseRepository.On("CheckCategories",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(categoryDomain, nil).Once()

		courseRepository.On("CheckTeacher",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(teacherDomain, nil).Once()

		crs, err := courseService.Update(context.Background(),
			"1",
			courseDomain)

		assert.NotNil(t, err)
		assert.Equal(t, course.Domain{}, crs)
	})
	t.Run("Test case 11 | Update - failed || data teacher empty", func(t *testing.T) {
		courseRepository.On("Update",
			mock.Anything,
			mock.AnythingOfType("course.Domain")).Return(courseDomain, nil).Once()

		courseRepository.On("GetCourseById",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(courseDomain, err.ErrDataEmpty).Once()

		courseRepository.On("CheckCategories",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(categoryDomain, nil).Once()

		courseRepository.On("CheckTeacher",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(teacher.Domain{}, err.ErrDataEmpty).Once()

		crs, err := courseService.Update(context.Background(),
			"1",
			courseDomain)

		assert.NotNil(t, err)
		assert.Equal(t, course.Domain{}, crs)
	})
}

func TestDelete(t *testing.T) {
	setup()

	t.Run("Test case 1 | Delete - Success", func(t *testing.T) {
		courseRepository.On("GetCourseById", mock.Anything, mock.AnythingOfType("uint")).Return(courseDomain, nil).Once()
		courseRepository.On("Delete", mock.Anything, mock.AnythingOfType("uint")).Return(nil).Once()
		data, errr := courseService.Delete(context.Background(), "1")

		assert.NoError(t, errr)
		assert.NotNil(t, data)
	})

	t.Run("Test case 2 | Delete - Failed || Error Get Course By Id", func(t *testing.T) {

		courseRepository.On("GetCourseById",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(courseDomain, err.ErrDataEmpty).Once()

		courseRepository.On("Delete",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(nil).Once()

		crs, err := courseService.Delete(context.Background(), "1")

		assert.Error(t, err)
		assert.Equal(t, crs, course.Domain{})
	})

	t.Run("Test case 3 | Delete - Failed || Invalid Id Empty", func(t *testing.T) {
		_, err := courseService.Delete(context.Background(), "")

		assert.NotNil(t, err)
	})

	t.Run("Test case 4 | Delete - Failed || cant convert Id ", func(t *testing.T) {
		_, err := courseService.Delete(context.Background(), "a")

		assert.NotNil(t, err)
	})

}

func TestGetCourseByStudentId(t *testing.T) {
	setup()

	t.Run("Test Case 1 | GetCourseByStudentId - Success", func(t *testing.T) {
		courseRepository.On("GetEnrollmentsByStudentId", mock.Anything, mock.AnythingOfType("uint")).Return(enrollmentsDomain, nil).Once()
		courseRepository.On("GetCoursesByCourseIds", mock.Anything, mock.Anything).Return(batchedsDomain, nil).Once()
		_, errr := courseService.GetCourseByStudentId(context.Background(), 1)

		assert.NoError(t, errr)
	})
	t.Run("Test Case 2 | GetCourseByStudentId - Failed || student id empty", func(t *testing.T) {
		courseRepository.On("GetEnrollmentsByStudentId", mock.Anything, mock.AnythingOfType("uint")).Return([]course.CourseEnrollmentDomain{}, err.ErrDataEmpty).Once()
		_, errr := courseService.GetCourseByStudentId(context.Background(), 0)

		assert.Error(t, errr)
	})
	t.Run("Test Case 3 | GetCourseByStudentId - Failed || course empty", func(t *testing.T) {
		courseRepository.On("GetEnrollmentsByStudentId", mock.Anything, mock.AnythingOfType("uint")).Return(enrollmentsDomain, nil).Once()
		courseRepository.On("GetCoursesByCourseIds", mock.Anything, mock.Anything).Return([]course.BatchesDomain{}, err.ErrDataEmpty).Once()
		_, errr := courseService.GetCourseByStudentId(context.Background(), 0)

		assert.Error(t, errr)
	})

}

func TestGetCourseByTeacherId(t *testing.T) {
	setup()

	t.Run("Test Case 1 | GetCourseByTeacherId - Success", func(t *testing.T) {
		courseRepository.On("GetCourseByTeacherId", mock.Anything, mock.AnythingOfType("uint")).Return(batchedsDomain, nil).Once()
		_, errr := courseService.GetCourseByTeacherId(context.Background(), 1)

		assert.NoError(t, errr)
	})
	t.Run("Test Case 2 | GetCourseByTeacherId - failed || teacherid empty", func(t *testing.T) {
		courseRepository.On("GetCourseByTeacherId", mock.Anything, mock.AnythingOfType("uint")).Return([]course.BatchesDomain{}, err.ErrDataEmpty).Once()
		data, errr := courseService.GetCourseByTeacherId(context.Background(), 1)

		assert.Error(t, errr)
		assert.Equal(t, data, []course.BatchesDomain{})
	})

}
