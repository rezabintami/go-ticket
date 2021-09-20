package topup

import (
	"context"
	"ticketing/business/payments"
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

func (repository *mysqlTopUpRepository) Store(ctx context.Context, topupDomain *topup.Domain) (payments.Domain, error) {
	rec := fromDomain(*topupDomain)

	result := repository.Conn.Create(rec)
	if result.Error != nil {
		return payments.Domain{}, result.Error
	}

	err := repository.Conn.Preload("User").First(&rec, rec.ID).Error
	if err != nil {
		return payments.Domain{}, result.Error
	}

	return rec.toPaymentDomain(), nil
}

func (repository *mysqlTopUpRepository) Update(ctx context.Context, topupDomain *topup.Domain) error {
	rec := fromDomain(*topupDomain)

	result := repository.Conn.Where("order_id = ?", rec.OrderID).Updates(&rec)
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

func (repository *mysqlTopUpRepository) GetByOrder(ctx context.Context, orderId string) (topup.Domain, error) {
	top := Topup{}
	result := repository.Conn.Where("order_id = ?", orderId).Find(&top)
	if result.Error != nil {
		return topup.Domain{}, result.Error
	}
	return top.toDomain(), nil
}
