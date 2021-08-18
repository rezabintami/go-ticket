package movies_test

import (
	"context"
	"errors"
	"net/url"
	"os"
	"testing"
	"ticketing/business/moviedb"
	moviedbMock "ticketing/business/moviedb/mocks"
	movies "ticketing/business/movies"
	moviesMock "ticketing/business/movies/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	moviesRepository  moviesMock.Repository
	moviesUsecase     movies.Usecase
	moviedbRepository moviedbMock.Repository
)

func setup() {
	moviesUsecase = movies.NewMoviesUsecase(&moviesRepository, 2, &moviedbRepository)
}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestGetByID(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		domain := movies.Domain{
			ID:          2,
			Title:       "Tinker Bell and the Lost Treasure",
			Language:    "en",
			MovieID:     25475,
			Description: "A blue harvest moon will rise, allowing the fairies to use a precious moonstone to restore the Pixie Dust Tree, the source of all their magic. But when Tinker Bell accidentally puts all of Pixie Hollow in jeopardy, she must venture out across the sea on a secret quest to set things right.",
			Path:        "/hg1959yuBkHb4BKbIvETQSfxGCT.jpg",
			VoteAverage: 6.8,
			VoteCount:   721,
		}
		moviesRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(domain, nil).Once()

		result, err := moviesUsecase.GetByID(context.Background(), 1)

		assert.Nil(t, err)
		assert.Equal(t, domain.Title, result.Title)
	})

	t.Run("test case 2, data not found", func(t *testing.T) {
		errRepository := errors.New("data not found")
		moviesRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(movies.Domain{}, errRepository).Once()

		result, err := moviesUsecase.GetByID(context.Background(), -1)

		assert.Equal(t, errRepository, err)
		assert.Equal(t, result, movies.Domain{})
	})
}

func TestFetch(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		moviedbDomain := []moviedb.Domain{
			{
				ID:          2,
				Title:       "Tinker Bell and the Lost Treasure",
				Language:    "en",
				MovieID:     25475,
				Description: "A blue harvest moon will rise, allowing the fairies to use a precious moonstone to restore the Pixie Dust Tree, the source of all their magic. But when Tinker Bell accidentally puts all of Pixie Hollow in jeopardy, she must venture out across the sea on a secret quest to set things right.",
				Path:        "/hg1959yuBkHb4BKbIvETQSfxGCT.jpg",
				VoteAverage: 6.8,
				VoteCount:   721,
			},
			{
				ID:          3,
				Title:       "The Treasure of the Sierra Madre",
				Language:    "en",
				MovieID:     3090,
				Description: "Fred C. Dobbs and Bob Curtin, both down on their luck in Tampico, Mexico in 1925, meet up with a grizzled prospector named Howard and decide to join with him in search of gold in the wilds of central Mexico. Through enormous difficulties, they eventually succeed in finding gold, but bandits, the elements, and most especially greed threaten to turn their success into disaster.",
				Path:        "/58PUObt4eQTczaevgxJT8iGKoMa.jpg",
				VoteAverage: 8.1,
				VoteCount:   771,
			},
		}
		movieDomain := []movies.Domain{
			{
				ID:          2,
				Title:       "Tinker Bell and the Lost Treasure",
				Language:    "en",
				MovieID:     25475,
				Description: "A blue harvest moon will rise, allowing the fairies to use a precious moonstone to restore the Pixie Dust Tree, the source of all their magic. But when Tinker Bell accidentally puts all of Pixie Hollow in jeopardy, she must venture out across the sea on a secret quest to set things right.",
				Path:        "/hg1959yuBkHb4BKbIvETQSfxGCT.jpg",
				VoteAverage: 6.8,
				VoteCount:   721,
			},
			{
				ID:          3,
				Title:       "The Treasure of the Sierra Madre",
				Language:    "en",
				MovieID:     3090,
				Description: "Fred C. Dobbs and Bob Curtin, both down on their luck in Tampico, Mexico in 1925, meet up with a grizzled prospector named Howard and decide to join with him in search of gold in the wilds of central Mexico. Through enormous difficulties, they eventually succeed in finding gold, but bandits, the elements, and most especially greed threaten to turn their success into disaster.",
				Path:        "/58PUObt4eQTczaevgxJT8iGKoMa.jpg",
				VoteAverage: 8.1,
				VoteCount:   771,
			},
		}

		errRepository := errors.New("data duplicate found")

		moviedbRepository.On("GetMovies", mock.Anything, mock.AnythingOfType("string")).Return(moviedbDomain, nil).Once()
		moviesRepository.On("Check", mock.Anything, mock.Anything).Return(errRepository)
		moviesRepository.On("Store", mock.Anything, mock.Anything).Return(nil)
		moviesRepository.On("Search", mock.Anything, mock.AnythingOfType("string")).Return(movieDomain, nil).Once()

		result, err := moviesUsecase.Fetch(context.Background(), url.QueryEscape("Treasure"), "Treasure")

		assert.Nil(t, err)
		assert.Equal(t, 2, len(result))
	})

	t.Run("test case 2, data not found", func(t *testing.T) {
		errRepository := errors.New("data not found")
		moviedbRepository.On("GetMovies", mock.Anything, mock.AnythingOfType("string")).Return([]moviedb.Domain{}, errRepository).Once()

		result, err := moviesUsecase.Fetch(context.Background(), url.QueryEscape("Treasure"), "Treasure")

		assert.Equal(t, errRepository, err)
		assert.Equal(t, result, []movies.Domain{})
	})

	t.Run("test case 3, error searching", func(t *testing.T) {
		moviedbDomain := []moviedb.Domain{
			{
				ID:          2,
				Title:       "Tinker Bell and the Lost Treasure",
				Language:    "en",
				MovieID:     25475,
				Description: "A blue harvest moon will rise, allowing the fairies to use a precious moonstone to restore the Pixie Dust Tree, the source of all their magic. But when Tinker Bell accidentally puts all of Pixie Hollow in jeopardy, she must venture out across the sea on a secret quest to set things right.",
				Path:        "/hg1959yuBkHb4BKbIvETQSfxGCT.jpg",
				VoteAverage: 6.8,
				VoteCount:   721,
			},
			{
				ID:          3,
				Title:       "The Treasure of the Sierra Madre",
				Language:    "en",
				MovieID:     3090,
				Description: "Fred C. Dobbs and Bob Curtin, both down on their luck in Tampico, Mexico in 1925, meet up with a grizzled prospector named Howard and decide to join with him in search of gold in the wilds of central Mexico. Through enormous difficulties, they eventually succeed in finding gold, but bandits, the elements, and most especially greed threaten to turn their success into disaster.",
				Path:        "/58PUObt4eQTczaevgxJT8iGKoMa.jpg",
				VoteAverage: 8.1,
				VoteCount:   771,
			},
		}
		movieDomain := []movies.Domain{
			{
				ID:          2,
				Title:       "Tinker Bell and the Lost Treasure",
				Language:    "en",
				MovieID:     25475,
				Description: "A blue harvest moon will rise, allowing the fairies to use a precious moonstone to restore the Pixie Dust Tree, the source of all their magic. But when Tinker Bell accidentally puts all of Pixie Hollow in jeopardy, she must venture out across the sea on a secret quest to set things right.",
				Path:        "/hg1959yuBkHb4BKbIvETQSfxGCT.jpg",
				VoteAverage: 6.8,
				VoteCount:   721,
			},
			{
				ID:          3,
				Title:       "The Treasure of the Sierra Madre",
				Language:    "en",
				MovieID:     3090,
				Description: "Fred C. Dobbs and Bob Curtin, both down on their luck in Tampico, Mexico in 1925, meet up with a grizzled prospector named Howard and decide to join with him in search of gold in the wilds of central Mexico. Through enormous difficulties, they eventually succeed in finding gold, but bandits, the elements, and most especially greed threaten to turn their success into disaster.",
				Path:        "/58PUObt4eQTczaevgxJT8iGKoMa.jpg",
				VoteAverage: 8.1,
				VoteCount:   771,
			},
		}

		errRepository := errors.New("error searching")

		moviedbRepository.On("GetMovies", mock.Anything, mock.AnythingOfType("string")).Return(moviedbDomain, nil).Once()
		moviesRepository.On("Check", mock.Anything, mock.Anything).Return(errRepository)
		moviesRepository.On("Store", mock.Anything, mock.Anything).Return(nil)
		moviesRepository.On("Search", mock.Anything, mock.AnythingOfType("string")).Return(movieDomain, errRepository).Once()

		result, err := moviesUsecase.Fetch(context.Background(), url.QueryEscape("Treasure"), "Treasure")

		assert.Equal(t, errRepository, err)
		assert.Equal(t, result, []movies.Domain{})
	})
}
