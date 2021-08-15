package topup

import (
	"context"
	"ticketing/business/topup"

	"gorm.io/gorm"
)

type mysqlTopUpRepository struct {
	Conn *gorm.DB
}

func NewMySQLTopUpRepository(conn *gorm.DB) topup.Repository {
	return &mysqlTopUpRepository{
		Conn: conn,
	}
}

func (repository *mysqlTopUpRepository) Store(ctx context.Context, topupDomain *topup.Domain) error {
	rec := fromDomain(*topupDomain)

	result := repository.Conn.Create(rec)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository *mysqlTopUpRepository) GetByID(ctx context.Context, id int) ([]topup.Domain, error) {
	top := []Topup{}
	result := repository.Conn.Where("user_id = ?", id).Find(&top)
	if result.Error != nil {
		return []topup.Domain{}, result.Error
	}
	var historyTopup []topup.Domain
	for _, value := range top {
		historyTopup = append(historyTopup, value.toDomain())
	}
	return historyTopup, nil
}
