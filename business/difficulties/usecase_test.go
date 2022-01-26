package difficulties_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"backend/business/difficulties"
	_mockDifficultyRepository "backend/business/difficulties/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var difficultyRepository _mockDifficultyRepository.Repository
var difficultyService difficulties.Usecase
var difficultyDomain difficulties.Domain
var difficultiesDomain []difficulties.Domain

func setup() {
	difficultyService = difficulties.NewDifficultyUsecase(time.Hour*1, &difficultyRepository /*, configJWT */)
	difficultyDomain = difficulties.Domain{
		Id:    1,
		Title: "Science and Engineering",
	}
	difficultiesDomain = append(difficultiesDomain, difficultyDomain)
}

func TestGetAll(t *testing.T) {
	setup()

	t.Run("Test case 1", func(t *testing.T) {
		difficultyRepository.On("GetAll",
			mock.Anything).Return(difficultiesDomain, nil).Once()

		difficulty, err := difficultyService.GetAll(context.Background())

		assert.NoError(t, err)
		assert.NotNil(t, difficulty)

	})

	t.Run("Test case 2 | Error", func(t *testing.T) {
		difficultyRepository.On("GetAll",
			mock.Anything).Return([]difficulties.Domain{}, errors.New("Error at Get All Categories")).Once()

		difficulty, err := difficultyService.GetAll(context.Background())

		assert.Error(t, err)
		assert.Equal(t, difficulty, []difficulties.Domain{})
	})
}
