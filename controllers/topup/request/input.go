package request

import (
	"ticketing/business/payments"
	"ticketing/business/topup"
)

type TopUp struct {
	PaymentName string  `json:"payment_name"`
	OrderID     string  `json:"order_id"`
	Amount      float64 `json:"amount"`
}

func (req *TopUp) ToDomain() *topup.Domain {
	return &topup.Domain{
		OrderID:     req.OrderID,
		PaymentName: req.PaymentName,
		Amount:      req.Amount,
	}
}

func (req *TopUp) ToPaymentDomain() *payments.Domain {
	return &payments.Domain{
		OrderID:     req.OrderID,
		PaymentName: req.PaymentName,
		Amount:      req.Amount,
	}
}
