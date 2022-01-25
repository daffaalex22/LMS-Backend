package types_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"backend/business/types"
	_mockTypeRepository "backend/business/types/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var typeRepository _mockTypeRepository.Repository
var typeService types.Usecase
var typeDomain types.Domain
var typesDomain []types.Domain

func setup() {
	typeService = types.NewTypeUsecase(time.Hour*1, &typeRepository /*, configJWT */)
	typeDomain = types.Domain{
		Id:    1,
		Title: "Counselling",
	}
	typesDomain = append(typesDomain, typeDomain)
}

func TestGetAll(t *testing.T) {
	setup()

	t.Run("Test case 1", func(t *testing.T) {
		typeRepository.On("GetAll",
			mock.Anything).Return(typesDomain, nil).Once()

		typ, err := typeService.GetAll(context.Background())

		assert.NoError(t, err)
		assert.NotNil(t, typ)

	})

	t.Run("Test case 2 | Error", func(t *testing.T) {
		typeRepository.On("GetAll",
			mock.Anything).Return([]types.Domain{}, errors.New("Error at Get All Types")).Once()

		typ, err := typeService.GetAll(context.Background())

		assert.Error(t, err)
		assert.Equal(t, typ, []types.Domain{})
	})
}
