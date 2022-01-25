package enrollments_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"backend/business/course"
	"backend/business/enrollments"
	_mockEnrollmentRepository "backend/business/enrollments/mocks"
	"backend/business/student"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var enrollmentRepository _mockEnrollmentRepository.EnrollmentsRepoInterface
var enrollmentService enrollments.EnrollmentsUseCaseInterface
var enrollmentDomain enrollments.Domain
var enrollmentsDomain []enrollments.Domain
var studentDomain student.Domain
var courseDomain course.Domain

func setup() {
	enrollmentService = enrollments.NewUseCase(&enrollmentRepository, time.Hour*1 /*, configJWT */)
	enrollmentDomain = enrollments.Domain{
		StudentId: 1,
		CourseId:  1,
		Rating:    4,
		Review:    "Mantap Jiwa Coursenya",
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
	enrollmentsDomain = append(enrollmentsDomain, enrollmentDomain)
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
		enrollmentRepository.On("EnrollmentGetAll",
			mock.Anything).Return(enrollmentsDomain, nil).Once()

		enrollment, err := enrollmentService.EnrollmentGetAll(context.Background())

		assert.NoError(t, err)
		assert.NotNil(t, enrollment)

	})

	t.Run("Test case 2 | Error", func(t *testing.T) {
		enrollmentRepository.On("EnrollmentGetAll",
			mock.Anything).Return([]enrollments.Domain{}, errors.New("Error at Get All Categories")).Once()

		enrollment, err := enrollmentService.EnrollmentGetAll(context.Background())

		assert.Error(t, err)
		assert.Equal(t, enrollment, []enrollments.Domain{})
	})
}

func TestAdd(t *testing.T) {
	setup()

	t.Run("Test case 1 | Valid EnrollmentAdd", func(t *testing.T) {
		enrollmentRepository.On("EnrollmentAdd",
			mock.Anything,
			mock.AnythingOfType("enrollments.Domain")).Return(enrollmentDomain, nil).Once()

		enrollmentRepository.On("CheckStudent",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(studentDomain, nil).Once()

		enrollmentRepository.On("CheckCourse",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(courseDomain, nil).Once()

		enroll, err := enrollmentService.EnrollmentAdd(context.Background(), enrollmentDomain)

		assert.Nil(t, err)
		assert.Equal(t, uint(1), enroll.StudentId)
		assert.Equal(t, uint(1), enroll.CourseId)
	})

	t.Run("Test case 2 | Error EnrollmentAdd", func(t *testing.T) {
		enrollmentRepository.On("EnrollmentAdd",
			mock.Anything,
			mock.AnythingOfType("enrollments.Domain")).Return(enrollments.Domain{}, errors.New("Unexpected Error")).Once()

		enrollmentRepository.On("CheckStudent",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(studentDomain, nil).Once()

		enrollmentRepository.On("CheckCourse",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(courseDomain, nil).Once()

		enroll, err := enrollmentService.EnrollmentAdd(context.Background(), enrollmentDomain)

		assert.Error(t, err)
		assert.Equal(t, enroll, enrollments.Domain{})
	})

	t.Run("Test case 3 | Invalid StudentId Empty", func(t *testing.T) {
		_, err := enrollmentService.EnrollmentAdd(context.Background(), enrollments.Domain{
			StudentId: 0,
			CourseId:  1,
		})

		enrollmentRepository.On("CheckStudent",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(studentDomain, nil).Once()

		enrollmentRepository.On("CheckCourse",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(courseDomain, nil).Once()

		assert.NotNil(t, err)
	})

	t.Run("Test case 4 | Invalid CategoryId Empty", func(t *testing.T) {
		_, err := enrollmentService.EnrollmentAdd(context.Background(), enrollments.Domain{
			StudentId: 1,
			CourseId:  0,
		})

		enrollmentRepository.On("CheckStudent",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(studentDomain, nil).Once()

		enrollmentRepository.On("CheckCourse",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(courseDomain, nil).Once()

		assert.NotNil(t, err)
	})

}

func TestUpdate(t *testing.T) {
	setup()

	t.Run("Test case 1 | Valid EnrollUpdate", func(t *testing.T) {
		enrollmentRepository.On("EnrollUpdate",
			mock.Anything,
			mock.AnythingOfType("enrollments.Domain"),
			mock.AnythingOfType("uint"),
			mock.AnythingOfType("uint")).Return(enrollmentDomain, nil).Once()

		enrollmentRepository.On("CheckStudent",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(studentDomain, nil).Once()

		enrollmentRepository.On("CheckCourse",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(courseDomain, nil).Once()

		enroll, err := enrollmentService.EnrollUpdate(context.Background(), enrollmentDomain)

		assert.Nil(t, err)
		assert.Equal(t, uint(1), enroll.StudentId)
		assert.Equal(t, uint(1), enroll.CourseId)
	})

	t.Run("Test case 2 | Error EnrollUpdate", func(t *testing.T) {
		enrollmentRepository.On("EnrollUpdate",
			mock.Anything,
			mock.AnythingOfType("enrollments.Domain"),
			mock.AnythingOfType("uint"),
			mock.AnythingOfType("uint")).Return(enrollments.Domain{}, errors.New("Unexpected Error")).Once()

		enrollmentRepository.On("CheckStudent",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(studentDomain, nil).Once()

		enrollmentRepository.On("CheckCourse",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(courseDomain, nil).Once()

		enroll, err := enrollmentService.EnrollUpdate(context.Background(), enrollmentDomain)

		assert.Error(t, err)
		assert.Equal(t, enroll, enrollments.Domain{})
	})

	t.Run("Test case 3 | Invalid StudentId Empty", func(t *testing.T) {
		_, err := enrollmentService.EnrollUpdate(context.Background(), enrollments.Domain{
			StudentId: 0,
			CourseId:  1,
			Rating:    1,
			Review:    "Mantap Jiwa",
		})

		enrollmentRepository.On("CheckStudent",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(studentDomain, nil).Once()

		enrollmentRepository.On("CheckCourse",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(courseDomain, nil).Once()

		assert.NotNil(t, err)
	})

	t.Run("Test case 4 | Invalid CategoryId Empty", func(t *testing.T) {
		_, err := enrollmentService.EnrollUpdate(context.Background(), enrollments.Domain{
			StudentId: 1,
			CourseId:  0,
			Rating:    1,
			Review:    "Mantap Jiwa",
		})

		enrollmentRepository.On("CheckStudent",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(studentDomain, nil).Once()

		enrollmentRepository.On("CheckCourse",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(courseDomain, nil).Once()

		assert.NotNil(t, err)
	})

	t.Run("Test case 5 | Invalid Rating Empty", func(t *testing.T) {
		_, err := enrollmentService.EnrollUpdate(context.Background(), enrollments.Domain{
			StudentId: 1,
			CourseId:  1,
			Rating:    0,
			Review:    "Mantap Jiwa",
		})

		enrollmentRepository.On("CheckStudent",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(studentDomain, nil).Once()

		enrollmentRepository.On("CheckCourse",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(courseDomain, nil).Once()

		assert.NotNil(t, err)
	})

	t.Run("Test case 6 | Invalid Review Empty", func(t *testing.T) {
		_, err := enrollmentService.EnrollUpdate(context.Background(), enrollments.Domain{
			StudentId: 1,
			CourseId:  1,
			Rating:    1,
			Review:    "",
		})

		enrollmentRepository.On("CheckStudent",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(studentDomain, nil).Once()

		enrollmentRepository.On("CheckCourse",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(courseDomain, nil).Once()

		assert.NotNil(t, err)
	})

}

func TestGetById(t *testing.T) {
	setup()

	t.Run("Test case 1", func(t *testing.T) {
		enrollmentRepository.On("EnrollGetByCourseId",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(enrollmentsDomain, nil).Once()

		enrollment, err := enrollmentService.EnrollGetByCourseId(context.Background(), 1)

		assert.NoError(t, err)
		assert.NotNil(t, enrollment)

	})

	t.Run("Test case 2 | Error", func(t *testing.T) {
		enrollmentRepository.On("EnrollGetByCourseId",
			mock.Anything,
			mock.AnythingOfType("uint")).Return([]enrollments.Domain{}, errors.New("Error at Get All Categories")).Once()

		enrollment, err := enrollmentService.EnrollGetByCourseId(context.Background(), 1)

		assert.Error(t, err)
		assert.Equal(t, enrollment, []enrollments.Domain{})
	})
}
