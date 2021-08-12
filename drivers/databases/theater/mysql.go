package theater

import (
	"context"
	"ticketing/business/theater"

	"gorm.io/gorm"
)

type mysqlTheaterRepository struct {
	Conn *gorm.DB
}

func NewMySQLTheaterRepository(conn *gorm.DB) theater.Repository {
	return &mysqlTheaterRepository{
		Conn: conn,
	}
}

func (repository *mysqlTheaterRepository) Store(ctx context.Context, theaterDomain *theater.Domain) error {
	rec := fromDomain(*theaterDomain)

	result := repository.Conn.Create(rec)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository *mysqlTheaterRepository) Delete(ctx context.Context, id int) error {
	theaterDelete := Theater{}
	result := repository.Conn.Where("id = ?", id).Delete(&theaterDelete)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository *mysqlTheaterRepository) Update(ctx context.Context, theaterDomain *theater.Domain, id int) error {
	rec := fromDomain(*theaterDomain)

	result := repository.Conn.Where("id = ?", id).Updates(&rec)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository *mysqlTheaterRepository) GetAll(ctx context.Context, theaterDomain *theater.Domain) ([]theater.Domain, error) {
	var rec []Theater

	result := repository.Conn.Find(&rec)
	if result.Error != nil {
		return []theater.Domain{}, result.Error
	}
	var allTheater []theater.Domain
	for _, value := range rec {
		allTheater = append(allTheater, value.toDomain())
	}
	return allTheater, nil
}
