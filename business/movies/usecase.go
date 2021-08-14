package movies

import (
	"context"
	"ticketing/business/moviedb"
	"time"
)

type MoviesUsecase struct {
	moviesRepository  Repository
	contextTimeout    time.Duration
	movieDBRepository moviedb.Repository
}

func NewMoviesUsecase(mr Repository, timeout time.Duration, dbr moviedb.Repository) Usecase {
	return &MoviesUsecase{
		moviesRepository:  mr,
		contextTimeout:    timeout,
		movieDBRepository: dbr,
	}
}

func (mu *MoviesUsecase) Fetch(ctx context.Context, urlsearch string, search string) ([]Domain, error) {
	//! Ambil Data Dari MOVIEDB
	//! Taruh Dalam Array
	//! Check Jika ada di DB
	//! Store DB
	result, err := mu.movieDBRepository.GetMovies(ctx, urlsearch)
	if err != nil {
		return []Domain{}, err
	}
	for _, value := range result {
		//! Check
		err := mu.moviesRepository.Check(ctx, value.MovieID)
		if err != nil {
			//! Store
			mu.moviesRepository.Store(ctx, &value)
		}
	}
	res, err := mu.moviesRepository.Search(ctx, search)
	if err != nil {
		return []Domain{}, err
	}
	return res, nil
}


func (mu *MoviesUsecase) GetByID(ctx context.Context, id int) (Domain, error) {
	result, err := mu.moviesRepository.GetByID(ctx, id)
	if err != nil {
		return Domain{}, err
	}
	return result, nil
}