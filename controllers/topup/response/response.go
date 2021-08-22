package response

import (
	"ticketing/business/topup"
)

type TopUp struct {
	ID      int     `json:"id"`
	UserID  int     `json:"user_id"`
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
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
