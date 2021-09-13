package users_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"ticketing/app/middleware"
	"ticketing/business"
	"ticketing/business/users"
	usersMock "ticketing/business/users/mocks"
	"ticketing/helper/encrypt"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	usersRepository usersMock.Repository
	usersUsecase    users.Usecase
	jwtAuth         *middleware.ConfigJWT
)

func setup() {
	jwtAuth = &middleware.ConfigJWT{SecretJWT: "abc123", ExpiresDuration: 2}
	usersUsecase = users.NewUserUsecase(&usersRepository, jwtAuth, 2)
}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestGetByID(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		domain := users.Domain{
			ID:    1,
			Name:  "reza bintami",
			Email: "rezabintami@gmail.com",
		}
		usersRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(domain, nil).Once()

		result, err := usersUsecase.GetByID(context.Background(), 1)

		assert.Nil(t, err)
		assert.Equal(t, domain.ID, result.ID)
	})

	t.Run("test case 2, data not found", func(t *testing.T) {
		errRepository := errors.New("data not found")
		usersRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(users.Domain{}, errRepository).Once()
		result, err := usersUsecase.GetByID(context.Background(), -1)
		assert.Equal(t, result, users.Domain{})
		assert.Equal(t, err, errRepository)
	})
}

func TestRegister(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		domain := users.Domain{
			ID:       1,
			Password: "asyudasd820aisd",
			Name:     "reza bintami",
			Email:    "rezabintami@gmail.com",
			Balance:  0,
			Language: "en",
		}
		usersRepository.On("GetByEmail", mock.Anything, mock.Anything).Return(users.Domain{}, nil).Once()
		usersRepository.On("Register", mock.Anything, mock.Anything).Return(nil).Once()

		err := usersUsecase.Register(context.Background(), &domain, false)

		assert.Nil(t, err)
	})

	t.Run("test case 2, duplicate data", func(t *testing.T) {
		domain := users.Domain{
			ID:       1,
			Password: "asyudasd820aisd",
			Name:     "reza bintami",
			Email:    "rezabintami@gmail.com",
			Balance:  0,
			Language: "en",
		}
		errRepository := errors.New("duplicate data")
		usersRepository.On("GetByEmail", mock.Anything, mock.Anything).Return(domain, errRepository).Once()

		err := usersUsecase.Register(context.Background(), &domain, false)

		assert.Equal(t, err, business.ErrDuplicateData)
	})

	t.Run("test case 3, data has exist", func(t *testing.T) {
		domain := users.Domain{
			ID:       1,
			Password: "asyudasd820aisd",
			Name:     "reza bintami",
			Email:    "rezabintami@gmail.com",
			Balance:  0,
			Language: "en",
		}
		usersRepository.On("GetByEmail", mock.Anything, mock.Anything).Return(domain, nil).Once()

		err := usersUsecase.Register(context.Background(), &domain, false)

		assert.Equal(t, err, business.ErrDuplicateData)
	})

	// t.Run("test case 4, hashing password error", func(t *testing.T) {
	// 	domain := users.Domain{
	// 		ID:       1,
	// 		Password: "",
	// 		Name:     "reza bintami",
	// 		Email:    "rezabintami@gmail.com",
	// 		Balance:  0,
	// 		Language: "en",
	// 	}
	// 	usersRepository.On("GetByEmail", mock.Anything, mock.Anything).Return(users.Domain{}, nil).Once()

	// 	err := usersUsecase.Register(context.Background(), &domain)

	// 	assert.Equal(t, err, business.ErrInternalServer)
	// })

	t.Run("test case 4, register failed", func(t *testing.T) {
		domain := users.Domain{
			ID:       1,
			Name:     "reza bintami",
			Email:    "rezabintami@gmail.com",
			Balance:  0,
			Language: "en",
		}
		errRepository := errors.New("register failed")
		usersRepository.On("GetByEmail", mock.Anything, mock.Anything).Return(users.Domain{}, nil).Once()
		usersRepository.On("Register", mock.Anything, mock.Anything).Return(errRepository).Once()

		err := usersUsecase.Register(context.Background(), &domain, false)

		assert.Equal(t, err, errRepository)
	})
}

func TestUpdateUser(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		domain := users.Domain{
			ID:       1,
			Password: "asyudasd820aisd",
			Name:     "reza bintami",
			Email:    "rezabintami@gmail.com",
			Balance:  0,
			Language: "en",
		}
		usersRepository.On("UpdateUser", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(nil).Once()

		err := usersUsecase.UpdateUser(context.Background(), &domain, 1)

		assert.Nil(t, err)
	})

	t.Run("test case 2, id not found", func(t *testing.T) {
		domain := users.Domain{
			ID:       1,
			Password: "asyudasd820aisd",
			Name:     "reza bintami",
			Email:    "rezabintami@gmail.com",
			Balance:  0,
			Language: "en",
		}
		errRepository := errors.New("id not found")
		usersRepository.On("UpdateUser", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(errRepository).Once()

		err := usersUsecase.UpdateUser(context.Background(), &domain, -1)

		assert.Equal(t, err, errRepository)
	})
}

func TestLogin(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		pass, _ := encrypt.Hash("123123")
		usersDomain := users.Domain{
			ID:       1,
			Password: pass,
			Name:     "zaza",
			Email:    "zaza@gmail.com",
			Balance:  0,
			Language: "en",
		}

		usersRepository.On("GetByEmail", mock.Anything, mock.AnythingOfType("string")).Return(usersDomain, nil).Once()

		_, err := usersUsecase.Login(context.Background(), "zaza@gmail.com", "123123", false)
		assert.Nil(t, err)
	})
	t.Run("test case 2, password error", func(t *testing.T) {
		pass, _ := encrypt.Hash("1231231")
		usersDomain := users.Domain{
			ID:       1,
			Password: pass,
			Name:     "zaza",
			Email:    "zaza@gmail.com",
			Balance:  0,
			Language: "en",
		}

		usersRepository.On("GetByEmail", mock.Anything, mock.AnythingOfType("string")).Return(usersDomain, nil).Once()

		_, err := usersUsecase.Login(context.Background(), "zaza@gmail.com", "123123", false)
		assert.Equal(t, err, business.ErrEmailPasswordNotFound)

	})

	t.Run("test case 3, error record", func(t *testing.T) {

		errRepository := errors.New("error record")
		usersRepository.On("GetByEmail", mock.Anything, mock.AnythingOfType("string")).Return(users.Domain{}, errRepository).Once()

		result, err := usersUsecase.Login(context.Background(), "rezabintami@gmail.com", "123123", false)

		assert.Equal(t, err, errRepository)
		assert.Equal(t, "", result)
	})
}
