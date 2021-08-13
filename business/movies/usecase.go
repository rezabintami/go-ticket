package movies

import (
	"context"
	"fmt"
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

func (mu *MoviesUsecase) Store(ctx context.Context, search string) error {
	//! Ambil Data Dari MOVIEDB
	//! Taruh Dalam Array
	//! Check Jika ada di DB
	//! Store DB
	result, err := mu.movieDBRepository.GetMovies(ctx, search)
	if err != nil {
		return err
	}
	fmt.Println(result)
	for _, value := range result {
		//! Check
		err := mu.moviesRepository.Check(ctx, value.Title)
		if err != nil {
			//! Store
			go mu.moviesRepository.Store(ctx, &value)
		}
	}
	// err = mu.moviesRepository.CheckAndStoreMovies(ctx, result)
	// if err != nil {
	// 	return err
	// }
	// mu.moviesRepository.CheckMovies(ctx, result)
	return nil
}
