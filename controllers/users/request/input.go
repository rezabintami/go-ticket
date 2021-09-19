package request

import "ticketing/business/users"

type Users struct {
	Name     string  `json:"name"`
	Password string  `json:"password,omitempty"`
	Email    string  `json:"email"`
	Amount  float64 `json:"amount"`
	Language string  `json:"language"`
}

func (req *Users) ToDomain() *users.Domain {
	return &users.Domain{
		Name:     req.Name,
		Password: req.Password,
		Email:    req.Email,
		Amount:  req.Amount,
		Language: req.Language,
	}
}
