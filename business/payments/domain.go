package payments

import "context"

type Domain struct {
	ID          int
	Title       string
	MovieID		int64
	Language    string
	Description string
	Path        string
	VoteAverage float64
	VoteCount   int64
}

type DomainResponse struct {
	Token         string   
	RedirectURL   string   
	StatusCode    string  
	ErrorMessages []string 
}

type Repository interface {
	Transactions(ctx context.Context, data *Domain) (DomainResponse, error)
}
