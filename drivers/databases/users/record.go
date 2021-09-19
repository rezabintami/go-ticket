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
	Amount    float64   `json:"amount"`
	Language  string    `json:"language"`
	Sso       bool      `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (rec *Users) toDomain() users.Domain {
	return users.Domain{
		ID:        rec.ID,
		Name:      rec.Name,
		Password:  rec.Password,
		Amount:    rec.Amount,
		Email:     rec.Email,
		Language:  rec.Language,
		Sso:       rec.Sso,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func fromDomain(userDomain users.Domain) *Users {
	return &Users{
		ID:        userDomain.ID,
		Name:      userDomain.Name,
		Password:  userDomain.Password,
		Amount:    userDomain.Amount,
		Email:     userDomain.Email,
		Language:  userDomain.Language,
		Sso:       userDomain.Sso,
		CreatedAt: userDomain.CreatedAt,
		UpdatedAt: userDomain.UpdatedAt,
	}
}
