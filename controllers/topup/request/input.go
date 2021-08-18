package request

import "ticketing/business/topup"

type TopUp struct {
	UserID int     `json:"user_id"`
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
}

func (req *TopUp) ToDomain() *topup.Domain {
	return &topup.Domain{
		UserID: req.UserID,
		Name:    req.Name,
		Balance: req.Balance,
	}
}
