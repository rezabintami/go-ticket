package movies

import (
	"context"
	"ticketing/business/moviedb"
	"ticketing/business/movies"

	"gorm.io/gorm"
)

type mysqlMoviesRepository struct {
	Conn *gorm.DB
}

func NewMySQLMoviesRepository(conn *gorm.DB) movies.Repository {
	return &mysqlMoviesRepository{
		Conn: conn,
	}
}

func (repository *mysqlMoviesRepository) Check(ctx context.Context, id int64) error {
	// for _, value := range data {
	// 	rec := fromDomain(value)
	// 	movie := Movie{}
	// 	result := repository.Conn.Where("title = ?", rec.Title).First(&movie)
	// 	fmt.Println(result.Error)
	// 	if result.Error != nil {
	// 		result := repository.Conn.Create(rec)
	// 		if result.Error != nil {
	// 			return result.Error
	// 		}
	// 	}
	// }
	movie := Movie{}
	result := repository.Conn.Where("movie_id = ?", id).First(&movie)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository *mysqlMoviesRepository) Store(ctx context.Context, data *moviedb.Domain) error {
	rec := fromDomain(*data)
	result := repository.Conn.Create(rec)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repository *mysqlMoviesRepository) Search(ctx context.Context, title string) ([]movies.Domain, error) {
	movie := []Movie{}
	result := repository.Conn.Where("title like ?", "%"+title+"%").Find(&movie)
	if result.Error != nil {
		return []movies.Domain{}, result.Error
	}
	// fmt.Println(movie)
	var domainMovies []movies.Domain
	for _, value := range movie {
		domainMovies = append(domainMovies, value.toDomain())
	}
	return domainMovies, nil
}

func (repository *mysqlMoviesRepository) GetByID(ctx context.Context, id int) (movies.Domain, error) {
	movie := Movie{}
	result := repository.Conn.Where("id = ?", id).First(&movie)
	if result.Error != nil {
		return movies.Domain{}, result.Error
	}
	return movie.toDomain(), nil
}