package movies_test

// import (
// 	"context"
// 	"errors"
// 	"os"
// 	"testing"
// 	movies "ticketing/business/movies"
// 	moviesMock "ticketing/business/movies/mocks"
// 	"ticketing/business/users"
// 	usersMock "ticketing/business/users/mocks"

// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// var (
// 	moviesRepository moviesMock.Repository
// 	moviesUsecase   movies.Usecase
// 	usersRepository usersMock.Repository
// )

// func setup() {
// 	moviesUsecase = movies.NewMoviesUsecase(&moviesRepository, 2, &usersRepository)
// }

// func TestMain(m *testing.M) {
// 	setup()
// 	os.Exit(m.Run())
// }

// func TestGetByID(t *testing.T) {
// 	t.Run("test case 1, valid test", func(t *testing.T) {
// 		domain := []movies.Domain{
// 			{
// 				ID:      1,
// 				Name:    "OVO",
// 				UserID:  2,
// 				Balance: 20000,
// 			},
// 			{
// 				ID:      2,
// 				Name:    "GOPAY",
// 				UserID:  2,
// 				Balance: 25000,
// 			},
// 		}
// 		moviesRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(domain, nil).Once()

// 		result, err := moviesUsecase.GetByID(context.Background(), 1)

// 		assert.Nil(t, err)
// 		assert.Equal(t, 2, len(result))
// 	})

// 	t.Run("test case 2, repository error", func(t *testing.T) {
// 		errRepository := errors.New("data not found")
// 		moviesRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return([]movies.Domain{}, errRepository).Once()

// 		result, err := moviesUsecase.GetByID(context.Background(), -1)

// 		assert.Equal(t, 0, len(result))
// 		assert.Equal(t, errRepository, err)
// 	})
// }

// func TestStore(t *testing.T) {
// 	t.Run("test case 1, valid test", func(t *testing.T) {
// 		theaterdomain := movies.Domain{
// 			ID:      1,
// 			Name:    "OVO",
// 			UserID:  2,
// 			Balance: 20000,
// 		}
// 		userDomain := users.UserDomain{
// 			ID:    1,
// 			Name:  "reza bintami",
// 			Email: "rezabintami@gmail.com",
// 		}
// 		moviesRepository.On("Store", mock.Anything, mock.Anything).Return(nil).Once()
// 		usersRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(userDomain, nil).Once()
// 		usersRepository.On("UpdateUser", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(nil).Once()

// 		err := moviesUsecase.Store(context.Background(), &theaterdomain)

// 		assert.Nil(t, err)
// 	})

// 	t.Run("test case 2, data is null", func(t *testing.T) {
// 		errRepository := errors.New("data is null")
// 		moviesRepository.On("Store", mock.Anything, mock.Anything).Return(errRepository).Once()

// 		err := moviesUsecase.Store(context.Background(), &movies.Domain{})

// 		assert.Equal(t, errRepository, err)
// 	})
// 	t.Run("test case 3, id not found", func(t *testing.T) {
// 		errRepository := errors.New("id not found")
// 		moviesRepository.On("Store", mock.Anything, mock.Anything).Return(nil).Once()
// 		usersRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(users.UserDomain{}, errRepository).Once()

// 		err := moviesUsecase.Store(context.Background(), &movies.Domain{})

// 		assert.Equal(t, errRepository, err)
// 	})
// 	t.Run("test case 4, repository error", func(t *testing.T) {
// 		userDomain := users.UserDomain{
// 			ID:    1,
// 			Name:  "reza bintami",
// 			Email: "rezabintami@gmail.com",
// 		}
// 		errRepository := errors.New("mysql not running")
// 		moviesRepository.On("Store", mock.Anything, mock.Anything).Return(nil).Once()
// 		usersRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(userDomain, nil).Once()
// 		usersRepository.On("UpdateUser", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(errRepository).Once()

// 		err := moviesUsecase.Store(context.Background(), &movies.Domain{})

// 		assert.Equal(t, errRepository, err)
// 	})
// }
