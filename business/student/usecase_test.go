package student_test

import (
	"backend/business/student"
	_mockStudentRepository "backend/business/student/mocks"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var studentRepository _mockStudentRepository.StudentRepoInterface
var studentService student.StudentUseCaseInterface
var studentDomainPointer *student.Domain
var studentsDomain []student.Domain
var studentDomain student.Domain

func setup() {
	studentService = student.NewUseCase(&studentRepository, time.Hour*1 /*, configJWT */)
	studentDomainPointer = &student.Domain{
		Id:       1,
		Name:     "Daffa' Alexander",
		Password: "$2a$10$6UPRGYrF4cQCO4aHLcGmouO8jXLWW.KkAYiSlU1p3IGSUwZf/kjnq",
		Email:    "drivealex22@gmail.com",
		Avatar:   "https://firebasestorage.googleapis.com/v0/b/learning-management-syst-bf947.appspot.com/o/profile%2FS__66502658.jpg?alt=media&token=46ea7317-0cad-44d9-af22-6453fc26ee30",
		Phone:    "089602903213",
		Address:  "Bandung",
		Token:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiZXhwIjoxNjQyMDYzNzE3fQ.wOON-dfYvFHmUM5wVVNTeaOf1spuHYnkDiur45jnyJ4",
	}
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
	studentsDomain = append(studentsDomain, studentDomain)
}

func TestRegister(t *testing.T) {
	setup()

	t.Run("Test case 1 | Valid Register", func(t *testing.T) {
		studentRepository.On("Register",
			mock.AnythingOfType("*student.Domain"),
			mock.Anything).Return(studentDomain, nil).Once()

		stdt, err := studentService.Register(studentDomainPointer, context.Background())

		assert.Nil(t, err)
		assert.Equal(t, stdt, studentDomain)
	})

	t.Run("Test case 2 | Error Register", func(t *testing.T) {
		studentRepository.On("Register",
			mock.AnythingOfType("*student.Domain"),
			mock.Anything).Return(student.Domain{}, errors.New("Unexpected Error")).Once()

		stdt, err := studentService.Register(studentDomainPointer, context.Background())

		assert.Error(t, err)
		assert.Equal(t, stdt, student.Domain{})
	})

	t.Run("Test case 3 | Invalid Name Empty", func(t *testing.T) {
		_, err := studentService.Register(&student.Domain{
			Name:     "",
			Email:    "alexander@alterra.id",
			Password: "alterra123",
		}, context.Background())

		assert.NotNil(t, err)
	})

	t.Run("Test case 4 | Invalid Email Empty", func(t *testing.T) {
		_, err := studentService.Register(&student.Domain{
			Name:     "Daffa' Alexander",
			Email:    "",
			Password: "alterra123",
		}, context.Background())

		assert.NotNil(t, err)
	})

	t.Run("Test case 5 | Invalid Password Empty", func(t *testing.T) {
		_, err := studentService.Register(&student.Domain{
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
		studentRepository.On("StudentUpdate",
			mock.Anything,
			studentDomain,
			mock.AnythingOfType("uint")).Return(studentDomain, nil).Once()

		studentRepository.On("GetProfile",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(studentDomain, nil).Once()

		stdt, err := studentService.StudentUpdate(context.Background(), student.Domain{
			Id:       1,
			Name:     "Daffa' Alexander",
			Password: "alterra",
			Email:    "drivealex22@gmail.com",
			Avatar:   "https://firebasestorage.googleapis.com/v0/b/learning-management-syst-bf947.appspot.com/o/profile%2FS__66502658.jpg?alt=media&token=46ea7317-0cad-44d9-af22-6453fc26ee30",
			Phone:    "089602903213",
			Address:  "Bandung",
			Token:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiZXhwIjoxNjQyMDYzNzE3fQ.wOON-dfYvFHmUM5wVVNTeaOf1spuHYnkDiur45jnyJ4",
		}, 1)

		assert.Nil(t, err)
		assert.Equal(t, studentDomain, stdt)
	})

	t.Run("Test case 2 | Error Update", func(t *testing.T) {
		studentRepository.On("StudentUpdate",
			mock.Anything,
			studentDomain,
			mock.AnythingOfType("uint")).Return(studentDomain, errors.New("Unexpected Error")).Once()

		studentRepository.On("GetProfile",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(studentDomain, nil).Once()

		stdt, err := studentService.StudentUpdate(context.Background(), student.Domain{
			Id:       1,
			Name:     "Daffa' Alexander",
			Password: "alterra",
			Email:    "drivealex22@gmail.com",
			Avatar:   "https://firebasestorage.googleapis.com/v0/b/learning-management-syst-bf947.appspot.com/o/profile%2FS__66502658.jpg?alt=media&token=46ea7317-0cad-44d9-af22-6453fc26ee30",
			Phone:    "089602903213",
			Address:  "Bandung",
			Token:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiZXhwIjoxNjQyMDYzNzE3fQ.wOON-dfYvFHmUM5wVVNTeaOf1spuHYnkDiur45jnyJ4",
		}, 1)

		assert.Error(t, err)
		assert.Equal(t, stdt, student.Domain{})
	})

	t.Run("Test case 3 | Error GetProfile", func(t *testing.T) {
		studentRepository.On("StudentUpdate",
			mock.Anything,
			studentDomain,
			mock.AnythingOfType("uint")).Return(studentDomain, nil).Once()

		studentRepository.On("GetProfile",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(studentDomain, errors.New("Unexpected Error")).Once()

		stdt, err := studentService.StudentUpdate(context.Background(), student.Domain{
			Id:       1,
			Name:     "Daffa' Alexander",
			Password: "alterra",
			Email:    "drivealex22@gmail.com",
			Avatar:   "https://firebasestorage.googleapis.com/v0/b/learning-management-syst-bf947.appspot.com/o/profile%2FS__66502658.jpg?alt=media&token=46ea7317-0cad-44d9-af22-6453fc26ee30",
			Phone:    "089602903213",
			Address:  "Bandung",
			Token:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiZXhwIjoxNjQyMDYzNzE3fQ.wOON-dfYvFHmUM5wVVNTeaOf1spuHYnkDiur45jnyJ4",
		}, 1)

		assert.Error(t, err)
		assert.Equal(t, stdt, student.Domain{})
	})

	t.Run("Test case 4 | Error Wrong Password", func(t *testing.T) {
		studentRepository.On("StudentUpdate",
			mock.Anything,
			studentDomain,
			mock.AnythingOfType("uint")).Return(studentDomain, nil).Once()

		studentRepository.On("GetProfile",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(studentDomain, nil).Once()

		stdt, err := studentService.StudentUpdate(context.Background(), student.Domain{
			Id:       1,
			Name:     "Daffa' Alexander",
			Password: "Mantab",
			Email:    "drivealex22@gmail.com",
			Avatar:   "https://firebasestorage.googleapis.com/v0/b/learning-management-syst-bf947.appspot.com/o/profile%2FS__66502658.jpg?alt=media&token=46ea7317-0cad-44d9-af22-6453fc26ee30",
			Phone:    "089602903213",
			Address:  "Bandung",
			Token:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiZXhwIjoxNjQyMDYzNzE3fQ.wOON-dfYvFHmUM5wVVNTeaOf1spuHYnkDiur45jnyJ4",
		}, 1)

		assert.Error(t, err)
		assert.Equal(t, stdt, student.Domain{})
	})

	t.Run("Test case 5 | Error Empty Password", func(t *testing.T) {

		stdt, err := studentService.StudentUpdate(context.Background(), student.Domain{
			Id:       1,
			Name:     "Daffa' Alexander",
			Password: "",
			Email:    "drivealex22@gmail.com",
			Avatar:   "https://firebasestorage.googleapis.com/v0/b/learning-management-syst-bf947.appspot.com/o/profile%2FS__66502658.jpg?alt=media&token=46ea7317-0cad-44d9-af22-6453fc26ee30",
			Phone:    "089602903213",
			Address:  "Bandung",
			Token:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiZXhwIjoxNjQyMDYzNzE3fQ.wOON-dfYvFHmUM5wVVNTeaOf1spuHYnkDiur45jnyJ4",
		}, 1)

		assert.Error(t, err)
		assert.Equal(t, stdt, student.Domain{})
	})

	t.Run("Test case 6 | Error Empty Name", func(t *testing.T) {

		studentRepository.On("GetProfile",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(studentDomain, nil).Once()

		stdt, err := studentService.StudentUpdate(context.Background(), student.Domain{
			Id:       1,
			Name:     "",
			Password: "alterra",
			Email:    "drivealex22@gmail.com",
			Avatar:   "https://firebasestorage.googleapis.com/v0/b/learning-management-syst-bf947.appspot.com/o/profile%2FS__66502658.jpg?alt=media&token=46ea7317-0cad-44d9-af22-6453fc26ee30",
			Phone:    "089602903213",
			Address:  "Bandung",
			Token:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiZXhwIjoxNjQyMDYzNzE3fQ.wOON-dfYvFHmUM5wVVNTeaOf1spuHYnkDiur45jnyJ4",
		}, 1)

		assert.Error(t, err)
		assert.Equal(t, stdt, student.Domain{})
	})

	t.Run("Test case 7 | Error Empty Email", func(t *testing.T) {

		studentRepository.On("GetProfile",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(studentDomain, nil).Once()

		stdt, err := studentService.StudentUpdate(context.Background(), student.Domain{
			Id:       1,
			Name:     "Daffa' Alex",
			Password: "alterra",
			Email:    "",
			Avatar:   "https://firebasestorage.googleapis.com/v0/b/learning-management-syst-bf947.appspot.com/o/profile%2FS__66502658.jpg?alt=media&token=46ea7317-0cad-44d9-af22-6453fc26ee30",
			Phone:    "089602903213",
			Address:  "Bandung",
			Token:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiZXhwIjoxNjQyMDYzNzE3fQ.wOON-dfYvFHmUM5wVVNTeaOf1spuHYnkDiur45jnyJ4",
		}, 1)

		assert.Error(t, err)
		assert.Equal(t, stdt, student.Domain{})
	})

	t.Run("Test case 8 | Error Empty Avatar", func(t *testing.T) {

		studentRepository.On("GetProfile",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(studentDomain, nil).Once()

		stdt, err := studentService.StudentUpdate(context.Background(), student.Domain{
			Id:       1,
			Name:     "Daffa' Alex",
			Password: "alterra",
			Email:    "drivealex22@gmail.com",
			Avatar:   "",
			Phone:    "089602903213",
			Address:  "Bandung",
			Token:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiZXhwIjoxNjQyMDYzNzE3fQ.wOON-dfYvFHmUM5wVVNTeaOf1spuHYnkDiur45jnyJ4",
		}, 1)

		assert.Error(t, err)
		assert.Equal(t, stdt, student.Domain{})
	})

	t.Run("Test case 9 | Error Empty Address", func(t *testing.T) {

		studentRepository.On("GetProfile",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(studentDomain, nil).Once()

		stdt, err := studentService.StudentUpdate(context.Background(), student.Domain{
			Id:       1,
			Name:     "Daffa' Alex",
			Password: "alterra",
			Email:    "drivealex22@gmail.com",
			Avatar:   "https://firebasestorage.googleapis.com/v0/b/learning-management-syst-bf947.appspot.com/o/profile%2FS__66502658.jpg?alt=media&token=46ea7317-0cad-44d9-af22-6453fc26ee30",
			Phone:    "089602903213",
			Address:  "",
			Token:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiZXhwIjoxNjQyMDYzNzE3fQ.wOON-dfYvFHmUM5wVVNTeaOf1spuHYnkDiur45jnyJ4",
		}, 1)

		assert.Error(t, err)
		assert.Equal(t, stdt, student.Domain{})
	})

	t.Run("Test case 10 | Error Empty Phone Number", func(t *testing.T) {

		studentRepository.On("GetProfile",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(studentDomain, nil).Once()

		stdt, err := studentService.StudentUpdate(context.Background(), student.Domain{
			Id:       1,
			Name:     "Daffa' Alex",
			Password: "alterra",
			Email:    "drivealex22@gmail.com",
			Avatar:   "https://firebasestorage.googleapis.com/v0/b/learning-management-syst-bf947.appspot.com/o/profile%2FS__66502658.jpg?alt=media&token=46ea7317-0cad-44d9-af22-6453fc26ee30",
			Phone:    "",
			Address:  "Bandung",
			Token:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiZXhwIjoxNjQyMDYzNzE3fQ.wOON-dfYvFHmUM5wVVNTeaOf1spuHYnkDiur45jnyJ4",
		}, 1)

		assert.Error(t, err)
		assert.Equal(t, stdt, student.Domain{})
	})
}

func TestGetById(t *testing.T) {
	setup()

	t.Run("Test case 1", func(t *testing.T) {
		studentRepository.On("GetProfile",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(studentDomain, nil).Once()

		stdt, err := studentService.GetProfile(context.Background(), 1)

		assert.NoError(t, err)
		assert.NotNil(t, stdt)

	})

	t.Run("Test case 2 | Error", func(t *testing.T) {
		studentRepository.On("GetProfile",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(student.Domain{}, errors.New("Error at Get All Categories")).Once()

		stdt, err := studentService.GetProfile(context.Background(), 1)

		assert.Error(t, err)
		assert.Equal(t, stdt, student.Domain{})
	})
}

func TestLogin(t *testing.T) {
	setup()

	t.Run("Test case 1", func(t *testing.T) {
		studentRepository.On("Login",
			mock.AnythingOfType("student.Domain"),
			mock.Anything).Return(studentDomain, nil).Once()

		stdt, err := studentService.Login(studentDomain, context.Background())

		assert.NoError(t, err)
		assert.NotNil(t, stdt)

	})

	t.Run("Test case 2 | Error", func(t *testing.T) {
		studentRepository.On("Login",
			mock.AnythingOfType("student.Domain"),
			mock.Anything).Return(student.Domain{}, errors.New("Error at Get All Categories")).Once()

		stdt, err := studentService.Login(studentDomain, context.Background())

		assert.Error(t, err)
		assert.Equal(t, stdt, student.Domain{})
	})

	t.Run("Test case 3 | Error Emmpty Password", func(t *testing.T) {

		stdt, err := studentService.Login(student.Domain{
			Password: "",
			Email:    "drivealex22@gmail.com",
		}, context.Background())

		assert.Error(t, err)
		assert.Equal(t, stdt, student.Domain{})
	})

	t.Run("Test case 4 | Error Emmpty Email", func(t *testing.T) {

		stdt, err := studentService.Login(student.Domain{
			Password: "alterra",
			Email:    "",
		}, context.Background())

		assert.Error(t, err)
		assert.Equal(t, stdt, student.Domain{})
	})
}
