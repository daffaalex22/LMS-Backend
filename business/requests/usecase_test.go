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
var coursesDomain []course.Domain

func setup() {
	requestService = requests.NewUseCase(&requestRepository, time.Hour*1 /*, configJWT */)
	requestDomain = requests.Domain{
		StudentId: 1,
		CourseId:  1,
		TypeId:    1,
		Status:    "Pending",
		Message:   "Request Bro",
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
	coursesDomain = append(coursesDomain, courseDomain)
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

func TestGetById(t *testing.T) {
	setup()

	t.Run("Test case 1", func(t *testing.T) {
		requestRepository.On("RequestGetById",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(requestDomain, nil).Once()

		request, err := requestService.RequestGetById(context.Background(), 1)

		assert.NoError(t, err)
		assert.NotNil(t, request)

	})

	t.Run("Test case 2 | Error", func(t *testing.T) {
		requestRepository.On("RequestGetById",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(requests.Domain{}, errors.New("Error at Get All Categories")).Once()

		request, err := requestService.RequestGetById(context.Background(), 1)

		assert.Error(t, err)
		assert.Equal(t, request, requests.Domain{})
	})
}

func TestGetByStudentId(t *testing.T) {
	setup()

	t.Run("Test case 1", func(t *testing.T) {
		requestRepository.On("RequestsGetByStudentId",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(requestsDomain, nil).Once()

		request, err := requestService.RequestsGetByStudentId(context.Background(), 1)

		assert.NoError(t, err)
		assert.NotNil(t, request)

	})

	t.Run("Test case 2 | Error", func(t *testing.T) {
		requestRepository.On("RequestsGetByStudentId",
			mock.Anything,
			mock.AnythingOfType("uint")).Return([]requests.Domain{}, errors.New("Error at Get All Categories")).Once()

		request, err := requestService.RequestsGetByStudentId(context.Background(), 1)

		assert.Error(t, err)
		assert.Equal(t, request, []requests.Domain{})
	})
}

func TestGetByCourseId(t *testing.T) {
	setup()

	t.Run("Test case 1", func(t *testing.T) {
		requestRepository.On("RequestsGetByCourseId",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(requestsDomain, nil).Once()

		request, err := requestService.RequestsGetByCourseId(context.Background(), 1)

		assert.NoError(t, err)
		assert.NotNil(t, request)

	})

	t.Run("Test case 2 | Error", func(t *testing.T) {
		requestRepository.On("RequestsGetByCourseId",
			mock.Anything,
			mock.AnythingOfType("uint")).Return([]requests.Domain{}, errors.New("Error at Get All Categories")).Once()

		request, err := requestService.RequestsGetByCourseId(context.Background(), 1)

		assert.Error(t, err)
		assert.Equal(t, request, []requests.Domain{})
	})
}

func TestGetByTeacherId(t *testing.T) {
	setup()

	t.Run("Test case 1", func(t *testing.T) {
		requestRepository.On("RequestsGetByTeacherId",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(requestsDomain, nil).Once()

		requestRepository.On("GetCoursesByTeacherId",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(coursesDomain, nil).Once()

		requestRepository.On("RequestsGetByCourseIds",
			mock.Anything,
			mock.AnythingOfType("[]uint")).Return(requestsDomain, nil).Once()

		request, err := requestService.RequestsGetByTeacherId(context.Background(), 1)

		assert.NoError(t, err)
		assert.NotNil(t, request)

	})

	t.Run("Test case 2 | Error", func(t *testing.T) {
		requestRepository.On("RequestsGetByTeacherId",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(requestsDomain, nil).Once()

		requestRepository.On("GetCoursesByTeacherId",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(coursesDomain, nil).Once()

		requestRepository.On("RequestsGetByCourseIds",
			mock.Anything,
			mock.AnythingOfType("[]uint")).Return([]requests.Domain{}, errors.New("Error at Get All Categories")).Once()

		request, err := requestService.RequestsGetByTeacherId(context.Background(), 1)

		assert.Error(t, err)
		assert.Equal(t, request, []requests.Domain{})
	})

	t.Run("Test case 3 | Error", func(t *testing.T) {
		requestRepository.On("RequestsGetByTeacherId",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(requestsDomain, nil).Once()

		requestRepository.On("GetCoursesByTeacherId",
			mock.Anything,
			mock.AnythingOfType("uint")).Return([]course.Domain{}, errors.New("Unexpected Error")).Once()

		request, err := requestService.RequestsGetByTeacherId(context.Background(), 1)

		assert.Error(t, err)
		assert.Equal(t, request, []requests.Domain{})
	})
}

func TestUpdate(t *testing.T) {
	setup()

	t.Run("Test case 1", func(t *testing.T) {
		requestRepository.On("RequestsUpdate",
			mock.Anything,
			mock.AnythingOfType("requests.Domain"),
			mock.AnythingOfType("uint")).Return(requestDomain, nil).Once()

		request, err := requestService.RequestsUpdate(context.Background(), requestDomain, 1)

		assert.NoError(t, err)
		assert.NotNil(t, request)

	})

	t.Run("Test case 2 | Error", func(t *testing.T) {
		requestRepository.On("RequestsUpdate",
			mock.Anything,
			mock.AnythingOfType("requests.Domain"),
			mock.AnythingOfType("uint")).Return(requests.Domain{}, errors.New("Error at Get All Categories")).Once()

		request, err := requestService.RequestsUpdate(context.Background(), requestDomain, 1)

		assert.Error(t, err)
		assert.Equal(t, request, requests.Domain{})
	})
}

func TestAdd(t *testing.T) {
	setup()

	t.Run("Test case 1 | Valid RequestsAdd", func(t *testing.T) {
		requestRepository.On("RequestsAdd",
			mock.Anything,
			mock.AnythingOfType("requests.Domain")).Return(requestDomain, nil).Once()

		requestRepository.On("CheckStudent",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(studentDomain, nil).Once()

		requestRepository.On("CheckCourse",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(courseDomain, nil).Once()

		enroll, err := requestService.RequestsAdd(context.Background(), requestDomain)

		assert.Nil(t, err)
		assert.Equal(t, uint(1), enroll.StudentId)
		assert.Equal(t, uint(1), enroll.CourseId)
	})

	t.Run("Test case 2 | Error RequestsAdd", func(t *testing.T) {
		requestRepository.On("RequestsAdd",
			mock.Anything,
			mock.AnythingOfType("requests.Domain")).Return(requests.Domain{}, errors.New("Unexpected Error")).Once()

		requestRepository.On("CheckStudent",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(studentDomain, nil).Once()

		requestRepository.On("CheckCourse",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(courseDomain, nil).Once()

		enroll, err := requestService.RequestsAdd(context.Background(), requestDomain)

		assert.Error(t, err)
		assert.Equal(t, enroll, requests.Domain{})
	})

	t.Run("Test case 3 | Invalid StudentId Empty", func(t *testing.T) {
		_, err := requestService.RequestsAdd(context.Background(), requests.Domain{
			StudentId: 0,
			CourseId:  1,
			TypeId:    1,
			Status:    "Pending",
			Message:   "Request Bro",
		})

		requestRepository.On("CheckStudent",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(studentDomain, nil).Once()

		requestRepository.On("CheckCourse",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(courseDomain, nil).Once()

		assert.NotNil(t, err)
	})

	t.Run("Test case 4 | Invalid CategoryId Empty", func(t *testing.T) {
		_, err := requestService.RequestsAdd(context.Background(), requests.Domain{
			StudentId: 1,
			CourseId:  0,
			TypeId:    1,
			Status:    "Pending",
			Message:   "Request Bro",
		})

		requestRepository.On("CheckStudent",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(studentDomain, nil).Once()

		requestRepository.On("CheckCourse",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(courseDomain, nil).Once()

		assert.NotNil(t, err)
	})

	// t.Run("Test case 5| Invalid TypeId Empty", func(t *testing.T) {
	// 	_, err := requestService.RequestsAdd(context.Background(), requests.Domain{
	// 		StudentId: 1,
	// 		CourseId:  1,
	// 		TypeId:    0,
	// 		Status:    "Pending",
	// 		Message:   "Request Bro",
	// 	})

	// 	requestRepository.On("CheckStudent",
	// 		mock.Anything,
	// 		mock.AnythingOfType("uint")).Return(studentDomain, nil).Once()

	// 	requestRepository.On("CheckCourse",
	// 		mock.Anything,
	// 		mock.AnythingOfType("uint")).Return(courseDomain, nil).Once()

	// 	assert.NotNil(t, err)
	// })

	t.Run("Test case 6 | Invalid Status Empty", func(t *testing.T) {
		_, err := requestService.RequestsAdd(context.Background(), requests.Domain{
			StudentId: 1,
			CourseId:  1,
			TypeId:    1,
			Status:    "",
			Message:   "Request Bro",
		})

		requestRepository.On("CheckStudent",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(studentDomain, nil).Once()

		requestRepository.On("CheckCourse",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(courseDomain, nil).Once()

		assert.NotNil(t, err)
	})

	t.Run("Test case 7 | Invalid Message Empty", func(t *testing.T) {
		_, err := requestService.RequestsAdd(context.Background(), requests.Domain{
			StudentId: 1,
			CourseId:  1,
			TypeId:    1,
			Status:    "Pending",
			Message:   "",
		})

		requestRepository.On("CheckStudent",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(studentDomain, nil).Once()

		requestRepository.On("CheckCourse",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(courseDomain, nil).Once()

		assert.NotNil(t, err)
	})

}
