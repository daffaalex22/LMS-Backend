package readings_test

import (
	"backend/business/modules"
	"backend/business/readings"
	_mockReadingRepository "backend/business/readings/mocks"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var readingRepository _mockReadingRepository.ReadingsRepoInterface
var readingService readings.ReadingsUseCaseInterface
var readingDomain readings.Domain
var readingsDomain []readings.Domain
var moduleDomain modules.Domain

func setup() {
	readingService = readings.NewUseCase(&readingRepository, time.Hour*1 /*, configJWT */)
	readingDomain = readings.Domain{
		Id:         1,
		ModuleId:   1,
		Module:     moduleDomain,
		Title:      "Next.js Tutorial #2 - Pages & Routes",
		Content:    "Hey gang, in this Next.js tutorial series you'll learn how to create a website with Next (& React) - including pages, routes, layouts, fetching data & deployment.",
		Order:      1,
		Quiz:       "https://firebasestorage.googleapis.com/v0/b/learning-management-syst-bf947.appspot.com/o/profile%2FTHT%20Transient%20Vibration%20Getmek.docx?alt=media&token=d8bf8d8a-3e36-439c-bc0c-963c4552e295",
		Attachment: "https://firebasestorage.googleapis.com/v0/b/learning-management-syst-bf947.appspot.com/o/profile%2FTHT%20Transient%20Vibration%20Getmek.docx?alt=media&token=d8bf8d8a-3e36-439c-bc0c-963c4552e295",
	}
	readingsDomain = append(readingsDomain, readingDomain)
	moduleDomain = modules.Domain{
		Id:       1,
		Title:    "Frontend Development",
		CourseId: 1,
		Order:    1,
	}
}

func TestGetById(t *testing.T) {
	setup()

	t.Run("Test case 1", func(t *testing.T) {
		readingRepository.On("ReadingsGetById",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(readingDomain, nil).Once()

		video, err := readingService.ReadingsGetById(context.Background(), 1)

		assert.NoError(t, err)
		assert.NotNil(t, video)

	})

	t.Run("Test case 2 | Error", func(t *testing.T) {
		readingRepository.On("ReadingsGetById",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(readings.Domain{}, errors.New("Error at Get All Categories")).Once()

		video, err := readingService.ReadingsGetById(context.Background(), 1)

		assert.Error(t, err)
		assert.Equal(t, video, readings.Domain{})
	})
}

func TestDelete(t *testing.T) {
	setup()

	t.Run("Test case 1", func(t *testing.T) {
		readingRepository.On("ReadingsDelete",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(nil).Once()

		err := readingService.ReadingsDelete(context.Background(), 1)

		assert.NoError(t, err)
	})

	t.Run("Test case 2 | Error", func(t *testing.T) {
		readingRepository.On("ReadingsDelete",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(errors.New("Error at Delete Readings")).Once()

		err := readingService.ReadingsDelete(context.Background(), 1)

		assert.Error(t, err)
	})
}

func TestGetByModuleId(t *testing.T) {
	setup()

	t.Run("Test case 1", func(t *testing.T) {
		readingRepository.On("ReadingsGetByModuleId",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(readingsDomain, nil).Once()

		vds, err := readingService.ReadingsGetByModuleId(context.Background(), 1)

		assert.NoError(t, err)
		assert.NotNil(t, vds)
	})

	t.Run("Test case 2 | Error", func(t *testing.T) {
		readingRepository.On("ReadingsGetByModuleId",
			mock.Anything,
			mock.AnythingOfType("uint")).Return([]readings.Domain{}, errors.New("Error at Get Readings By Module Id")).Once()

		vds, err := readingService.ReadingsGetByModuleId(context.Background(), 1)

		assert.Error(t, err)
		assert.Equal(t, vds, []readings.Domain{})
	})
}

func TestAdd(t *testing.T) {
	setup()

	t.Run("Test case 1 | Valid ReadingsAdd", func(t *testing.T) {
		readingRepository.On("ReadingsAdd",
			mock.Anything,
			mock.AnythingOfType("readings.Domain")).Return(readingDomain, nil).Once()

		readingRepository.On("CheckModule",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(moduleDomain, nil).Once()

		vdo, err := readingService.ReadingsAdd(context.Background(), readingDomain)

		assert.Nil(t, err)
		assert.Equal(t, uint(1), vdo.ModuleId)
		assert.Equal(t, 1, vdo.Order)
	})

	t.Run("Test case 2 | Error ReadingsAdd", func(t *testing.T) {
		readingRepository.On("ReadingsAdd",
			mock.Anything,
			mock.AnythingOfType("readings.Domain")).Return(readings.Domain{}, errors.New("Unexpected Error")).Once()

		readingRepository.On("CheckModule",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(moduleDomain, nil).Once()

		vdo, err := readingService.ReadingsAdd(context.Background(), readingDomain)

		assert.Error(t, err)
		assert.Equal(t, vdo, readings.Domain{})
	})

	t.Run("Test case 3 | Invalid ModuleId Empty", func(t *testing.T) {
		_, err := readingService.ReadingsAdd(context.Background(), readings.Domain{
			ModuleId: 0,
			Title:    "Ilustrasi Hk Newton",
			Content:  "Sebuah Kotak",
			Order:    1,
		})

		readingRepository.On("CheckModule",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(moduleDomain, nil).Once()

		assert.NotNil(t, err)
	})

	t.Run("Test case 4 | Invalid Title Empty", func(t *testing.T) {
		_, err := readingService.ReadingsAdd(context.Background(), readings.Domain{
			ModuleId: 1,
			Title:    "",
			Content:  "Sebuah Kotak",
			Order:    1,
		})

		readingRepository.On("CheckModule",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(moduleDomain, nil).Once()

		assert.NotNil(t, err)
	})

	t.Run("Test case 5 | Invalid Content Empty", func(t *testing.T) {
		_, err := readingService.ReadingsAdd(context.Background(), readings.Domain{
			ModuleId: 1,
			Title:    "Ilustrasi Hk Newton",
			Content:  "",
			Order:    1,
		})

		readingRepository.On("CheckModule",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(moduleDomain, nil).Once()

		assert.NotNil(t, err)
	})

	t.Run("Test case 6 | Invalid Order Empty", func(t *testing.T) {
		_, err := readingService.ReadingsAdd(context.Background(), readings.Domain{
			ModuleId: 1,
			Title:    "Ilustrasi Hk Newton",
			Content:  "Sebuah Kotak",
			Order:    0,
		})

		readingRepository.On("CheckModule",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(moduleDomain, nil).Once()

		assert.NotNil(t, err)
	})

	t.Run("Test case 7 | Invalid CheckModule Failed", func(t *testing.T) {
		readingRepository.On("ReadingsAdd",
			mock.Anything,
			mock.AnythingOfType("readings.Domain")).Return(readingDomain, errors.New("Unexpected Error")).Once()

		readingRepository.On("CheckModule",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(modules.Domain{}, errors.New("Unexpected Error")).Once()

		vdo, err := readingService.ReadingsAdd(context.Background(), readingDomain)

		assert.Error(t, err)
		assert.Equal(t, vdo, readings.Domain{})
	})
}

func TestUpdate(t *testing.T) {
	setup()

	t.Run("Test case 1 | Valid ReadingsUpdate", func(t *testing.T) {
		readingRepository.On("ReadingsUpdate",
			mock.Anything,
			mock.AnythingOfType("readings.Domain"),
			mock.AnythingOfType("uint")).Return(readingDomain, nil).Once()

		readingRepository.On("CheckModule",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(moduleDomain, nil).Once()

		vdo, err := readingService.ReadingsUpdate(context.Background(), readingDomain, 1)

		assert.Nil(t, err)
		assert.Equal(t, uint(1), vdo.ModuleId)
		assert.Equal(t, 1, vdo.Order)
	})

	t.Run("Test case 2 | Error ReadingsUpdate", func(t *testing.T) {
		readingRepository.On("ReadingsUpdate",
			mock.Anything,
			mock.AnythingOfType("readings.Domain"),
			mock.AnythingOfType("uint")).Return(readings.Domain{}, errors.New("Unexpected Error")).Once()

		readingRepository.On("CheckModule",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(moduleDomain, nil).Once()

		vdo, err := readingService.ReadingsUpdate(context.Background(), readingDomain, 1)

		assert.Error(t, err)
		assert.Equal(t, vdo, readings.Domain{})
	})

	t.Run("Test case 3 | Invalid ModuleId Empty", func(t *testing.T) {
		_, err := readingService.ReadingsUpdate(context.Background(), readings.Domain{
			ModuleId: 0,
			Title:    "Ilustrasi Hk Newton",
			Content:  "Sebuah Kotak",
			Order:    1,
		}, 1)

		readingRepository.On("CheckModule",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(moduleDomain, nil).Once()

		assert.NotNil(t, err)
	})

	t.Run("Test case 4 | Invalid Title Empty", func(t *testing.T) {
		_, err := readingService.ReadingsUpdate(context.Background(), readings.Domain{
			ModuleId: 1,
			Title:    "",
			Content:  "Sebuah Kotak",
			Order:    1,
		}, 1)

		readingRepository.On("CheckModule",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(moduleDomain, nil).Once()

		assert.NotNil(t, err)
	})

	t.Run("Test case 5 | Invalid Content Empty", func(t *testing.T) {
		_, err := readingService.ReadingsUpdate(context.Background(), readings.Domain{
			ModuleId: 1,
			Title:    "Ilustrasi Hk Newton",
			Content:  "",
			Order:    1,
		}, 1)

		readingRepository.On("CheckModule",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(moduleDomain, nil).Once()

		assert.NotNil(t, err)
	})

	t.Run("Test case 6 | Invalid Order Empty", func(t *testing.T) {
		_, err := readingService.ReadingsUpdate(context.Background(), readings.Domain{
			ModuleId: 1,
			Title:    "Ilustrasi Hk Newton",
			Content:  "Sebuah Kotak",
			Order:    0,
		}, 1)

		readingRepository.On("CheckModule",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(moduleDomain, nil).Once()

		assert.NotNil(t, err)
	})

	t.Run("Test case 7 | Invalid CheckModule Failed", func(t *testing.T) {
		readingRepository.On("ReadingsUpdate",
			mock.Anything,
			mock.AnythingOfType("readings.Domain"),
			mock.AnythingOfType("uint")).Return(readingDomain, errors.New("Unexpected Error")).Once()

		readingRepository.On("CheckModule",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(modules.Domain{}, errors.New("Unexpected Error")).Once()

		vdo, err := readingService.ReadingsUpdate(context.Background(), readingDomain, 1)

		assert.Error(t, err)
		assert.Equal(t, vdo, readings.Domain{})
	})
}
