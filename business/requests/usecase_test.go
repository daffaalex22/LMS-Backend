package requests_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"backend/business/course"
	"backend/business/requests"
	_mockRequestRepository "backend/business/requests/mocks"
	"backend/business/student"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var requestRepository _mockRequestRepository.RequestsRepoInterface
var requestService requests.RequestsUseCaseInterface
var requestDomain requests.Domain
var requestsDomain []requests.Domain
var studentDomain student.Domain
var courseDomain course.Domain

func setup() {
	requestService = requests.NewUseCase(&requestRepository, time.Hour*1 /*, configJWT */)
	requestDomain = requests.Domain{
		StudentId: 1,
		CourseId:  1,
		TypeId:    1,
		Status:    "Pending",
		Course: course.Domain{
			Id:           1,
			Title:        "Frontend Development",
			Thumbnail:    "https://www.nawpic.com/media/2020/cool-nawpic-5.jpg",
			Description:  "backend Developer",
			Rating:       3.6,
			CategoryId:   1,
			TeacherId:    1,
			DifficultyId: 1,
		},
	}
	requestsDomain = append(requestsDomain, requestDomain)
	studentDomain = student.Domain{
		Id:       1,
		Name:     "Daffa' Alexander",
		Password: "$2a$10$6UPRGYrF4cQCO4aHLcGmouO8jXLWW.KkAYiSlU1p3IGSUwZf/kjnq",
		Email:    "drivealex22@gmail.com",
		Avatar:   "https://firebasestorage.googleapis.com/v0/b/learning-management-syst-bf947.appspot.com/o/profile%2FS__66502658.jpg?alt=media&token=46ea7317-0cad-44d9-af22-6453fc26ee30",
		Phone:    "089602903213",
		Address:  "Bandung",
		Token:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiZXhwIjoxNjQyMDYzNzE3fQ.wOON-dfYvFHmUM5wVVNTeaOf1spuHYnkDiur45jnyJ4",
	}
	courseDomain = course.Domain{
		Id:           1,
		Title:        "Frontend Development",
		Thumbnail:    "https://www.nawpic.com/media/2020/cool-nawpic-5.jpg",
		Description:  "backend Developer",
		Rating:       3.6,
		CategoryId:   1,
		TeacherId:    1,
		DifficultyId: 1,
	}
}

func TestGetAll(t *testing.T) {
	setup()

	t.Run("Test case 1", func(t *testing.T) {
		requestRepository.On("RequestsGetAll",
			mock.Anything).Return(requestsDomain, nil).Once()

		request, err := requestService.RequestsGetAll(context.Background())

		assert.NoError(t, err)
		assert.NotNil(t, request)

	})

	t.Run("Test case 2 | Error", func(t *testing.T) {
		requestRepository.On("RequestsGetAll",
			mock.Anything).Return([]requests.Domain{}, errors.New("Error at Get All Categories")).Once()

		request, err := requestService.RequestsGetAll(context.Background())

		assert.Error(t, err)
		assert.Equal(t, request, []requests.Domain{})
	})
}

func TestGetByCourseId(t *testing.T) {
	setup()

	t.Run("Test case 1", func(t *testing.T) {
		requestRepository.On("EnrollGetByCourseId",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(requestsDomain, nil).Once()

		request, err := requestService.RequestGetById(context.Background(), 1)

		assert.NoError(t, err)
		assert.NotNil(t, request)

	})

	t.Run("Test case 2 | Error", func(t *testing.T) {
		requestRepository.On("EnrollGetByCourseId",
			mock.Anything,
			mock.AnythingOfType("uint")).Return([]requests.Domain{}, errors.New("Error at Get All Categories")).Once()

		request, err := requestService.RequestGetById(context.Background(), 1)

		assert.Error(t, err)
		assert.Equal(t, request, []requests.Domain{})
	})
}
