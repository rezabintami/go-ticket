package users

import (
	"ticketing/business/users"
	"time"
)

type Users struct {
	ID        int       `gorm:"primary_key" json:"id"`
	Name      string    `json:"name"`
	Password  string    `json:"-"`
	Email     string    `json:"email"`
	Balance   float64   `json:"balance"`
	Language  string    `json:"language"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (rec *Users) toDomain() users.Domain {
	return users.Domain{
		ID:        rec.ID,
		Name:      rec.Name,
		Password:  rec.Password,
		Balance:   rec.Balance,
		Email:     rec.Email,
		Language:  rec.Language,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func fromDomain(userDomain users.Domain) *Users {
	return &Users{
		ID:        userDomain.ID,
		Name:      userDomain.Name,
		Password:  userDomain.Password,
		Balance:   userDomain.Balance,
		Email:     userDomain.Email,
		Language:  userDomain.Language,
		CreatedAt: userDomain.CreatedAt,
		UpdatedAt: userDomain.UpdatedAt,
	}
}
