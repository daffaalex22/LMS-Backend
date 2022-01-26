package teacher_test

import (
	"backend/business/teacher"
	_mockTeacherRepository "backend/business/teacher/mocks"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var teacherRepository _mockTeacherRepository.TeacherRepoInterface
var teacherService teacher.TeacherUseCaseInterface
var teacherDomainPointer *teacher.Domain
var teachersDomain []teacher.Domain
var teacherDomain teacher.Domain

func setup() {
	teacherService = teacher.NewUseCase(&teacherRepository, time.Hour*1 /*, configJWT */)
	teacherDomainPointer = &teacher.Domain{
		Id:         1,
		Name:       "Daffa' Alexander",
		Password:   "$2a$10$6UPRGYrF4cQCO4aHLcGmouO8jXLWW.KkAYiSlU1p3IGSUwZf/kjnq",
		Email:      "drivealex22@gmail.com",
		Avatar:     "https://firebasestorage.googleapis.com/v0/b/learning-management-syst-bf947.appspot.com/o/profile%2FS__66502658.jpg?alt=media&token=46ea7317-0cad-44d9-af22-6453fc26ee30",
		Phone:      "089602903213",
		Address:    "Bandung",
		Token:      "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiZXhwIjoxNjQyMDYzNzE3fQ.wOON-dfYvFHmUM5wVVNTeaOf1spuHYnkDiur45jnyJ4",
		BackGround: "Computer Science Teacher from Harvard",
	}
	teacherDomain = teacher.Domain{
		Id:         1,
		Name:       "Daffa' Alexander",
		Password:   "$2a$10$6UPRGYrF4cQCO4aHLcGmouO8jXLWW.KkAYiSlU1p3IGSUwZf/kjnq",
		Email:      "drivealex22@gmail.com",
		Avatar:     "https://firebasestorage.googleapis.com/v0/b/learning-management-syst-bf947.appspot.com/o/profile%2FS__66502658.jpg?alt=media&token=46ea7317-0cad-44d9-af22-6453fc26ee30",
		Phone:      "089602903213",
		Address:    "Bandung",
		Token:      "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiZXhwIjoxNjQyMDYzNzE3fQ.wOON-dfYvFHmUM5wVVNTeaOf1spuHYnkDiur45jnyJ4",
		BackGround: "Computer Science Teacher from Harvard",
	}
	teachersDomain = append(teachersDomain, teacherDomain)
}

func TestTeacherRegister(t *testing.T) {
	setup()

	t.Run("Test case 1 | Valid TeacherRegister", func(t *testing.T) {
		teacherRepository.On("TeacherRegister",
			mock.AnythingOfType("*teacher.Domain"),
			mock.Anything).Return(teacherDomain, nil).Once()

		tchr, err := teacherService.TeacherRegister(teacherDomainPointer, context.Background())

		assert.Nil(t, err)
		assert.Equal(t, tchr, teacherDomain)
	})

	t.Run("Test case 2 | Error TeacherRegister", func(t *testing.T) {
		teacherRepository.On("TeacherRegister",
			mock.AnythingOfType("*teacher.Domain"),
			mock.Anything).Return(teacher.Domain{}, errors.New("Unexpected Error")).Once()

		tchr, err := teacherService.TeacherRegister(teacherDomainPointer, context.Background())

		assert.Error(t, err)
		assert.Equal(t, tchr, teacher.Domain{})
	})

	t.Run("Test case 3 | Invalid Name Empty", func(t *testing.T) {
		_, err := teacherService.TeacherRegister(&teacher.Domain{
			Name:     "",
			Email:    "alexander@alterra.id",
			Password: "alterra123",
		}, context.Background())

		assert.NotNil(t, err)
	})

	t.Run("Test case 4 | Invalid Email Empty", func(t *testing.T) {
		_, err := teacherService.TeacherRegister(&teacher.Domain{
			Name:     "Daffa' Alexander",
			Email:    "",
			Password: "alterra123",
		}, context.Background())

		assert.NotNil(t, err)
	})

	t.Run("Test case 5 | Invalid Password Empty", func(t *testing.T) {
		_, err := teacherService.TeacherRegister(&teacher.Domain{
			Name:     "Daffa' Alexander",
			Email:    "alexander@alterra.id",
			Password: "",
		}, context.Background())

		assert.NotNil(t, err)
	})
}

func TestUpdate(t *testing.T) {
	setup()

	t.Run("Test case 1 | Valid Update", func(t *testing.T) {
		teacherRepository.On("TeacherUpdate",
			mock.Anything,
			teacherDomain,
			mock.AnythingOfType("uint")).Return(teacherDomain, nil).Once()

		teacherRepository.On("TeacherGetProfile",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(teacherDomain, nil).Once()

		tchr, err := teacherService.TeacherUpdate(context.Background(), teacher.Domain{
			Id:         1,
			Name:       "Daffa' Alexander",
			Password:   "alterra",
			Email:      "drivealex22@gmail.com",
			Avatar:     "https://firebasestorage.googleapis.com/v0/b/learning-management-syst-bf947.appspot.com/o/profile%2FS__66502658.jpg?alt=media&token=46ea7317-0cad-44d9-af22-6453fc26ee30",
			Phone:      "089602903213",
			Address:    "Bandung",
			Token:      "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiZXhwIjoxNjQyMDYzNzE3fQ.wOON-dfYvFHmUM5wVVNTeaOf1spuHYnkDiur45jnyJ4",
			BackGround: "Computer Science Teacher from Harvard",
		}, 1)

		assert.Nil(t, err)
		assert.Equal(t, teacherDomain, tchr)
	})

	t.Run("Test case 2 | Error Update", func(t *testing.T) {
		teacherRepository.On("TeacherUpdate",
			mock.Anything,
			teacherDomain,
			mock.AnythingOfType("uint")).Return(teacherDomain, errors.New("Unexpected Error")).Once()

		teacherRepository.On("TeacherGetProfile",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(teacherDomain, nil).Once()

		tchr, err := teacherService.TeacherUpdate(context.Background(), teacher.Domain{
			Id:         1,
			Name:       "Daffa' Alexander",
			Password:   "alterra",
			Email:      "drivealex22@gmail.com",
			Avatar:     "https://firebasestorage.googleapis.com/v0/b/learning-management-syst-bf947.appspot.com/o/profile%2FS__66502658.jpg?alt=media&token=46ea7317-0cad-44d9-af22-6453fc26ee30",
			Phone:      "089602903213",
			Address:    "Bandung",
			Token:      "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiZXhwIjoxNjQyMDYzNzE3fQ.wOON-dfYvFHmUM5wVVNTeaOf1spuHYnkDiur45jnyJ4",
			BackGround: "Computer Science Teacher from Harvard",
		}, 1)

		assert.Error(t, err)
		assert.Equal(t, tchr, teacher.Domain{})
	})

	t.Run("Test case 3 | Error TeacherGetProfile", func(t *testing.T) {
		teacherRepository.On("TeacherUpdate",
			mock.Anything,
			teacherDomain,
			mock.AnythingOfType("uint")).Return(teacherDomain, nil).Once()

		teacherRepository.On("TeacherGetProfile",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(teacherDomain, errors.New("Unexpected Error")).Once()

		tchr, err := teacherService.TeacherUpdate(context.Background(), teacher.Domain{
			Id:         1,
			Name:       "Daffa' Alexander",
			Password:   "alterra",
			Email:      "drivealex22@gmail.com",
			Avatar:     "https://firebasestorage.googleapis.com/v0/b/learning-management-syst-bf947.appspot.com/o/profile%2FS__66502658.jpg?alt=media&token=46ea7317-0cad-44d9-af22-6453fc26ee30",
			Phone:      "089602903213",
			Address:    "Bandung",
			Token:      "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiZXhwIjoxNjQyMDYzNzE3fQ.wOON-dfYvFHmUM5wVVNTeaOf1spuHYnkDiur45jnyJ4",
			BackGround: "Computer Science Teacher from Harvard",
		}, 1)

		assert.Error(t, err)
		assert.Equal(t, tchr, teacher.Domain{})
	})

	t.Run("Test case 4 | Error Wrong Password", func(t *testing.T) {
		teacherRepository.On("TeacherUpdate",
			mock.Anything,
			teacherDomain,
			mock.AnythingOfType("uint")).Return(teacherDomain, nil).Once()

		teacherRepository.On("TeacherGetProfile",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(teacherDomain, nil).Once()

		tchr, err := teacherService.TeacherUpdate(context.Background(), teacher.Domain{
			Id:         1,
			Name:       "Daffa' Alexander",
			Password:   "Mantab",
			Email:      "drivealex22@gmail.com",
			Avatar:     "https://firebasestorage.googleapis.com/v0/b/learning-management-syst-bf947.appspot.com/o/profile%2FS__66502658.jpg?alt=media&token=46ea7317-0cad-44d9-af22-6453fc26ee30",
			Phone:      "089602903213",
			Address:    "Bandung",
			Token:      "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiZXhwIjoxNjQyMDYzNzE3fQ.wOON-dfYvFHmUM5wVVNTeaOf1spuHYnkDiur45jnyJ4",
			BackGround: "Computer Science Teacher from Harvard",
		}, 1)

		assert.Error(t, err)
		assert.Equal(t, tchr, teacher.Domain{})
	})

	t.Run("Test case 5 | Error Empty Password", func(t *testing.T) {

		tchr, err := teacherService.TeacherUpdate(context.Background(), teacher.Domain{
			Id:         1,
			Name:       "Daffa' Alexander",
			Password:   "",
			Email:      "drivealex22@gmail.com",
			Avatar:     "https://firebasestorage.googleapis.com/v0/b/learning-management-syst-bf947.appspot.com/o/profile%2FS__66502658.jpg?alt=media&token=46ea7317-0cad-44d9-af22-6453fc26ee30",
			Phone:      "089602903213",
			Address:    "Bandung",
			Token:      "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiZXhwIjoxNjQyMDYzNzE3fQ.wOON-dfYvFHmUM5wVVNTeaOf1spuHYnkDiur45jnyJ4",
			BackGround: "Computer Science Teacher from Harvard",
		}, 1)

		assert.Error(t, err)
		assert.Equal(t, tchr, teacher.Domain{})
	})

	t.Run("Test case 6 | Error Empty Name", func(t *testing.T) {

		teacherRepository.On("TeacherGetProfile",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(teacherDomain, nil).Once()

		tchr, err := teacherService.TeacherUpdate(context.Background(), teacher.Domain{
			Id:         1,
			Name:       "",
			Password:   "alterra",
			Email:      "drivealex22@gmail.com",
			Avatar:     "https://firebasestorage.googleapis.com/v0/b/learning-management-syst-bf947.appspot.com/o/profile%2FS__66502658.jpg?alt=media&token=46ea7317-0cad-44d9-af22-6453fc26ee30",
			Phone:      "089602903213",
			Address:    "Bandung",
			Token:      "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiZXhwIjoxNjQyMDYzNzE3fQ.wOON-dfYvFHmUM5wVVNTeaOf1spuHYnkDiur45jnyJ4",
			BackGround: "Computer Science Teacher from Harvard",
		}, 1)

		assert.Error(t, err)
		assert.Equal(t, tchr, teacher.Domain{})
	})

	t.Run("Test case 7 | Error Empty Email", func(t *testing.T) {

		teacherRepository.On("TeacherGetProfile",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(teacherDomain, nil).Once()

		tchr, err := teacherService.TeacherUpdate(context.Background(), teacher.Domain{
			Id:         1,
			Name:       "Daffa' Alex",
			Password:   "alterra",
			Email:      "",
			Avatar:     "https://firebasestorage.googleapis.com/v0/b/learning-management-syst-bf947.appspot.com/o/profile%2FS__66502658.jpg?alt=media&token=46ea7317-0cad-44d9-af22-6453fc26ee30",
			Phone:      "089602903213",
			Address:    "Bandung",
			Token:      "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiZXhwIjoxNjQyMDYzNzE3fQ.wOON-dfYvFHmUM5wVVNTeaOf1spuHYnkDiur45jnyJ4",
			BackGround: "Computer Science Teacher from Harvard",
		}, 1)

		assert.Error(t, err)
		assert.Equal(t, tchr, teacher.Domain{})
	})

	t.Run("Test case 8 | Error Empty Avatar", func(t *testing.T) {

		teacherRepository.On("TeacherGetProfile",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(teacherDomain, nil).Once()

		tchr, err := teacherService.TeacherUpdate(context.Background(), teacher.Domain{
			Id:         1,
			Name:       "Daffa' Alex",
			Password:   "alterra",
			Email:      "drivealex22@gmail.com",
			Avatar:     "",
			Phone:      "089602903213",
			Address:    "Bandung",
			Token:      "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiZXhwIjoxNjQyMDYzNzE3fQ.wOON-dfYvFHmUM5wVVNTeaOf1spuHYnkDiur45jnyJ4",
			BackGround: "Computer Science Teacher from Harvard",
		}, 1)

		assert.Error(t, err)
		assert.Equal(t, tchr, teacher.Domain{})
	})

	t.Run("Test case 9 | Error Empty Address", func(t *testing.T) {

		teacherRepository.On("TeacherGetProfile",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(teacherDomain, nil).Once()

		tchr, err := teacherService.TeacherUpdate(context.Background(), teacher.Domain{
			Id:         1,
			Name:       "Daffa' Alex",
			Password:   "alterra",
			Email:      "drivealex22@gmail.com",
			Avatar:     "https://firebasestorage.googleapis.com/v0/b/learning-management-syst-bf947.appspot.com/o/profile%2FS__66502658.jpg?alt=media&token=46ea7317-0cad-44d9-af22-6453fc26ee30",
			Phone:      "089602903213",
			Address:    "",
			Token:      "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiZXhwIjoxNjQyMDYzNzE3fQ.wOON-dfYvFHmUM5wVVNTeaOf1spuHYnkDiur45jnyJ4",
			BackGround: "Computer Science Teacher from Harvard",
		}, 1)

		assert.Error(t, err)
		assert.Equal(t, tchr, teacher.Domain{})
	})

	t.Run("Test case 10 | Error Empty Phone Number", func(t *testing.T) {

		teacherRepository.On("TeacherGetProfile",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(teacherDomain, nil).Once()

		tchr, err := teacherService.TeacherUpdate(context.Background(), teacher.Domain{
			Id:         1,
			Name:       "Daffa' Alex",
			Password:   "alterra",
			Email:      "drivealex22@gmail.com",
			Avatar:     "https://firebasestorage.googleapis.com/v0/b/learning-management-syst-bf947.appspot.com/o/profile%2FS__66502658.jpg?alt=media&token=46ea7317-0cad-44d9-af22-6453fc26ee30",
			Phone:      "",
			Address:    "Bandung",
			Token:      "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiZXhwIjoxNjQyMDYzNzE3fQ.wOON-dfYvFHmUM5wVVNTeaOf1spuHYnkDiur45jnyJ4",
			BackGround: "Computer Science Teacher from Harvard",
		}, 1)

		assert.Error(t, err)
		assert.Equal(t, tchr, teacher.Domain{})
	})
}

func TestGetById(t *testing.T) {
	setup()

	t.Run("Test case 1", func(t *testing.T) {
		teacherRepository.On("TeacherGetProfile",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(teacherDomain, nil).Once()

		tchr, err := teacherService.TeacherGetProfile(context.Background(), 1)

		assert.NoError(t, err)
		assert.NotNil(t, tchr)

	})

	t.Run("Test case 2 | Error", func(t *testing.T) {
		teacherRepository.On("TeacherGetProfile",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(teacher.Domain{}, errors.New("Error at Get All Categories")).Once()

		tchr, err := teacherService.TeacherGetProfile(context.Background(), 1)

		assert.Error(t, err)
		assert.Equal(t, tchr, teacher.Domain{})
	})
}

func TestLogin(t *testing.T) {
	setup()

	t.Run("Test case 1", func(t *testing.T) {
		teacherRepository.On("TeacherLogin",
			mock.AnythingOfType("teacher.Domain"),
			mock.Anything).Return(teacherDomain, nil).Once()

		tchr, err := teacherService.TeacherLogin(teacherDomain, context.Background())

		assert.NoError(t, err)
		assert.NotNil(t, tchr)

	})

	t.Run("Test case 2 | Error", func(t *testing.T) {
		teacherRepository.On("TeacherLogin",
			mock.AnythingOfType("teacher.Domain"),
			mock.Anything).Return(teacher.Domain{}, errors.New("Error at Get All Categories")).Once()

		tchr, err := teacherService.TeacherLogin(teacherDomain, context.Background())

		assert.Error(t, err)
		assert.Equal(t, tchr, teacher.Domain{})
	})

	t.Run("Test case 3 | Error Emmpty Password", func(t *testing.T) {

		tchr, err := teacherService.TeacherLogin(teacher.Domain{
			Password: "",
			Email:    "drivealex22@gmail.com",
		}, context.Background())

		assert.Error(t, err)
		assert.Equal(t, tchr, teacher.Domain{})
	})

	t.Run("Test case 4 | Error Emmpty Email", func(t *testing.T) {

		tchr, err := teacherService.TeacherLogin(teacher.Domain{
			Password: "alterra",
			Email:    "",
		}, context.Background())

		assert.Error(t, err)
		assert.Equal(t, tchr, teacher.Domain{})
	})
}
