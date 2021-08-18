package theater_test

import (
	"context"
	"errors"
	"os"
	"testing"
	theater "ticketing/business/theater"
	theaterMock "ticketing/business/theater/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	theaterRepository theaterMock.Repository
	theaterUsecase   theater.Usecase
)

func setup() {
	theaterUsecase = theater.NewTheaterUsecase(&theaterRepository, 2)
}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestGetAll(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		domain := []theater.Domain{
			{
				ID:    1,
				Name:  "Sutos Surabaya",
				Place: "Surabaya",
			},
			{
				ID:    2,
				Name:  "TP Surabaya",
				Place: "Surabaya",
			},
		}
		theaterRepository.On("GetAll", mock.Anything).Return(domain, nil).Once()

		result, err := theaterUsecase.GetAll(context.Background())

		assert.Nil(t, err)
		assert.Equal(t, 2, len(result))
	})

	t.Run("test case 2, repository error", func(t *testing.T) {
		errRepository := errors.New("data not found")
		theaterRepository.On("GetAll", mock.Anything).Return([]theater.Domain{}, errRepository).Once()

		result, err := theaterUsecase.GetAll(context.Background())

		assert.Equal(t, 0, len(result))
		assert.Equal(t, errRepository, err)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		domain := theater.Domain{
			ID:    1,
			Name:  "Royal Surabaya",
			Place: "Surabaya",
		}
		theaterRepository.On("Update", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(nil).Once()

		err := theaterUsecase.Update(context.Background(), &domain, 1)

		assert.Nil(t, err)
	})

	t.Run("test case 2, id not found", func(t *testing.T) {
		errRepository := errors.New("id not found")
		theaterRepository.On("Update", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(errRepository).Once()

		err := theaterUsecase.Update(context.Background(), &theater.Domain{}, -1)

		assert.Equal(t, errRepository, err)
	})
}


func TestDelete(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		theaterRepository.On("Delete", mock.Anything,  mock.AnythingOfType("int")).Return(nil).Once()

		err := theaterUsecase.Delete(context.Background(), 1)

		assert.Nil(t, err)
	})

	t.Run("test case 2, id not found", func(t *testing.T) {
		errRepository := errors.New("id not found")
		theaterRepository.On("Delete", mock.Anything, mock.AnythingOfType("int")).Return(errRepository).Once()

		err := theaterUsecase.Delete(context.Background(), -1)

		assert.Equal(t, errRepository, err)
	})
}

func TestStore(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		domain := theater.Domain{
			ID:    1,
			Name:  "Royal Surabaya",
			Place: "Surabaya",
		}
		theaterRepository.On("Store", mock.Anything, mock.Anything).Return(nil).Once()

		err := theaterUsecase.Store(context.Background(), &domain)

		assert.Nil(t, err)
	})

	t.Run("test case 2, id not found", func(t *testing.T) {
		errRepository := errors.New("data is null")
		theaterRepository.On("Store", mock.Anything, mock.Anything).Return(errRepository).Once()

		err := theaterUsecase.Store(context.Background(), &theater.Domain{})

		assert.Equal(t, errRepository, err)
	})
}