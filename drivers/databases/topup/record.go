package topup

import (
	"ticketing/business/payments"
	"ticketing/business/topup"
	"ticketing/drivers/databases/users"
	"time"
)

type Topup struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	OrderID     string `json:"order_id"`
	User        users.Users
	PaymentName string    `json:"payment_name"`
	Amount      float64   `json:"amount"`
	Status      string    `json:"status"`
	PaymentUrl  string    `json:"payment_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func fromDomain(topupDomain topup.Domain) *Topup {
	return &Topup{
		ID:         topupDomain.ID,
		UserID:     topupDomain.UserID,
		OrderID:    topupDomain.OrderID,
		Amount:     topupDomain.Amount,
		Status:     topupDomain.Status,
		PaymentUrl: topupDomain.PaymentUrl,
		CreatedAt:  topupDomain.CreatedAt,
		UpdatedAt:  topupDomain.UpdatedAt,
	}
}

func (rec *Topup) toDomain() topup.Domain {
	return topup.Domain{
		ID:          rec.ID,
		UserID:      rec.UserID,
		OrderID:     rec.OrderID,
		PaymentName: rec.PaymentName,
		Name:        rec.User.Name,
		Amount:      rec.Amount,
		Status:      rec.Status,
		PaymentUrl:  rec.PaymentUrl,
		CreatedAt:   rec.CreatedAt,
		UpdatedAt:   rec.UpdatedAt,
	}
}

func (rec *Topup) toPaymentDomain() payments.Domain {
	return payments.Domain{
		ID:          rec.ID,
		UserID:      rec.UserID,
		OrderID:     rec.OrderID,
		PaymentName: rec.PaymentName,
		Name:        rec.User.Name,
		Email:       rec.User.Email,
		Amount:      rec.Amount,
	}
}
