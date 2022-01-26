package modules_test

import (
	"backend/business/course"
	"backend/business/modules"
	_mockModuleRepository "backend/business/modules/mocks"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var moduleRepository _mockModuleRepository.ModulesRepoInterface
var moduleService modules.ModulesUseCaseInterface
var moduleDomain modules.Domain
var modulesDomain []modules.Domain
var courseDomain course.Domain

func setup() {
	moduleService = modules.NewUseCase(&moduleRepository, time.Hour*1 /*, configJWT */)
	moduleDomain = modules.Domain{
		Id:       1,
		Title:    "Frontend Development",
		CourseId: 1,
		Order:    1,
		Course:   courseDomain,
	}
	modulesDomain = append(modulesDomain, moduleDomain)
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

func TestGetById(t *testing.T) {
	setup()

	t.Run("Test case 1", func(t *testing.T) {
		moduleRepository.On("ModulesGetById",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(moduleDomain, nil).Once()

		module, err := moduleService.ModulesGetById(context.Background(), 1)

		assert.NoError(t, err)
		assert.NotNil(t, module)

	})

	t.Run("Test case 2 | Error", func(t *testing.T) {
		moduleRepository.On("ModulesGetById",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(modules.Domain{}, errors.New("Error at Get All Categories")).Once()

		module, err := moduleService.ModulesGetById(context.Background(), 1)

		assert.Error(t, err)
		assert.Equal(t, module, modules.Domain{})
	})
}

func TestDelete(t *testing.T) {
	setup()

	t.Run("Test case 1", func(t *testing.T) {
		moduleRepository.On("ModulesDelete",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(nil).Once()

		err := moduleService.ModulesDelete(context.Background(), 1)

		assert.NoError(t, err)
	})

	t.Run("Test case 2 | Error", func(t *testing.T) {
		moduleRepository.On("ModulesDelete",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(errors.New("Error at Delete Modules")).Once()

		err := moduleService.ModulesDelete(context.Background(), 1)

		assert.Error(t, err)
	})
}

func TestGetByCourseId(t *testing.T) {
	setup()

	t.Run("Test case 1", func(t *testing.T) {
		moduleRepository.On("ModulesGetByCourseId",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(modulesDomain, nil).Once()

		vds, err := moduleService.ModulesGetByCourseId(context.Background(), 1)

		assert.NoError(t, err)
		assert.NotNil(t, vds)
	})

	t.Run("Test case 2 | Error", func(t *testing.T) {
		moduleRepository.On("ModulesGetByCourseId",
			mock.Anything,
			mock.AnythingOfType("uint")).Return([]modules.Domain{}, errors.New("Error at Get Modules By Module Id")).Once()

		vds, err := moduleService.ModulesGetByCourseId(context.Background(), 1)

		assert.Error(t, err)
		assert.Equal(t, vds, []modules.Domain{})
	})
}

func TestGetAll(t *testing.T) {
	setup()

	t.Run("Test case 1", func(t *testing.T) {
		moduleRepository.On("ModulesGetAll",
			mock.Anything).Return(modulesDomain, nil).Once()

		module, err := moduleService.ModulesGetAll(context.Background())

		assert.NoError(t, err)
		assert.NotNil(t, module)

	})

	t.Run("Test case 2 | Error", func(t *testing.T) {
		moduleRepository.On("ModulesGetAll",
			mock.Anything).Return([]modules.Domain{}, errors.New("Error at Get All Categories")).Once()

		module, err := moduleService.ModulesGetAll(context.Background())

		assert.Error(t, err)
		assert.Equal(t, module, []modules.Domain{})
	})
}

func TestAdd(t *testing.T) {
	setup()

	t.Run("Test case 1 | Valid ModulesAdd", func(t *testing.T) {
		moduleRepository.On("ModulesAdd",
			mock.Anything,
			mock.AnythingOfType("modules.Domain")).Return(moduleDomain, nil).Once()

		moduleRepository.On("CheckCourse",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(courseDomain, nil).Once()

		mdl, err := moduleService.ModulesAdd(context.Background(), moduleDomain)

		assert.Nil(t, err)
		assert.Equal(t, uint(1), mdl.CourseId)
		assert.Equal(t, 1, mdl.Order)
	})

	t.Run("Test case 2 | Error ModulesAdd", func(t *testing.T) {
		moduleRepository.On("ModulesAdd",
			mock.Anything,
			mock.AnythingOfType("modules.Domain")).Return(modules.Domain{}, errors.New("Unexpected Error")).Once()

		moduleRepository.On("CheckCourse",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(courseDomain, nil).Once()

		mdl, err := moduleService.ModulesAdd(context.Background(), moduleDomain)

		assert.Error(t, err)
		assert.Equal(t, mdl, modules.Domain{})
	})

	t.Run("Test case 3 | Invalid CourseId Empty", func(t *testing.T) {
		_, err := moduleService.ModulesAdd(context.Background(), modules.Domain{
			Title:    "Frontend Development",
			CourseId: 0,
			Order:    1,
		})

		moduleRepository.On("CheckCourse",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(moduleDomain, nil).Once()

		assert.NotNil(t, err)
	})

	t.Run("Test case 4 | Invalid Title Empty", func(t *testing.T) {
		_, err := moduleService.ModulesAdd(context.Background(), modules.Domain{
			Title:    "",
			CourseId: 1,
			Order:    1,
		})

		moduleRepository.On("CheckCourse",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(moduleDomain, nil).Once()

		assert.NotNil(t, err)
	})

	t.Run("Test case 5 | Invalid Order Empty", func(t *testing.T) {
		_, err := moduleService.ModulesAdd(context.Background(), modules.Domain{
			Title:    "Frontend Development",
			CourseId: 1,
			Order:    0,
		})

		moduleRepository.On("CheckCourse",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(moduleDomain, nil).Once()

		assert.NotNil(t, err)
	})
}

func TestUpdate(t *testing.T) {
	setup()

	// t.Run("Test case 1 | Valid ModulesUpdate", func(t *testing.T) {
	// 	// moduleRepository.On("ModulesUpdate",
	// 	// 	mock.Anything,
	// 	// 	mock.AnythingOfType("modules.Domain"),
	// 	// 	mock.AnythingOfType("uint")).Return(moduleDomain, nil).Once()

	// 	moduleRepository.On("CheckCourse",
	// 		mock.Anything,
	// 		mock.AnythingOfType("uint")).Return(courseDomain, errors.New("Unexpected Error")).Once()

	// 	mdl, err := moduleService.ModulesUpdate(context.Background(), moduleDomain, uint(1))

	// 	assert.Nil(t, err)
	// 	assert.Equal(t, uint(1), mdl.CourseId)
	// 	assert.Equal(t, 1, mdl.Order)
	// 	assert.Equal(t, mdl.Course, courseDomain)
	// })

	// t.Run("Test case 2 | Error ModulesUpdate", func(t *testing.T) {

	// 	moduleRepository.On("CheckCourse",
	// 		mock.Anything,
	// 		mock.AnythingOfType("uint")).Return(courseDomain, nil).Once()

	// 	moduleRepository.On("ModulesUpdate",
	// 		mock.Anything,
	// 		mock.AnythingOfType("modules.Domain"),
	// 		mock.AnythingOfType("uint")).Return(modules.Domain{}, errors.New("Unexpected Error")).Once()

	// 	mdl, err := moduleService.ModulesUpdate(context.Background(), moduleDomain, uint(1))

	// 	assert.NotNil(t, err)
	// 	assert.Equal(t, mdl, modules.Domain{})
	// 	assert.IsType(t, course.Domain{}, courseDomain)
	// })

	t.Run("Test case 3 | Invalid CourseId Empty", func(t *testing.T) {
		_, err := moduleService.ModulesUpdate(context.Background(), modules.Domain{
			Title:    "Frontend Development",
			CourseId: 0,
			Order:    1,
		}, uint(1))

		moduleRepository.On("CheckCourse",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(courseDomain, nil).Once()

		assert.NotNil(t, err)
	})

	t.Run("Test case 4 | Invalid Title Empty", func(t *testing.T) {
		_, err := moduleService.ModulesUpdate(context.Background(), modules.Domain{
			Title:    "",
			CourseId: 1,
			Order:    1,
		}, uint(1))

		moduleRepository.On("CheckCourse",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(courseDomain, nil).Once()

		assert.NotNil(t, err)
	})

	t.Run("Test case 5 | Invalid Order Empty", func(t *testing.T) {
		_, err := moduleService.ModulesUpdate(context.Background(), modules.Domain{
			Title:    "Frontend Development",
			CourseId: 1,
			Order:    0,
		}, uint(1))

		moduleRepository.On("CheckCourse",
			mock.Anything,
			mock.AnythingOfType("uint")).Return(courseDomain, nil).Once()

		assert.NotNil(t, err)
	})
}
