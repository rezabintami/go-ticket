package response

import (
	"ticketing/business/payments"
	"ticketing/business/topup"
)

type TopUp struct {
	ID      int     `json:"id"`
	UserID  int     `json:"user_id"`
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
}

type Payment struct {
	Token         string   `json:"token"`
	RedirectURL   string   `json:"redirect_url"`
	StatusCode    string   `json:"status_code"`
	ErrorMessages []string `json:"error"`
}

func FromDomain(topupDomain []topup.Domain) []TopUp {
	topup := []TopUp{}
	for _, value := range topupDomain {
		top := TopUp{
			ID:     value.ID,
			Name:   value.Name,
			UserID: value.UserID,
		}
		topup = append(topup, top)
	}
	return topup
}

func FromPaymentDomain(paymentsDomain payments.DomainResponse) Payment {
	return Payment{
		Token:         paymentsDomain.Token,
		RedirectURL:   paymentsDomain.RedirectURL,
		StatusCode:    paymentsDomain.StatusCode,
		ErrorMessages: paymentsDomain.ErrorMessages,
	}
}
