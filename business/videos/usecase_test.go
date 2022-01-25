package videos_test

import (
	"backend/business/modules"
	"backend/business/videos"
	_mockVideoRepository "backend/business/videos/mocks"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var videoRepository _mockVideoRepository.VideosRepoInterface
var videoService videos.VideosUseCaseInterface
var videoDomain videos.Domain
var videosDomain []videos.Domain
var moduleDomain modules.Domain

func setup() {
	videoService = videos.NewUseCase(&videoRepository, time.Hour*1 /*, configJWT */)
	videoDomain = videos.Domain{
		Id:         1,
		ModuleId:   1,
		Module:     moduleDomain,
		Title:      "Next.js Tutorial #2 - Pages & Routes",
		Caption:    "Hey gang, in this Next.js tutorial series you'll learn how to create a website with Next (& React) - including pages, routes, layouts, fetching data & deployment.",
		Url:        "zktJ8-k0JDc",
		Order:      1,
		Quiz:       "https://firebasestorage.googleapis.com/v0/b/learning-management-syst-bf947.appspot.com/o/profile%2FTHT%20Transient%20Vibration%20Getmek.docx?alt=media&token=d8bf8d8a-3e36-439c-bc0c-963c4552e295",
		Attachment: "https://firebasestorage.googleapis.com/v0/b/learning-management-syst-bf947.appspot.com/o/profile%2FTHT%20Transient%20Vibration%20Getmek.docx?alt=media&token=d8bf8d8a-3e36-439c-bc0c-963c4552e295",
	}
	videosDomain = append(videosDomain, videoDomain)
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
		videoRepository.On("VideosGetById",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(videoDomain, nil).Once()

		video, err := videoService.VideosGetById(context.Background(), 1)

		assert.NoError(t, err)
		assert.NotNil(t, video)

	})

	t.Run("Test case 2 | Error", func(t *testing.T) {
		videoRepository.On("VideosGetById",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(videos.Domain{}, errors.New("Error at Get All Categories")).Once()

		video, err := videoService.VideosGetById(context.Background(), 1)

		assert.Error(t, err)
		assert.Equal(t, video, videos.Domain{})
	})
}

func TestDelete(t *testing.T) {
	setup()

	t.Run("Test case 1", func(t *testing.T) {
		videoRepository.On("VideosDelete",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(nil).Once()

		err := videoService.VideosDelete(context.Background(), 1)

		assert.NoError(t, err)
	})

	t.Run("Test case 2 | Error", func(t *testing.T) {
		videoRepository.On("VideosDelete",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(errors.New("Error at Delete Videos")).Once()

		err := videoService.VideosDelete(context.Background(), 1)

		assert.Error(t, err)
	})
}

func TestGetByModuleId(t *testing.T) {
	setup()

	t.Run("Test case 1", func(t *testing.T) {
		videoRepository.On("VideosGetByModuleId",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(videosDomain, nil).Once()

		vds, err := videoService.VideosGetByModuleId(context.Background(), 1)

		assert.NoError(t, err)
		assert.NotNil(t, vds)
	})

	t.Run("Test case 2 | Error", func(t *testing.T) {
		videoRepository.On("VideosGetByModuleId",
			mock.Anything,
			mock.AnythingOfType("uint")).Return([]videos.Domain{}, errors.New("Error at Get Videos By Module Id")).Once()

		vds, err := videoService.VideosGetByModuleId(context.Background(), 1)

		assert.Error(t, err)
		assert.Equal(t, vds, []videos.Domain{})
	})
}

func TestAdd(t *testing.T) {
	setup()

	t.Run("Test case 1 | Valid VideosAdd", func(t *testing.T) {
		videoRepository.On("VideosAdd",
			mock.Anything,
			mock.AnythingOfType("videos.Domain")).Return(videoDomain, nil).Once()

		videoRepository.On("CheckModule",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(moduleDomain, nil).Once()

		vdo, err := videoService.VideosAdd(context.Background(), videoDomain)

		assert.Nil(t, err)
		assert.Equal(t, uint(1), vdo.ModuleId)
		assert.Equal(t, 1, vdo.Order)
	})

	t.Run("Test case 2 | Error VideosAdd", func(t *testing.T) {
		videoRepository.On("VideosAdd",
			mock.Anything,
			mock.AnythingOfType("videos.Domain")).Return(videos.Domain{}, errors.New("Unexpected Error")).Once()

		videoRepository.On("CheckModule",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(moduleDomain, nil).Once()

		vdo, err := videoService.VideosAdd(context.Background(), videoDomain)

		assert.Error(t, err)
		assert.Equal(t, vdo, videos.Domain{})
	})

	t.Run("Test case 3 | Invalid ModuleId Empty", func(t *testing.T) {
		_, err := videoService.VideosAdd(context.Background(), videos.Domain{
			ModuleId: 0,
			Title:    "Ilustrasi Hk Newton",
			Caption:  "Sebuah Kotak",
			Url:      "google.com",
			Order:    1,
		})

		videoRepository.On("CheckModule",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(moduleDomain, nil).Once()

		assert.NotNil(t, err)
	})

	t.Run("Test case 4 | Invalid Title Empty", func(t *testing.T) {
		_, err := videoService.VideosAdd(context.Background(), videos.Domain{
			ModuleId: 1,
			Title:    "",
			Caption:  "Sebuah Kotak",
			Url:      "google.com",
			Order:    1,
		})

		videoRepository.On("CheckModule",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(moduleDomain, nil).Once()

		assert.NotNil(t, err)
	})

	t.Run("Test case 5 | Invalid Caption Empty", func(t *testing.T) {
		_, err := videoService.VideosAdd(context.Background(), videos.Domain{
			ModuleId: 1,
			Title:    "Ilustrasi Hk Newton",
			Caption:  "",
			Url:      "google.com",
			Order:    1,
		})

		videoRepository.On("CheckModule",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(moduleDomain, nil).Once()

		assert.NotNil(t, err)
	})

	t.Run("Test case 6 | Invalid Url Empty", func(t *testing.T) {
		_, err := videoService.VideosAdd(context.Background(), videos.Domain{
			ModuleId: 1,
			Title:    "Ilustrasi Hk Newton",
			Caption:  "Sebuah Kotak",
			Url:      "",
			Order:    1,
		})

		videoRepository.On("CheckModule",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(moduleDomain, nil).Once()

		assert.NotNil(t, err)
	})

	t.Run("Test case 7 | Invalid Order Empty", func(t *testing.T) {
		_, err := videoService.VideosAdd(context.Background(), videos.Domain{
			ModuleId: 1,
			Title:    "Ilustrasi Hk Newton",
			Caption:  "Sebuah Kotak",
			Url:      "https://www.google.com",
			Order:    0,
		})

		videoRepository.On("CheckModule",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(moduleDomain, nil).Once()

		assert.NotNil(t, err)
	})

	t.Run("Test case 8 | Invalid CheckModule Failed", func(t *testing.T) {
		videoRepository.On("VideosAdd",
			mock.Anything,
			mock.AnythingOfType("videos.Domain")).Return(videoDomain, errors.New("Unexpected Error")).Once()

		videoRepository.On("CheckModule",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(modules.Domain{}, errors.New("Unexpected Error")).Once()

		vdo, err := videoService.VideosAdd(context.Background(), videoDomain)

		assert.Error(t, err)
		assert.Equal(t, vdo, videos.Domain{})
	})
}

func TestUpdate(t *testing.T) {
	setup()

	t.Run("Test case 1 | Valid VideosUpdate", func(t *testing.T) {
		videoRepository.On("VideosUpdate",
			mock.Anything,
			mock.AnythingOfType("videos.Domain"),
			mock.AnythingOfType("uint")).Return(videoDomain, nil).Once()

		videoRepository.On("CheckModule",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(moduleDomain, nil).Once()

		vdo, err := videoService.VideosUpdate(context.Background(), videoDomain, 1)

		assert.Nil(t, err)
		assert.Equal(t, uint(1), vdo.ModuleId)
		assert.Equal(t, 1, vdo.Order)
	})

	t.Run("Test case 2 | Error VideosUpdate", func(t *testing.T) {
		videoRepository.On("VideosUpdate",
			mock.Anything,
			mock.AnythingOfType("videos.Domain"),
			mock.AnythingOfType("uint")).Return(videos.Domain{}, errors.New("Unexpected Error")).Once()

		videoRepository.On("CheckModule",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(moduleDomain, nil).Once()

		vdo, err := videoService.VideosUpdate(context.Background(), videoDomain, 1)

		assert.Error(t, err)
		assert.Equal(t, vdo, videos.Domain{})
	})

	t.Run("Test case 3 | Invalid ModuleId Empty", func(t *testing.T) {
		_, err := videoService.VideosUpdate(context.Background(), videos.Domain{
			ModuleId: 0,
			Title:    "Ilustrasi Hk Newton",
			Caption:  "Sebuah Kotak",
			Url:      "google.com",
			Order:    1,
		}, 1)

		videoRepository.On("CheckModule",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(moduleDomain, nil).Once()

		assert.NotNil(t, err)
	})

	t.Run("Test case 4 | Invalid Title Empty", func(t *testing.T) {
		_, err := videoService.VideosUpdate(context.Background(), videos.Domain{
			ModuleId: 1,
			Title:    "",
			Caption:  "Sebuah Kotak",
			Url:      "google.com",
			Order:    1,
		}, 1)

		videoRepository.On("CheckModule",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(moduleDomain, nil).Once()

		assert.NotNil(t, err)
	})

	t.Run("Test case 5 | Invalid Caption Empty", func(t *testing.T) {
		_, err := videoService.VideosUpdate(context.Background(), videos.Domain{
			ModuleId: 1,
			Title:    "Ilustrasi Hk Newton",
			Caption:  "",
			Url:      "google.com",
			Order:    1,
		}, 1)

		videoRepository.On("CheckModule",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(moduleDomain, nil).Once()

		assert.NotNil(t, err)
	})

	t.Run("Test case 6 | Invalid Url Empty", func(t *testing.T) {
		_, err := videoService.VideosUpdate(context.Background(), videos.Domain{
			ModuleId: 1,
			Title:    "Ilustrasi Hk Newton",
			Caption:  "Sebuah Kotak",
			Url:      "",
			Order:    1,
		}, 1)

		videoRepository.On("CheckModule",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(moduleDomain, nil).Once()

		assert.NotNil(t, err)
	})

	t.Run("Test case 7 | Invalid Order Empty", func(t *testing.T) {
		_, err := videoService.VideosUpdate(context.Background(), videos.Domain{
			ModuleId: 1,
			Title:    "Ilustrasi Hk Newton",
			Caption:  "Sebuah Kotak",
			Url:      "https://www.google.com",
			Order:    0,
		}, 1)

		videoRepository.On("CheckModule",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(moduleDomain, nil).Once()

		assert.NotNil(t, err)
	})

	t.Run("Test case 8 | Invalid CheckModule Failed", func(t *testing.T) {
		videoRepository.On("VideosUpdate",
			mock.Anything,
			mock.AnythingOfType("videos.Domain"),
			mock.AnythingOfType("uint")).Return(videoDomain, errors.New("Unexpected Error")).Once()

		videoRepository.On("CheckModule",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(modules.Domain{}, errors.New("Unexpected Error")).Once()

		vdo, err := videoService.VideosUpdate(context.Background(), videoDomain, 1)

		assert.Error(t, err)
		assert.Equal(t, vdo, videos.Domain{})
	})
}
