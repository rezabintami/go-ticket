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

func (repository *mysqlMoviesRepository) Check(ctx context.Context, title string) error {
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
	result := repository.Conn.Where("title = ?", title).First(&movie)
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
