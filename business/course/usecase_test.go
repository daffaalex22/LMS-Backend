package course_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"backend/business/categories"
	"backend/business/course"
	_mockCourseRepository "backend/business/course/mocks"
	"backend/business/teacher"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var courseRepository _mockCourseRepository.Repository
var courseService course.Usecase
var courseDomain course.Domain
var coursesDomain []course.Domain
var categoryDomain categories.Domain
var teacherDomain teacher.Domain

func setup() {
	courseService = course.NewCourseUsecase(time.Hour*1, &courseRepository /*, configJWT */)
	courseDomain = course.Domain{
		Id:           1,
		Title:        "Physics Fundamentals",
		Thumbnail:    "https://www.google.com",
		Description:  "Ini adalah course",
		CategoryId:   1,
		TeacherId:    1,
		DifficultyId: 1,
	}
	categoryDomain = categories.Domain{
		Id:    1,
		Title: "Science and Engineering",
	}
	teacherDomain = teacher.Domain{
		Id:         1,
		Name:       "Bambang",
		Password:   "alterra123",
		Email:      "daffa@alterra.id",
		Avatar:     "https://www.google.com",
		Phone:      "6281123456789",
		Address:    "Bandung",
		BackGround: "Jago Ngoding",
		Token:      "alterra123",
	}
	coursesDomain = append(coursesDomain, courseDomain)
}

func TestCreate(t *testing.T) {
	setup()

	t.Run("Test case 1 | Valid Create", func(t *testing.T) {
		courseRepository.On("Create",
			mock.Anything,
			mock.AnythingOfType("course.Domain")).Return(courseDomain, nil).Once()

		crs, err := courseService.Create(context.Background(), courseDomain)

		assert.Nil(t, err)
		assert.Equal(t, "Physics Fundamentals", crs.Title)
	})

	t.Run("Test case 2 | Error Create", func(t *testing.T) {
		courseRepository.On("Create",
			mock.Anything,
			mock.AnythingOfType("course.Domain")).Return(course.Domain{}, errors.New("Unexpected Error")).Once()

		crs, err := courseService.Create(context.Background(), courseDomain)

		assert.Error(t, err)
		assert.Equal(t, crs, course.Domain{})
	})

	t.Run("Test case 3 | Invalid Title Empty", func(t *testing.T) {
		_, err := courseService.Create(context.Background(), course.Domain{
			Title:        "",
			Thumbnail:    "https://www.google.com",
			Description:  "Ini adalah course",
			CategoryId:   1,
			TeacherId:    1,
			DifficultyId: 1,
		})

		assert.NotNil(t, err)
	})

	t.Run("Test case 4 | Invalid CategoryId Empty", func(t *testing.T) {
		_, err := courseService.Create(context.Background(), course.Domain{
			Title:        "Physics Fundamentals",
			Thumbnail:    "https://www.google.com",
			Description:  "Ini adalah course",
			TeacherId:    1,
			DifficultyId: 1,
		})

		assert.NotNil(t, err)
	})

	t.Run("Test case 4 | Invalid TeacherId Empty", func(t *testing.T) {
		_, err := courseService.Create(context.Background(), course.Domain{
			Title:        "Physics Fundamentals",
			Thumbnail:    "https://www.google.com",
			Description:  "Ini adalah course",
			CategoryId:   1,
			DifficultyId: 1,
		})

		assert.NotNil(t, err)
	})
}

func TestGetAll(t *testing.T) {
	setup()

	t.Run("Test case 1", func(t *testing.T) {
		courseRepository.On("GetAll",
			mock.Anything).Return(coursesDomain, nil).Once()

		crs, err := courseService.GetAll(context.Background())

		assert.NoError(t, err)
		assert.NotNil(t, crs)

	})

	t.Run("Test case 2 | Error", func(t *testing.T) {
		courseRepository.On("GetAll",
			mock.Anything).Return([]course.Domain{}, errors.New("Error at Get All course")).Once()

		crs, err := courseService.GetAll(context.Background())

		assert.Error(t, err)
		assert.Equal(t, crs, []course.Domain{})
	})
}

func TestGetCourseById(t *testing.T) {
	setup()

	t.Run("Test case 1", func(t *testing.T) {
		courseRepository.On("GetCourseById",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(courseDomain, nil).Once()

		crs, err := courseService.GetCourseById(context.Background(), "1")

		assert.NoError(t, err)
		assert.NotNil(t, crs)
	})

	t.Run("Test case 2 | Error", func(t *testing.T) {
		courseRepository.On("GetCourseById",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(course.Domain{}, errors.New("Unexpected Error")).Once()

		crs, err := courseService.GetCourseById(context.Background(), "1")

		assert.Error(t, err)
		assert.Equal(t, crs, course.Domain{})
	})

	t.Run("Test case 4 | Invalid Id Empty", func(t *testing.T) {
		_, err := courseService.GetCourseById(context.Background(), "")

		assert.NotNil(t, err)
	})
}

func TestUpdate(t *testing.T) {
	setup()

	t.Run("Test case 1 | Valid Update", func(t *testing.T) {
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
		assert.Equal(t, "Physics Fundamentals", crs.Title)
	})

	t.Run("Test case 2 | Error Update", func(t *testing.T) {
		courseRepository.On("Update",
			mock.Anything,
			mock.AnythingOfType("course.Domain")).Return(course.Domain{}, errors.New("Unexpected Error")).Once()

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

		assert.Error(t, err)
		assert.Equal(t, crs, course.Domain{})
	})

	t.Run("Test case 3 | Invalid Title Empty", func(t *testing.T) {
		_, err := courseService.Update(context.Background(),
			"1",
			course.Domain{
				Title:        "",
				Thumbnail:    "https://www.google.com",
				Description:  "Ini adalah course",
				CategoryId:   1,
				TeacherId:    1,
				DifficultyId: 1,
			})

		assert.NotNil(t, err)
	})

	t.Run("Test case 4 | Invalid CategoryId Empty", func(t *testing.T) {
		_, err := courseService.Update(context.Background(),
			"1",
			course.Domain{
				Title:        "Physics Fundamentals",
				Thumbnail:    "https://www.google.com",
				Description:  "Ini adalah course",
				TeacherId:    1,
				DifficultyId: 1,
			})

		assert.NotNil(t, err)
	})

	t.Run("Test case 4 | Invalid TeacherId Empty", func(t *testing.T) {
		_, err := courseService.Update(context.Background(),
			"1",
			course.Domain{
				Title:        "Physics Fundamentals",
				Thumbnail:    "https://www.google.com",
				Description:  "Ini adalah course",
				CategoryId:   1,
				DifficultyId: 1,
			})

		assert.NotNil(t, err)
	})

	t.Run("Test case 5 | Invalid Id Empty", func(t *testing.T) {
		_, err := courseService.Update(context.Background(),
			"",
			course.Domain{
				Title:        "Physics Fundamentals",
				Thumbnail:    "https://www.google.com",
				Description:  "Ini adalah course",
				CategoryId:   1,
				DifficultyId: 1,
			})

		assert.NotNil(t, err)
	})

	t.Run("Test case 6 | Error CheckCategory", func(t *testing.T) {
		courseRepository.On("Update",
			mock.Anything,
			mock.AnythingOfType("course.Domain")).Return(course.Domain{}, errors.New("Unexpected Error")).Once()

		courseRepository.On("GetCourseById",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(courseDomain, nil).Once()

		courseRepository.On("CheckCategories",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(categoryDomain, errors.New("Unexpected Error")).Once()

		courseRepository.On("CheckTeacher",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(teacherDomain, nil).Once()

		crs, err := courseService.Update(context.Background(),
			"1",
			courseDomain)

		assert.Error(t, err)
		assert.Equal(t, crs, course.Domain{})
	})

	t.Run("Test case 7 | Error CheckTeacher", func(t *testing.T) {
		courseRepository.On("Update",
			mock.Anything,
			mock.AnythingOfType("course.Domain")).Return(course.Domain{}, errors.New("Unexpected Error")).Once()

		courseRepository.On("GetCourseById",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(courseDomain, nil).Once()

		courseRepository.On("CheckCategories",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(categoryDomain, nil).Once()

		courseRepository.On("CheckTeacher",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(teacherDomain, errors.New("Unexpected Error")).Once()

		crs, err := courseService.Update(context.Background(),
			"1",
			courseDomain)

		assert.Error(t, err)
		assert.Equal(t, crs, course.Domain{})
	})

	t.Run("Test case 8 | Error GetCourseById", func(t *testing.T) {
		courseRepository.On("Update",
			mock.Anything,
			mock.AnythingOfType("course.Domain")).Return(course.Domain{}, errors.New("Unexpected Error")).Once()

		courseRepository.On("GetCourseById",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(courseDomain, errors.New("Unexpected Error")).Once()

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
}

func TestDelete(t *testing.T) {
	setup()

	t.Run("Test case 1", func(t *testing.T) {
		courseRepository.On("GetCourseById",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(courseDomain, nil).Once()

		courseRepository.On("Delete",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(nil).Once()

		crs, err := courseService.Delete(context.Background(), "1")

		assert.NoError(t, err)
		assert.NotNil(t, crs)
	})

	t.Run("Test case 2 | Error Get Course By Id", func(t *testing.T) {

		courseRepository.On("GetCourseById",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(courseDomain, errors.New("Unexpected Error")).Once()

		courseRepository.On("Delete",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(nil).Once()

		crs, err := courseService.Delete(context.Background(), "1")

		assert.Error(t, err)
		assert.Equal(t, crs, course.Domain{})
	})

	t.Run("Test case 3 | Invalid Id Empty", func(t *testing.T) {
		_, err := courseService.Delete(context.Background(), "")

		assert.NotNil(t, err)
	})
}
