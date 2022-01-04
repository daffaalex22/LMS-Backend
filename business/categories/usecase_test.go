package categories_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"backend/business/categories"
	_mockCategoryRepository "backend/business/categories/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository _mockCategoryRepository.Repository
var categorieService categories.Usecase
var categoryDomain categories.Domain
var categoriesDomain []categories.Domain

func setup() {
	categorieService = categories.NewCategoryUsecase(time.Hour*1, &categoryRepository /*, configJWT */)
	categoryDomain = categories.Domain{
		Id:    1,
		Title: "Science and Engineering",
	}
	categoriesDomain = append(categoriesDomain, categoryDomain)
}

func TestGetAll(t *testing.T) {
	setup()

	t.Run("Test case 1", func(t *testing.T) {
		categoryRepository.On("GetAll",
			mock.Anything).Return(categoriesDomain, nil).Once()

		category, err := categorieService.GetAll(context.Background())

		assert.NoError(t, err)
		assert.NotNil(t, category)

	})

	t.Run("Test case 2 | Error", func(t *testing.T) {
		categoryRepository.On("GetAll",
			mock.Anything).Return([]categories.Domain{}, errors.New("Error at Get All Categories")).Once()

		category, err := categorieService.GetAll(context.Background())

		assert.Error(t, err)
		assert.Equal(t, category, []categories.Domain{})
	})
}
