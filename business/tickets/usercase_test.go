package tickets_test

import (
	"context"
	"errors"
	"os"
	"testing"
	tickets "ticketing/business/tickets"
	ticketsMock "ticketing/business/tickets/mocks"
	"ticketing/business/users"
	usersMock "ticketing/business/users/mocks"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	ticketsRepository ticketsMock.Repository
	ticketsUsecase    tickets.Usecase
	usersRepository   usersMock.Repository
)

func setup() {
	ticketsUsecase = tickets.NewTicketsUsecase(&ticketsRepository, &usersRepository, 2)
}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestGetByID(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		domain := []tickets.Domain{
			{
				ID:          9,
				BookingCode: "3cc278cf-a381-45ad-9377-afeebfdb8536",
				MovieID:     3,
				UserID:      2,
				TheaterID:   2,
				Seats:       "H9",
				TotalPrice:  120000,
				Time:        time.Now(),
			},
			{
				ID:          10,
				BookingCode: "3cc278cf-a381-45ad-9377-afeebfdb1234",
				MovieID:     3,
				UserID:      2,
				TheaterID:   2,
				Seats:       "H9",
				TotalPrice:  120000,
				Time:        time.Now(),
			},
		}
		ticketsRepository.On("GetAllByID", mock.Anything, mock.AnythingOfType("int")).Return(domain, nil).Once()

		result, err := ticketsUsecase.GetByID(context.Background(), 1)

		assert.Nil(t, err)
		assert.Equal(t, 2, len(result))
	})

	t.Run("test case 2, repository error", func(t *testing.T) {
		errRepository := errors.New("data not found")
		ticketsRepository.On("GetAllByID", mock.Anything, mock.AnythingOfType("int")).Return([]tickets.Domain{}, errRepository).Once()

		result, err := ticketsUsecase.GetByID(context.Background(), -1)

		assert.Equal(t, 0, len(result))
		assert.Equal(t, errRepository, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		ticketDomain := tickets.Domain{
			ID:          9,
			BookingCode: "3cc278cf-a381-45ad-9377-afeebfdb8536",
			MovieID:     3,
			UserID:      2,
			TheaterID:   2,
			Seats:       "H9",
			TotalPrice:  120000,
			Time:        time.Now(),
		}
		userDomain := users.Domain{
			ID:    1,
			Name:  "reza bintami",
			Email: "rezabintami@gmail.com",
		}
		ticketsRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(ticketDomain, nil).Once()
		ticketsRepository.On("Delete", mock.Anything, mock.AnythingOfType("int")).Return(nil).Once()
		usersRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(userDomain, nil).Once()
		usersRepository.On("UpdateUser", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(nil).Once()

		err := ticketsUsecase.Delete(context.Background(), 1, 1)

		assert.Nil(t, err)
	})

	t.Run("test case 2, id ticket not found", func(t *testing.T) {
		errRepository := errors.New("id not found")

		ticketsRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(tickets.Domain{}, errRepository).Once()

		err := ticketsUsecase.Delete(context.Background(), -1, 1)

		assert.Equal(t, errRepository, err)
	})

	t.Run("test case 3, delete ticket failed", func(t *testing.T) {
		errRepository := errors.New("delete ticket failed")
		ticketDomain := tickets.Domain{
			ID:          9,
			BookingCode: "3cc278cf-a381-45ad-9377-afeebfdb8536",
			MovieID:     3,
			UserID:      2,
			TheaterID:   2,
			Seats:       "H9",
			TotalPrice:  120000,
			Time:        time.Now(),
		}
		ticketsRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(ticketDomain, nil).Once()
		ticketsRepository.On("Delete", mock.Anything, mock.AnythingOfType("int")).Return(errRepository).Once()

		err := ticketsUsecase.Delete(context.Background(), -1, 1)

		assert.Equal(t, errRepository, err)
	})

	t.Run("test case 4, id user not found", func(t *testing.T) {
		errRepository := errors.New("id not found")
		ticketDomain := tickets.Domain{
			ID:          9,
			BookingCode: "3cc278cf-a381-45ad-9377-afeebfdb8536",
			MovieID:     3,
			UserID:      2,
			TheaterID:   2,
			Seats:       "H9",
			TotalPrice:  120000,
			Time:        time.Now(),
		}
		ticketsRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(ticketDomain, nil).Once()
		ticketsRepository.On("Delete", mock.Anything, mock.AnythingOfType("int")).Return(nil).Once()
		usersRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(users.Domain{}, errRepository).Once()

		err := ticketsUsecase.Delete(context.Background(), -1, 1)

		assert.Equal(t, errRepository, err)
	})

	t.Run("test case 5, update user failed", func(t *testing.T) {
		errRepository := errors.New("update user failed")
		ticketDomain := tickets.Domain{
			ID:          9,
			BookingCode: "3cc278cf-a381-45ad-9377-afeebfdb8536",
			MovieID:     3,
			UserID:      2,
			TheaterID:   2,
			Seats:       "H9",
			TotalPrice:  120000,
			Time:        time.Now(),
		}
		userDomain := users.Domain{
			ID:    1,
			Name:  "reza bintami",
			Email: "rezabintami@gmail.com",
		}
		ticketsRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(ticketDomain, nil).Once()
		ticketsRepository.On("Delete", mock.Anything, mock.AnythingOfType("int")).Return(nil).Once()
		usersRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(userDomain, nil).Once()
		usersRepository.On("UpdateUser", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(errRepository).Once()

		err := ticketsUsecase.Delete(context.Background(), 1, 1)

		assert.Equal(t, errRepository, err)
	})
}

func TestStore(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		ticketDomain := tickets.Domain{
			ID:          9,
			BookingCode: "3cc278cf-a381-45ad-9377-afeebfdb8536",
			MovieID:     3,
			UserID:      2,
			TheaterID:   2,
			Seats:       "H9",
			TotalPrice:  120000,
			Time:        time.Now(),
		}
		userDomain := users.Domain{
			ID:    1,
			Name:  "reza bintami",
			Email: "rezabintami@gmail.com",
		}
		ticketsRepository.On("Store", mock.Anything, mock.Anything).Return(nil).Once()
		usersRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(userDomain, nil).Once()
		usersRepository.On("UpdateUser", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(nil).Once()

		err := ticketsUsecase.Store(context.Background(), &ticketDomain, 1)

		assert.Nil(t, err)
	})
	t.Run("test case 2, store ticket error", func(t *testing.T) {

		errRepository := errors.New(" store ticket error")
		ticketsRepository.On("Store", mock.Anything, mock.Anything).Return(errRepository).Once()

		err := ticketsUsecase.Store(context.Background(), &tickets.Domain{}, 1)

		assert.Equal(t, errRepository, err)
	})

	t.Run("test case 3, id user not found", func(t *testing.T) {
		ticketDomain := tickets.Domain{
			ID:          9,
			BookingCode: "3cc278cf-a381-45ad-9377-afeebfdb8536",
			MovieID:     3,
			UserID:      2,
			TheaterID:   2,
			Seats:       "H9",
			TotalPrice:  120000,
			Time:        time.Now(),
		}

		errRepository := errors.New("id user not found")
		ticketsRepository.On("Store", mock.Anything, mock.Anything).Return(nil).Once()
		usersRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(users.Domain{}, errRepository).Once()

		err := ticketsUsecase.Store(context.Background(), &ticketDomain, 1)

		assert.Equal(t, errRepository, err)
	})

	t.Run("test case 4, update user failed", func(t *testing.T) {
		ticketDomain := tickets.Domain{
			ID:          9,
			BookingCode: "3cc278cf-a381-45ad-9377-afeebfdb8536",
			MovieID:     3,
			UserID:      2,
			TheaterID:   2,
			Seats:       "H9",
			TotalPrice:  120000,
			Time:        time.Now(),
		}
		userDomain := users.Domain{
			ID:    1,
			Name:  "reza bintami",
			Email: "rezabintami@gmail.com",
		}
		errRepository := errors.New("update user failed")
		ticketsRepository.On("Store", mock.Anything, mock.Anything).Return(nil).Once()
		usersRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(userDomain, nil).Once()
		usersRepository.On("UpdateUser", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(errRepository).Once()

		err := ticketsUsecase.Store(context.Background(), &ticketDomain, 1)

		assert.Equal(t, errRepository, err)
	})

	// t.Run("test case 3, delete ticket failed", func(t *testing.T) {
	// 	errRepository := errors.New("delete ticket failed")
	// 	ticketDomain := tickets.Domain{
	// 		ID:          9,
	// 		BookingCode: "3cc278cf-a381-45ad-9377-afeebfdb8536",
	// 		MovieID:     3,
	// 		UserID:      2,
	// 		TheaterID:   2,
	// 		Seats:       "H9",
	// 		TotalPrice:  120000,
	// 		Time:        time.Now(),
	// 	}
	// 	ticketsRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(ticketDomain, nil).Once()
	// 	ticketsRepository.On("Delete", mock.Anything, mock.AnythingOfType("int")).Return(errRepository).Once()

	// 	err := ticketsUsecase.Delete(context.Background(), -1, 1)

	// 	assert.Equal(t, errRepository, err)
	// })

	// t.Run("test case 4, id user not found", func(t *testing.T) {
	// 	errRepository := errors.New("id not found")
	// 	ticketDomain := tickets.Domain{
	// 		ID:          9,
	// 		BookingCode: "3cc278cf-a381-45ad-9377-afeebfdb8536",
	// 		MovieID:     3,
	// 		UserID:      2,
	// 		TheaterID:   2,
	// 		Seats:       "H9",
	// 		TotalPrice:  120000,
	// 		Time:        time.Now(),
	// 	}
	// 	ticketsRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(ticketDomain, nil).Once()
	// 	ticketsRepository.On("Delete", mock.Anything, mock.AnythingOfType("int")).Return(nil).Once()
	// 	usersRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(users.UserDomain{}, errRepository).Once()

	// 	err := ticketsUsecase.Delete(context.Background(), -1, 1)

	// 	assert.Equal(t, errRepository, err)
	// })

	// t.Run("test case 5, update user failed", func(t *testing.T) {
	// 	errRepository := errors.New("update user failed")
	// 	ticketDomain := tickets.Domain{
	// 		ID:          9,
	// 		BookingCode: "3cc278cf-a381-45ad-9377-afeebfdb8536",
	// 		MovieID:     3,
	// 		UserID:      2,
	// 		TheaterID:   2,
	// 		Seats:       "H9",
	// 		TotalPrice:  120000,
	// 		Time:        time.Now(),
	// 	}
	// 	userDomain := users.UserDomain{
	// 		ID:    1,
	// 		Name:  "reza bintami",
	// 		Email: "rezabintami@gmail.com",
	// 	}
	// 	ticketsRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(ticketDomain, nil).Once()
	// 	ticketsRepository.On("Delete", mock.Anything, mock.AnythingOfType("int")).Return(nil).Once()
	// 	usersRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(userDomain, nil).Once()
	// 	usersRepository.On("UpdateUser", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(errRepository).Once()

	// 	err := ticketsUsecase.Delete(context.Background(), 1, 1)

	// 	assert.Equal(t, errRepository, err)
	// })
}
