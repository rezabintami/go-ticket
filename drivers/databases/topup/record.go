package topup

import (
	"ticketing/business/topup"
	"ticketing/drivers/databases/users"
	"time"
)

type Topup struct {
	ID        int `json:"id"`
	UserID    int `json:"user_id"`
	User      users.Users
	Name      string    `json:"name"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (rec *Topup) toDomain() topup.Domain {
	return topup.Domain{
		ID:        rec.ID,
		UserID:    rec.UserID,
		Name:      rec.Name,
		Balance:   rec.Balance,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func fromDomain(topupDomain topup.Domain) *Topup {
	return &Topup{
		ID:        topupDomain.ID,
		UserID:    topupDomain.UserID,
		Name:      topupDomain.Name,
		Balance:   topupDomain.Balance,
		CreatedAt: topupDomain.CreatedAt,
		UpdatedAt: topupDomain.UpdatedAt,
	}
}
