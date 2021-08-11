package request

import "ticketing/business/topup"

type TopUp struct {
	User_ID int     `json:"user_id"`
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
}

func (req *TopUp) ToDomain() *topup.Domain {
	return &topup.Domain{
		User_ID: req.User_ID,
		Name:    req.Name,
		Balance: req.Balance,
	}
}
