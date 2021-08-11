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

func (repository *mysqlTopUpRepository) Payment(ctx context.Context, topupDomain *topup.Domain) error {
	rec := fromDomain(*topupDomain)

	result := repository.Conn.Create(rec)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
