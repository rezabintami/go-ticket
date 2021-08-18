package request

import "ticketing/business/theater"

type Theater struct {
	Name  string `json:"name"`
	Place string `json:"place"`
}

func (req *Theater) ToDomain() *theater.Domain {
	return &theater.Domain{
		Name:  req.Name,
		Place: req.Place,
	}
}
