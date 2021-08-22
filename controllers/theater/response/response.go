package response

import (
	"ticketing/business/theater"
)

type Theater struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Place string `json:"place"`
}

func FromDomain(theaterDomain []theater.Domain) []Theater {
	theaters := []Theater{}
	for _, value := range theaterDomain {
		theat := Theater{
			ID:    value.ID,
			Name:  value.Name,
			Place: value.Place,
		}
		theaters = append(theaters, theat)
	}
	return theaters
}
