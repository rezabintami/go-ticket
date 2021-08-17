package topup_test

import (
	"context"
	"errors"
	"os"
	"testing"
	topup "ticketing/business/topup"
	topupMock "ticketing/business/topup/mocks"
	"ticketing/business/users"
	usersMock "ticketing/business/users/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	topupRepository topupMock.Repository
	topupUsecase   topup.Usecase
	usersRepository usersMock.Repository
)

func setup() {
	topupUsecase = topup.NewTopUpUsecase(&topupRepository, 2, &usersRepository)
}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestGetByID(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		domain := []topup.Domain{
			{
				ID:      1,
				Name:    "OVO",
				UserID:  2,
				Balance: 20000,
			},
			{
				ID:      2,
				Name:    "GOPAY",
				UserID:  2,
				Balance: 25000,
			},
		}
		topupRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(domain, nil).Once()

		result, err := topupUsecase.GetByID(context.Background(), 1)

		assert.Nil(t, err)
		assert.Equal(t, 2, len(result))
	})

	t.Run("test case 2, repository error", func(t *testing.T) {
		errRepository := errors.New("data not found")
		topupRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return([]topup.Domain{}, errRepository).Once()

		result, err := topupUsecase.GetByID(context.Background(), -1)

		assert.Equal(t, 0, len(result))
		assert.Equal(t, errRepository, err)
	})
}

func TestStore(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		theaterdomain := topup.Domain{
			ID:      1,
			Name:    "OVO",
			UserID:  2,
			Balance: 20000,
		}
		userDomain := users.Domain{
			ID:    1,
			Name:  "reza bintami",
			Email: "rezabintami@gmail.com",
		}
		topupRepository.On("Store", mock.Anything, mock.Anything).Return(nil).Once()
		usersRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(userDomain, nil).Once()
		usersRepository.On("UpdateUser", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(nil).Once()

		err := topupUsecase.Store(context.Background(), &theaterdomain)

		assert.Nil(t, err)
	})

	t.Run("test case 2, data is null", func(t *testing.T) {
		errRepository := errors.New("data is null")
		topupRepository.On("Store", mock.Anything, mock.Anything).Return(errRepository).Once()

		err := topupUsecase.Store(context.Background(), &topup.Domain{})

		assert.Equal(t, errRepository, err)
	})
	t.Run("test case 3, id not found", func(t *testing.T) {
		errRepository := errors.New("id not found")
		topupRepository.On("Store", mock.Anything, mock.Anything).Return(nil).Once()
		usersRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(users.Domain{}, errRepository).Once()

		err := topupUsecase.Store(context.Background(), &topup.Domain{})

		assert.Equal(t, errRepository, err)
	})
	t.Run("test case 4, repository error", func(t *testing.T) {
		userDomain := users.Domain{
			ID:    1,
			Name:  "reza bintami",
			Email: "rezabintami@gmail.com",
		}
		errRepository := errors.New("mysql not running")
		topupRepository.On("Store", mock.Anything, mock.Anything).Return(nil).Once()
		usersRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(userDomain, nil).Once()
		usersRepository.On("UpdateUser", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(errRepository).Once()

		err := topupUsecase.Store(context.Background(), &topup.Domain{})

		assert.Equal(t, errRepository, err)
	})
}
